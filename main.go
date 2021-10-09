package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"Backend/Controller"
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

	log.Fatal(http.ListenAndServe(":8080", router))
	
}
