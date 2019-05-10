/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)


func mudarFrase(texto string) string{
	flag := false
	aux := ""
	for k:=0;k<len(texto);k++{
		if(!(string(texto[k]) == " ")){
			if(flag == false){
				aux += strings.ToUpper(string(texto[k]))
				flag = true
			}else{
				aux += strings.ToLower(string(texto[k]))
				flag = false
			}
		}else{
			aux += " "
		}
	}
	return aux
}



func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() != false{
	    entrada := scanner.Text()
	    frase := mudarFrase(entrada)
	    fmt.Println(frase)
	}
}