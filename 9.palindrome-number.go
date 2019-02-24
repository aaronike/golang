/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (41.83%)
 * Total Accepted:    507.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 * 
 * Example 1:
 * 
 * 
 * Input: 121
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 * 
 * 
 * Follow up:
 * 
 * Coud you solve it without converting the integer to a string?
 * 
 */
package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var a, b int
	for a = x; x > 0; x = x / 10{
		n := x % 10
		b = b * 10 + n
	}
	return a == b
}

func main() {
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(0))
}
