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
	var numStr string
	var strLen, count int
	for i := 0; i < testCases; i++ {
		numStr = nextString(scanner)
		strLen = len(numStr)
		count = 0

		for j := 0; j < strLen; j++ {
			if numStr[j] == '4' {
				count++;
			}
		} 

		fmt.Printf("%d\n", count)
	}
}