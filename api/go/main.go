package main

import (
    "log"
    "net/http"
)

func main() {
	log.Print("Starting Hockey Player Generator API")
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}