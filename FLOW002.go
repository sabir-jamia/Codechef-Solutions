package main

import(
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	testCases := nextInt(scanner)
	for i := 0; i < testCases; i++ {
		numA := nextInt(scanner)
		numB := nextInt(scanner)

		fmt.Printf("%d\n", numA % numB)
	}
}