package main

import (
	"fmt"
	"net/http"
	"os"
)

var srvLocation string = os.Getenv("srvLocation")
var stat int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/favicon.ico" {
		stat++
	}
	fmt.Println("[Handled] " + r.RequestURI)
	fmt.Fprintf(w, "%d번째 방문자!", stat)
}

func main() {
	if len(srvLocation) < 1 {
		srvLocation = ":8080"
	}

	http.HandleFunc("/", handler)
	fmt.Println("Server is now on " + srvLocation + "!")
	http.ListenAndServe(srvLocation, nil)
}
