package Restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func SampleAPI() {
	fmt.Println("Sample Rest API")

	//create a new router
	r := mux.NewRouter()

	//Specify endpoints, handler functions and HTTP method
	r.HandleFunc("/health-check", HealthCheck).Methods("GET")
	r.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", r)

	//start and listen to requests
	http.ListenAndServe(":8080", r)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	//declare response variable
	var response Response

	//Retrieve person details
	persons := prepareResponse()

	//assign person details to response
	response.Persons = persons

	//update content type
	w.Header().Set("Content-Type", "application/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	//update response
	w.Write(jsonResponse)
}

func prepareResponse() []Person {
	var persons []Person

	var person Person
	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "Newton"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "Einstein"
	persons = append(persons, person)

	return persons
}
