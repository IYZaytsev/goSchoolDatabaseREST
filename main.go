package main

import (
    "log"
    "net/http"
    "fmt"

    "./schoolLibrary"
)

var School schoolLib.School

func main() {

    School = schoolLib.NewSchool("oak grove")

    fmt.Println("starting router")

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
