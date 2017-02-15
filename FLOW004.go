package main

import (
	"bufio"
	"os"
	"fmt"
)


func countDigits(num int) int{
	count := 0
	
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func power(num int) int {
	result := 1
	for i := 0; i < num; i++ {
		result *= 10
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var numDigits int
	testCases := nextInt(scanner)

	for i := 0; i < testCases; i++ {
		num := nextInt(scanner)
		numDigits = countDigits(num)
		fmt.Printf("%d\n", (num / power(numDigits - 1)) + (num % 10))
	}
}