package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)
	max := -1
	for i := 0; i < t; i++ {
		n := nextInt(scanner)
		for j := 0; j < n; j ++ {
			res := 0
			if num := nextInt(scanner); max < num {
				max = num
			}
		}
	}
}
