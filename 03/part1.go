package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isZeroMax(column []int64) bool {
	var result int
	for _, elm := range column {
		if elm == 0 {
			result += 1
		}
	}
	return result > len(column)/2
}

func main1() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	var table [][]int64

	scanner := bufio.NewScanner(file)
	var lineCounter int
	for scanner.Scan() {
		txt := scanner.Text()
		if lineCounter == 0 {
			for i := 1; i <= len(txt); i++ {
				arr := make([]int64, 0)
				table = append(table, arr)
			}
		}
		lineCounter += 1

		for i, elm := range txt {
			arr := table[i]
			bit, err := strconv.ParseInt(string(elm), 0, 64)
			if err != nil {
				panic(err)
			}
			arr = append(arr, bit)
			table[i] = arr
		}
	}
	var gammaRate = ""
	var epsilonRate = ""

	for _, line := range table {
		if isZeroMax(line) {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}
	gamma, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("gammaRate: %v\n", gamma)
	fmt.Printf("epsilonRate: %v\n", epsilon)
	fmt.Println("result:", gamma*epsilon)

}
