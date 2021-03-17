package main

import "testing"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	// 双指针向右滑动
	i := 0
	j := i + 1
	for j < len(nums) {
		// 如果i和j指向的元素相等，j向右滑动
		if nums[i] == nums[j] {
			j++
			continue
		}
		// 如果i和j指向的元素不等
		if j != i+1 {
			// 如果不是相邻的两个元素，互换j和i+1位置的两个元素
			nums[i+1], nums[j] = nums[j], nums[i+1]
		}
		// 向右滑动i和j
		i++
		j++
	}
	return i + 1
}

func sliceEqual(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
		len    int
	}{
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: []int{0, 1, 2, 3, 4, 0, 2, 1, 3, 1},
			len:    5,
		},
	}
	for _, tc := range tests {
		len := removeDuplicates(tc.input)
		if !sliceEqual(tc.input, tc.output) || len != tc.len {
			t.Errorf("got %d %v, want %d %v", len, tc.input, tc.len, tc.output)
		}
	}
}
