// Simple program which prints the hostname of your webserver
// And return code500 after 10 hits
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

type counter struct {
	count int
}

var c counter

func (c *counter) init(initCount int) {
	c.count = initCount
}

func (c *counter) incr() {
	c.count++
}

func (c *counter) get() int {
	return c.count
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	c.incr()
	if c.count > 10 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "The hostname is: "+hostname+"\nThe Version is: V2\n")
	fmt.Printf("Count: %v\n", c.count)
}

func main() {
	c.init(0)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen+":"+port, nil))
}
