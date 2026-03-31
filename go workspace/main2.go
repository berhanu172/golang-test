package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("hello world")

	godotenv.Load()
	
	

	portstring := os.Getenv("port")
	if portstring == "" { 
		log.Fatal("port is not found")
	}
		r:=chi.NewRouter()
 serv:= &http.server{
		Handler:r,
		Addr:   ":" +   portstring,
	}
	log.Printf("server starting on port"%v,portstring)
	err := serv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	fmt.Println("port:", portstring)

	}
