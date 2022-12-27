package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type Students struct {
	Name  string `json:"name"`
	Study string `json:"study"`
}

var Student = make([]Students, 0)

func createStudent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var prodigy Students
	err := json.NewDecoder(request.Body).Decode(&prodigy)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	writer.WriteHeader(http.StatusOK)

	Student = append(Student, prodigy)
	err = json.NewEncoder(writer).Encode(&prodigy)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}

}

func updateStudent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var prodigy Students
	err := json.NewDecoder(request.Body).Decode(&prodigy)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	for index, structs := range Student {
		if structs.Name == prodigy.Name {
			Student = append(Student[:index], Student[index+1:]...)
		}
	}
	Student = append(Student, prodigy)
	err = json.NewEncoder(writer).Encode(&prodigy)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
}

func readStudent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(request, "name")
	for _, structs := range Student {
		if structs.Name == params {
			err := json.NewEncoder(writer).Encode(&structs)
			if err != nil {
				log.Fatalln("There was an error encoding the initialized struct")
			}
		}
	}

}

func deleteStudent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(request, "name")
	indexChoice := 0
	for index, structs := range Student {
		if structs.Name == params {
			indexChoice = index
		}
	}
	Student = append(Student[:indexChoice], Student[indexChoice+1:]...)
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/get/{name}", readStudent)
	router.Post("/", createStudent)
	router.Put("/update/{name}", updateStudent)
	router.Delete("/delete/{name}", deleteStudent)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}
