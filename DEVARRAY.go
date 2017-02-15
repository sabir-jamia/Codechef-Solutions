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
	N := nextInt(scanner)
	Q := nextInt(scanner)
	min := 1000000000
	max := 1
	for i:=0; i < N; i++ {
		input := nextInt(scanner)
		if input < min {
			min = input
		}
		if input > max {
			max = input
		}
	}
	for i:=0; i < Q; i++ {
		input := nextInt(scanner)
		if min <= input && input <= max {
			fmt.Printf("%s\n", "Yes")
		} else {
			fmt.Printf("%s\n", "No")
		}
	}
}
