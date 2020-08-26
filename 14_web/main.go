package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}

func main() {
	fmt.Println("WEB")
	port := ":3000"
	http.HandleFunc("/", index)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)
}
