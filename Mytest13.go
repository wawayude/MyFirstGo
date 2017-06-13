package main

import "fmt"

func main() {
	p:=new(int)
	fmt.Println(*p)
	fmt.Println(p)

	var u uint8=254
	var i int8 =127
	fmt.Println(u,u+1,u+2,u*u)
	fmt.Printf("   the u %b,%b,%b\n",u,u+1,u+2)
	fmt.Printf("the i %#b,%#[1]b,%#[1]d,%#[1]x,%#[1]o",i,)

	fmt.Println("hello world")

	var z float64
	fmt.Println(z,-z,1/z,-1/z,z/z)

}
