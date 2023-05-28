package main
import (
	"fmt"
)
func main(){
	//Taking input of string
	fmt.Println("Enter the String:")
	var str string
	fmt.Scanln(&str)
	//Converting Sring to rune array
	arr:=[]rune(str)
	//calling isvalid function taking array of rune as a parameter 
	value:=isValid(arr)
	//display message based on return type value
	if(value){
		fmt.Println("String is Valid")
	} else{
		fmt.Println("String is invalid")
	}
	
}
//Function for checking validation
func isValid(arr []rune) bool{
	//creating new slice
	slice:=make([]int,len(arr))
	//iterating array of rune and maintain count of equal number of characters and adding these count in slice
	for i:=0;i<len(arr);i++ {
		var count int
		for j:=0;j<len(arr);j++{
			if(arr[i]==arr[j]){
				
				count++
			}	
		}
		slice[i]=count
	}
	//checking for odd and even count by maintaining flag (if even=>invalid/false <-> if odd=>valid/true)
	flag:=true
	for i:=0;i<len(slice);i++{
		if(slice[i]%2==0){
			flag=false
		}
	}
	return flag
}