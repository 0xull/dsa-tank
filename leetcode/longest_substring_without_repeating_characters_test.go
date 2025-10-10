package leetcode

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
    substr := []string{}
    indices := map[string]int{}

    for i, r := range s {
        val := string(r)
        if _, ok := indices[val]; ok {
            index := indices[val]
            fmt.Println(index)
            substr1, substr2 := substr[:index], substr[index:]
            if len(substr1) > len(substr2) {
                substr = substr1
            } else {
                substr = substr2
            }

            delete(indices, val)
            continue
        }

        indices[val] = i
        substr = append(substr, val)
    }

    return len(substr)
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