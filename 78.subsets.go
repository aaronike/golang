package main

import "fmt"

func subsets(nums []int) [][]int {

  var l uint = uint(len(nums))
  var k uint;
  var result = [][]int{}
  var bitsect = []uint{}
  var count uint = 1 << l
  for k = 0; k < l; k++{
    bitsect = append(bitsect, 1 << k)
  }
  for ; count > 0; count-- {
    list := []int{}
    for i, num := range(bitsect) {
      if num & count != 0 {
        list = append(list, nums[i])
      }
    }
    result = append(result, list)
  }
  return result
}


func main() {
  nums := []int{1,2,3}
  fmt.Println(subsets(nums))
}
