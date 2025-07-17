package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "ok")
	})

	server.ListenAndServe()
}
