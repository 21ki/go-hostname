// Simple program which prints the hostname of your webserver
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
const (
	listen = "0.0.0.0"
	port   = "8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if hostname, err := os.Hostname(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "Your hostname is: " + hostname)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen + ":" + port, nil))
}
