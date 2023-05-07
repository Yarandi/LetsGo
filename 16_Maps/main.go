package main

import "fmt"

func main(){

	//Creating map data structure
	languages := make(map[string]string)
	languages["js"] = "java Script"
	languages["RB"] = "Ruby"
	languages["Py"] = "Phyton"
	//fmt.Printf("%T \n",languages)
	//fmt.Println(languages)
    fmt.Println("Js shorts for", languages["js"])
	
	//delete a key 
	delete(languages, "js")
	//fmt.Println(languages)

	//How to Loop Through Maps in Go
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}
}