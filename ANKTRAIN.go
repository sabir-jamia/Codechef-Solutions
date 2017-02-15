package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var testCases, birthNumber, nextBirthnuber int
	seats := [8]string{"LB", "MB", "UB", "LB", "MB", "UB", "SL", "SU"}
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	testCases = nextInt(scanner)

	for i := 0; i < testCases; i++ { 
		birthNumber = nextInt(scanner)
		c := birthNumber % 8;
		
		switch c {
			case 0 :
				nextBirthnuber = 7
			case 1 : 
				nextBirthnuber = 4
			case 2 :
				nextBirthnuber = 5
			case 3 :
				nextBirthnuber = 6
			case 4 :
				nextBirthnuber = 1
			case 5 :
				nextBirthnuber = 2
			case 6 :
				nextBirthnuber = 3
			case 7 :
				nextBirthnuber = 8
		}
		fmt.Printf("%d%s\n", ((birthNumber - 1 )/ 8) * 8 + nextBirthnuber, seats[nextBirthnuber - 1]);	
	}
}