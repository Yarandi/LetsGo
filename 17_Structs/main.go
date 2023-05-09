package main

import "fmt"

func main() {

	hamed := User{
		Name:   "Hamed",
		Email:  "hamed.yarandi@gmail.com",
		Status: true,
		Age:    35,
	}

	fmt.Println(hamed)
	fmt.Printf("details: %+v \n", hamed)
	fmt.Printf("name is: %+v \n", hamed)
	fmt.Printf("Name is %v and email is %v \n", hamed.Name, hamed.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}
