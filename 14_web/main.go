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

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
