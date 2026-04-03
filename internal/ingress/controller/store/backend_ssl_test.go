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

package store

import (
	"testing"

	"github.com/eapache/channels"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
)

func TestIsErrSecretForAuth(t *testing.T) {
	assert.True(t, isErrSecretForAuth(ErrSecretForAuth))
	assert.False(t, isErrSecretForAuth(nil))
	assert.False(t, isErrSecretForAuth(assert.AnError))
}

func TestGetPemCertificateBranches(t *testing.T) {
	tests := []struct {
		name        string
		secret      *v1.Secret
		expectError bool
		errorMsg    string
	}{
		{
			name:        "secret with no tls keys and no ca or auth",
			secret:      &v1.Secret{Data: map[string][]byte{}},
			expectError: true,
			errorMsg:    "contains no keypair or CA certificate",
		},
		{
			name: "secret with only tls.crt missing tls.key",
			secret: &v1.Secret{Data: map[string][]byte{
				"tls.crt": []byte("cert-data"),
			}},
			expectError: true,
			errorMsg:    "contains no keypair or CA certificate",
		},
		{
			name: "secret with only tls.key missing tls.crt",
			secret: &v1.Secret{Data: map[string][]byte{
				"tls.key": []byte("key-data"),
			}},
			expectError: true,
			errorMsg:    "contains no keypair or CA certificate",
		},
		{
			name: "secret with only auth key",
			secret: &v1.Secret{Data: map[string][]byte{
				"auth": []byte("auth-data"),
			}},
			expectError: true,
			errorMsg:    "secret is used for authentication",
		},
		{
			name: "secret with ca.crt only",
			secret: &v1.Secret{Data: map[string][]byte{
				"ca.crt": []byte("ca-data"),
			}},
			expectError: true,
			errorMsg:    "unexpected error creating SSL Cert",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &k8sStore{
				listers: &Lister{
					Secret: SecretLister{
						Store: newFakeSecretStore(map[string]*v1.Secret{
							"default/test-secret": tt.secret,
						}),
					},
				},
			}

			cert, err := store.getPemCertificate("default/test-secret")

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				assert.Nil(t, cert)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, cert)
			}
		})
	}
}

type fakeSecretStore struct {
	secrets map[string]*v1.Secret
}

func newFakeSecretStore(secrets map[string]*v1.Secret) *fakeSecretStore {
	return &fakeSecretStore{secrets: secrets}
}

func (f *fakeSecretStore) GetByKey(key string) (interface{}, bool, error) {
	secret, ok := f.secrets[key]
	if !ok {
		return nil, false, nil
	}
	return secret, true, nil
}

func (f *fakeSecretStore) Add(obj interface{}) error {
	return nil
}

func (f *fakeSecretStore) Update(obj interface{}) error {
	return nil
}

func (f *fakeSecretStore) Delete(obj interface{}) error {
	return nil
}

func (f *fakeSecretStore) List() []interface{} {
	return nil
}

func (f *fakeSecretStore) ListKeys() []string {
	return nil
}

func (f *fakeSecretStore) Get(obj interface{}) (interface{}, bool, error) {
	return nil, false, nil
}

func (f *fakeSecretStore) Replace(list []interface{}, _ string) error {
	return nil
}

func (f *fakeSecretStore) Resync() error {
	return nil
}

func TestSendDummyEvent(t *testing.T) {
	ch := channels.NewRingChannel(10)
	store := &k8sStore{
		updateCh: ch,
	}

	store.sendDummyEvent()

	event := (<-ch.Out()).(Event)
	assert.Equal(t, UpdateEvent, event.Type)

	ing := event.Obj.(*networking.Ingress)
	assert.Equal(t, "dummy", ing.Name)
	assert.Equal(t, "dummy", ing.Namespace)
}
