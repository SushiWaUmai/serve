package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	log.Println("Starting File Server at port 3000")
	http.ListenAndServe(":3000", fs)
}
