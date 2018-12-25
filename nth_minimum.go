// Order-Statistics: Find nth minimum element of the sorted array, given an
// unsorted array.

package main

import (
  "fmt"
)

// Entry point of the program.
func main() {
  inputList := []int{10, 3, 8, 0, 4, 6}
  var num int
  fmt.Print("Enter a number to check in the list: ")
  fmt.Scan(&num)
  isPresent := isNumberPresent(inputList, num)
  if isPresent {
    fmt.Println("Number", num, "is present in the list")
  } else {
    fmt.Println("Number", num, "is not present in the list")
  }
}

func isNumberPresent(arr []int, num int) bool {
  for _, value := range arr {
    if value == num {
      return true
    }
  }
  return false
}
