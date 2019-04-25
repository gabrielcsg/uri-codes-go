/*
	By:
		Gabriel Castro
*/

package main

import (
	"fmt"
)

func checarPomekon(vetor []string, valor string) bool {
	for i := 0; i < len(vetor); i++ {
		if valor == vetor[i] {
			return true
		}
	}
	return false
}

func main() {
	var n int
	pomekons := make([]string, 0)
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var entrada string
		fmt.Scanln(&entrada)
		if !(checarPomekon(pomekons, entrada)) {
			pomekons = append(pomekons, entrada)
		}
	}
	capturados := len(pomekons)
	resultado := 151 - capturados
	fmt.Printf("Falta(m) %d pomekon(s).\n",resultado)
}
