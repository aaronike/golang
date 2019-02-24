package main

import "fmt"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.10%)
 * Total Accepted:    606.2K
 * Total Submissions: 2.4M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 123
 * Output: 321
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -123
 * Output: -321
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 120
 * Output: 21
 * 
 * 
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 * 
 */
const (
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
)

func reverse(x int) int {
	var result int
	var sign int = 1
	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		n := x % 10
		x = x / 10
		result = result * 10 + n
		if result * sign > MaxInt32 {
			return 0
		} else if result * sign < MinInt32 {
			return 0
		}
	}
	return result * sign
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}
