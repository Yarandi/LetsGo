package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//PerformGetRequest()
      PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Lenght: ", response.ContentLength)

	//first way
	//content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))

	//second way using strings
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is: ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	request := strings.NewReader(`
		{
			"name":"hamed",
			"coursename":"golang",
			"website":"yarandi.com"
		}
	`)
	response, err := http.Post(myUrl, "application/json", request)
	if err != nil {
		panic(err)
	}
	//should close the body
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
