package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req http.Request) {
	fmt.Fprint(w)
}

func main() {
	fmt.Println("WEB")

	http.HandleFunc("/", index)
}
