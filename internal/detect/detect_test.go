package detect

import (
	"strings"
	"testing"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{"test.go", "package main", "go"},
		{"test.py", "import os", "python"},
		{"script.sh", "#!/bin/bash", "bash"},
		{"unknown", "some random text", "unknown"},
		{"shebang.py", "#!/usr/bin/env python3\nprint('hi')", "python"},
	}

	for _, tt := range tests {
		r := strings.NewReader(tt.content)
		got := Detect(tt.name, r)
		if got != tt.expected {
			t.Errorf("Detect(%q, %q) = %q, want %q", tt.name, tt.content, got, tt.expected)
		}
	}
}
