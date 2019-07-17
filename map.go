package main

import (
  "fmt"
)

func main() {
  // 定义&初始化&增
  ages := make(map[string]int) 
  ages2 := map[string]int{
    "alice":   31,
    "charlie": 34,
  }
  ages["lih"]=4
  fmt.Print(ages, ages2)

  var ages3 map[string]int
  ages3 = make(map[string]int)
  ages3["la"] = 2
  fmt.Print(ages3)

  var ages4 = make(map[string]int)
  ages4["qianmian"] = 5
  fmt.Print(ages4)
  
  // 删
  delete(ages2,"alice")
  fmt.Print(ages2)

  // 改
  ages2["charlie"] = 9999
  fmt.Print(ages2)

  // 查
  for key, name := range ages2 {
    fmt.Print(key, name)
  }
  age, ok := ages2["hhh"]
  if !ok {
    fmt.Print("\n no have \n")
  }
  fmt.Print(age)

  // 比较
  // 不能直接比较，同slice一样，只能和nil比较，否则只能通过for循环来手动比较

}