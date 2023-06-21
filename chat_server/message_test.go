package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Link(t *testing.T) {
	userId := []string{"1", "2", "3", "4", "5"}
	msg := []string{"a", "b", "c", "d", "e"}

	list := make([]*Message, 5)
	for i := 0; i < 5; i++ {
		list[i] = NewMessage("1", userId[i], msg[i])
	}

	link := NewMessageLink()

	for i := 0; i < 5; i++ {
		link.Add(&MessageLink{msg: list[i]})
	}

	for link.next != nil {
		fmt.Println(link.next.msg)
		link = link.next
	}
}

var abc string = "----kTfLmS::bOuNdArY::tEsT:1128973598ksdjfhsdk\nContent-Type: text/plain; charset=\"euc-kr\"; name=\"text_1591728039780.txt\"\nContent-Transfer-Encoding: 8bit\n\n�\u05FD�Ʈ�Դϴ�.\n\n----kTfLmS::bOuNdArY::tEsT:1128973598ksdjfhsdk\nContent-Type: image/png; name=\"test.png\"\nContent-Transfer-Encoding: base64\n\n/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsK\nCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQU\nFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wAARCAGRAyADASIA\nAhEBAxEB/8QAHgAAAQMFAQEAAAAAAAAAAAAAAwIEBQABBggJBwr/xABpEAABAgQDBAUHAw0KCAsH\nAwUBAgMABAURBhIhBwgxQRMiUWFxCRQVMoGR0kJWkhcZIzM0UlNygpSVobEWGCRDVGKEwdHTREZ0\ndYWisrM2NzhVY3ODk7Th8CUmNWRlo8NHdsIoOUVYxP/EABwBAQACAwEBAQAAAAAAAAAAAAABAgME\nBQYHCP/EADARAAIBAwQCAgMAAQMEAwEAAAABAgMREgQTMVEFIQZBIjJhFCMzUhVxgZFCYqGx/9oA\nDAMBAAIRAxEAPwDdT1oUnjFgIVliLsoX4ACEqF9QYUrrQ2mZtEqtpCkqPSGwI4QuwLKere8JghT1\nbQK+sLlS8FSk9toFBUqIiGShcHHIwCCpVYQTZIvMecKFlaHhCE6woetE5FSymx3e6BKbF/VBg6uE\nDUq3hDKXZZWfKAKZHNKYC4yjsHuhypWbwgLg6vthlLsWX0MnkDiEiGyk87Wh+tPZDZ1vTWIbfY9f\n\n----kTfLmS::bOuNdArY::tEsT:1128973598ksdjfhsdk--\n\n----kTfLmS::bOuNdArY::tEsT:1128973598ksdjfhsda\nContent-Type: image/png; name=\"test11.png\"\nContent-Transfer-Encoding: base64\n\n/9j/4AAQSkZJRgABAQAAAQABAAD/11111111111111111111111111111111BQoHBwYIDAoMDAsK\nCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQU\nFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wAARCAGRAyADASIA\nAhEBAxEB/8QAHgAAAQMFAQEAAAAAAAAAAAAAAwIEBQABBggJBwr/xABpEAABAgQDBAUHAw0KCAsH\nAwUBAgMABAURBhIhBwgxQRMiUWFxCRQVMoGR0kJWkhcZIzM0UlNygpSVobEWGCRDVGKEwdHTREZ0\ndYWisrM2NzhVY3ODk7Th8CUmNWRlo8NHdsIoOUVYxP/EABwBAQACAwEBAQAAAAAAAAAAAAABAgME\nBQYHCP/EADARAAIBAwQCAgMAAQMEAwEAAAABAgMREgQTMVEFIQZBIjJhFCMzUhVxgZFCYqGx/9oA\nDAMBAAIRAxEAPwDdT1oUnjFgIVliLsoX4ACEqF9QYUrrQ2mZtEqtpCkqPSGwI4QuwLKere8JghT1\nbQK+sLlS8FSk9toFBUqIiGShcHHIwCCpVYQTZIvMecKFlaHhCE6woetE5FSymx3e6BKbF/VBg6uE\nDUq3hDKXZZWfKAKZHNKYC4yjsHuhypWbwgLg6vthlLsWX0MnkDiEiGyk87Wh+tPZDZ1vTWIbfY9f\nQhtXYdYeMu9axN4j8oSr1rwZlXWiMmT6+0S6XOUFbSFaHhDJpVodNuRbJ9kNR6CqZbOqmkk94ECc\nlWFD7Si/4ogwcz8OEVYRZTkito9DNUnLn+IR7opMnL5tWU28IeKbBEA9UxbckMI9F/MZZWnQoPiI\nSaPJq4sJPhBEOW8YLnJiu5LsjGPQ09DSX8nT7zCTRZP8APeYf8RFsoi27InCPRGGiSf4H9cIVRZT\n8EPeYk1DiITlEN2XZXGPRGehJLmyPef7Yv6Hkk6Bkfr/ALYfqSIQrjFdyXZO3Hoaei5P+TpHvi4p\ncnw6BP64cxUW3JE4R6Aikyg4sJPvi/oqSP8AgyR7/wC2DhRMKzRXcl2MI9DcUuTFv4On3mL+i5P+\nTog4VCVPBLqUZVEq520huS7GMegfoyS/k6feYumlyalW83TaD5QYuCGzxjHU1KoRzqSsi0aSk7JA\nPREkn/BkfriyqZJDjLIHvgc5WGpa4Op7oxypYiK7gryJ8Y+b+V+cUdLenpvykdzT+HdT3JeiemDS\n5biy2T3D/wA4jXqpKMm7Uu2D2qEYfPYkS3mKbBPaTHmWONvGGMHJUapWWULH8WhV1fqj5vW895ry\n07U72f0j0NLxenpe2j2+ZxUGwQX8iexNhEa/ihkJuFqc5xphirfap8qhRo9Kdm2+UxNOBls+Gbj7\n\n----kTfLmS::bOuNdArY::tEsT:1128973598ksdjfhsda--"

func recur(s string, rs []string) []string {
	base := "base64"
	target := "----"

	idx := strings.Index(s, base)
	if idx == -1 {
		return rs
	}
	s = s[idx+len(base):]
	idx = strings.Index(s, target)
	if idx == -1 {
		return rs
	}
	// recur call by call stack
	rs = append(rs, strings.TrimSpace(s[:idx]))
	return recur(s[idx+1:], rs)
}

func Test_Thing(t *testing.T) {
	rs := make([]string, 0)
	rs = recur(abc, rs)

	fmt.Println(len(rs))

	for _, v := range rs {
		fmt.Println(v)
		fmt.Println()
	}
}

func Test_slicing(t *testing.T) {
	idx := 11
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a := s[:idx]
	s = append(a, s[idx+1:]...)

	fmt.Println(s)
}

func slicer[T any](arr []T, target int) []T {
	var answer []T
	switch {
	case target < 0:
	case target == len(arr):
		answer = arr[:target]
	default:
		answer = arr[:target]
		answer = append(answer, arr[target+1:]...)
	}
	return answer
}

func Test_Thing2(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	arr = slicer(arr, 4)
	fmt.Println(arr)
}
