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
	var num, sum int
	for i := 0; i < testCases; i++ {
		num = nextInt(scanner)
		sum = 0

		for num != 0 {
			sum += num % 10
			num = num / 10
		}

		fmt.Printf("%d\n", sum)
	}
}