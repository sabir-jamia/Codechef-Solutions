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
	for i := 0; i < t; i++ { 
		numString := nextString(scanner) 
		count := len(numString) 
		flag := true 
		countZero := 0 
		countOne := 0
		for j := 0; j < count; j ++ { 
			
			if numString[j] == '0' { 
				countZero ++ 
			} else { 
				countOne ++ 
			} 
		}

		if count != 1 && (countZero == 0 || countOne == 0 || (countZero > 1 && countOne > 1)) {		
					flag = false
		}  

		if flag == true { 
			fmt.Printf("%s\n", "Yes") 
		} else { 
			fmt.Printf("%s\n", "No") 
		} 
	} 
}