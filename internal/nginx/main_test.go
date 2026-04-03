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

package nginx

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetServerBlock(t *testing.T) {
	sampleConf := `# server block for example.com
## start server example.com
server {
    server_name example.com;
    listen 80;
    location / {
        proxy_pass http://backend;
    }
}
## end server example.com

## start server other.com
server {
    server_name other.com;
    listen 80;
}
## end server other.com
`

	tests := []struct {
		name        string
		conf        string
		host        string
		expectBlock string
		expectError bool
	}{
		{
			name:        "find existing server block",
			conf:        sampleConf,
			host:        "example.com",
			expectBlock: "server {\n    server_name example.com;\n    listen 80;\n    location / {\n        proxy_pass http://backend;\n    }\n}\n",
			expectError: false,
		},
		{
			name:        "find second server block",
			conf:        sampleConf,
			host:        "other.com",
			expectBlock: "server {\n    server_name other.com;\n    listen 80;\n}\n",
			expectError: false,
		},
		{
			name:        "host not found",
			conf:        sampleConf,
			host:        "missing.com",
			expectBlock: "",
			expectError: true,
		},
		{
			name:        "empty config",
			conf:        "",
			host:        "example.com",
			expectBlock: "",
			expectError: true,
		},
		{
			name:        "start marker found but end marker missing",
			conf:        "## start server broken.com\nserver {}\n",
			host:        "broken.com",
			expectBlock: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block, err := GetServerBlock(tt.conf, tt.host)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectBlock, block)
			}
		})
	}
}

func TestReadFileToString(t *testing.T) {
	dir := t.TempDir()

	tests := []struct {
		name        string
		content     string
		createFile  bool
		expectError bool
	}{
		{
			name:       "read existing file",
			content:    "test content\nwith multiple lines",
			createFile: true,
		},
		{
			name:        "read non-existent file",
			content:     "",
			createFile:  false,
			expectError: true,
		},
		{
			name:       "read empty file",
			content:    "",
			createFile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.createFile {
				filePath = filepath.Join(dir, tt.name+".txt")
				err := os.WriteFile(filePath, []byte(tt.content), 0o644)
				require.NoError(t, err)
			} else {
				filePath = filepath.Join(dir, "nonexistent.txt")
			}

			content, err := readFileToString(filePath)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.content, content)
			}
		})
	}
}

func TestVersion(t *testing.T) {
	result := Version()
	assert.NotEmpty(t, result)
}

func TestIsRunning(t *testing.T) {
	result := IsRunning()
	assert.IsType(t, true, result)
}
