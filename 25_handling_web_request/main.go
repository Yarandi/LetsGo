package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	const url = "https://lco.dev"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("the type of response is: %T \n", response) //the type of response is: *http.Response
	defer response.Body.Close() //its so important and it our responcibility to close the request

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}
