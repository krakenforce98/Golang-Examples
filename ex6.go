package main 

import "fmt"

func main () {
	var distance = 56000000
	var time = 28 * 60
	var speed = distance / time

	fmt.Println(speed)

}