/*
Copyright 2020 The Kubernetes Authors.

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
	"testing"

	"github.com/stretchr/testify/assert"
	networking "k8s.io/api/networking/v1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/rewrite"
	"k8s.io/ingress-nginx/pkg/apis/ingress"
)

func TestNormalizePrefixPath(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		expected string
	}{
		{"root path", "/", "/"},
		{"path without trailing slash", "/foo", "/foo/"},
		{"path with trailing slash", "/foo/", "/foo/"},
		{"multi-segment without trailing slash", "/foo/bar", "/foo/bar/"},
		{"multi-segment with trailing slash", "/foo/bar/", "/foo/bar/"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizePrefixPath(tc.path)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestNeedsRewrite(t *testing.T) {
	testCases := []struct {
		name     string
		location *ingress.Location
		expected bool
	}{
		{
			"no rewrite target",
			&ingress.Location{
				Path:    "/foo",
				Rewrite: makeRewriteConfig("", false),
			},
			false,
		},
		{
			"rewrite target same as path",
			&ingress.Location{
				Path:    "/foo",
				Rewrite: makeRewriteConfig("/foo", false),
			},
			false,
		},
		{
			"rewrite target different from path",
			&ingress.Location{
				Path:    "/foo",
				Rewrite: makeRewriteConfig("/bar", false),
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := needsRewrite(tc.location)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestUpdateServerLocations(t *testing.T) {
	t.Run("root location is not modified", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/",
				PathType: &pathTypePrefix,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 1)
		assert.Equal(t, "/", result[0].Path)
	})

	t.Run("exact location is not modified", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/foo",
				PathType: &pathTypeExact,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 1)
		assert.Equal(t, "/foo", result[0].Path)
		assert.Equal(t, pathTypeExact, *result[0].PathType)
	})

	t.Run("prefix location without rewrite gets exact location added", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/foo",
				PathType: &pathTypePrefix,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 2)

		assert.Equal(t, "/foo/", result[0].Path)
		assert.Equal(t, pathTypePrefix, *result[0].PathType)
		assert.Equal(t, "/foo", result[0].IngressPath)

		assert.Equal(t, "/foo", result[1].Path)
		assert.Equal(t, pathTypeExact, *result[1].PathType)
		assert.Equal(t, "/foo", result[1].IngressPath)
	})

	t.Run("prefix location with rewrite does not get exact location", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/foo",
				PathType: &pathTypePrefix,
				Rewrite:  makeRewriteConfig("/bar", false),
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 1)
		assert.Equal(t, "/foo", result[0].Path)
		assert.Equal(t, "/foo", result[0].IngressPath)
	})

	t.Run("prefix location with UseRegex does not get exact location", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/foo",
				PathType: &pathTypePrefix,
				Rewrite:  makeRewriteConfig("", true),
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 1)
		assert.Equal(t, "/foo", result[0].Path)
	})

	t.Run("prefix location with existing exact location normalizes path only", func(t *testing.T) {
		exactType := networking.PathTypeExact
		locations := []*ingress.Location{
			{
				Path:     "/foo",
				PathType: &pathTypeExact,
			},
			{
				Path:     "/foo",
				PathType: &pathTypePrefix,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 2)

		assert.Equal(t, "/foo", result[0].Path)
		assert.Equal(t, exactType, *result[0].PathType)

		assert.Equal(t, "/foo/", result[1].Path)
		assert.Equal(t, pathTypePrefix, *result[1].PathType)
	})

	t.Run("prefix location with trailing slash gets exact location added", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/foo/",
				PathType: &pathTypePrefix,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 2)

		assert.Equal(t, "/foo/", result[0].Path)
		assert.Equal(t, "/foo/", result[0].IngressPath)

		assert.Equal(t, "/foo/", result[1].Path)
		assert.Equal(t, pathTypeExact, *result[1].PathType)
	})

	t.Run("implementation specific path type does not get exact location", func(t *testing.T) {
		implSpecific := networking.PathType("ImplementationSpecific")
		locations := []*ingress.Location{
			{
				Path:     "/foo",
				PathType: &implSpecific,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 1)
		assert.Equal(t, "/foo", result[0].Path)
		assert.Equal(t, "/foo", result[0].IngressPath)
	})

	t.Run("multiple locations mixed types", func(t *testing.T) {
		locations := []*ingress.Location{
			{
				Path:     "/",
				PathType: &pathTypePrefix,
			},
			{
				Path:     "/api",
				PathType: &pathTypePrefix,
			},
			{
				Path:     "/static",
				PathType: &pathTypeExact,
			},
		}

		result := updateServerLocations(locations)
		assert.Len(t, result, 4)

		assert.Equal(t, "/", result[0].Path)

		assert.Equal(t, "/api/", result[1].Path)
		assert.Equal(t, pathTypePrefix, *result[1].PathType)

		assert.Equal(t, "/api", result[2].Path)
		assert.Equal(t, pathTypeExact, *result[2].PathType)

		assert.Equal(t, "/static", result[3].Path)
		assert.Equal(t, pathTypeExact, *result[3].PathType)
	})

	t.Run("empty locations returns empty slice", func(t *testing.T) {
		result := updateServerLocations([]*ingress.Location{})
		assert.Empty(t, result)
	})
}

func makeRewriteConfig(target string, useRegex bool) rewrite.Config {
	return rewrite.Config{
		Target:   target,
		UseRegex: useRegex,
	}
}
