package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	testCases := nextInt(scanner)
	for i := 0; i < testCases; i++ {
		fmt.Printf("%d\n", revers(nextInt(scanner)))
	}
}

func revers(num int) int{
	res := 0

	for num != 0 {
		res = res * 10 +  (num % 10)
		num /= 10 
	}
	return res;
}