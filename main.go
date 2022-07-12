package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello")
	})
	http.HandleFunc("/goodby", func(rw http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(rw, "hello %s", d)
	})
	http.ListenAndServe(":9000", nil)
}
