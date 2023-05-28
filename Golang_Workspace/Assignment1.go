package main

import "fmt"

func main(){
	var num int
	fmt.Scanln(&num)
	for i:=1;i<=num;i++ {

			for l:=i;l<num;l++{
				fmt.Print(" ")
			}

			for j:=1;j<=i;j++ {
				fmt.Print("*")
			}
		
			for k:=i-1;k>=1;k-- {
				fmt.Print("*")
			}
			fmt.Println(" ")

	}
}