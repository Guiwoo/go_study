package template

import (
	"fmt"
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

type ServiceHandler struct {
	service ServiceTemplate
}

func (s *ServiceHandler) Init(service ServiceTemplate) *ServiceHandler {
	s.service = service
	return s
}
func (s *ServiceHandler) ParentsFeature01() {
	fmt.Println("부모의 기능01")
}
func (s *ServiceHandler) ParentsFeature02() {
	fmt.Println("부모의 기능02")
}

func (s *ServiceHandler) Run(w http.ResponseWriter, r *http.Request) {

	s.ParentsFeature01()
	s.ParentsFeature02()

	if s.service.IsRequireAuth() {
		fmt.Println("인증이 필요한경우 여기서 구현")
	}
	if err := s.service.GetRequest(); err != nil {
		fmt.Println("GetRequest 에러 처리")
	}
	if err := s.service.GetParam(); err != nil {
		fmt.Println("GetParam 에러 처리")
	}
	if err := s.service.Process(); err != nil {
		fmt.Println("Process 에러 처리")
	}
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

type MyContext struct {
	mux     *http.ServeMux
	service *ServiceHandler
}

func (m *MyContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}
func (m *MyContext) InitRouting() {
	m.mux.HandleFunc("/get", m.service.Init(ServiceGetTest{}).Run)
	m.mux.HandleFunc("/post", m.service.Init(ServicePostTest{}).Run)
}

func ProcessService() {

	myHandler := &MyContext{
		mux:     http.NewServeMux(),
		service: &ServiceHandler{},
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
