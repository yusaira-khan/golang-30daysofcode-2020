package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
func factorial(n int32) int32 {
	if n<2 {
		return 1;
	}
	return n * factorial(n-1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	result := factorial(n)
	fmt.Printf("%d\n", result)
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
