package main

import (
	"fmt"
	"bufio"
	"os"
)

func main () {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	testCases := nextInt(scanner)

	var  numArr[] int
	//len := 1
	//maxlen := 1 
	for i := 0; i < testCases; i ++ {
		numElements := nextInt(scanner)
		numArr = make([]int,	 numElements)
		for j := 0; j < numElements; j ++ {
			numArr[j] = nextInt(scanner)
		}

		fmt.Printf("%+v\n", numArr)
	}
}

func gcd(a,b int) int {
	for b != 0 {
		a, b = b, a %b
	}
	return a
}