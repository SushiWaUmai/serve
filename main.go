package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./files"))
	log.Println("Starting File Server at port 3000")
	err := http.ListenAndServe(":3000", fs)
	log.Fatal(err);
}
