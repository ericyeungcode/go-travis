package main

import "fmt"

func CalcSum(x, y int) int {
	return x + y
}

func main() {
	val := CalcSum(100, 8)
	fmt.Printf("sum = %v\n", val)
}
