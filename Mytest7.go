package main

import "fmt"

func main() {
	var nameCountryMap map[string] string
	nameCountryMap=make(map[string] string)
	
	nameCountryMap["liming"]="jinan"
	nameCountryMap["youlun"]="haidi"
	nameCountryMap["xiaoxian"]="hangzhou"
	nameCountryMap["huihui"]="newyork"

	for name := range nameCountryMap  {
		fmt.Println("name is ",name ,"home is ",nameCountryMap[name])
		
	}

	count, ok:=nameCountryMap["xiao jue"]
	if(ok ){
		fmt.Println("name is ",ok)
	}else  {
		fmt.Println("name is not",ok ,count)
	}
}
