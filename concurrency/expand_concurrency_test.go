package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime/debug"
	"testing"
	"time"
)

/**
확장에서의 동시성
1. 에러전파 발생한 사건
	- 디스크가 가득찼다, 커넥션이 끊겼다, 자격이 만료되었다 등
2. 발생한 장소 및 시점
	- 전체 스택트레이스 가 포함되어야 한다. 스택 위쪽에서 에러에 접근할때는 쉽게 접근할수 있어야 한다.
	- 에러에는 실행 중인 컨택스트 와 고나련된 정보가 포함돼야 한다.
3. 사용자 친화적인 메세지
	- 시스템의 사용자에 맞춰 조정해야 한다. 간략하면서도 유의미한 정보가필요하다.
4. 사용자가  추가적인 정보를 얻을 수 있는 방법
	- 사용자에게 표시되는 에러는 에러가 기록된 시간이 아닌 에러의 발생 시간,스택 트레이스 등 참조될수 있는 아이디가 포함 되어야 한다.

잘알려진 에러로 는2 가지가 존재한다
- 버그
- 알려진 예외적인 경우 (네트워크 의 단절, 쓰기 실패, 버퍼 읽기실패등이 될수 있다.
- 올바른 형식으로 포장되어야 한다.
*/

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}
func (m MyError) Error() string {
	return m.Message
}

type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(wrapError(err, err.Error()))}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/some/where"
	isExcutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return IntermediateErr{wrapError(err, "cannot run job %q: requisite binaries not available", id)}
	} else if isExcutable == false {
		return wrapError(nil, "Job binary is not executable")
	}

	return exec.Command(jobBinPath, "--id="+id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logId : %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

func Test_expand01(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	err := runJob("1")
	if err != nil {
	}
	msg := "There was an unexpected issue; please report this as a bug"
	if _, ok := err.(IntermediateErr); ok {
		msg = err.Error()
	}
	handleError(1, err, msg)
}

/**
result := add(2,3)
writeTotallyToState(result)
result := add(2,3)
writeTotallyToState(result)
result := add(2,3)
writeTotallyToState(result)
result := add(2,3)
writeTotallyToState(result)
*/

/**
하트비트
일정 시간간격으로 발생하는 하트 비트 , 하나의 작업단위를 ㅓㅊ라히가 위해 다른일이 얼어날때까지 댁ㅣ중인 동시코드에 유용하다.
고루틴은 이 일이 언제 시작될지 알지못하기 때문에 뭔가 일어나기를 기다리는 동안 잠시 빈둥거리고 있을 것이다.
하트비트는 모든 것이 잘되고 있으며, 그 침묵이 예상된 것이라는 사실을 리스너에게 알리는 방법이다.
*/

func Test_HeartBeat(t *testing.T) {
	doWork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartBeat := make(chan interface{})
		result := make(chan time.Time)

		go func() {
			defer close(heartBeat)
			defer close(result)

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartBeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case result <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()

		return heartBeat, result
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 2 * time.Second

	heartbeat, rs := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-rs:
			if ok == false {
				return
			}
			fmt.Printf("result %+v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}
}

func TestHeartBeat_Panic(t *testing.T) {
	doWork := func(done <-chan interface{}, tick time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		result := make(chan time.Time)

		go func() {
			pulse := time.Tick(tick)
			workGen := time.Tick(2 * tick)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-pulse:
						sendPulse()
					case result <- r:
						return
					}
				}
			}

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, result
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 2 * time.Second
	heartbeat, rs := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-rs:
			if ok == false {
				return
			}
			fmt.Printf("result %+v", r.Second())
		case <-time.After(timeout):
			fmt.Println("worker go routine is not healthy")
			return
		}
	}
}

func Test_HeartBeat02(t *testing.T) {
	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		heartBeatStream := make(chan interface{}, 1)
		workStream := make(chan int)

		go func() {
			defer close(heartBeatStream)
			defer close(workStream)

			for i := 0; i < 10; i++ {
				select {
				case heartBeatStream <- struct{}{}:
				default:
				}
				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}

		}()
		return heartBeatStream, workStream
	}

	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			} else {
				fmt.Printf("result is %d\n", r)
			}
		}
	}
}
func DokWorkVer1(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{})
	result := make(chan int)
	go func() {
		defer close(result)
		defer close(heartbeat)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case result <- n:
			}
		}
	}()

	return heartbeat, result
}
func Test_HeartBeat_Writing(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{1, 2, 3, 4, 5}
	_, result := DokWorkVer1(done, intSlice...)
	for i, expected := range intSlice {
		select {
		case r := <-result:
			if r != expected {
				t.Errorf("index %d : expected %d but received %d", i, expected, r)
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("test time out")
		}
	}
}

func Test_HeartBeat_for_delay(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{1, 2, 3, 4, 5, 6}
	heartbeat, results := DokWorkVer1(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; expected != r {
			t.Errorf("index %d expected %d but got %d", i, expected, r)
		}
		i++
	}
}
