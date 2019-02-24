package main

import "fmt"
/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (30.57%)
 * Total Accepted:    285.3K
 * Total Submissions: 933.2K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 * 
 * 
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * 
 * 
 * And then read line by line: "PAHNAPLSIIGYIR"
 * 
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 * 
 * 
 * string convert(string s, int numRows);
 * 
 * Example 1:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 * 
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 * 
 */
func convert2(s string, numRows int) string {
	var rows [][]byte
	slen := len(s)
	for index := 0; index < slen; {
		var row = make([]byte, numRows)
		for i := 0; i < numRows && index < slen; i++ {
			row[i] = s[index]
			index++
		}
		rows = append(rows, row)
		
		for i := numRows - 2; i > 0 && index < slen; i-- {
			var row = make([]byte, numRows)
			row[i] = s[index]
			index++
			rows = append(rows, row)
		}
	}

	result := make([]byte, slen)
	index := 0
	for i := 0; i < numRows; i++ {
		for _, row := range rows {
			if ch := row[i]; ch > 0 {
				result[index] = ch
				index++
			}
		}
	}
	return string(result)
}

func convert(s string, n int) string {
	slen := len(s)
	result := make([]byte, slen)
	index := 0
	for row := 0; row < n; row++ {
		for col, pos := 0, 0; pos < slen && index < slen; col++{
			if row != 0 && row != n - 1 {
				pos = (n + n - 2) * col - row
				if pos >= 0 && pos < slen {
					result[index] = s[pos]
					index++
				}
			}

			if n == 1 {
				pos = col
			} else {
				pos = (n + n - 2) * col + row
			}
			if pos < slen {
				result[index] = s[pos]
				index++
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("AB", 1))
	return
	fmt.Println(convert("A", 1))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
