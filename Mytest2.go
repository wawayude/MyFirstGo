package main

import "fmt"
import "math"

func main() {

	getRoot:=func(x float64) float64{
		return math.Sqrt(x)
	}

	fmt.Println(getRoot(16))
	var n[10] int
	for i:=0;i<=9;i++ {
		n[i]=i+100
	}

	for j:=0;j<=9;j++{
		fmt.Printf("a[%d]= %d\n",j ,n[j])
	}

	var name  = [] int {1000,3,2,50,17}
	var avg  float64
	avg=getAverage(name,5)
	fmt.Printf("balance %f",avg)
}

func getAverage( arr [] int ,size int ) float64{
	var i,sum int
	var avg float64
	for i=0;i<size;i++{
	sum += arr[i]
	}
	avg = float64(sum /size)
	return avg
}