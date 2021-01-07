package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func swap(a *int32, b *int32) {
	t := *b
	*b = *a
	*a = t
}
func reverse(arr []int32) {
	l := len(arr) / 2
	e := len(arr) - 1
	i := 0
	for i < l {
		swap(&arr[i], &arr[e])
		i++
		e--
	}
}
func printWSpace(arr []int32){
	l:= len(arr)
	for i:=0; i < l; {
		fmt.Print(arr[i])
		i++
		if(i==l){
			fmt.Print("\n")
		}else{
			fmt.Print(" ")
		}
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}
	reverse(arr)
	printWSpace(arr)

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
