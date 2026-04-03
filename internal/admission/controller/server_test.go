/*
Copyright 2019 The Kubernetes Authors.

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

package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	admissionv1 "k8s.io/api/admission/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/apimachinery/pkg/types"
	jsonutil "k8s.io/apimachinery/pkg/util/json"
)

type mockAdmissionController struct {
	response runtime.Object
	err      error
}

func (m *mockAdmissionController) HandleAdmission(_ runtime.Object) (runtime.Object, error) {
	return m.response, m.err
}

func buildAdmissionReviewJSON(t *testing.T, uid string, ingress *networking.Ingress) []byte {
	uidTyped := types.UID(uid)
	t.Helper()

	var raw []byte
	if ingress != nil {
		var err error
		raw, err = jsonutil.Marshal(ingress)
		assert.NoError(t, err)
	}

	review := admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Request: &admissionv1.AdmissionRequest{
			UID:       uidTyped,
			Kind:      metav1.GroupVersionKind{Group: networking.GroupName, Version: "v1", Kind: "Ingress"},
			Name:      "test-ingress",
			Namespace: "default",
			Object: runtime.RawExtension{
				Raw: raw,
			},
		},
	}

	codec := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme, scheme, json.SerializerOptions{
		Pretty: true,
	})
	var buf bytes.Buffer
	err := codec.Encode(&review, &buf)
	assert.NoError(t, err)
	return buf.Bytes()
}

func decodeAdmissionReview(t *testing.T, data []byte) *admissionv1.AdmissionReview {
	t.Helper()

	codec := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme, scheme, json.SerializerOptions{
		Pretty: true,
	})
	obj, _, err := codec.Decode(data, nil, nil)
	assert.NoError(t, err)
	review, ok := obj.(*admissionv1.AdmissionReview)
	assert.True(t, ok, "expected AdmissionReview, got %T", obj)
	return review
}

func TestServeHTTP(t *testing.T) {
	ing := &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testIngressName,
			Namespace: "default",
		},
	}

	tests := []struct {
		name               string
		body               []byte
		controllerResponse runtime.Object
		controllerErr      error
		wantStatus         int
		wantAllowed        bool
		wantUID            string
		wantContentType    bool
	}{
		{
			name:               "valid request - checker passes",
			body:               buildAdmissionReviewJSON(t, "test-uid-1", ing),
			controllerResponse: buildAllowedReview("test-uid-1"),
			controllerErr:      nil,
			wantStatus:         http.StatusOK,
			wantAllowed:        true,
			wantUID:            "test-uid-1",
			wantContentType:    true,
		},
		{
			name:               "valid request - checker fails",
			body:               buildAdmissionReviewJSON(t, "test-uid-2", ing),
			controllerResponse: buildDeniedReview("test-uid-2"),
			controllerErr:      fmt.Errorf("admission failed"),
			wantStatus:         http.StatusInternalServerError,
			wantAllowed:        false,
			wantContentType:    false,
		},
		{
			name:               "malformed JSON body",
			body:               []byte("{not valid json}"),
			controllerResponse: nil,
			controllerErr:      nil,
			wantStatus:         http.StatusBadRequest,
			wantContentType:    false,
		},
		{
			name:               "empty body",
			body:               []byte{},
			controllerResponse: nil,
			controllerErr:      nil,
			wantStatus:         http.StatusBadRequest,
			wantContentType:    false,
		},
		{
			name:               "non-JSON content",
			body:               []byte("this is plain text, not json"),
			controllerResponse: nil,
			controllerErr:      nil,
			wantStatus:         http.StatusBadRequest,
			wantContentType:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := NewAdmissionControllerServer(&mockAdmissionController{
				response: tt.controllerResponse,
				err:      tt.controllerErr,
			})

			req := httptest.NewRequest(http.MethodPost, "/admission", bytes.NewReader(tt.body))
			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantContentType {
				review := decodeAdmissionReview(t, rec.Body.Bytes())
				assert.Equal(t, tt.wantUID, string(review.Response.UID))
				assert.Equal(t, tt.wantAllowed, review.Response.Allowed)
			}

			if tt.wantStatus == http.StatusOK && tt.wantContentType {
				body := rec.Body.Bytes()
				assert.True(t, len(body) > 0, "expected non-empty response body")
				review := decodeAdmissionReview(t, body)
				assert.NotNil(t, review.Response)
			}
		})
	}
}

func TestServeHTTPBodyTooLarge(t *testing.T) {
	server := NewAdmissionControllerServer(&mockAdmissionController{})

	bigBody := make([]byte, maxBodySizeBytes)
	for i := range bigBody {
		bigBody[i] = 'a'
	}

	req := httptest.NewRequest(http.MethodPost, "/admission", bytes.NewReader(bigBody))
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusRequestEntityTooLarge, rec.Code)
}

func TestServeHTTPValidReviewAllowedTrue(t *testing.T) {
	ing := &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testIngressName,
			Namespace: "default",
		},
	}

	body := buildAdmissionReviewJSON(t, "uid-allow-true", ing)
	expectedReview := buildAllowedReview("uid-allow-true")

	server := NewAdmissionControllerServer(&mockAdmissionController{
		response: expectedReview,
		err:      nil,
	})

	req := httptest.NewRequest(http.MethodPost, "/admission", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	review := decodeAdmissionReview(t, rec.Body.Bytes())
	assert.Equal(t, "uid-allow-true", string(review.Response.UID))
	assert.True(t, review.Response.Allowed)
}

func TestServeHTTPValidReviewAllowedFalse(t *testing.T) {
	ing := &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testIngressName,
			Namespace: "default",
		},
	}

	body := buildAdmissionReviewJSON(t, "uid-allow-false", ing)
	expectedReview := buildDeniedReview("uid-allow-false")

	server := NewAdmissionControllerServer(&mockAdmissionController{
		response: expectedReview,
		err:      nil,
	})

	req := httptest.NewRequest(http.MethodPost, "/admission", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	review := decodeAdmissionReview(t, rec.Body.Bytes())
	assert.Equal(t, "uid-allow-false", string(review.Response.UID))
	assert.False(t, review.Response.Allowed)
}

func buildAllowedReview(uid string) *admissionv1.AdmissionReview {
	return &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Response: &admissionv1.AdmissionResponse{
			UID:     types.UID(uid),
			Allowed: true,
		},
	}
}

func buildDeniedReview(uid string) *admissionv1.AdmissionReview {
	return &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Response: &admissionv1.AdmissionResponse{
			UID:     types.UID(uid),
			Allowed: false,
			Result: &metav1.Status{
				Status:  metav1.StatusFailure,
				Code:    http.StatusBadRequest,
				Reason:  metav1.StatusReasonBadRequest,
				Message: "admission denied",
			},
		},
	}
}
