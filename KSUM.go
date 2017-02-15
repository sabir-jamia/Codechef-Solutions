package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p - 1)
		quickSort(arr, p + 1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high - 1; j ++ {
		if arr[j] < pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i + 1], &arr[high])
	return i + 1
}

func swap(a,b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	var numArr []int
	reader :=  bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	n := nextInt(scanner)
	k := nextInt(scanner)
	numArr = make([]int, n)
	for i := 0; i < n; i ++ {
		numArr[i] = nextInt(scanner)
	}
	quickSort(numArr, 0, n - 1)

	fmt.Printf("%+v\n", numArr)
}
