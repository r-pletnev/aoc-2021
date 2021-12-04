package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cur, prev, sum int64
	for scanner.Scan() {
		cur, err = strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			panic(err)
		}
		if prev > 0 && cur > prev {
			sum += 1
		}
		prev = cur
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)

}
