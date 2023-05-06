package main

import "fmt"

func main(){

	//define an int variable named one and 2 value
	//var one = 2 

	//define an int variable named one and 2 value
	//var one int = 2

    //define an int variable named one and 0 value
	 //var one int

	//define a ponter named ptr and <nill> value
	//var prt *int

	// fmt.Println("the value is", prt)
	// fmt.Printf("the type of value is: %T", prt)
 
	myNumber := 32

	var pointer = &myNumber

	fmt.Println("Valu of actual pointer is:", pointer) //0xc0000ac008
	fmt.Println("Valu of actual pointer is:", *pointer) //32

	*pointer = *pointer + 2

	fmt.Println("New Value is : ", myNumber) //34

}