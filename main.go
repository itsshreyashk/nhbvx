package main

import (
    "fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func main ()  {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Listening...")
    http.ListenAndServe(":8080", r)
}