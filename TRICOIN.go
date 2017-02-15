package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func getHeight(n int) int {
	i := 1
	for ; i <= n; i++ {
		sum := i * (i + 1) / 2
		if sum == n {
			return i
		} else if sum > n {
			return i - 1
		}
	}

	return i
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nexInt(scanner)

	for i := 0; i < t; i++ {
		fmt.Printf("%d\n", getHeight(nexInt(scanner)))
	}
}