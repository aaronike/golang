package main

import "fmt"

func longestPalindrome_basic(s string) string {
  var max, max_i, max_j int
  var count int
  l := len(s)
  for k, _ := range s {
    for n := 0; k - n >= 0 && n + k < l; n++ {
      i := k - n
      j := k + n
      if s[i] != s[j] {
        break
      }
      count = n * 2 + 1
      if count > max {
        max, max_i, max_j = count, i, j
      }
    }
    for n := 1; k - n + 1 >= 0 && k + n < l; n++ {
      i := k - n + 1
      j := k + n
      if s[i] != s[j] {
        break
      }
      count = n * 2
      if count > max {
        max, max_i, max_j = count, i, j
        max = count
      }
    }

  }
  return s[max_i:max_j+1]
}

func longestPalindrome(s string) string {
  var max, max_i, max_j int = 1, 0, 0
  var l = len(s)
  var f = [1000][1000]bool{}

  if l <= 1 {
    return s
  }

  for j := 0; j < l; j++ {
    f[j][j] = true
    for i := 0; i <= j; i++{
      x := i + 1
      y := j - 1
      if s[i] != s[j] {
        continue
      }
      if x > y { // 相邻
        f[i][j] = true
      } else if f[x][y] { // 子串相同
        f[i][j] = true
      }
      //fmt.Println(i, j, f[i][j])
      if !f[i][j] {
        continue
      }
      count := j - i + 1
      if count > max {
        max, max_i, max_j = count, i, j
        max = count
      }
    }
  }

  return s[max_i:max_j+1]
}

func main() {
  a := "xababbay"
  fmt.Println(a, longestPalindrome(a))
  b := "bbbbb"
  fmt.Println(b, longestPalindrome(b))
  c := "pwwkew"
  fmt.Println(c, longestPalindrome(c))
  d := "abababababa"
  fmt.Println(d, longestPalindrome(d))
}
