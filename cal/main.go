package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getIpAddress(reader *bufio.Reader) []string {
	var ip string
	if _, err := fmt.Fscanln(reader, &ip); err != nil {
		log.Fatalf("wrong ip address input %+v", ip)
	}

	ipArr := strings.Split(ip, ".")
	if len(ipArr) != 4 {
		log.Fatalf("wrong ip address %+v", ipArr)
	}
	return ipArr
}

func toBinaryArr(arr []string) []int {
	binaryArr := make([]int, len(arr))
	for i := range arr {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			log.Fatalf("worng ip address number %+v", arr)
		}
		binaryArr[i] = num
	}
	return binaryArr
}

func printArr(arr []int) {
	for i := range arr {
		fmt.Printf("%08b  ", arr[i])
	}
	fmt.Println()
}

func getAndPrintIPAddress(reader *bufio.Reader, data string) []int {
	ask := fmt.Sprintf("==%s 를 입력해주세요==", data)
	fmt.Println(ask)
	ipArr := getIpAddress(reader)
	ipBinary := toBinaryArr(ipArr)
	fmt.Printf("%s %+v\n", data, ipArr)
	printArr(ipBinary)
	return ipBinary
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ipBinary := getAndPrintIPAddress(reader, "IP 주소")

	subnetMaskBinary := getAndPrintIPAddress(reader, "서브넷 마스크")

	fmt.Println("==== and ====")
	fmt.Println("And 연산 결과")

	binary := ""
	ip := ""
	for i := range subnetMaskBinary {
		num := ipBinary[i] & subnetMaskBinary[i]
		binary += fmt.Sprintf("%08b  ", num)
		ip += fmt.Sprintf("%d  ", num)
	}

	fmt.Println(binary)
	fmt.Println(ip)
	fmt.Println()
}
