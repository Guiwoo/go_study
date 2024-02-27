package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getEnv(key string) string {
	k, e := os.LookupEnv(key)
	if e == false {
		log.Panicf("fail to find env %+v", key)
	}
	return k
}

func main() {
	token := getEnv("token")
	//id := getEnv("id")
	//table := "HSAD-DATA2"
	//host := "https://api.airtable.com/v0"
	host := fmt.Sprintf("https://api.airtable.com/v0/appwE4RkhMda9jPuS/tblYcCQ01U58eE3Eo/listRecords")
	fmt.Println(host)
	req, err := http.NewRequest(http.MethodPost, host, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
