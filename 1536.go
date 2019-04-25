/*
	By:
		Gabriel Castro
*/

package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Scanln(&n)
	for n > 0{
		time1 := 0
		time2 := 0
		var casa1,casa2,fora1,fora2 int
		var x string
		fmt.Scan(&casa1,&x,&fora2)
		fmt.Scan(&casa2,&x,&fora1)
		time1 += casa1 + fora1
		time2 += fora2 + casa2

		if(time1 > time2){
			fmt.Println("Time 1")
		}else if(time2 > time1){
			fmt.Println("Time 2")
		}else{
			if(fora1 > fora2){
				fmt.Println("Time 1")
			}else if(fora2 > fora1){
				fmt.Println("Time 2")
			}else{
				fmt.Println("Penaltis")
			}
		}
		n--
	}
}