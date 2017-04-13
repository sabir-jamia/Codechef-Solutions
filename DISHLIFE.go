package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func nextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)
	for i := 0; i < t; i++ {
		n := nextInt(scanner)
		k := nextInt(scanner)
		mapIngrediensts := make([]int, k+1)

		some := false
		all := false
		count := 0
		for j := 1; j <= n; j++ {
			p := nextInt(scanner)
			for l := 1; l <= p; l++ {
				num := nextInt(scanner)
				if mapIngrediensts[num] != 0 {

				} else {
					mapIngrediensts[num] = 1
					count++
				}
			}

			if count == k && !all && !some {
				if j == n {
					all = true
				} else {
					some = true
				}
			}
		}

		if all {
			fmt.Println("all")
		} else if some {
			fmt.Println("some")
		} else {
			fmt.Println("sad")
		}
	}
}