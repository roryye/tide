package main

import "sort"

func binarySearchRecursive(target int, nums []int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if target == nums[mid] {
		return mid
	}
	if nums[mid] < target {
		return binarySearchRecursive(target, nums, mid+1, right)
	}
	if nums[mid] > target {
		return binarySearchRecursive(target, nums, left, mid-1)
	}
	return -1
}

func findIndex(target int, nums []int) []int {
	index := binarySearchRecursive(target, nums, 0, len(nums)-1)
	if index == -1 {
		return nil
	}

	res := []int{index}
	left, right := index-1, index+1
	for left >= 0 && right < len(nums) && nums[left] == target && nums[right] == target {
		res = append(res, left)
		res = append(res, right)
		left--
		right++
	}
	for left >= 0 && nums[left] == target {
		res = append(res, left)
		left--
	}
	for right < len(nums) && nums[right] == target {
		res = append(res, right)
		right++
	}

	sort.Ints(res)
	return res
}
