package romantointeger

// https://leetcode.com/problems/roman-to-integer/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func romanToInt(s string) int {
	count := 0

	sLen := len(s)

	notAtEnd := func(i int) bool { return i+1 != sLen }

	for i := 0; i < sLen; {
		switch s[i] {
		case 'I':

			if notAtEnd(i) && s[i+1] == 'V' {
				count += 4
				i = i + 2
				continue
			} else if notAtEnd(i) && s[i+1] == 'X' {
				count += 9
				i = i + 2
				continue
			} else {
				count += 1
			}
		case 'V':
			count += 5
		case 'X':
			if notAtEnd(i) && s[i+1] == 'L' {
				count += 40
				i = i + 2
				continue
			} else if notAtEnd(i) && s[i+1] == 'C' {
				count += 90
				i = i + 2
				continue
			} else {
				count += 10
			}
		case 'L':
			count += 50
		case 'C':
			if notAtEnd(i) && s[i+1] == 'D' {
				count += 400
				i = i + 2
				continue
			} else if notAtEnd(i) && s[i+1] == 'M' {
				count += 900
				i = i + 2
				continue
			} else {
				count += 100
			}
		case 'D':
			count += 500
		case 'M':
			count += 1000
		default:
			panic("Unexpected char" + string(s[i]))
		}
		i++
	}
	return count
}

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		s           string
		expectedInt int
	}{
		{s: "I", expectedInt: 1},
		{s: "XI", expectedInt: 11},
		{s: "IV", expectedInt: 4},
		{s: "IXI", expectedInt: 10},
		{s: "III", expectedInt: 3},
		{s: "LVIII", expectedInt: 58},
		{s: "MCMXCIV", expectedInt: 1994},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			actualInt := romanToInt(tt.s)
			require.Equal(t, tt.expectedInt, actualInt)
		})
	}
}
