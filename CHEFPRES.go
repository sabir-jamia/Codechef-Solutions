package main

import(
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	numCities := nextInt(scanner)
	fmt.Printf("%d\n", numCities)
	numProducts := nextInt(scanner)
	fmt.Printf("%d\n", numProducts)
	capitalCity := nextInt(scanner)
	fmt.Printf("%d\n", capitalCity)
	var cityDistanceA, cityDistanceB int
	cityDistArr := make([]int, numCities - 1) 
	prodcutArr := make([]int, numProducts)

	for i := 1; i < numCities; i++ {
		cityDistanceA = nextInt(scanner)
		fmt.Printf("%d\n", cityDistanceA)
		cityDistanceB = nextInt(scanner)
		fmt.Printf("%d\n", cityDistanceB)
		cityDistArr[cityDistanceA] = cityDistanceB
	}

	for i := 1; i <= numCities; i++ {
		prodcutArr[i] = nextInt(scanner)
	}
	fmt.Printf("%d\n", prodcutArr)

	numQuery := nextInt(scanner)
	var queryA, queryB int
	for i := 0; i < numQuery; i++ {
		queryA = nextInt(scanner)
		fmt.Printf("%d\n", queryA)
		queryB = nextInt(scanner)
		fmt.Printf("%d\n", queryB)
	}
}