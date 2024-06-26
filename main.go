package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"context"

	"github.com/gorilla/mux"

	// "encoding/json"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	/*Ext.*/
	"math/rand"
	"math/random"
	"time"
)

type urlPair struct {
	OriginalURL string `bson:"original_url"`
	ShortenedURL string `bson:"shortened_url"`
}

func gRT(length int)  {
    randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)


	rd_chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQ"


}
func main() {
	/*Database*/
	err := godotenv.Load(".env")
	if err!= nil {
        log.Fatal("Error loading.env file")
    }
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}


	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err!= nil {
        log.Panic(err)
    }


	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()


	coll := client.Database("nhbvx").Collection("mappings")


	fmt.Printf("Coll Type : %T", coll)


	/* Creating Router */
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
	}).Methods("GET")
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", r)
}