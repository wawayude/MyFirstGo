package main

import "fmt"

func main() {

	var a int =7
	var io *int
	io=&a
	fmt.Printf("the value of a %d \n",a)
	fmt.Printf("the value of a %d \n",&a)
	fmt.Printf("the value of io %d \n",io)
	fmt.Printf("the value of io %d \n",*io)




}
