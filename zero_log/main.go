package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func openFile(path int) (*os.File, error) {
	return os.OpenFile(
		fmt.Sprintf("%d", path),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
}

type FileWriter struct {
	log  chan byte
	cur  time.Time
	file *os.File
}

func (f FileWriter) Run(e *zerolog.Event, level zerolog.Level, message string) {
	// 날짜 비교 해서 날짜가 지났으면 카피 후 새파일 생성
	// 파일 열고 없으면 만들고
	// 파일 에 message를 작성한다.
}

func main() {
	idx := 1
	//file, _ := openFile(idx)
	console := zerolog.ConsoleWriter{Out: os.Stderr}
	zerolog.MultiLevelWriter(console)
	log := zerolog.New(console).With().Timestamp().Caller().Logger()
	log = log.Hook(FileWriter{})

	for {
		time.Sleep(1 * time.Second)
		log.Info().Msgf("1초 %d", idx)
		time.Sleep(1 * time.Second)
		log.Info().Msgf("2초 %d", idx)
		time.Sleep(1 * time.Second)
		log.Info().Msgf("3초 %d", idx)
		time.Sleep(1 * time.Second)
		log.Info().Msgf("4초 %d", idx)

		idx++
		break
		time.Sleep(5 * time.Second)
	}
}
