package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	target := "/Users/guiwoopark/Desktop/OAS"
	suffix := ".cpp"

	fileList, err := getFilePaths(target, suffix)
	if err != nil {
		panic(err)
	}

	for _, v := range fileList {
		file, _ := getFile(v)
		rs := findWord(file, "SELECT")
		fmt.Println(rs)
	}
}
func findWord(file *os.File, text string) []string {
	var found []string
	name := "✅" + file.Name() + "✅\n"
	found = append(found, name)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		sb := strings.Builder{}
		line := reader.Text()
		if strings.Contains(line, text) {
			sb.WriteString(line + "\n")
			found = append(found, sb.String()+"\n")
		}
	}
	return found
}

func getFile(path string) (*os.File, error) {
	return os.Open(path)
}

func getFilePaths(target, suffix string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}
