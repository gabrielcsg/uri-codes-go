/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
)

func checarPercentual(totalS, totalB, totalA, saqueOk, blockOk, attackOk int)(float64,float64,float64){
	SaquePercentual := float64(saqueOk*100)/float64(totalS)
	BlockPercentual := float64(blockOk*100)/float64(totalB)
	AtaquePercentual := float64(attackOk*100)/float64(totalA)

	return SaquePercentual, BlockPercentual, AtaquePercentual
}

func main(){
	var n int
	var totalS, totalB, totalA int
	var saqueOk, blockOk, attackOk int
	var nomeJogador string
	var s,b,a,s1,b1,a1 int
	fmt.Scan(&n)
	for n>0{
		fmt.Scan(&nomeJogador)
		fmt.Scan(&s,&b,&a)
		fmt.Scan(&s1,&b1,&a1)
		totalS += s
		totalB += b
		totalA += a
		saqueOk += s1
		blockOk += b1
		attackOk += a1
		n--
	}
	p1, p2, p3 := checarPercentual(totalS,totalB,totalA,saqueOk,blockOk,attackOk)
	fmt.Printf("Pontos de Saque: %.2f %%.\n",p1)
	fmt.Printf("Pontos de Bloqueio: %.2f %%.\n",p2)
	fmt.Printf("Pontos de Ataque: %.2f %%.\n",p3)
}