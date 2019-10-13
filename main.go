package main

import (
	"fmt"
	"log"
	"net/http"

	schoollib "./schoolLibrary"
)

//School is captial here so it can be accesed within the handlers.go file
var School *schoollib.School

func main() {

	School = schoollib.NewSchool("oak grove")

	fmt.Println("starting router")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
