package main

import (
	"Instagram/Controller"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Printf("Server started")

	// Routing using mux---------------------------------------------------------

	router := mux.NewRouter()
	router.HandleFunc("/posts", Controller.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", Controller.FetchPost).Methods("GET")
	router.HandleFunc("posts/user/{userPostId}", Controller.ListUserPosts).Methods("POST")
	router.HandleFunc("/users", Controller.ListUserPosts).Methods("POST")
	router.HandleFunc("/users/{userId}", Controller.ListUserPosts).Methods("GET")

	// --------------------------------------------------------------------------

	log.Fatal(http.ListenAndServe(":8079", router))

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27016")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	}

