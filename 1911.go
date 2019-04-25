/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
)

func checarAssinaturaFalsa(vetor [][]string, nome, assinatura string) bool{
	cont := 0
	for i:=0;i<len(vetor);i++{
		if(nome == vetor[i][0]){
			for k:=0;k<len(assinatura);k++{
				if(assinatura[k]!=vetor[i][1][k]){
					cont++
				}
			}
			if(cont > 1){
				return true
			}
		}
	}
	return false
}

func main(){
	var n int
	fmt.Scan(&n)
	for n != 0{
		turma := make([][]string,0)
		for i:=0;i<n;i++{
			var nome, assinatura string
			fmt.Scan(&nome, &assinatura)
			aluno := []string{nome,assinatura}
			turma = append(turma,aluno)
		}

		var m int
		fmt.Scan(&m)
		quant := 0
		for j:=0;j<m;j++{
			var nome, assinatura string
			fmt.Scan(&nome, &assinatura)
			if(checarAssinaturaFalsa(turma,nome,assinatura)){
				quant++
			}
		}
		fmt.Println(quant)
		fmt.Scan(&n)
	}
}