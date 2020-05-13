package main

import "fmt"

func main(){
	var collections []string

	collections = []string{
		"Lorem 1",
		"Lorem 2",
	}

	fmt.Println(collections[0] + " " + collections[1])

	fmt.Println()

	for _,text := range collections {
		fmt.Println(text)
	}

	fmt.Println()

	var slice1 []string
	slice1 = append(slice1,"Fauzan")

	fmt.Println(slice1[0])
}