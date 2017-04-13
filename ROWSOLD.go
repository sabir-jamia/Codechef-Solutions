package main

import (
	"bufio"
	"os"
	"fmt"
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
	var i, j int
	for i = 0; i < t; i++ {
		str := nextString(scanner);
		strLen := len(str)
		arr := make([]int, strLen)

		countOnes := 0
		for j = 0; j < strLen; j++ {
			if   str[j] == '1' {
				arr[j] = 1
				countOnes ++
			} else {
				arr[j] = 0
			}
		}

		count := 0
		//leftCountOnes := 0
		rightCountOnes := countOnes - 1
		for j := 0; j < strLen - 1; j++ {
			if arr[j] ==  1 {
				//fmt.Printf("%d\n", strLen - j - 1 - rightCountOnes)
				count += strLen - j - 1 - rightCountOnes
				if arr[j + 1] == 0 {
					//count++
					count += countOnes - rightCountOnes
				}
				rightCountOnes--
			}
		}
		fmt.Printf("%d\n", count)
	}
}