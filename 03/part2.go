package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isOneMax(column []int64) bool {
	var result int
	for _, elm := range column {
		if elm == 1 {
			result += 1
		}
	}
	return result > len(column)/2
}

func getRareCommonValue(column []int64) int64 {
	if isZeroMax(column) {
		return 1
	}
	if isOneMax(column) {
		return 0
	}

	// zero & one is equals: select 0
	return 0
}

func getMostCommonValue(column []int64) int64 {
	if isZeroMax(column) {
		return 0
	}
	if isOneMax(column) {
		return 1
	}
	// zero & one is equals: select 1
	return 1
}

func filterTableByMCV(table [][]int64, currentIndex int, columnLength int) [][]int64 {
	if currentIndex == columnLength {
		return table
	}
	line := table[currentIndex]
	mcv := getMostCommonValue(line)

	acc := make([][]int64, 0)
	for i, bit := range line {
		if bit == mcv {
			arr := make([]int64, columnLength)
			for j := 0; j < columnLength; j++ {
				current := table[j][i]
				arr[j] = current
			}
			acc = append(acc, arr)
		}
	}
	var result [][]int64
	for i, row := range acc {
		for j, bit := range row {
			if i == 0 {
				arr := make([]int64, 0)
				result = append(result, arr)
			}
			arr := result[j]
			arr = append(arr, bit)
			result[j] = arr
		}
	}
	return filterTableByMCV(result, currentIndex+1, columnLength)
}

func filterTableByRCV(table [][]int64, currentIndex, columnLength int) [][]int64 {
	// fmt.Printf("curIndex = %+v\n", currentIndex)
	// fmt.Printf("table	 = %+v\n", table)
	// if currentIndex == columnLength {
	// 	return table
	// }
	line := table[currentIndex]
	mcv := getRareCommonValue(line)
	// fmt.Printf("line = %+v\n", line)
	// fmt.Printf("mcv = %+v\n", mcv)

	acc := make([][]int64, 0)
	for i, bit := range line {
		if bit == mcv {
			arr := make([]int64, columnLength)
			for j := 0; j < columnLength; j++ {
				current := table[j][i]
				arr[j] = current
			}
			acc = append(acc, arr)
		}
	}
	if len(acc) == 1 {
		return acc
	}

	var result [][]int64
	for i, row := range acc {
		for j, bit := range row {
			if i == 0 {
				arr := make([]int64, 0)
				result = append(result, arr)
			}
			arr := result[j]
			arr = append(arr, bit)
			result[j] = arr
		}
	}
	return filterTableByRCV(result, currentIndex+1, columnLength)

}

func findOxygenGeneratorRating(table [][]int64, columnLength int) int64 {
	result := filterTableByMCV(table, 0, columnLength)
	var rate string
	for _, row := range result {
		for _, bit := range row {
			rate += strconv.Itoa(int(bit))
		}
	}
	oxygenRate, err := strconv.ParseInt(rate, 2, 64)
	if err != nil {
		panic(err)
	}

	return oxygenRate
}

func findC02ScrubberRating(table [][]int64, columnLength int) int64 {
	result := filterTableByRCV(table, 0, columnLength)
	var rate string
	for _, row := range result {
		for _, bit := range row {
			rate += strconv.Itoa(int(bit))
		}
	}
	CO2Rate, err := strconv.ParseInt(rate, 2, 64)
	if err != nil {
		panic(err)
	}

	return CO2Rate

}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	var table [][]int64
	var lineCounter, columnCounter int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		columnCounter = len(txt)
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
	o2 := findOxygenGeneratorRating(table, columnCounter)
	co2 := findC02ScrubberRating(table, columnCounter)
	fmt.Printf("O2 = %+v\n", o2)

	fmt.Printf("C02 = %+v\n", co2)
	fmt.Printf("result  = %+v\n", o2*co2)
}
