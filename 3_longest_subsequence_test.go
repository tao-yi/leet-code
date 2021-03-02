package main

import "testing"

/*
Given a string s, find the length of the longest substring without repeating characters.

Input: s = "abcabcbb"
Output: 3
*/

func lcs(s string) int {
	runes := []rune(s)
	var start, end, max, windowSize int
	m := make(map[rune]int, len(runes))
	for end < len(runes) {
		if pos, ok := m[runes[end]]; ok && start <= pos {
			start = pos + 1
		}
		windowSize = end - start + 1
		if windowSize > max {
			max = windowSize
		}
		m[runes[end]] = end
		end++
	}
	return max
}

func TestLcs(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{input: "abcabcbb", output: 3},
		{input: "bbbbb", output: 1},
		{input: "pwwkew", output: 3},
		{input: "abba", output: 2},
		{input: "", output: 0},
	}

	for _, test := range tests {
		actual := lcs(test.input)
		if actual != test.output {
			t.Errorf("want %d, got %d", test.output, actual)
		}
	}
}
