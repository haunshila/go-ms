package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello WOrld")
		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data %s\n", d)
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)

}
