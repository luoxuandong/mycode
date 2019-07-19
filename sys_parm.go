package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Print(os.Args)
  for no, value := range os.Args {
    fmt.Print("\n")
    fmt.Print(no, value)
  }
}