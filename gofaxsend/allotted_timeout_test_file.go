package main

import (
	"fmt"
	"github.com/timeout"
)

func main() {

	testCalculateAllottedTimeoutWhenTotalPagesIsLessThanOne()
	testCalculateAllottedTimeoutWhenTotalPagesIsEqualToOne()
	testCalculateAllottedTimeoutWhenTotalPagesIsGreaterThanOne()
}

func testCalculateAllottedTimeoutWhenTotalPagesIsLessThanOne() {
	fmt.Println("Given that the value of TotalPages is 0")
	fmt.Println("When CalculateAllottedTimeout is executed")
	fmt.Println("Then the allotted timeout is 400 which is", 400 == timeout.CalculateAllottedTimeout(0))
}

func testCalculateAllottedTimeoutWhenTotalPagesIsEqualToOne() {
	fmt.Println("")
	fmt.Println("Given that the value of TotalPages is 1")
	fmt.Println("When CalculateAllottedTimeout is executed")
	fmt.Println("Then the allotted timeout is 400 which is", 400 == timeout.CalculateAllottedTimeout(1))
}

func testCalculateAllottedTimeoutWhenTotalPagesIsGreaterThanOne() {
	fmt.Println("")
	fmt.Println("Given that the value of TotalPages is 2")
	fmt.Println("When CalculateAllottedTimeout is executed")
	fmt.Println("Then the allotted timeout is 800 which is", 800 == timeout.CalculateAllottedTimeout(2))
}



