package main

import (
  "fmt"
  "strconv"
)

// go 中没有集合，需要用Map来模拟集合（set）类型
func main() {
  myset1 := map[string]bool{}
  myset2 := map[string]bool{}

  for i := 0 ; i<=5 ; i++ {
    myset1["set"+strconv.Itoa(i)] = true
    myset2["set"+strconv.Itoa(i+2)] = true
  }
  fmt.Print(myset1)
  fmt.Print(myset2)
  // 增删改查同map

  // 常用操作 & | -
  fmt.Print("\n")
  fmt.Print(setUnion(myset1, myset2))
  fmt.Print("\n")
  fmt.Print(setIntersection(myset1, myset2))
  fmt.Print("\n")
  fmt.Print(setDifference(myset1, myset2))
}

// |
func setUnion(myset1, myset2 map[string]bool) map[string]bool {
  myset3 := map[string]bool{}
  for key,value := range myset1 {
    myset3[key] = value
  }
  for key,value := range myset2 {
    myset3[key] = value
  }
  return myset3
}

// &
func setIntersection(myset1, myset2 map[string]bool) map[string]bool {
  myset3 := map[string]bool{}
  for key, value := range myset1 {
    if _, ok := myset2[key]; ok {
      myset3[key] = value
    }
  }
  return myset3
}

// -
func setDifference(myset1, myset2 map[string]bool) map[string]bool {
  myset3 := map[string]bool{}
  for key, value := range myset1 {
    if _, ok := myset2[key]; !ok {
      myset3[key] = value
    }
  }
  return myset3
}