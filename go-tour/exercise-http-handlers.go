package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str String) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, str)
}

func (str *Struct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s %s", str.Greeting, str.Punct, str.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe(":4000", nil))
}
