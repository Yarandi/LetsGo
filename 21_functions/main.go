package main

import "fmt"

func main() {
	//greeter()
	//result := adder(70, 90)
	//fmt.Println("result is: ", result)
	proRes, myMessage  := proAdder(12,13,34,5465)
	fmt.Println(proRes)
	fmt.Println("pro message is:", myMessage)

}

func greeter() {
	fmt.Print("Hamed from golang \n")
}

func adder(a int, b int) int {
	return a+b
}

func proAdder(values ...int) (int,string) {
	total := 0
	for _,val:= range values{
		total += val
	}
	return total, "Hi i want to pass two item"
}