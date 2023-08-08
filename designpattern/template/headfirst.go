package template

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/**
무언가 만드는 방법이 유사하다.
*/

type ServiceTemplate interface {
	IsRequireAuth() bool
	GetRequest() error
	GetParam() error
	Process() error
}

type ServiceGetTest struct{}

func (s ServiceGetTest) IsRequireAuth() bool {
	//TODO implement me
	panic("implement me")
}

func (s ServiceGetTest) GetRequest() error {
	//TODO implement me
	panic("implement me")
}

func (s ServiceGetTest) GetParam() error {
	//TODO implement me
	panic("implement me")
}

func (s ServiceGetTest) Process() error {
	//TODO implement me
	panic("implement me")
}

var _ ServiceTemplate = (*ServiceGetTest)(nil)

type ServicePostTest struct{}

func (s ServicePostTest) IsRequireAuth() bool {
	//TODO implement me
	panic("implement me")
}

func (s ServicePostTest) GetRequest() error {
	//TODO implement me
	panic("implement me")
}

func (s ServicePostTest) GetParam() error {
	//TODO implement me
	panic("implement me")
}

func (s ServicePostTest) Process() error {
	//TODO implement me
	panic("implement me")
}

var _ ServiceTemplate = (*ServicePostTest)(nil)

type ServiceStruct struct {
	GetTest  ServiceTemplate
	PostTest ServiceTemplate
}

func (s *ServiceStruct) HandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if err := s.GetTest.Process(); err != nil {
			log.Println(err)
			value, err := json.Marshal(struct {
				Code int    `json:"code"`
				Err  string `json:"err"`
			}{
				Code: 500,
				Err:  err.Error(),
			})
			if err == nil {
				panic(err)
			}
			w.Write(value)
		}
		w.Write([]byte("success"))
	case "POST":
		if err := s.PostTest.Process(); err != nil {
			log.Println(err)
			value, err := json.Marshal(struct {
				Code int    `json:"code"`
				Err  string `json:"err"`
			}{
				Code: 500,
				Err:  err.Error(),
			})
			if err == nil {
				panic(err)
			}
			w.Write(value)
		}
		w.Write([]byte("success"))
	}
}

func newServiceStruct() *ServiceStruct {
	return &ServiceStruct{
		GetTest:  &ServiceGetTest{},
		PostTest: &ServicePostTest{},
	}
}

type MyContext struct {
	ctx     context.Context
	mux     *http.ServeMux
	service *ServiceStruct
}

func (m *MyContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}
func (m *MyContext) InitRouting() {
	m.mux.HandleFunc("/get", m.service.HandleFunc)
	m.mux.HandleFunc("/post", m.service.HandleFunc)
}

func ProcessService() {

	myHandler := &MyContext{
		ctx:     context.Background(),
		mux:     http.NewServeMux(),
		service: newServiceStruct(),
	}

	http.NewServeMux()
	myHandler.InitRouting()

	s := &http.Server{
		Addr:           ":9300",
		Handler:        myHandler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
