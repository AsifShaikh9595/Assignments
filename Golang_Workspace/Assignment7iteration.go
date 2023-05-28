package main
import "fmt"
func main(){
	arr,key:=AcceptElement()
	index:=search(arr,key)
	if(index<0){
		fmt.Print("Element is not present in Array")
	} else {
		fmt.Print("Element is present at Position ",index)
	}

}
//accepting array elements
func AcceptElement()([]int,int){
	fmt.Println("Enter the size of an array:")
	var size int
	var key int
	//accepting size of an array
	 fmt.Scanln(&size)
	//allocating memory for specified size
	 arr:=make([]int,size)
	//accepting array elements
	fmt.Println("Please Enter the element in Ascending order:")
	 for i:=0;i<len(arr);i++ {
		fmt.Println("Enter element no",i+1," :")
		fmt.Scanln(&arr[i])
	 }
	 fmt.Println("Enter the key for searching:")
	 fmt.Scanln(&key)
	 return arr,key
}
//Binary Searching
func search(arr []int , key int)(int){
	l:=1
	h:=len(arr)
	for l<=h {
		mid:=(l+h)/2
		if(arr[mid]==key){
			return mid		
		}
		if(arr[mid]>key){
			h=mid-1
		}
		if(arr[mid]<key){
			l=mid+1
		}
	}
	return -1
}