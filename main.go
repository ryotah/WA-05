package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Item struct {
	ID   string
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	i1 := Item{}
	log.Printf("Name = %v\n", i1.Name)
	i2 := &Item{}
	log.Printf("Name = %v\n", i2.Name)
	var i3 *Item
	log.Printf("Name = %v\n", i3.Name)

	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/panic", panicHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
