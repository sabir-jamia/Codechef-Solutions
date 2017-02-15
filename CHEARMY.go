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
	for i := 0; i < t; i++ {
		num := nextInt(scanner)
		n := num / 5
		res := 0
		for ; n > 4; {
			res = res + (n/5)*200+(n%5)*20
			n = n/5
		}
		res = res + ((num-1)%5)*2
		fmt.Printf("%d\n", res)
	}
}
