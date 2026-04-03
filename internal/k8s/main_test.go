/*
Copyright 2015 The Kubernetes Authors.

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

package k8s

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	fakediscovery "k8s.io/client-go/discovery/fake"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestParseNameNS(t *testing.T) {
	tests := []struct {
		title  string
		input  string
		ns     string
		name   string
		expErr bool
	}{
		{"empty string", "", "", "", true},
		{"demo", "demo", "", "", true},
		{"kube-system", "kube-system", "", "", true},
		{"default/kube-system", "default/kube-system", "default", "kube-system", false},
	}

	for _, test := range tests {
		ns, name, err := ParseNameNS(test.input)
		if test.expErr {
			if err == nil {
				t.Errorf("%v: expected error but returned nil", test.title)
			}
			continue
		}
		if test.ns != ns {
			t.Errorf("%v: expected %v but returned %v", test.title, test.ns, ns)
		}
		if test.name != name {
			t.Errorf("%v: expected %v but returned %v", test.title, test.name, name)
		}
	}
}

func TestGetNodeIP(t *testing.T) {
	fKNodes := []struct {
		name          string
		cs            *testclient.Clientset
		nodeName      string
		ea            []string
		useInternalIP bool
	}{
		{
			"empty node list",
			testclient.NewSimpleClientset(),
			"demo",
			[]string{},
			true,
		},
		{
			"node does not exist",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						},
					},
				},
			}}}), "notexistnode",
			[]string{},
			true,
		},
		{
			"node exists and only has an internal IP address (useInternalIP=false)",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						},
					},
				},
			}}}), "demo",
			[]string{"10.0.0.1"},
			false,
		},
		{
			"node exists has an internal IP address and an empty external IP address (useInternalIP=false)",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeExternalIP,
							Address: "",
						},
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						},
					},
				},
			}}}), "demo",
			[]string{"10.0.0.1"},
			false,
		},
		{
			"node exists and has two internal IP address (useInternalIP=false)",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						},
						{
							Type:    apiv1.NodeInternalIP,
							Address: "fd00::1",
						},
					},
				},
			}}}), "demo",
			[]string{"10.0.0.1", "fd00::1"},
			false,
		},
		{
			"node exists and only has an internal IP address",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						},
					},
				},
			}}}), "demo",
			[]string{"10.0.0.1"},
			true,
		},
		{
			"node exists and has two internal IP address",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						},
						{
							Type:    apiv1.NodeInternalIP,
							Address: "fd00::1",
						},
					},
				},
			}}}), "demo",
			[]string{"10.0.0.1", "fd00::1"},
			true,
		},
		{
			"node exist and only has an external IP address",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeExternalIP,
							Address: "10.0.0.1",
						},
					},
				},
			}}}), "demo",
			[]string{"10.0.0.1"},
			false,
		},
		{
			"node exist and only has an external IP address (useInternalIP=true)",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeExternalIP,
							Address: "10.0.0.1",
						},
					},
				},
			}}}), "demo",
			[]string{},
			true,
		},
		{
			"multiple nodes - choose the right one",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "demo1",
					},
					Status: apiv1.NodeStatus{
						Addresses: []apiv1.NodeAddress{
							{
								Type:    apiv1.NodeInternalIP,
								Address: "10.0.0.1",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "demo2",
					},
					Status: apiv1.NodeStatus{
						Addresses: []apiv1.NodeAddress{
							{
								Type:    apiv1.NodeInternalIP,
								Address: "10.0.0.2",
							},
						},
					},
				},
			}}),
			"demo2",
			[]string{"10.0.0.2"},
			true,
		},
		{
			"node with both internal and external IP address - returns external IP",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.1",
						}, {
							Type:    apiv1.NodeExternalIP,
							Address: "10.0.0.2",
						},
					},
				},
			}}}),
			"demo",
			[]string{"10.0.0.2"},
			false,
		},
		{
			"node with both internal and external IP address - returns internal IP",
			testclient.NewSimpleClientset(&apiv1.NodeList{Items: []apiv1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "demo",
				},
				Status: apiv1.NodeStatus{
					Addresses: []apiv1.NodeAddress{
						{
							Type:    apiv1.NodeExternalIP,
							Address: "",
						}, {
							Type:    apiv1.NodeInternalIP,
							Address: "10.0.0.2",
						},
					},
				},
			}}}),
			"demo",
			[]string{"10.0.0.2"},
			true,
		},
	}

	for _, fk := range fKNodes {
		addresses := GetNodeIPs(fk.cs, fk.nodeName, fk.useInternalIP)
		if !slices.Equal(addresses, fk.ea) {
			t.Errorf("%v - expected %v, but returned %v", fk.name, fk.ea, addresses)
		}
	}
}

func TestGetIngressPod(t *testing.T) {
	// POD_NAME & POD_NAMESPACE not exist
	t.Setenv("POD_NAME", "")
	t.Setenv("POD_NAMESPACE", "")
	err := GetIngressPod(testclient.NewSimpleClientset())
	if err == nil {
		t.Errorf("expected an error but returned nil")
	}

	// POD_NAME not exist
	t.Setenv("POD_NAME", "")
	t.Setenv("POD_NAMESPACE", apiv1.NamespaceDefault)
	err = GetIngressPod(testclient.NewSimpleClientset())
	if err == nil {
		t.Errorf("expected an error but returned nil")
	}

	// POD_NAMESPACE not exist
	t.Setenv("POD_NAME", "testpod")
	t.Setenv("POD_NAMESPACE", "")
	err = GetIngressPod(testclient.NewSimpleClientset())
	if err == nil {
		t.Errorf("expected an error but returned nil")
	}

	// POD not exist
	t.Setenv("POD_NAME", "testpod")
	t.Setenv("POD_NAMESPACE", apiv1.NamespaceDefault)
	err = GetIngressPod(testclient.NewSimpleClientset())
	if err == nil {
		t.Errorf("expected an error but returned nil")
	}

	// success to get PodInfo
	fkClient := testclient.NewSimpleClientset(
		&apiv1.PodList{Items: []apiv1.Pod{{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "testpod",
				Namespace: apiv1.NamespaceDefault,
				Labels: map[string]string{
					"first":                       "first_label",
					"second":                      "second_label",
					"app.kubernetes.io/component": "controller",
					"app.kubernetes.io/instance":  "ingress-nginx",
					"app.kubernetes.io/name":      "ingress-nginx",
				},
			},
		}}},
		&apiv1.NodeList{Items: []apiv1.Node{{
			ObjectMeta: metav1.ObjectMeta{
				Name: "demo",
			},
			Status: apiv1.NodeStatus{
				Addresses: []apiv1.NodeAddress{
					{
						Type:    apiv1.NodeInternalIP,
						Address: "10.0.0.1",
					},
				},
			},
		}}})

	err = GetIngressPod(fkClient)
	if err != nil {
		t.Errorf("expected a PodInfo but returned error: %v", err)
		return
	}
}

func TestSetDefaultNGINXPathType(t *testing.T) {
	prefixType := networkingv1.PathTypePrefix
	exactType := networkingv1.PathTypeExact
	implSpecificType := networkingv1.PathTypeImplementationSpecific

	tests := []struct {
		name          string
		ingress       *networkingv1.Ingress
		expectedTypes []networkingv1.PathType
	}{
		{
			name: "nil PathType should default to Prefix",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/",
											PathType: nil,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedTypes: []networkingv1.PathType{prefixType},
		},
		{
			name: "ImplementationSpecific PathType should be changed to Prefix",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/",
											PathType: &implSpecificType,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedTypes: []networkingv1.PathType{prefixType},
		},
		{
			name: "Exact PathType should remain unchanged",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/exact",
											PathType: &exactType,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedTypes: []networkingv1.PathType{exactType},
		},
		{
			name: "Prefix PathType should remain unchanged",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/prefix",
											PathType: &prefixType,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedTypes: []networkingv1.PathType{prefixType},
		},
		{
			name: "multiple paths with mixed PathTypes",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/nil",
											PathType: nil,
										},
										{
											Path:     "/exact",
											PathType: &exactType,
										},
										{
											Path:     "/impl",
											PathType: &implSpecificType,
										},
										{
											Path:     "/prefix",
											PathType: &prefixType,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedTypes: []networkingv1.PathType{prefixType, exactType, prefixType, prefixType},
		},
		{
			name: "multiple rules with mixed HTTP and nil HTTP",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							Host: "host1.example.com",
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/",
											PathType: nil,
										},
									},
								},
							},
						},
						{
							Host:             "host2.example.com",
							IngressRuleValue: networkingv1.IngressRuleValue{},
						},
						{
							Host: "host3.example.com",
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/api",
											PathType: nil,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedTypes: []networkingv1.PathType{prefixType, prefixType},
		},
		{
			name: "empty rules should not panic",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{},
				},
			},
			expectedTypes: nil,
		},
		{
			name: "rule with HTTP but empty paths should not panic",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{},
								},
							},
						},
					},
				},
			},
			expectedTypes: nil,
		},
		{
			name: "nil rules should not panic",
			ingress: &networkingv1.Ingress{
				Spec: networkingv1.IngressSpec{},
			},
			expectedTypes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetDefaultNGINXPathType(tt.ingress)

			var gotTypes []networkingv1.PathType
			for _, rule := range tt.ingress.Spec.Rules {
				if rule.HTTP == nil {
					continue
				}
				for _, path := range rule.HTTP.Paths {
					if path.PathType != nil {
						gotTypes = append(gotTypes, *path.PathType)
					}
				}
			}

			assert.Equal(t, tt.expectedTypes, gotTypes)
		})
	}
}

func TestNetworkingIngressAvailable(t *testing.T) {
	tests := []struct {
		name          string
		serverVersion *version.Info
		expected      bool
	}{
		{
			name: "kubernetes v1.22.0 supports networking v1",
			serverVersion: &version.Info{
				Major:      "1",
				Minor:      "22",
				GitVersion: "v1.22.0",
			},
			expected: true,
		},
		{
			name: "kubernetes v1.19.0 supports networking v1",
			serverVersion: &version.Info{
				Major:      "1",
				Minor:      "19",
				GitVersion: "v1.19.0",
			},
			expected: true,
		},
		{
			name: "kubernetes v1.18.0 does not support networking v1",
			serverVersion: &version.Info{
				Major:      "1",
				Minor:      "18",
				GitVersion: "v1.18.0",
			},
			expected: false,
		},
		{
			name: "kubernetes v1.21.3 supports networking v1",
			serverVersion: &version.Info{
				Major:      "1",
				Minor:      "21",
				GitVersion: "v1.21.3",
			},
			expected: true,
		},
		{
			name: "kubernetes v1.30.0 supports networking v1",
			serverVersion: &version.Info{
				Major:      "1",
				Minor:      "30",
				GitVersion: "v1.30.0",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := testclient.NewSimpleClientset()
			client.Discovery().(*fakediscovery.FakeDiscovery).FakedServerVersion = tt.serverVersion

			result := NetworkingIngressAvailable(client)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMetaNamespaceKey(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		expected string
	}{
		{
			name: "namespaced object returns namespace/name",
			obj: &apiv1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "test-pod",
				},
			},
			expected: "default/test-pod",
		},
		{
			name: "non-namespaced object returns just name",
			obj: &apiv1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-node",
				},
			},
			expected: "test-node",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := MetaNamespaceKey(tt.obj)
			assert.Equal(t, tt.expected, key)
		})
	}
}
