package main

import (
	"testing"
)

func TestCheckPatterns(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantOutput bool
	}{
		{
			name:       "too short",
			input:      "131",
			wantOutput: false,
		},
		{
			name:       "too long",
			input:      "1334234234234231",
			wantOutput: false,
		},
		{
			name:       "invalid1",
			input:      "1334234231",
			wantOutput: false,
		},
		{
			name:       "invalid2",
			input:      "abscjoi234",
			wantOutput: false,
		},
		{
			name:       "valid1",
			input:      "abcdA1#1",
			wantOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, _ := checkPatterns(tt.input)
			if gotOutput != tt.wantOutput {
				t.Errorf("codeToString() = (%v), want (%v)",
					gotOutput, tt.wantOutput)
			}
		})
	}
}
