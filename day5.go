package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func loop(n int32){
	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n",n,i,n*int32(i))
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	loop(n)

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
