package main

import "fmt"

func main(){

		var courses = []string{"php", "go", "pascsal", "Ruby", "nodeJs"}

		//we can using append to remove go from the slice
		courses = append(courses[:1], courses[2:]...)
        fmt.Println(courses)

}