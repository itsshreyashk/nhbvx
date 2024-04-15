package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"

	"context"
	"encoding/json"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/tinyurl/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Printf("as %v", id)
	}).Methods("GET")
	r.HandleFunc("/shorten/{url}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
        url := vars["url"]
        fmt.Printf("as %v", url)
	})
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", r)
}
