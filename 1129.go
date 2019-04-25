/*
	By:
		Gabriel Castro
*/

package main

import (
	"fmt"
)

func leitura(a,b,c,d,e int) string {
	/* Valor <= 127 == preto
	   Valor > 127 == branco
	*/
	if(a <= 127 && b > 127 && c > 127 && d > 127 && e > 127 ){
		return "A"
	}
	if (a > 127 && b <= 127 && c > 127 && d > 127 && e > 127){
		return "B"
	}
	if (a > 127 && b > 127 && c <= 127 && d > 127 && e > 127){
		return "C"
	}
	if (a > 127 && b > 127 && c > 127 && d <= 127 && e > 127){
		return "D"
	}
	if (a > 127 && b > 127 && c > 127 && d > 127 && e <= 127){
		return "E"
	}
	return "*"
}

func main(){
	var n int
	fmt.Scanln(&n)
	for(n != 0){
		for i:=0; i < n; i++{
		var a,b,c,d,e int
		fmt.Scan(&a, &b, &c, &d, &e)
		fmt.Println(leitura(a,b,c,d,e))
		}
		fmt.Scanln(&n)
	}
}