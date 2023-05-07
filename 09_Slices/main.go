package main

import (
	"fmt"
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









	// myFruitList = append(myFruitList, "benana", "mongo", "cherry")
	// fmt.Println(myFruitList)

	// myFruitList = append(myFruitList[:3])
	// fmt.Println(myFruitList)
  
	// //make 
	// highScores := make([]int, 4)

	// highScores[0] = 947
	// highScores[1] = 342
	// highScores[2] = 231
	// highScores[3] = 657
	// //highScores[4] = 654
	// fmt.Println(highScores)

	// highScores = append(highScores, 367, 789)

	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// //check ints is sorted or not
	// fmt.Println(sort.IntsAreSorted(highScores))

}