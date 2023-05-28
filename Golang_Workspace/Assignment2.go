package main

import "fmt"

func main(){
	var num int
	fmt.Scanln(&num)
	for i:=0;i<=num;i++ {

		for k:=0;k<i;k++ {
			fmt.Print(" ")
		}

		for j:=i;j<num;j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}