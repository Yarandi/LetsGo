package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{
			Name:     "ReactJS",
			Price:    299,
			Platform: "t.on",
			Password: "abc123",
			Tags:     []string{"web-dev", "js"},
		},
		{
			Name:     "Ruby",
			Price:    299,
			Platform: "t.on",
			Password: "abc123",
			Tags:     []string{"web-dev", "ruby"},
		},
		{
			Name:     "Go",
			Price:    340,
			Platform: "t.on",
			Password: "abc123",
			Tags:     []string{"web-dev", "js"},
		},
		{
			Name:     "PHP",
			Price:    2300,
			Platform: "t.on",
			Password: "abc123",
		},
	}

	//package this data as json data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {

	jsonDatafromWeb := []byte(`
	
	{
		"coursename": "ReactJS",
		"Price": 299,
		"website": "t.on",
		"tags": ["web-dev","js"]
    }
`)
	var lCourse course

	checkValidityOfJson := json.Valid(jsonDatafromWeb)

	if checkValidityOfJson {
		fmt.Println("json is valid")
		json.Unmarshal(jsonDatafromWeb, &lCourse)
		fmt.Printf("%#v\n", lCourse)
		//main.course{Name:"ReactJS", Price:299, Platform:"t.on", Password:"", Tags:[]string{"web-dev", "js"}}

	} else {
		fmt.Println("json is not valid")
	}

	//some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	//map[string]interface {}{"Price":299, "coursename":"ReactJS", "tags":[]interface {}{"web-dev", "js"}, "website":"t.on"}
	for key, value := range myOnlineData{
		fmt.Printf("key is %v and value is %v and Type is: %T \n", key, value, value)
	}
	//key is Price and value is 299 and Type is: float64
	//key is website and value is t.on and Type is: string
	//key is tags and value is [web-dev js] and Type is: []interface {}
	//key is coursename and value is ReactJS and Type is: string
}
