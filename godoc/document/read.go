// Package document 문서에 대한 패키지 읽고,쓰기 함수가 있습니다.
package document

import "fmt"

// Document Struct with title and content
type Document struct {
	title, content string
}

// Read a Document File return data with error
func (d *Document) Read() {
	fmt.Println("read something")
}
