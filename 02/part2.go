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

	var h, d, a int64

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
			a += i
		case "u":
			a -= i
		case "f":
			h += i
			d += a * i
		}
	}
	fmt.Println(h * d)
}
