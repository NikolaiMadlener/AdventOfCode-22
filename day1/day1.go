package main

import (
    "fmt"
	"bufio"
    "os"
	"strconv"
	"sort"
)

func main() {
	topElve()
	topThreeElves()
}

func topElve() {
	biggestNumber := 0
    filePath := "input.txt"
	readFile, err := os.Open(filePath)
 
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	
	counter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			number, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			counter += number
		} else {
			if counter > biggestNumber {
				biggestNumber = counter
			}
			// reset at empty newline
			counter = 0
		}
	}
 
	readFile.Close()

	
	fmt.Println(biggestNumber)
}

func topThreeElves() {
	totals := []int{}
	totalThree := 0

    filePath := "input.txt"
	readFile, err := os.Open(filePath)
 
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	
	counter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			number, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			counter += number
		} else {
			totals = append(totals, counter)
			// reset at empty newline
			counter = 0
		}
	}
	readFile.Close()

	sort.Ints(totals)
	// reverse array (big to small now)
	for i, j := 0, len(totals)-1; i < j; i, j = i+1, j-1 {
        totals[i], totals[j] = totals[j], totals[i]
    }
 
	for i := 0; i < 3; i++ {
		totalThree += totals[i]
	}
	
	fmt.Println(totalThree)
}