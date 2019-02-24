/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (24.94%)
 * Total Accepted:    274.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 * 
 * 
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * 
 * 
 * The matching should cover the entire input string (not partial).
 * 
 * Note:
 * 
 * 
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the precedeng element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore
 * it matches "aab".
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 * 
 * 
 */

package main

import "fmt"

type Node struct  {
	Char byte
	Glob bool
	Next *Node
}

func compilePattern(p string) *Node {
	var prev, node *Node

	root := new(Node)
	prev = root

	plen := len(p)
	for i := 0; i < plen; i++ {
		ch := p[i]
		switch  ch {
		case '*':
			node.Glob = true
		default:
			node = new(Node)
			node.Char = ch
			prev.Next = node
			prev = node
		}
	}
	return root.Next
}

func matchNode(s string, index int, node *Node) bool {
	slen := len(s)

	// fmt.Println(index, node, slen)
	if node == nil {
		return index == slen
	}

	if index == slen && !node.Glob{
		return false
	}

	if node.Glob {
		if result := matchNode(s, index, node.Next); result {
			return true
		}
	} else if index == slen {
		return false
	}

	for i := index; i < slen; i++ {
		if node.Char != '.' && node.Char != s[i] {
			return false
		}
		if result := matchNode(s, i + 1, node.Next); result {
			return true
		}
		if !node.Glob {
			break
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	fmt.Println(s, p)
	root := compilePattern(p)
	return matchNode(s, 0, root)
}

func main() {
	fmt.Println(isMatch("", "a*"))
	fmt.Println(isMatch("a", "ab*"))
	fmt.Println(isMatch("sipi", "s*p*."))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", "a."))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*"))
}
