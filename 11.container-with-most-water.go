/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (42.41%)
 * Total Accepted:    315.7K
 * Total Submissions: 744.5K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 * 
 * Note: You may not slant the container and n is at least 2.
 * 
 * 
 * 
 * 
 * 
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49. 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * 
 */
package main

import "fmt"

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

func maxAreaForce(height []int) int {
	l := len(height)
	var result int
	for i, x := range height {

		for j := i + 1; j < l; j++ {
			area := (j - i) * min(x, height[j])
			result = max(result, area)
		}
	}
	return result
}

func maxAreaGreedy(height []int) int {
	l := len(height)
	var result, area int
	for i, j := 0, l - 1; i < j; {
		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}
		if area > result {
			result = area
		}
	}
	return result
}

func maxArea(height []int) int {
	return maxAreaGreedy(height)
}

func main() {
	var a []int
 	a = []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(a))

 	a = []int{1,8}
	fmt.Println(maxArea(a))
}
