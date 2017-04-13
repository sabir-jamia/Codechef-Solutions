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

func generateGrundy() {
	totalNum := [1<<64]
	fmt.Printf("%d", totalNum);
}

func main() { 
	reader := bufio.NewReaderSize(os.Stdin, 1000000) 
	scanner := bufio.NewScanner(reader) 
	scanner.Split(bufio.ScanWords)
	generateGrundy()
	os.Exit(1)
	t := nextInt(scanner)
	for i := 0; i < t; i++ { 
		//numCakes := nextInt(scanner)
		//numQueries := nextInt(scanner)
		var numArr [4][4]int

		for j := 0; j < 4; j++ {
			numStr := nextString(scanner)
			for index, ch:= range numStr {
				numArr[j][index] = int(ch - '0')
			}
		}
		generateGrundy()
	}
}