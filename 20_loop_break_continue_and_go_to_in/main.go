package main

import "fmt"

func main() {

	//days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday", "Monday"}

	//  for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	//  }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("index is %v and the value is %v \n", index, day)
	// }

	// rougueValue := 1
	// for rougueValue < 10{

	// 	if rougueValue == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// }

	// rougueValue := 1
	// for rougueValue < 10{

	// 	if rougueValue == 5 {
	// 		rougueValue++
	// 		continue
	// 	}
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			//jump to lco
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	// lets jump, hamed do you remember ? we had jumb in assemblly too :)
lco:
	fmt.Println("Jumping to here")

}
