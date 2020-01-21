package main

import (
  "fmt"
  "math"
)

func main () {
  v := 42
  if v >= 0 && v <= math.MaxUint8 {
    v8 := uint(8)
    fmt.Println("converted", v8)
  }
}
