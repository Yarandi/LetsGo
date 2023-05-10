package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	goDefer()

	//Defer will reverse the program flow (last in first out)
}

func goDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//--------------------
//this program will return Hello
//--------------------

// 4
// 3
// 2
// 1
// 0
// three
// two
// one
