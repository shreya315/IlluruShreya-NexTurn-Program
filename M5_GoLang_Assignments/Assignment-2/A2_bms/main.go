package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	InitDB()

	// Register HTTP routes
	http.Handle("/blog", Logging(http.HandlerFunc(CreateBlog)))
	http.Handle("/blog/", Logging(http.HandlerFunc(BlogByID)))
	http.Handle("/blogs", Logging(http.HandlerFunc(GetAllBlogs)))

	// Start the server
	log.Println("Starting server on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
