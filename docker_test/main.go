package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /hello",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "Here is Hello Hello")
		},
	)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("fail to run server")
	}
}
