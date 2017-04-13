package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func nextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)
	var i, j int
	for i = 0; i < t; i++ {
		n := nextInt(scanner)
		numOnes := 0
		result := int64(0)
		str := nextString(scanner);
		for j = 0; j < n; j++ {
			if   str[j] == '1' {
				numOnes += 1
				result += int64(numOnes)
			}
		}
		fmt.Println(result)
	}
}