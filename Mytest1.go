package main

import "fmt"
func main() {
  var  j int
	for a:=1 ;a<100; a++{
		for j:=2;j<=(a/j);j++{

			if(a%j==0){break; }
		}


 	  if( j >(a/j)){
		fmt.Println("the value is %d ",a)
		}

	}


}
