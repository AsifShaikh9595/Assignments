package main
import "fmt"
func main(){
	arr,target:=AcceptElement()
	index:= LinearSearch(arr,target)
	if index>=0{
		fmt.Println("Index position of ",target," is ",index)
	} else {
		fmt.Println("Element is not present ",index)
	}
	
}
//accepting array elements and target
func AcceptElement()([]int,int){
	fmt.Println("Enter the size of an array:")
	var size int
	var target int
	 fmt.Scanln(&size)

	//allocating memory location
	 arr:=make([]int,size)
	
	 //accepting array elements
	 for i:=0;i<len(arr);i++ {
		fmt.Println("Enter element no",i+1," :")
		fmt.Scanln(&arr[i])
	 }
	 //accepting target
	 fmt.Println("Enter the Target Element::")
	 fmt.Scanln(&target)
	 return arr,target
}
//linear searching for the target element
func LinearSearch(arr []int,target int) int {
	flag:=false
	var index int
	//checking for target element is present at which index
	for i:=0;i<len(arr);i++{
		if(arr[i]==target){
			flag=true
			index= i 
		}
	}
	if flag==false{
		index= -1
	}
	return index	
}