package main

import "fmt"

func twoSum(nums []int, target int) []int {
  l := len(nums)
  for i, a := range nums {
    for j := i + 1; j < l; j++ {
      if a + nums[j] == target {
        return []int{i, j}
      }
    }
  }
  return nil
}

func main() {
  nums := []int{1,2,3,4}
  result := twoSum(nums, 7)

  fmt.Println("hello world:", result[0], result[1])
}
