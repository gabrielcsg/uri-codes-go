/*
	By:
		Gabriel Castro
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func eTaut(s string) bool{
	s = strings.ToUpper(s)
	vetor := strings.Split(s, " ")
	
	letra := vetor[0][0]
	for i:=0;i<len(vetor);i++{
    	if(vetor[i][0] != letra){
    		return false
    	}	
    }
    return true
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		text := scanner.Text()
		if(text == "*"){
			break
		}
	    if(eTaut(text)){
	    	fmt.Println("Y")
	    }else{
	    	fmt.Println("N")
	    }
	}
}