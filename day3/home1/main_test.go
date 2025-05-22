package main

import (
	"testing"
)

func TestCodeToString(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOutput string
	}{
		{
			name:       "sample",
			input:      "220411112603141304",
			wantOutput: "well done",
		},
		{
			name:       "my example",
			input:      "0704111114262214171103",
			wantOutput: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := codeToString(tt.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("codeToString() = (%v), want (%v)",
					gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestStringToCode(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOutput string
	}{
		{
			name:       "sample",
			input:      "well done",
			wantOutput: "220411112603141304",
		},
		{
			name:       "my example",
			input:      "hello world",
			wantOutput: "0704111114262214171103",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := stringToCode(tt.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("codeToString() = (%v), want (%v)",
					gotOutput, tt.wantOutput)
			}
		})
	}
}
