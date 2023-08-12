package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"testing"
	"time"
)

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

func TestHeartBeat04(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	heartbeat, results := DokWorkVer1(done, intSlice...)
	<-heartbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; expected != r {
			t.Errorf("index %d expected %d got %d\n", i, expected, r)
		}
		i++
	}
}

func DoWorkVer2(done <-chan interface{}, tick time.Duration, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(tick)

	numLoop:
		for _, n := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					continue numLoop
				}
			}
		}

	}()

	return heartbeat, intStream
}

func TestHeartbeatAllNums(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	const timeout = 2 * time.Second
	hearbeat, results := DoWorkVer2(done, timeout/2, intSlice...)
	<-hearbeat

	i := 0
	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf("expected %d got %d", expected, r)
			}
			i++
		case <-hearbeat:
		case <-time.After(timeout):
		}
	}
}

func TestHeartBeatCopyRequest(t *testing.T) {
	doWork := func(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {
		started := time.Now()
		defer wg.Done()

		simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		}

		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("took %v %v \n", took, id)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg *sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go doWork(done, i, wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	fmt.Printf("All go answered from %v\n", firstReturned)
}

/*
*
루틴 여러개를 만들어서 요청을 처리하고 그중 가장 빠르게 응답오는 답을 돌려주는 방법 이러면 ? 리소스를 많이사용하게 되잖아.... 부자들만 할수 있는 패턴이네..
메모리 상에서의 복제는 괜찮지만 , 핸들러 복제, 프로세스,서버 데이터 센터 복제는 ? 비용이 많이든다 ? 메모리 상의 복제가 무엇을 의미하는지 잘 모르겠다.
*/
func TestCopiedRequest(t *testing.T) {
	dowork := func(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {
		started := time.Now()
		defer wg.Done()

		simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		}

		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)

		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("%v took %v \n", id, took)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go dowork(done, i, &wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	fmt.Printf("Received an answer from %v \n", firstReturned)
}

// 지금까지 본것중에 제일 어이없ㄱ는 패턴

/*
 */
type RateLimiter interface {
	Wait(ctx context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

func Per(i int, t time.Duration) rate.Limit {
	return rate.Every(t / time.Duration(i))
}
func Open() *APIConnection {
	return &APIConnection{
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 1),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
		diskLimit:    MultiLimiter(rate.NewLimiter(rate.Limit(1), 1)),
		networkLimit: MultiLimiter(rate.NewLimiter(Per(3, time.Second), 3)),
	}
}

type APIConnection struct {
	networkLimit,
	diskLimit,
	apiLimit RateLimiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	err := MultiLimiter(a.apiLimit, a.networkLimit, a.diskLimit).Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	err := MultiLimiter(a.diskLimit, a.apiLimit, a.networkLimit).Wait(ctx)
	if err != nil {
		return err
	}
	return nil
}

func TestLimitSpeed(t *testing.T) {
	defer log.Printf("Done \n")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot readFile : %v\n", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot resolve address : %v\n", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}

/**
비정상 고루틴 치료
*/

type startGoruotine func(done <-chan interface{}, pluseInterval time.Duration) (heartbeat <-chan interface{})

func TestRecoverGoroutine(t *testing.T) {
	or := func(one <-chan interface{}, two <-chan interface{}) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			select {
			case stream <- one:
			case stream <- two:
			}
		}()
		return stream
	}
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
	newSteward := func(timeout time.Duration, goruotine startGoruotine) startGoruotine {
		return func(done <-chan interface{}, pluseInterval time.Duration) <-chan interface{} {
			heartbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)

				var wardDone chan interface{}
				var wardHeartbeat <-chan interface{}

				startWard := func() {
					wardDone = make(chan interface{})
					wardHeartbeat = goruotine(or(wardDone, done), timeout/2)
				}
				startWard()
				pulse := time.Tick(pluseInterval)

			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)
					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat:
							continue monitorLoop
						case <-timeoutSignal:
							log.Println("stward: ward unhealthy; restarting")
							close(wardDone)
							startWard()
							continue monitorLoop
						case <-done:
							return
						}
					}
				}
			}()

			return heartbeat
		}
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	bridge := func(done <-chan interface{}, intStream <-chan <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-intStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Printf("ward : Hello, I'm irreponsible!")
		go func() {
			<-done
			log.Println("ward : I am halting.")
		}()
		return nil
	}
	doWorkFn := func(done <-chan interface{}, intList ...int) (startGoruotine, <-chan interface{}) {
		intChanStream := make(chan (<-chan interface{}))
		intStream := bridge(done, intChanStream)
		doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
			intStream := make(chan interface{})
			heartbeat := make(chan interface{})
			go func() {
				defer close(intStream)
				select {
				case intChanStream <- intStream:
				case <-done:
					return
				}
				pluse := time.Tick(pulseInterval)
				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Printf("negative value : %v\n", intVal)
							return
						}
						for {
							select {
							case <-pluse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}
						}
					}

				}
			}()
			return heartbeat
		}
		return doWork, intStream
	}

	done := make(chan interface{})
	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 1, 2)
	doWorkWIthSteward := newSteward(1*time.Millisecond, doWork)

	time.AfterFunc(9*time.Second, func() {
		log.Println("main : halting steward and ward.")
		close(done)
	})
	doWorkWIthSteward(done, 1*time.Hour)
	for intVal := range take(done, intStream, 6) {
		fmt.Printf("Received : %v\n", intVal)
	}
	log.Println("Done")
}
