package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var result, last *ListNode
  var a, b, c int
  for l1 != nil || l2 != nil || c != 0{
    if l1 ==  nil {
      a = 0
    } else {
      a = l1.Val
      l1 = l1.Next
    }
    
    if l2 ==  nil {
      b = 0
    } else {
      b = l2.Val
      l2 = l2.Next
    }

    c = a + b + c
    node := new(ListNode)

    if result == nil {
      result = node
    }

    if c >= 10 {
      node.Val = c - 10
      c = 1
    } else {
      node.Val = c
      c = 0
    }

    if last != nil {
      last.Next = node
    }
    last = node;
  }
  return result;
}

func makeListNode(nums []int) *ListNode {
  var result *ListNode
  var last *ListNode
  for _, num := range nums {
    node := new(ListNode)
    if result == nil {
      result = node
    }
    node.Val = num
    if last != nil {
      last.Next = node;
    }
    last = node;
  }
  return result
}

func printListNode(l *ListNode) {
  for node := l; node != nil; {
    fmt.Print(node.Val, "->")
    node = node.Next
  }
  fmt.Println("");
}

func main() {
  a := makeListNode([]int{2,4,3})
  b := makeListNode([]int{5,6,4})
  result := addTwoNumbers(a, b)
  printListNode(a)
  printListNode(b)
  printListNode(result)
}
