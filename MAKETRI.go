package main

import (
	"fmt"
	"os"
	"bufio"
)

func checkTriangle(a int64, b int64, x int64) bool {
	longest := a
	flag := 1
	
	if b > longest {
		longest = b
		flag = 2
	}

	if c > longest {
		longest = c
		flag = 3
	}

	if flag == 1 && (longest - b) <  c {
		return true
	} else if flag == 2 && (longest -c) <  a {
		return true
	} else if flag == 3 && (longest - a) <  b {
		return true
	}

	return false
}

func sumOfTwoMaximum(arr []int64) (int64, int64) {
	max1 := arr[0];
	max2 := arr[0];

	for i:= 0; i < len(arr); i++ {
		if max1 < arr[i] {
			max2 = max1
			max1 = arr[i]
		} else if max2 < arr[i] {
			max2 = arr[i]
		}
	}

	return max1, max2
}

func main() {
	var numArr []int64 
	var i,j,k,n,l,r int64
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	n = nextInt64(scanner)
	l = nextInt64(scanner)
	r = nextInt64(scanner)
	numArr = make([]int64, n)

	for i = 0; i < n; i++ {
		numArr[i] = nextInt64(scanner)
	} 

    max1, max2 := sumOfTwoMaximum(numArr)
	var flag bool
	count := 0
	for i = l; i <= r; i++ {
		flag = false

		if max1 <=  i - max2 {
				break
		}
		for j = 0; j < n; j++ {
			for k = j + 1; k < n; k++ {
				if checkTriangle(numArr[j], numArr[k], i) {
					flag = true
					break
				} 
			}
			
			if flag == true {
				break
			}
		}

		if flag {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}