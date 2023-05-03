package main

import "fmt"

//we can not declare variable without var style out of main function, But he didnt say why, But Tommy said :)
//jwttoken := "w94852930437327429837429384"
var jwttoken string = "38497583947659834753489758"
const JwtToken string = "ldakjfaksljdkljasdaalksdj"

func main(){
	var student = true
	fmt.Println(student)
	fmt.Printf("this undefined variable type is : %T \n", student)

	// var intvalue int = 2456
	// var floatvalue float32 = 23.7464748589567696
	// var boolvalue bool = true
	// var undefinedValue = false

	intvalue := 2456
	floatvalue := 23.93488594305803485
	boolvalue := true
	undefinedValue := false

	fmt.Printf("value type is : %T \n", intvalue)
	fmt.Printf("value type is : %T \n", floatvalue)
	fmt.Printf("value type is : %T \n", boolvalue)
	fmt.Printf("value type is : %T \n", undefinedValue)
	
	fmt.Printf("value type of const is : %T \n", JwtToken)
}
