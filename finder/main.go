package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	target := "/Users/guiwoopark/Desktop/RLS"
	suffix := ".cpp"

	fileList, err := getFilePaths(target, suffix)
	if err != nil {
		panic(err)
	}

	for _, v := range fileList {
		file, err := getFile(v)
		if err != nil {
			panic(err)
		}
		found := findWord(file, "Delete")
		fmt.Println(found)

		r := bufio.NewReader(os.Stdin)
		var a string
		fmt.Fscanln(r, &a)
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
		if strings.Contains(line, "AxDBResult") && strings.Contains(line, text) {
			sb.WriteString(line + "\n")
			for reader.Scan() {
				line := reader.Text()
				sb.WriteString(line + "\n")
				if strings.Contains(line, "return") {
					break
				}
			}
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
