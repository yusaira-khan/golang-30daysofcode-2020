package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	text  := scanner.Text()
	i2, _ := strconv.ParseUint(text,10,64)

	scanner.Scan()
	text2  := scanner.Text()
	d2,_ := strconv.ParseFloat(text2,64)

	scanner.Scan()
	s2  := scanner.Text()

	fmt.Println(i+i2)
	fmt.Printf("%.1f\n",d+d2)
	fmt.Println(s+s2)

}
