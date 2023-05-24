package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	fmt.Println("server is getting Started...")
	
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}