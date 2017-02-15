package main

import (
	"bufio"
	"strconv"
)

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func nextInt64(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	var num int64
	num, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	return num
}

func nextUInt64(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	var num uint64
	num, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	return num
}

func nextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func nextChar(scanner *bufio.Scanner) []byte {
	scanner.Scan()
	return scanner.Bytes()
}


