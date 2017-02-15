package main

import (
	"fmt"
	"math/rand"
	"bufio"
	"os"
)

func main() {
	arrBalls := [10]int{1,1,2,2,3,3,4,4,5,5}
	flag := true
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	index := -1
	for flag {
		fmt.Println("1")
		numFirstScaleBalls := rand.Intn(4) + 1
		numSecondScaleBalls := rand.Intn(4) + 1
		arrFirstScaleBalls := make([]int, numFirstScaleBalls)
		arrSecondScaleBalls := make([]int, numSecondScaleBalls)
		fmt.Printf("%d ", numFirstScaleBalls)
		for i := 0; i < numFirstScaleBalls; i++ {
			if randomIndex := rand.Intn(9); index > 0 && randomIndex == index {
				index = rand.Intn(9)
			} else {
				index = randomIndex
			}

			arrFirstScaleBalls[i] = arrBalls[index]
			fmt.Printf("%d ", arrFirstScaleBalls[i])
		}
		fmt.Println("")
		fmt.Printf("%d ", numSecondScaleBalls)
		index = -1
		for i := 0; i < numSecondScaleBalls; i++ {
			if randomIndex := rand.Intn(9); index > 0 && randomIndex == index {
				index = rand.Intn(9)
			} else {
				index = randomIndex
			}
			arrSecondScaleBalls[i] = arrBalls[rand.Intn(9)]
			fmt.Printf("%d ", arrSecondScaleBalls[i])
		}
		fmt.Println("")
		difference := nextInt(scanner)

		if(difference == 0) {
			if(numFirstScaleBalls == 1 && numSecondScaleBalls == 2) {
				fmt.Printf("%d ", arrFirstScaleBalls[0])
			}
			if(numSecondScaleBalls == 1 && numFirstScaleBalls == 2) {
				fmt.Printf("%d ", arrSecondScaleBalls[0])
			}
		} else if(difference < 0 && numSecondScaleBalls == 1){
			flag = false
			fmt.Printf("%d ", arrSecondScaleBalls[0])
		} else if(difference > 0 && numFirstScaleBalls == 1){
			flag = false
			fmt.Printf("%d ", arrFirstScaleBalls[0])
		}
	}
	return
}