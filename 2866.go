/*
	By:
		Gabriel Castro
*/

package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main(){
	var c int
	fmt.Scanln(&c)

	for i:=0;i<c;i++{
		var palavra string
		fmt.Scanln(&palavra)
		palavra = Reverse(palavra)
		palavraUp := strings.ToLower(palavra)
		result := ""
		for i := 0; i < len(palavra); i++ {
			if(palavra[i] == palavraUp[i]){
				result += string(palavra[i])
			}
		}

		fmt.Println(result)
	}
}