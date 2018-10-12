package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK test"))
	fmt.Println("hello world")
}

func main() {
	fmt.Println("http.HandleFunc")
	http.HandleFunc("/", IndexHandler)
	fmt.Println("http.ListenAndServe")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Done")
}
