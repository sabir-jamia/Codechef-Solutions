package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	testCase := nextInt(scanner)

	for i := 0; i < testCase; i++ {
		numA := nextInt(scanner)
		numB := nextInt(scanner)

		fmt.Printf("%d\n", numA + numB)
	}
}