/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		
		entrada := scanner.Text()		

		if(entrada == "esquerda"){
			fmt.Println("ingles")
		}else if(entrada == "direita"){
			fmt.Println("frances")
		}else if(entrada == "as duas"){
			fmt.Println("caiu")
		}else{
			fmt.Println("portugues")
		}
	}
}