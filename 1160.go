/*
	By:
		Gabriel Castro
*/

package main

import (
	"fmt"
)

func calculoCrescimento(a, b int, ta, tb float64) int {
	cont := 0
	for a <= b {
		a = a + int(float64(a)*(ta/100.0))
		b = b + int(float64(b)*(tb/100.0))
		cont++
	}
	return cont
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var popA, popB int
		var taxaA, taxaB float64
		fmt.Scan(&popA, &popB, &taxaA, &taxaB)
		result := calculoCrescimento(popA, popB, taxaA, taxaB)
		if result > 100 {
			fmt.Println("Mais de 1 seculo.")
		} else {
			fmt.Printf("%d anos.\n", result)
		}

	}
}
