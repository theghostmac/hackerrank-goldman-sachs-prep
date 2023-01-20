package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each socks,
// determine how many pairs of socks with matching colors there are.
/* --------> EXAMPLE <---------
n = 7
ar = [1,2,1,2,1,3,2]
There is one pair of color 1 and one of color 2. There are 3 odd socks left, one of each color. The number of pairs is 2.
*/

// sockMerchant accepts the number of socks in the pile, and the array of numbers representing the socks. It returns the number of pairs in the pile.
func sockMerchant(n int32, ar []int32) int32 {
	/* ---> STEPS <---
	1. Check if there are 2 socks of same type (numbers)
	2. If found, untag the sock type, and increment it
	3. Create a hashmap with int32 key and bool value
	4. Iterate through all array elements ar
	*/
	pair := map[int32]bool{}
	match := int32(0)
	for _, available := range ar {
		if pair[available] {
			pair[available] = false
			match++
		} else {
			pair[available] = true
		}
	}
	return match

	// Solution is O(n) time and O(c) space complexity.
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1-24*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)
	fmt.Println(writer, "%d\n", result)
	writer.Flush()
}

// checkError handles errors in the program.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// readLine receives a reader type variable and converts to a string.
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
