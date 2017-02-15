package main

import (
	"bufio"
	"os"
	"fmt"
)

const mod  = 1000000007

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)
	for i := 0; i < t; i++ {
		n := nextInt(scanner)
		m := nextInt(scanner)
		s := nextInt(scanner)
		s--
		arr := make([]int, m + 1)
		for j := 0; j < m; j++ {
			arr[j] = nextInt(scanner)
		}
		dpArr := make([][]int, n )
		for j := range dpArr {
			dpArr[j] = make([]int, m + 1)
		}
		dpArr[s][0] = 1
		for k := 1; k <= m; k++ {
			for j := 0; j < n; j ++ {
				if j - arr[k - 1] >= 0 {
					dpArr[j][k] += dpArr[j - arr[k - 1]][k - 1]
				}
				if j + arr[k - 1] < n {
					dpArr[j][k] += dpArr[j + arr[k - 1]][k - 1]
				}
			}
		}
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", dpArr[j][m])
		}
		fmt.Println("")
	}
}