package main

import(
	"fmt"
)

func func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	testCases := nextInt(scanner);
	arr []int;
	for i := 0; i < testCases; i++ {
		numElements := nextInt(scanner)
		arr = make([]int, numElements)
		for j := 0; j < numElements; j++ {
			arr[j] = nextInt(scanner)
			
		}
	}

}