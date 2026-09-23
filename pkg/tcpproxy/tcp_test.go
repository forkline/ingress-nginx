/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tcpproxy

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

const (
	passthroughHost = "passthrough.example.test"
	otherHost       = "other.example.test"
	sinkPassthrough = "passthrough"
	sinkDefault     = "default"
)

// clientHello captures the raw TLS ClientHello record a Go TLS client sends
// for the given server name.
func clientHello(t *testing.T, serverName string) []byte {
	t.Helper()
	clientSide, serverSide := net.Pipe()
	defer clientSide.Close()
	defer serverSide.Close()

	go func() {
		//nolint:errcheck,gosec // handshake is expected to fail once we stop reading
		tls.Client(clientSide, &tls.Config{ServerName: serverName, InsecureSkipVerify: true}).Handshake()
	}()

	record := make([]byte, 5, tlsRecordHeaderLen+maxTLSRecordLen)
	if _, err := io.ReadFull(serverSide, record); err != nil {
		t.Fatalf("reading record header: %v", err)
	}
	record = record[:5+binary.BigEndian.Uint16(record[3:5])]
	if _, err := io.ReadFull(serverSide, record[5:]); err != nil {
		t.Fatalf("reading record body: %v", err)
	}
	return record
}

type received struct {
	listener string
	data     []byte
}

// startSink listens on loopback and reports everything the first accepted
// connection sends (until the peer closes) on out, tagged with name.
func startSink(t *testing.T, name string, out chan<- received) *TCPServer {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	t.Cleanup(func() { ln.Close() })

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		//nolint:errcheck // best effort: the peer closes when Handle returns
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		//nolint:errcheck // a read deadline may cut the read short; forward whatever arrived
		data, _ := io.ReadAll(conn)
		out <- received{listener: name, data: data}
	}()

	addr, ok := ln.Addr().(*net.TCPAddr)
	if !ok {
		t.Fatalf("unexpected listener address type %T", ln.Addr())
	}
	return &TCPServer{Hostname: passthroughHost, IP: addr.IP.String(), Port: addr.Port}
}

func newTestProxy(t *testing.T) (*TCPProxy, <-chan received) {
	t.Helper()
	out := make(chan received, 2)
	passthrough := startSink(t, sinkPassthrough, out)
	def := startSink(t, sinkDefault, out)
	def.Hostname = ""
	return &TCPProxy{ServerList: []*TCPServer{passthrough}, Default: def}, out
}

// runHandle feeds hello to Handle in the given fragments (split offsets) and
// returns which sink received the bytes.
func runHandle(t *testing.T, p *TCPProxy, out <-chan received, hello []byte, splits ...int) received {
	t.Helper()
	clientSide, serverSide := net.Pipe()
	done := make(chan struct{})
	go func() {
		p.Handle(serverSide)
		close(done)
	}()

	prev := 0
	for _, s := range append(splits, len(hello)) {
		if _, err := clientSide.Write(hello[prev:s]); err != nil {
			t.Fatalf("write fragment: %v", err)
		}
		prev = s
		time.Sleep(50 * time.Millisecond)
	}
	clientSide.Close()

	select {
	case r := <-out:
		<-done
		return r
	case <-time.After(3 * time.Second):
		t.Fatal("no sink received data")
	}
	return received{}
}

func TestHandle_UnfragmentedClientHelloRoutesBySNI(t *testing.T) {
	p, out := newTestProxy(t)
	hello := clientHello(t, passthroughHost)
	r := runHandle(t, p, out, hello)
	if r.listener != sinkPassthrough || !bytes.Equal(r.data, hello) {
		t.Fatalf("got %s (%d bytes), want passthrough with full hello (%d bytes)", r.listener, len(r.data), len(hello))
	}
}

func TestHandle_FragmentedClientHelloRoutesBySNI(t *testing.T) {
	hello := clientHello(t, passthroughHost)
	for _, split := range []int{1, 3, 5, len(hello) / 2} {
		t.Run(fmt.Sprintf("split-at-%d", split), func(t *testing.T) {
			p, out := newTestProxy(t)
			r := runHandle(t, p, out, hello, split)
			if r.listener != sinkPassthrough {
				t.Fatalf("fragmented ClientHello misrouted to %s", r.listener)
			}
			if !bytes.Equal(r.data, hello) {
				t.Fatalf("passthrough received %d bytes, want %d", len(r.data), len(hello))
			}
		})
	}
}

func TestHandle_UnknownSNIFallsThroughToDefault(t *testing.T) {
	p, out := newTestProxy(t)
	hello := clientHello(t, otherHost)
	r := runHandle(t, p, out, hello, 1)
	if r.listener != sinkDefault || !bytes.Equal(r.data, hello) {
		t.Fatalf("got %s (%d bytes), want default with full hello (%d bytes)", r.listener, len(r.data), len(hello))
	}
}

func TestHandle_NonTLSBytesFallThroughToDefault(t *testing.T) {
	p, out := newTestProxy(t)
	r := runHandle(t, p, out, []byte("GET / HTTP/1.1\r\n"))
	if r.listener != sinkDefault || len(r.data) == 0 {
		t.Fatalf("got %s (%d bytes), want default", r.listener, len(r.data))
	}
}

func TestHandle_OversizedRecordFallsThroughToDefault(t *testing.T) {
	p, out := newTestProxy(t)
	header := []byte{0x16, 0x03, 0x01, 0xff, 0xff}
	r := runHandle(t, p, out, header)
	if r.listener != sinkDefault || !bytes.Equal(r.data, header) {
		t.Fatalf("got %s (%d bytes), want default with the 5 header bytes", r.listener, len(r.data))
	}
}

func TestHandle_ProxyProtocolHeader(t *testing.T) {
	p, out := newTestProxy(t)
	p.ServerList[0].ProxyProtocol = true
	hello := clientHello(t, passthroughHost)

	// Handle reads the peer addresses for the PROXY header, so this needs a
	// real TCP connection rather than net.Pipe.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	defer ln.Close()
	serverSideCh := make(chan net.Conn, 1)
	go func() {
		conn, err := ln.Accept()
		if err == nil {
			serverSideCh <- conn
		}
	}()
	clientSide, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	serverSide := <-serverSideCh

	done := make(chan struct{})
	go func() {
		p.Handle(serverSide)
		close(done)
	}()
	if _, err := clientSide.Write(hello); err != nil {
		t.Fatalf("write hello: %v", err)
	}
	clientSide.Close()

	var r received
	select {
	case r = <-out:
		<-done
	case <-time.After(3 * time.Second):
		t.Fatal("no sink received data")
	}
	if r.listener != sinkPassthrough {
		t.Fatalf("got %s, want passthrough", r.listener)
	}
	if !bytes.HasPrefix(r.data, []byte("PROXY TCP4 ")) {
		t.Fatalf("missing PROXY protocol header, got %q", r.data)
	}
	if !bytes.HasSuffix(r.data, hello) {
		t.Fatalf("payload does not end with the ClientHello (%d bytes)", len(hello))
	}
}

// runHandleClosed feeds fragments to Handle and expects it to return without
// any sink receiving data.
func runHandleClosed(t *testing.T, p *TCPProxy, out <-chan received, fragments ...[]byte) {
	t.Helper()
	clientSide, serverSide := net.Pipe()
	done := make(chan struct{})
	go func() {
		p.Handle(serverSide)
		close(done)
	}()

	for _, f := range fragments {
		if _, err := clientSide.Write(f); err != nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	clientSide.Close()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("Handle did not return")
	}
	select {
	case r := <-out:
		t.Fatalf("sink %s unexpectedly received %d bytes", r.listener, len(r.data))
	case <-time.After(200 * time.Millisecond):
	}
}

func TestHandle_ReadErrorReturns(t *testing.T) {
	p, out := newTestProxy(t)
	runHandleClosed(t, p, out)
}

func TestHandle_TruncatedRecordBodyReturns(t *testing.T) {
	p, out := newTestProxy(t)
	runHandleClosed(t, p, out, []byte{0x16, 0x03, 0x01, 0x00, 0x64})
}

func TestHandle_NoBackendForUnknownSNI(t *testing.T) {
	out := make(chan received, 1)
	passthrough := startSink(t, sinkPassthrough, out)
	p := &TCPProxy{ServerList: []*TCPServer{passthrough}}
	runHandleClosed(t, p, out, clientHello(t, otherHost))
}

func TestHandle_DialErrorReturns(t *testing.T) {
	out := make(chan received, 1)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	addr, ok := ln.Addr().(*net.TCPAddr)
	if !ok {
		t.Fatalf("unexpected listener address type %T", ln.Addr())
	}
	port := addr.Port
	ln.Close()
	p := &TCPProxy{Default: &TCPServer{IP: "127.0.0.1", Port: port}}
	runHandleClosed(t, p, out, []byte("GET / HTTP/1.1\r\n"))
}
