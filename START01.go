package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	testCases := nextInt(scanner)
	for i := 0; i < testCases; i++ {
		fmt.Printf("%d\n", nextInt(scanner))
	}

}