/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
)

func checarFamilia(p1, p2 string, vetor [][]string) [][]string{
	flag_1 := false
	flag_2 := false
	var indice_p1, indice_p2 int
	
	for i := 0; i < len(vetor); i++{
		for k:=0; k<len(vetor[i]); k++{
			if(vetor[i][k] == " "){
				break
			}else if(p1 == vetor[i][k]){
				indice_p1 = i
				flag_1 = true
				break
			}else if(p2 == vetor[i][k]){
				indice_p2 = i
				flag_2 = true
				break
			}
		}
	}
	if(flag_1){
		if(flag_2){
			if(indice_p1 != indice_p2){ 
				vet_aux := make([]string,0)
				novo_vet := make([][]string,0)
				for k:=0;k<len(vetor[indice_p1]);k++{
					vet_aux = append(vet_aux,vetor[indice_p1][k])
				}
				for k:=0;k<len(vetor[indice_p2]);k++{
					vet_aux = append(vet_aux,vetor[indice_p2][k])
				}
				for i:=0;i<len(vetor);i++{
					if(i != indice_p1 && i != indice_p2){
						novo_vet = append(novo_vet, vetor[i])
					}
				}
				novo_vet = append(novo_vet, vet_aux)

				vetor = novo_vet
			}
		}else{
			vetor[indice_p1] = append(vetor[indice_p1],p2)
		}
	}else{
		if(flag_2){
			vetor[indice_p2] = append(vetor[indice_p2],p1)
		}else{
			aux := make([]string,0)
			aux = append(aux,p1)
			aux = append(aux,p2)
			vetor = append(vetor,aux)
		}
	}
	return vetor
}


func main(){
	var m, n int
	fmt.Scan(&m,&n)
	familia := make([][]string,0)
	for i:=0; i<n; i++{
		var pessoa1, relacao, pessoa2 string
		fmt.Scan(&pessoa1, &relacao, &pessoa2)
		familia = checarFamilia(pessoa1,pessoa2,familia)
	}
	fmt.Println(len(familia))
}