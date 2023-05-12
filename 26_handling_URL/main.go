package main

import (
	"fmt"
	"net/url"
)

func main() {

	const myUrl = "https://test.dev:3000/learn?coursename=reactjs&paymentid=34567865"

	//parsing url
	result, _ := url.Parse(myUrl)

	//fmt.Println(result.Scheme)   // https
	//fmt.Println(result.Host)     // test.dev:3000
	//fmt.Println(result.Path)     // /learn
	//fmt.Println(result.Port())   // 3000
	//fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=34567865

	queryParam := result.Query()
	fmt.Printf("%T \n", queryParam)       // url.values
	fmt.Println(queryParam["coursename"]) // [reactjs]
	fmt.Println(queryParam["paymentid"])  // [34567865]

	for _, value := range queryParam{
		fmt.Println("the param is: ", value) //the param is:  [reactjs] the param is:  [34567865]
	}
	for key , val := range queryParam{
		fmt.Printf("param is %v and value is %v \n", key, val) //param is coursename and value is [reactjs] ,param is paymentid and value is [34567865]
	}


	partofUrl := &url.URL{
		Scheme: "https",
		Host: "test.dev",
		Path: "/abc",
		RawPath: "user=hamed",
	}
	
	anotherURL := partofUrl.String()
	fmt.Println(anotherURL)
}

