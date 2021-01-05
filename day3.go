package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func weirdStr(n int32) string {
	if n%2 == 1 { //is odd
		return "Weird"
	} else if 2 <= n && n <= 5 {
		return "Not Weird"
	} else if 6 <= n && n <= 20 {
		return "Weird"
	} else {
		return "Not Weird"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	fmt.Println(weirdStr(N))

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
