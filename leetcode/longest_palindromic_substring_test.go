package leetcode

import (
	"slices"
	"testing"
)

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	longestSubstr := ""
	for i := range len(s) {
		substr := expand(s, i, i)
		if len(substr) > len(longestSubstr) {
			longestSubstr = substr
		}
		
		substr = expand(s, i, i+1)
		if len(substr) > len(longestSubstr) {
			longestSubstr = substr
		}
	}
	
	return longestSubstr
}

func expand(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	
	return s[left+1: right]
}

func TestLongestPalindromicSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "first test case",
			input:    "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "second test case",
			input:    "cbbd",
			expected: []string{"bb"},
		},
		{
			name:     "third test case",
			input:    "racecar",
			expected: []string{"racecar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := longestPalindrome(tc.input)
			if !slices.Contains(tc.expected, got) {
				t.Errorf("expected a value from %#v; got %s", tc.expected, got)
			}
		})
	}
}
