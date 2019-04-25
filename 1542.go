/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
)

func main(){
	var q, d, p int
	fmt.Scanln(&q,&d,&p)
	for q != 0{
		dias := float64((d*q)) / float64((p - q))
		paginas := float64(p) * dias
		result := int(paginas)
		if(result == 1){
			fmt.Println(result,"pagina")
		}else{
			fmt.Println(result,"paginas")
		}
		fmt.Scanln(&q,&d,&p)
	}
}