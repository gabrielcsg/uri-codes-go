/*
	By:
		Gabriel Castro
*/

package main

import "fmt"

var num = 0

func fib(valor int) int {
	num++
	if valor == 0 {
		return 0
	}
	if valor == 1 {
		return 1
	}
	return fib(valor-1) + fib(valor-2)
}

func main() {
	var entrada1, entrada2 int
	fmt.Scanln(&entrada1)
	for entrada1 > 0 {
		fmt.Scanln(&entrada2)
		result := fib(entrada2)
		fmt.Printf("fib(%d) = %d calls = %d\n", entrada2, num-1, result)
		num = 0
		entrada1--
	}
}
