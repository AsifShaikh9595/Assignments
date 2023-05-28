package main
import "fmt"

func main(){
	arr := AcceptElement()
	SortNumbers(arr)
	fmt.Println(arr)

}
//accepting array elements
func AcceptElement()([]int){
	fmt.Println("Enter the size of an array:")
	var size int
	//accepting size of an array
	 fmt.Scanln(&size)
	//allocating memory for specified size
	 arr:=make([]int,size)
	//accepting array elements
	 for i:=0;i<len(arr);i++ {
		fmt.Println("Enter element no",i+1," :")
		fmt.Scanln(&arr[i])
	 }
	 return arr
}
//Sorting using Bubble sort
func SortNumbers(arr []int){
	flag :=false
	
	for i:=0;i<len(arr)-1;i++ {
		
		for j:=0;j<len(arr)-i-1;j++ {
			//if first element is greater then swap
			if(arr[j]>arr[j+1]){
				temp:=arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=temp
				flag=true
			}
		}
		//flag is maintain for avoid unnecessry iterations if array is already sorted
		if(flag==false){
			fmt.Println("Array is already sorted")
			break
		}
	}
	
}
