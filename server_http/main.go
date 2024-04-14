package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit Middleware One")
		next.ServeHTTP(w, r)
	})
}

type GGroup struct {
	prefix string
	*http.ServeMux
	handler http.Handler
}

func (g *GGroup) Group(prefix string) *GGroup {
	_g := &GGroup{prefix: prefix, ServeMux: http.NewServeMux()}
	handler := http.StripPrefix(prefix, _g.ServeMux)
	_g.handler = handler
	return _g
}

func (g *GGroup) Use(path string, target http.Handler, handler ...func(http.Handler) http.Handler) {
	finalHandler := target
	for i := len(handler) - 1; i >= 0; i-- {
		finalHandler = handler[i](finalHandler)
	}
	g.HandleFunc(path, finalHandler.ServeHTTP)
}

func main() {
	_handler := http.NewServeMux()

	group := &GGroup{}
	routeA := group.Group("/a")
	aOne := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is /a/1 Handler"))
	})

	routeA.Use("/1", aOne, middlewareOne)
	routeA.HandleFunc("/2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is /a/2"))
	})

	_handler.Handle("/", routeA.handler)

	_server := &http.Server{
		Addr:           ":8080",
		Handler:        _handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(_server.ListenAndServe())
}
