package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	// "context"
	// "encoding/json"
	// "github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err!= nil {
        log.Fatal("Error loading.env file")
    }
	url := os.Getenv("MONGODB_URI")
	if url == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
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
