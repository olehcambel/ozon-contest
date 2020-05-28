package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	set := make(map[int]int)
	file, _ := os.Open("input-201.txt")
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		n, _ := strconv.Atoi(scaner.Text())

		set[n]++
	}

	for n, count := range set {
		if count == 1 {
			fmt.Println(n)
			return
		}
	}

	fmt.Print(set)
}
