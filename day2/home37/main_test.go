package main

import (
	"testing"
)

func TestFindLongestWord(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantWord  string
		wantRunes int
	}{
		{
			name:      "basic test",
			input:     "1 55555 22 333 4444 666666 7777777 88888888",
			wantWord:  "55555",
			wantRunes: 5,
		},
		{
			name:      "unicode symbols",
			input:     "привет мир радость",
			wantWord:  "радость",
			wantRunes: 7,
		},
		{
			name:      "empty string",
			input:     "",
			wantWord:  "",
			wantRunes: 0,
		},
		{
			name:      "single word",
			input:     "hello",
			wantWord:  "hello",
			wantRunes: 5,
		},
		{
			name:      "with space limit",
			input:     "a bb ccc dddd eeeee ffffff",
			wantWord:  "eeeee",
			wantRunes: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWord, gotRunes := findLongestWord(tt.input)
			if gotWord != tt.wantWord || gotRunes != tt.wantRunes {
				t.Errorf("findLongestWord() = (%v, %v), want (%v, %v)",
					gotWord, gotRunes, tt.wantWord, tt.wantRunes)
			}
		})
	}
}
