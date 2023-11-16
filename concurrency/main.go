package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"syscall"
	"time"
)

type list struct {
	signal chan interface{}
	name   int
}
type Handler struct {
	list map[int]list
	sync sync.Mutex
}

func handlerStream(done <-chan interface{}) <-chan interface{} {
	stream := make(chan interface{})
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		defer func() {
			log.Println("handler stream closed")
			close(stream)
			ticker.Stop()
		}()
		// do something int stream handler
		for {
			select {
			case _, ok := <-done:
				if !ok {
					return
				}
				log.Println("got done signal close")
				return
			case <-ticker.C:
				time.Sleep(1 * time.Second)
				stream <- "something on your mind"
			}
		}
	}()
	return stream
}

func (h *Handler) Handle(a, b int) {
	log.Printf("got %d and %d", a, b)

	if b == 0 {
		log.Println("got 0")
		if handler, ok := h.list[a]; ok {
			h.sync.Lock()
			close(handler.signal)
			delete(h.list, a)
			h.sync.Unlock()
		}
	} else if b == -1 {
		for _, v := range h.list {
			fmt.Printf("go routine runngin :%d\n", v.name)
		}
	} else {
		//생성하는 로직
		if _, ok := h.list[a]; ok {
			return
		} else {
			log.Println("create go routine")
			h.list[a] = list{make(chan interface{}), a}
		}
		go func() {
			defer log.Println("go routine done")
			for _ = range handlerStream(h.list[a].signal) {
			}
		}()
	}
}

func NewHandler() *Handler {
	return &Handler{
		list: make(map[int]list),
		sync: sync.Mutex{},
	}
}

func mainTestRoutine() {
	go func() {
		http.ListenAndServe("localhost:4000", nil)
	}()

	handler := NewHandler()
	reader := bufio.NewReader(os.Stdin)
	for {
		var a, b int
		set := make(chan interface{})
		go func() {
			defer close(set)
			fmt.Fscanln(reader, &a, &b)
			set <- "done"
		}()
		<-set
		log.Println("got input")

		handler.Handle(a, b)

		log.Println("cycle done")
	}

	log.Println("go routine done")
}

func systemCall() {
	fmt.Println("system call ")
	file, _ := syscall.Open("./atask_user.log", syscall.O_RDONLY, uint32(0666))

	defer syscall.Close(file) // 파일 닫기 (defer를 사용하여 함수 종료 시 닫히도록 함)

	// 파일 읽기
	const bufferSize = 1024
	buf := make([]byte, bufferSize)
	for {
		// 파일에서 데이터 읽기
		n, err := syscall.Read(file, buf)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}

		// 더 이상 읽을 데이터가 없으면 종료
		if n == 0 {
			break
		}

		// 읽은 데이터 출력 또는 원하는 작업 수행
	}

	fmt.Println("system call done")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//
	//for i := 0; i < 1000; i++ {
	//	go systemCall()
	//}

	go func() {
		defer wg.Done()
		http.ListenAndServe("localhost:4000", nil)
	}()

	wg.Wait()
}
