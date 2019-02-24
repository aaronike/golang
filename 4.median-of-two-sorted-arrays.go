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

const (
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
)

func binarySearchRange(nums []int, num int, start int, end int) int {
	if len(nums) == 0 {
		return 0
	}
	for start < end {
		m := int((start + end) / 2)
		mid := nums[m]
		if (num == mid) {
			return m
		} else if (num < mid) {
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

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if x, y, found := findMedianSortedArraysHelper(nums1, nums2); found {
		return float64(x + y) / 2
	}
	if x, y, found := findMedianSortedArraysHelper(nums2, nums1); found {
		return float64(x + y) / 2
	}
	return 0
}

func findMedianSortedArraysHelper(nums1 []int, nums2 []int) (int, int, bool) {
	l1, l2 := len(nums1), len(nums2)
	start1, end1 := 0, l1 - 1
	start2, end2 := 0, l2 - 1

	for start1 <= end1 {
		i := (start1 + end1) / 2
		num := nums1[i]
		p := binarySearchRange(nums2, num, start2, end2)
		
		left := i + p
		right := l1 - i - 1 + l2 - p
		fmt.Printf("nums1(%d,%d)[%d]=%d in nums2[%d] left=%d,right=%d\n", start1, end1, i, num, p, left, right)
		if left == right {
			return num, num, true
		} else if left + 1 == right {
			var next = MaxInt64
			if i + 1 < l1 {
				next  = min(next, nums1[i + 1])
			}
			if 0 <= p && p < l2 {
				next  = min(next, nums2[p])
			}
			return num, next, true
		} else if left == right + 1 {
			var prev = MinInt64
			if i - 1 >= 0 {
				prev  = max(prev, nums1[i - 1])
			}
			if 0 <= p - 1 && p - 1 < l2 {
				prev  = max(prev, nums2[p - 1])
			}
			return prev, num, true
		} else if left < right {
			start1 = i + 1
			end2 = p
		} else {
			end1 = i - 1
			start2 = p
		}
	}
	return 0, 0, false
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 { // small first
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}

	start, end := 0, l1
	midLen := (l1 + l2 + 1) / 2

	var num, next int
	for start <= end {
		i := (start + end) / 2
		j := midLen - i
		
		fmt.Printf("(%d,%d) i=%d j=%d\n", start, end, i, j)
		if i < l1 && nums2[j - 1] > nums1[i] {
			start = i + 1
		} else if i > 0 && nums1[i - 1] > nums2[j] {
			end = i - 1
		} else {
			if i == 0 {
				num = nums2[j - 1]
			} else if j == 0 {
				num = nums1[i - 1]
			} else {
				num = max(nums1[i - 1], nums2[j - 1])
			}

			if (l1 + l2) % 2 == 1 {
				next = num
			} else {
				if i == l1 {
					next = nums2[j]
				} else if j == l2 {
					next = nums1[i]
				} else {
					next = min(nums1[i], nums2[j])
				}
			}
			return float64(num + next) / 2
		}
	}
	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Println(nums1, nums2)
	return findMedianSortedArrays2(nums1, nums2)
}

func main() {
	var nums1, nums2 []int

	nums1 = []int{1}
	nums2 = []int{2,3}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	return

	nums1 = []int{1,2}
	nums2 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2}
	nums2 = []int{1,2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{10000}
	nums2 = []int{10001}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{10,20,50,80,90}
	nums2 = []int{1,4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,5,8,9}
	nums2 = []int{7,14}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,3,4,5,6}
	nums2 = []int{7,8,9,10}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,5,8,9}
	nums2 = []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2}
	nums2 = []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,5,8,9}
	nums2 = []int{2,3}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,5,8,9}
	nums2 = []int{6,7}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,5,8,9}
	nums2 = []int{10,14}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}