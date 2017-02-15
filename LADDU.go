package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1000000)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	t := nextInt(scanner)

	for i := 0; i < t; i++ {
		numActivities := nextInt(scanner)
		redeemPoint := 400
		if(nextString(scanner) == "INDIAN") {
			redeemPoint = 200
		}
		laddu := 0
		for j := 0; j < numActivities; j++ {
			activity := nextString(scanner)
			if (activity == "CONTEST_WON") {
				laddu +=300
				if bonus := 20 - nextInt(scanner); bonus > 0 {
					laddu += bonus
				}
			} else if (activity == "TOP_CONTRIBUTOR") {
				laddu += 300
			} else if (activity == "BUG_FOUND") {
				laddu += nextInt(scanner)
			} else if (activity == "CONTEST_HOSTED") {
				laddu += 50
			}
		}
		fmt.Printf("%d\n", laddu/redeemPoint)
	}
}