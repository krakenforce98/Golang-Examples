package main

import (
  "fmt"
  "math/rand"
)

func main () {
  var number = 78
  var guess = 0
  for 0 == 0 {
    guess = rand.Intn(100)
    if (guess == number) {
      fmt.Println("Right!")
      break
    }else if (guess > number) {
      fmt.Println("Too big.")
    }else{
      fmt.Println("Too Small.")
    }
  }
}
