package main

import "fmt"

type  Books struct {
	title string;
	author string;
	subjects string

}

func main() {
	var book1 Books
	var book2 Books

	book1.author="Liming"
	book1.subjects="compute"
	book1.title="science"

	book2.author="wangming"
	book2.subjects="compute vff"
	book2.title="science adjd"
	//fmt.Printf(book1)
	//fmt.Printf(book2)
	printBooks(&book1)
	printBooks(&book2)
}

func printBooks( book *Books){
	fmt.Printf("author is %s\n",book.author)
	fmt.Printf("subjects is %s\n",book.subjects)
	fmt.Printf("title is %s\n",book.title)

}