package main

import (
  "fmt"
)

// go 中没有集合，需要用Map来模拟集合（set）类型
func main() {
  myset1 := map[string]bool{"hello":false}
  fmt.Print(myset1)
}