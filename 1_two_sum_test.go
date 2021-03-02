package main

import "testing"

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
*/

func TwoSum(input []int, target int) [2]int {
	// where key is residual, value is index
	m := make(map[int]int, len(input))
	for i, num := range input {
		if index, ok := m[num]; ok {
			return [2]int{index, i}
		}
		m[target-num] = i
	}
	return [2]int{}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output [2]int
	}{
		{input: []int{2, 7, 11, 15}, target: 9, output: [2]int{0, 1}},
		{input: []int{3, 2, 4}, target: 6, output: [2]int{1, 2}},
		{input: []int{3, 3}, target: 6, output: [2]int{0, 1}},
		{input: []int{3, 4}, target: 6, output: [2]int{}},
	}

	for _, test := range tests {
		actual := TwoSum(test.input, test.target)
		if actual != test.output {
			t.Errorf("want %v got %v", test.output, actual)
		}
	}
}
