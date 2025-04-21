package main

import (
	//"fmt"
	"io"
	"log"
	"net/http"
)

func myServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, " Hello, first sever\n")
}

func main() {
	http.HandleFunc("/hello", myServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
