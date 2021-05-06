package main

import (
  "fmt"
)


func twoSum(nums []int, target int) []int {
  hashMap := make(map[int]int)
  for idx, num := range nums {
    if v, ok := hashMap[target - num]; ok {
      return []int{v, idx}
    }
    hashMap[num] = idx
  }
    return []int{-1, -1}
}


func main() {
  nums := []int{2, 3, 1, 8, 7}
  target := 11
  fmt.Print(twoSum(nums, target))
}
