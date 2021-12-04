package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	var h, d int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		desc := scanner.Text()
		var firstLetter = string(desc[0])
		var i, err = strconv.ParseInt(string(desc[len(desc)-1]), 0, 64)
		if err != nil {
			panic(err)
		}
		switch firstLetter {
		case "d":
			d += i
		case "u":
			d -= i
		case "f":
			h += i
		}
	}
	fmt.Println(h * d)
}
