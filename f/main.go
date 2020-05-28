package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// FIXME: NOT WORKING
func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var target int
	// var seq []int

	scaner := bufio.NewScanner(file)
	if scaner.Scan() {
		target, _ = strconv.Atoi(scaner.Text())
	}

	reader := bufio.NewReader(file)
	buf := make([]byte, 0, 4*1024)

	for {
		// TODO: read from len(buf:cap)
		// and then append
		// check go doc. go tour
		n, err := reader.Read(buf[:cap(buf)])
		buf = append(buf, buf[:n])
		// buf = buf[:n]
		if n == 0 {
			// if err == nil {
			//   continue
			// }
			if err == io.EOF {
				break
			}
		}
	}

	seq, _ := sliceAtoi(strings.Split(string(buf), " "))

	// if scaner.Scan() {
	// 	seq, _ = sliceAtoi(strings.Split(scaner.Text(), " "))
	// }

	res := checkSequence(target, seq)
	ioutil.WriteFile("output.txt", []byte(b2s(res)), 0644)
}

func checkSequence(target int, seq []int) bool {
	y := 0
	idx := 0

	for {
		if idx >= len(seq) {
			idx = 0
			y++
		}

		if y >= len(seq) {
			return false
		}

		a := seq[y]
		b := seq[idx]

		if a+b == target {
			return true
		}

		idx++
	}
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func b2s(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
