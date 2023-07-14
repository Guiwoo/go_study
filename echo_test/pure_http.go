package main

import (
	"fmt"
	"log"
	"net/http"
)

type Command interface {
	Execute(w http.ResponseWriter,r *http.Request)
}

type MyHandler struct {
	routes map[string]Command
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if cmd,ok := m.routes[path];ok {
		cmd.Execute(w,r)
	}else{
		http.NotFound(w,r)
	}
}

type FooHandler struct{}
func (f *FooHandler) Execute(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"This is Foo Handler")
}

func callPureHttp(port string) {

	handler := &MyHandler{
		routes: map[string]{
			"/foo": &FooHandler{},
		},
	}

	log.Fatal(http.ListenAndServe(":4000",handler))
}

type Strategy interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StrategyA struct{}

func (s *StrategyA) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

type StrategyB struct{}

func (s *StrategyB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, you've requested: %s\n", r.URL.Path)
}

type Context struct {
	strategy Strategy
}

func (c *Context) ExecuteStrategy(w http.ResponseWriter, r *http.Request) {
	c.strategy.ServeHTTP(w, r)
}

func callStrategy(){
	strategyA := &StrategyA{}
	strategyB := &StrategyB{}

	context := &Context{}

	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		context.strategy = strategyA
		context.ExecuteStrategy(w, r)
	})

	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		context.strategy = strategyB
		context.ExecuteStrategy(w, r)
	})

	http.ListenAndServe(":8080", nil)
}