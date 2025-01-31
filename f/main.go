package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	inputPath  = "input.txt"
	outputPath = "output.txt"
)

func comparePair() error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var target int

	if scanner.Scan() {
		target, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
	}

	collect := make(map[int]int)
	var res bool

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		// no sense to go further if value is bigger than target (only positive)
		if x > target {
			continue
		}

		collect[x]++
		if count, ok := collect[target-x]; ok && (count > 1 || x != target-x) {
			res = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	err = ioutil.WriteFile(outputPath, []byte(b2s(res)), 0644)

	return err
}

func b2s(t bool) string {
	if t {
		return "1"
	}

	return "0"
}

func main() {
	err := comparePair()
	if err != nil {
		log.Fatal(err)
	}
}
