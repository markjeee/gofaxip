package main

import (
	"fmt"
	"github.com/timeout"
)

func main() {
	fmt.Println("Given the totalPages of 2")
	fmt.Println("When CalculateAllottedTimeout is executed")
	fmt.Println("Then the allotted timeout is 800 which is", 800 == timeout.CalculateAllottedTimeout(2))

	fmt.Println("")
	fmt.Println("Given that the value of TotalPages is 0")
	fmt.Println("When CalculateAllottedTimeout is executed")
	fmt.Println("Then the allotted timeout is 400 which is", 400 == timeout.CalculateAllottedTimeout(1))
}


