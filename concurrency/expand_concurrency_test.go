package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
	"testing"
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
