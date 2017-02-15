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
	n := nextInt(scanner)
	w := nextInt64(scanner)
	l := nextInt64(scanner)
	initialHeight := make([]int, n)
	rate := make([]int, n)
	for i := 0; i < n; i++ {
		initialHeight[i] = nextInt(scanner)
		rate[i] = nextInt(scanner)
	}
	min := int64(999999999999999999)
	sum := int64(0)
	for i := 0; i < n; i++ {
		if int64(initialHeight[i]) >= l {
			sum += int64(initialHeight[i])
		}
		if sum >= w {
			fmt.Printf("%d", 0)
			return
		}
	}
	for i := 0; i < n; i++ {
		d := l - int64(initialHeight[i])
		h := d / int64(rate[i])
		if h < min {
			min = h
		}
	}

	if min <= 0 {
		min = 1
	}
	var height int64
	for i := min;;i++ {
		sum = 0
		for j := 0; j < n; j++ {
			height = int64(initialHeight[j]) + (i - 1) * int64((rate[j]))
			if height >= l {
				sum += height
			}
		}
		if sum >= w {
			fmt.Printf("%d", i - 1)
			return
		}
	}
}