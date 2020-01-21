package main

import "fmt"

type sample func(x int) int

func sample2 (x int) int {
  return 0
}

func sample3 (fn sample) int {
  fmt.Println(fn(5))
  return 0
}

func main () {
  sample3(sample2)
}
