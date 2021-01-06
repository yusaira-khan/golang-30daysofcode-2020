package main

import (
	"bytes"
	"fmt"
	"bufio"
	"os"
)

func printSeparate(str string) {
	var evens, odds bytes.Buffer
	for pos, char := range str {
		if pos%2 == 0 {
			evens.WriteRune(char)
		} else {
			odds.WriteRune(char)
		}
	}
	fmt.Printf("%s %s\n", evens.String(), odds.String())
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	var str string
	fmt.Scan(&T)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < T; i++ {
		scanner.Scan()
		str = scanner.Text()
		printSeparate(str)
	}
}
