package main

import "fmt"

func main () {
  age := 41
  marsAge := float64(age)

  marsDays := 687.0
  earthDays := 365.2425
  marsAge = marsAge * earthDays / marsDays

  fmt.Println("I am", age, "years old on Mars.")
  fmt.Println(int(earthDays))
}