package main

import "fmt"

const MAX int   =3
func main() {
	a :=[] int {12,24,36}
	var i int
	var ptr [MAX] *int
	for i=0;i<MAX;i++{
		ptr[i]=&a[i]
	}
	for i=0;i< MAX;i++ {
		fmt.Printf("values of a[%d] =%d\n",i,*ptr[i])

	}

	var m int ;
	m=100;
	var m1 *int
	m1=&m;
	var m2 **int
	m2=&m1
	fmt.Printf("the value is %d \n",m)
	fmt.Printf("the value is %d \n",m1)
	fmt.Printf("the value is %d \n",*m1)
	//fmt.Printf("the value hello is %d ",**m1)
	fmt.Printf("the value is %d \n",m2)
	fmt.Printf("the value is %d \n",*m2)
	fmt.Printf("the value is %d\n ",**m2)



}
