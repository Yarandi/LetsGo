package main

import (
	"fmt"
	"sort"
	//"sort"
)

func main(){


	var myList = []string{"a","b","c"}
	fmt.Printf("%T \n", myList)
	fmt.Println(myList)

	//append to the list
	myList = append(myList, "d")
	fmt.Println(myList)

	//slice the slice
	myList = append(myList[1:])
	fmt.Println(myList)


	//make (default allocation of memory)
	highScores := make([]int, 4)
	highScores[0] = 12
	highScores[1] = 12
	highScores[2] = 12
	highScores[3] = 12
	//highScores[4] = 12 //makes error because this index in not in the range

	//but when we use append it is going to reallocate the memory and keep all the values
	highScores = append(highScores, 13, 14, 15, 16)
	fmt.Println(highScores)

	//sorting 
	sort.Ints(highScores)

	//check if this slice soreted or not
	fmt.Println(sort.IntsAreSorted(highScores))


}