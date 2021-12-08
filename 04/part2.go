package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var rounds = make([]int, 0)
	var matrixs = make([][]int, 0)
	for scanner.Scan() {
		if len(rounds) == 0 {
			rs := strings.Split(scanner.Text(), ",")
			for _, r := range rs {
				i, err := strconv.ParseInt(string(r), 0, 64)
				if err != nil {
					panic(err)
				}
				rounds = append(rounds, int(i))
			}
			continue
		}
		row := scanner.Text()
		if len(row) == 0 {
			matrix := make([]int, 0)
			matrixs = append(matrixs, matrix)
			continue
		}
		row1 := strings.Split(row, " ")
		matrix := matrixs[len(matrixs)-1]
		for _, elm := range row1 {
			if elm == "" {
				continue
			}
			bit, err := strconv.ParseInt(string(elm), 0, 64)
			if err != nil {
				panic(err)
			}
			matrix = append(matrix, int(bit))
			matrixs[len(matrixs)-1] = matrix
		}
	}
	boards := make([]Board, len(matrixs))
	for i := range matrixs {
		boards[i] = NewBoard(matrixs[i], 5)
	}
	var finishedBoards = make(map[int]bool, 0)
	for j, r := range rounds {
		for i := range boards {
			if _, boardDone := finishedBoards[i]; boardDone {
				continue
			}

			boards[i].markCell(r)
			if j >= 5 {
				score := boards[i].score()
				if score > 0 {
					fmt.Printf("bingo(%v) = %+v, step= %v\n", i, score*r, r)
					finishedBoards[i] = true
				}
			}
		}
	}
}
