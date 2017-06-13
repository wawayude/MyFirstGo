package main

import "fmt"

func feibo(i int ) (int ) {
	if i==0{
		return 0
	}
	if i==1{
		return 1
	}
	return  feibo(i-1) +feibo(i-2)
}

func main() {
	var i int
	for  i=0; i<10 ;i++  {
		fmt.Printf("%d ",feibo(i))

	}

	fmt.Printf("hello,world\n")

	var sum int =15
	var count int =5
	var mean float32
	mean =float32(sum)/float32(count)
	fmt.Printf("value of mean %f,",mean)


}
