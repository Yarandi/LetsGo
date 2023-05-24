package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

//controllers - file

//server home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")
	//grab id from request 
	params := mux.Vars(r)
	
	//loop through courses, find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
	json.NewEncoder(w).Encode("no course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create one course")

	w.Header().Set("Content-Type", "application/json")
	//what if : bosy is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}
	//what about  - {}
	var course Course 
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	//generate unique id, string
	//append mew course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse (w http.ResponseWriter, r *http.Request){

	fmt.Println("update one course")

	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request
	//params := mux.Vars(r)
	

	//loop, id, remove, add
}