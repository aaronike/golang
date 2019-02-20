package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
  var max int
  l := len(s)
  /*
  records := make([]int, l)
  dups := make([]int, l)
  */
  for index, _ := range s {
    m := make(map[int]int)
    for start := index; start < l; start++ {
      b := int(s[start])
      if m[b] != 0 {
        break
      }
      m[b] = 1
      count := start - index + 1
      if count > max {
        max = count
      }
    }
  }
  return max
}

func main() {
  a := "abcabcabc"
  fmt.Println(a, lengthOfLongestSubstring(a))
  b := "bbbbb"
  fmt.Println(b, lengthOfLongestSubstring(b))
  c := "pwwkew"
  fmt.Println(c, lengthOfLongestSubstring(c))
  d := "abcdefga"
  fmt.Println(d, lengthOfLongestSubstring(d))
}
