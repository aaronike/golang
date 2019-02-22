/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (25.42%)
 * Total Accepted:    379.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * 
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 * 
 * You may assume nums1 and nums2 cannot be both empty.
 * 
 * Example 1:
 * 
 * 
 * nums1 = [1, 3]
 * nums2 = [2]
 * 
 * The median is 2.0
 * 
 * 
 * Example 2:
 * 
 * 
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 
 * The median is (2 + 3)/2 = 2.5
 * 
 * 
 */

package main

import "fmt"

func binarySearch(nums []int, num int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		m := int((start + end) / 2)
		mid := nums[m]
		if (num < mid) {
			end = m - 1
		} else {
			start = m + 1
		}
	}
	return start
}

func findMedianSortedArrays_tmp(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)

	start1, start2 := 0, 0
	end1, end2 := l1, l2

	var p1, p2 int
	totalLen := l1 + l2
	if totalLen % 2 == 0 {

	}
	midPos := int((l1 + l2) / 2)
	for {
		m1, m2 := int((start1 + end1) / 2), int((start2 + end2) / 2)
		n1, n2 := nums1[m1], nums2[m2] // 各自中位数
		p1 = binarySearch(nums2, n1)
		p2 = binarySearch(nums1, n2)
		if (p1 + m1 == midPos) {
			return float64(n1)
		}
		if (p2 + m2 == midPos) {
			return float64(n2)
		}
	}
	return 0
}

func binarySearchRange(nums []int, num int, start int, end int) int {
	for start < end {
		m := int((start + end) / 2)
		mid := nums[m]
		if (num < mid) {
			end = m - 1
		} else {
			start = m + 1
		}
	}
	if (num > nums[start]) {
		return start + 1
	} else {
		return start
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	start1, start2 := 0, 0
	end1, end2 := l1 - 1, l2 - 1
	var pos, cpos int

	midPos := int((l1 + l2) / 2)
	fmt.Println(l1, l2, midPos)
	for count := 0; count < 5; count++ {
		if (start1 <= end1) {
			p1 := int((start1 + end1) / 2)
			n1 := nums1[p1]
			pos = binarySearchRange(nums2, n1, start2, end2)

			cpos = pos + p1
			fmt.Printf("nums1[%d,%d]=%d in nums2(%d, %d)=%d\n", start1, end1, n1, start2, end2, pos)
			if (cpos == midPos) {
				return float64(n1)
			} else if (cpos < midPos) { // 元素偏小
				start1 = p1 + 1
				start2 = pos + 1
			} else {
				end1 = p1 - 1
				end2 = pos - 1
			}

		}
		
		if (start2 <= end2) {
			p2 := int((start2 + end2) / 2)
			n2 := nums2[p2]
			pos = binarySearchRange(nums1, n2, start1, end1)

			fmt.Printf("nums2[%d]=%d in nums1(%d, %d)=%d\n", p2, n2, start1, end1, pos)
			cpos = pos + p2
			if (cpos == midPos) {
				return float64(n2)
			} else if (cpos < midPos) { // 元素偏小
				start1 = pos + 1
				start2 = p2 + 1
			} else {
				end1 = pos - 1
				end2 = p2 - 1
			}
		}
	}
	return 0
}

func main() {
	nums1 := []int{1,2,5,8,9}
	nums2 := []int{3,4}
	// fmt.Println(binarySearch(nums1, 9))
	// fmt.Println(binarySearch(nums2, 2))
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}