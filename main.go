package main

import (
	"fmt"
	"net/http"
)

func main() {

	//var x string

	//var handler func(ResponseWriter, *Request)
	http.HandleFunc("/", home)

	http.ListenAndServe(":8888", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `Hi! this is jubayer ahmed, welcome to my first golang webpage`)
}
