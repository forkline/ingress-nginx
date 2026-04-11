/*
Copyright 2016 The Kubernetes Authors.

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

package config

import (
	"encoding/json"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultSnapshot(t *testing.T) {
	cfg := NewDefault()
	cfg.WorkerProcesses = "8"

	data, err := json.MarshalIndent(cfg, "", "  ")
	assert.NoError(t, err, "serializing default config to JSON")

	goldenPath := "../../../../test/data/golden/default_config.json"

	if os.Getenv("UPDATE_GOLDEN") == "1" {
		err = os.MkdirAll("../../../../test/data/golden", 0o755)
		assert.NoError(t, err)
		err = os.WriteFile(goldenPath, data, 0o644)
		assert.NoError(t, err)
		t.Logf("updated golden file: %s", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if os.IsNotExist(err) {
		t.Fatalf("golden file %s does not exist. Run with UPDATE_GOLDEN=1 to create it.", goldenPath)
	}
	assert.NoError(t, err, "reading golden file")

	var expectedMap map[string]interface{}
	var actualMap map[string]interface{}

	err = json.Unmarshal(expected, &expectedMap)
	assert.NoError(t, err, "parsing expected golden JSON")

	err = json.Unmarshal(data, &actualMap)
	assert.NoError(t, err, "parsing actual config JSON")

	assert.Equal(t, expectedMap, actualMap, "NewDefault() config does not match golden snapshot")
}

func TestNewDefaultDeterministic(t *testing.T) {
	cfg1 := NewDefault()
	cfg2 := NewDefault()

	data1, err := json.Marshal(cfg1)
	assert.NoError(t, err)

	data2, err := json.Marshal(cfg2)
	assert.NoError(t, err)

	assert.Equal(t, string(data1), string(data2), "NewDefault() should produce deterministic output")
}

func TestNewDefaultContainsExpectedFields(t *testing.T) {
	cfg := NewDefault()

	assert.Equal(t, true, cfg.HSTS, "HSTS should be enabled by default")
	assert.Equal(t, "31536000", cfg.HSTSMaxAge, "HSTS max age default")
	assert.Equal(t, true, cfg.HSTSIncludeSubdomains, "HSTS include subdomains default")
	assert.Equal(t, false, cfg.HSTSPreload, "HSTS preload default")
	assert.Equal(t, "TLSv1.2 TLSv1.3", cfg.SSLProtocols, "SSL protocols default")
	assert.Equal(t, true, cfg.SSLSessionCache, "SSL session cache default")
	assert.Equal(t, "10m", cfg.SSLSessionCacheSize, "SSL session cache size default")
	assert.Equal(t, true, cfg.UseHTTP2, "HTTP2 should be enabled by default")
	assert.Equal(t, 75, cfg.KeepAlive, "keep-alive default")
	assert.Equal(t, 16384, cfg.MaxWorkerConnections, "max worker connections default")
	assert.Equal(t, true, cfg.EnableMultiAccept, "multi accept default")
	assert.Equal(t, true, cfg.ReusePort, "reuse port default")
	assert.Equal(t, true, cfg.SSLRedirect, "SSL redirect default (from Backend)")

	sort.Strings(cfg.ProxyRealIPCIDR)
	assert.Equal(t, []string{"0.0.0.0/0"}, cfg.ProxyRealIPCIDR, "proxy real IP CIDR default")

	assert.Equal(t, []string{"127.0.0.1"}, cfg.NginxStatusIpv4Whitelist, "nginx status IPv4 whitelist default")
	assert.Equal(t, []string{"::1"}, cfg.NginxStatusIpv6Whitelist, "nginx status IPv6 whitelist default")
}
