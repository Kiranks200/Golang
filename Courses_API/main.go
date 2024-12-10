package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for couse-file
type Course struct {
	CourseId    string  `josn:"courseid"`
	CourseName  string  `josn:"cousename"`
	CoursePrice float64 `josn:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middleware,helper-file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building API's")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React.js", CoursePrice: 2000,
		Author: &Author{Fullname: "John", Website: "udemy.com"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 5000,
		Author: &Author{Fullname: "Jack", Website: "go.com"}})

	//routing
	r.HandleFunc("/", ServerHome).Methods("GET")
	r.HandleFunc("/courses", getallCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createonecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOnecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//port listing
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers-file

// Serverhome
func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>API Building</h1>"))
}

func getallCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET one course")
	w.Header().Set("Content-Type", "application/json")

	//find the id of the request
	params := mux.Vars(r)

	for _, course1 := range courses {
		if course1.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course1)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "Application/json")

	//body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course1 Course
	_ = json.NewDecoder(r.Body).Decode(&course1)
	if course1.IsEmpty() {
		json.NewEncoder(w).Encode("No data within json")
		return
	}

	for _, course := range courses {
		if course1.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exists")
		}
	}

	//generate unique id for the course
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	course1.CourseId = strconv.Itoa(rng.Intn(100))
	courses = append(courses, course1)
	json.NewEncoder(w).Encode(course1)
	return
}

func updateOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "Application/json")

	//get id feom request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}
}
