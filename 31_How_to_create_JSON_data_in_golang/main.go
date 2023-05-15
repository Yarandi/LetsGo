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
	EncodeJson()
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
