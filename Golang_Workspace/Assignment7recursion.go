package main
import "fmt"
func main(){
	arr,key:=AcceptElement()
	index:=search(1,len(arr),key,arr)
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
//
func search(l int,h int,key int,arr []int)int{
	if(l==h){
		if(arr[l]==key){
			return l
		} else {
			return -1
		}
	} else {
		mid:=(l+h)/2
		if(arr[mid]==key){
			return mid
		}
		if(arr[mid]>key){
			return search(l,mid-1,key,arr)
		}
		if(arr[mid]<key){
			return search(mid+1,h,key,arr)
		}
	}
	return -1
}