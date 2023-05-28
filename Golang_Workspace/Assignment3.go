package main
import (
	"fmt"
	"sort"
)
func main(){
	var input string
	//taking input string
	fmt.Println("Enter the String:")
	fmt.Scanln(&input)

	mymap:= make(map[rune]int)
	//creating array of rune from string
	for _,ch :=range input{
		mymap[ch]++
	}
	
	arr := []int{}
	//creating array from counts
	for _,value :=range mymap{
		arr = append(arr, value)
	}
	//creating new map by reverting keys to values and values to key
	mymap2:=make(map[int]string)
	for key,value:=range mymap{
		val:=string(key)
		mymap2[value]=val
	}
	//sorting int array
	 sort.Ints(arr)
	ans := ""
	//iterating array and appending the character string to string
	for _,key :=range arr{
		for i:=0; i<key; i++ {
			ans += mymap2[key]
		}
	}

	fmt.Println("ans: ", ans)

}
