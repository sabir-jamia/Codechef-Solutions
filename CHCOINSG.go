package main

import (
	"bufio"
	"os"
	"math"
	"fmt"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)
	for i := 0; i < t; i++ {
		num := nextInt(scanner)
		count := getCountOfPrimeFactors(num)
		if count == 1 {
			fmt.Printf("%s\n", "Chef")
		} else {
			fmt.Printf("%s\n", "Misha")
		}
	}
}

func getCountOfPrimeFactors(num int) int {
	count := 0
	if num == 1 {
		count++
		return count
	}
	flag := false
	for ; num % 2 == 0 ; num = num / 2 {
		flag = true
	}
	if flag {
		count++
	}
	for i := 3; i <= int(math.Sqrt(float64(num))); i = i + 2{
		flag = false
		for ;num % i == 0; num = num / i {
			flag = true
		}
		if flag {
			count++
		}
	}
	if num > 2 {
		count++
	}
	return count
}