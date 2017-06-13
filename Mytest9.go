package main

import "math"
import "fmt"
type Shape interface {
	area() float64
}

type Circle struct {
	x,y,radius float64
}
type Rectangle struct {
	width,height float64
}
func ( circle Circle) area() float64{
	return math.Pi *circle.radius *circle.radius
}
func(rect Rectangle) area() float64{
	return  rect.width *rect.height
}

func getArea(shape Shape) float64{
	return shape.area()
}


func main() {

	circle:=Circle{x:0,y:0,radius:6}
	rectangle:=Rectangle{width:10,height:5}
	fmt.Printf("circle area :%f \n",getArea(circle))
	fmt.Printf("Rectangle area:%f\n",getArea(rectangle))
}
