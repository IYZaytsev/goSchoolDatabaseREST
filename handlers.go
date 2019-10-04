package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
func ClassIndex(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    classes := School.GetClasses()

    if err := json.NewEncoder(w).Encode(classes); err != nil {
        panic(err)
    }
}
