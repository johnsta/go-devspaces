
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	fmt.Println("Starting up on 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {	
	fmt.Fprintln(w, "Hello world from GO!")
}
