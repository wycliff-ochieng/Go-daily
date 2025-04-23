package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type customServeMux struct{}

func (p *customServeMux) ServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a random number: %f", rand.Float64())
}

func main() {
	mux := &customServeMux{}
	http.ListenAndServe(":/8000", mux)
}
