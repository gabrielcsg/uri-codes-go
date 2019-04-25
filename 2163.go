/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func procurarSabre(vet [][]string, n, m int) (int, int){
	x, y := 0, 0
	for i:=0;i<n;i++{
		for k:=0;k<m;k++{
			if(vet[i][k] == "42"){
				if(i != 0 && i != n-1 && k != 0 && k != m-1){
					if(vet[i-1][k-1] == "7" && vet[i-1][k] == "7" && vet[i-1][k+1] == "7" && vet[i][k-1] == "7" && 
						vet[i][k+1] == "7" && vet[i+1][k-1] == "7" && vet[i+1][k] == "7" && vet[i+1][k+1] == "7"){
							x = i+1
							y = k+1
							return x,y
						}
				}
			}
		}
	}
	return x,y
}

func main(){
	var n, m int
	vetor := make([][]string,0)
	fmt.Scanln(&n, &m)
	scanner := bufio.NewScanner(os.Stdin)
	for i:=0;i<n;i++{
		scanner.Scan()
		linha := scanner.Text()
		linhaSplit := strings.Split(linha," ")
		vetor = append(vetor, linhaSplit)
	}
	fmt.Println(procurarSabre(vetor,n,m))
}