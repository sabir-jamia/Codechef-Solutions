package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)
	var i, j int
	for i = 0; i < t; i++ {
		n := nextInt(scanner)
		numOnes := 0
		str := nextString(scanner);
		for j = 0; j < n; j++ {
			if   str[j] == '1' {
				numOnes += 1
			}
		}
		fmt.Printf("%d\n", (numOnes * (numOnes + 1)/2))
	}
}