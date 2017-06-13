package main

import "fmt"

func main() {
	//var numbers  =make([] int ,3, 5)
	var numbers1=[] int {1,2,3,4,5}
	//numbers ={1,2,1,5,6}
	printSlice(numbers1)
	numbers2:=numbers1[1:2]
	printSlice(numbers2)
}

func printSlice(x [] int ){

	fmt.Printf("len=%d\n,cap=%d \n slice=%v\n",len(x),cap(x),x)
}