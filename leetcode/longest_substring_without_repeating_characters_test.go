package leetcode

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	dups := map[string]struct{}{}
	substrIndex := 0
	substrLen := 0
	
	for i, r := range s {
		rs := string(r)
		if _, ok := dups[rs]; ok {
			l := len(s[substrIndex:i])
			if l > substrLen {
				substrLen = l
			}
			
			substrIndex = i
			dups = make(map[string]struct{})
			dups[rs] = struct{}{}
			continue
		}
		
		dups[rs] = struct{}{}
	}
	
	return substrLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expected int
	}{
		{
			name: "first test case",
			input: "abcabcbb",
			expected: 3,
		},
		{
			name: "second test case",
			input: "bbbbb",
			expected: 1,
		},
		{
			name: "third test case",
			input: "pwwkew",
			expected: 3,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.input)
			if got != tc.expected {
				t.Errorf("expected: %d; got: %d", tc.expected, got)
			}
		})
	}
}