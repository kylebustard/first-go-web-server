package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_POST = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_POST, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
