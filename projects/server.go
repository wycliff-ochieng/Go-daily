package main

import (
	//"fmt"
	"io"
	"log"

	//"log"
	//"math/rand"
	"net/http"
	//"os"
)

func main() {
	http.HandleFunc("/hello", myServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
	//mux
}

//handler : myServer( function that processes https request in a web server)
// w - response writer. ResponseWriter is an interface used by http handler handle http requests
// req - http ,method whih deals with all

func myServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, " Hello, first sever\n")
}

// serveMux - basic httpRouter in golang
// multipelxing - is able to handle many routes
/*
type customServeMux struct{}

func (p customServeMux) ServeHttp(w http.ResponseWriter, r *http.Request){
	if r.URL.Path;os.IsNotExist(err)
}

func giveRandom(w http.ResponseWriter, r *http.Request){
	fmt.Pritnln(w, "your random numer is %f", rand.Float32())
}*/
