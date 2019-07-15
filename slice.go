package main

import (
  "fmt"
)

func main() {

  // define , append
  var testSlice []int
  testSlice = append(testSlice, 22)
  testSlice = append(testSlice, 33)
  testSlice = append(testSlice, 44)
  testSlice = append(testSlice, 55)
  fmt.Print(testSlice)

  // 拼接
  slice1 := []int{22,33,44}
  slice2 := []int{33,55,66}
  slice3 := append(slice1,slice2...) // append 追加一个slice时要在后面加...
  fmt.Print("\n",slice3, "\n")

  // delete
  index:=1;
  testSlice=append(testSlice[:index],testSlice[index+1:]...)
  fmt.Print("after delete",testSlice, "\n")

  // for range
  for no,value := range testSlice {
    fmt.Print(no, value,"\n")
  }

  // 查询某个元素是否在slice中
  // slice没有这种方法，只能遍历，或者如果要常常使用，可以改变数据结构为map就容易进行这种查询

  // 比较两个slice是否相等
  //也没有如python的cmp方法，比较的话需要自己写，可以先Len比较长度，再遍历每个元素比较
  testSlice1 := []int{11,22,33}
  testSlice2 := []int{11,22,33}
  if len(testSlice1) != len(testSlice2) {
    fmt.Print("\n!=\n")
  }
  for no,value := range testSlice1 {
    if value != testSlice2[no] {
      fmt.Print("\n!=\n")
    }
  }
  fmt.Print("\n==\n")

  // max,min,len
  // go的slice也没有max与min，可以先排序然后获取，排序用sort包下的函数
  // 或者自己实现max与min函数
  fmt.Print("\n",len(testSlice), "\n")

}