/*
	By:
		Gabriel Castro
*/

package main

import (
	"fmt"
)

func raiz(valor float64) float64{
	z := 1.0
	for i:=0; i < 1000; i++{
		z -= (z*z - valor) / (2 * z)
	}
	return z
}


func descobrirMedida(a,b,c int) float64{
	result := float64((a*b*100)/c)
	return raiz(result)
}

func main(){
	var a,b,c int
	fmt.Scanln(&a,&b,&c)
	for a != 0{
		fmt.Println(int(descobrirMedida(a,b,c)))
		fmt.Scanln(&a,&b,&c)
	}
}