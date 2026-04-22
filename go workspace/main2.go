package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	go get github.com/go-chi/cors

	//"github.com/go-chi/cors"
)

func main() {

	fmt.Println("hello world")

	godotenv.Load()

	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("port is not found")
	}
	r := chi.NewRouter()

	serv := &http.Server{
		Handler: r,
		Addr:    ":" + portstring,
	}
	log.Printf("server starting on port %v", portstring)
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)

	}
}
