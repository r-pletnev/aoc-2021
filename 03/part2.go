package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findOxygenGeneratorRating(table [][]int64) [][]int64 {
	for _, line := range table {
		fmt.Printf("line = %+v\n", line)
	}
	return table
}

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		panic(err)
	}
	var table [][]int64
	var lineCounter int
	scanner := bufio.NewScanner(file)
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
	fmt.Printf("table = %+v\n", table)
	// fmt.Printf("table[1][0] = %+v\n", table[1][1])

}
