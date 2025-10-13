package leetcode

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
    start := 0
    maxLength := 0

    for i := range len(s) {
        char := s[i]
        if lastIndex, seen := lastSeen[char]; seen && lastIndex >= start {
            start = lastIndex + 1
        }

        lastSeen[char] = i

        l := i - start + 1
        if l > maxLength {
            maxLength = l
        }
    }
    return maxLength
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