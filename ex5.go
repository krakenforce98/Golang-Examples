package main

import (
	"fmt"
	"math/rand"
)

func main () {
	var num = rand.Intn(10) + 1 // random integer number
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}