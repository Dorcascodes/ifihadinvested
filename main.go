package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/saintmalik/ifihadinvested/handlers"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handlers.Ifihad)
	http.HandleFunc("/worthnow", handlers.Invested)
	http.HandleFunc("/hodl", handlers.Ifihadhodl)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
