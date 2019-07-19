package main

import (
  "fmt"
  "path/filepath"
)

func main() {

  // 构建
  pa := filepath.Join("/home/coding/","workspace","mycode","README.md")
  fmt.Println(pa)

  // 拆分
  p,f := filepath.Split(pa)
  fmt.Println(p)
  fmt.Println(f)
  fmt.Println(filepath.Ext(pa));

  // 
 // print(filepath.IsFile(pa)) //无判断path是文件还是目录的函数
 // print(filepath.IsDir(pa)) 
  fmt.Println(filepath.IsAbs(pa)) //
  abss,_ := filepath.Abs(pa)
  fmt.Println(abss) 
  fmt.Println(filepath.Base(pa)) 
  fmt.Println(filepath.Dir(pa)) 
  fmt.Println(filepath.Ext(pa)) 
}