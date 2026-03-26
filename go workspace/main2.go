package main

import (
	"fmt"
	"log"

	"os"
)

func main() {
	fmt.Println("hello world")

	portstring := os.Getenv("port")
	if portstring == "" {
		log.Fatal("port is not found")
	}
	fmt.Println("port:", portstring)

	//fmt.Println("listining  on port 8000")
	//http.ListenAndServe(":8000", nil)
}
