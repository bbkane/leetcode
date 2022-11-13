package ransomnote

// https://leetcode.com/problems/ransom-note/description/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func canConstruct(ransomNote string, magazine string) bool {
	// because magazine is *guaranteed* to contain only a-z,
	// I can go through each byte and use it's offset from 'a'
	// I don't need a range loop because that produces runes when I want bytes
	// https://go.dev/blog/strings

	magBytesCounts := [26]int{}
	for i := 0; i < len(magazine); i++ {
		magBytesCounts[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		magBytesCounts[ransomNote[i]-'a']--
	}

	for _, el := range magBytesCounts {
		if el < 0 {
			return false
		}
	}
	return true
}

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote  string
		magazine    string
		expectedRes bool
	}{
		{
			ransomNote:  "a",
			magazine:    "b",
			expectedRes: false,
		},
		{
			ransomNote:  "aa",
			magazine:    "ab",
			expectedRes: false,
		},
		{
			ransomNote:  "aa",
			magazine:    "aab",
			expectedRes: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.ransomNote+","+tt.magazine, func(t *testing.T) {
			actualRes := canConstruct(tt.ransomNote, tt.magazine)
			require.Equal(t, tt.expectedRes, actualRes)
		})
	}
}
