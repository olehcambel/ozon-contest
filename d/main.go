package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	var data string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}

	spl := strings.Split(data, " ")
	n := new(big.Int)
	n, _ = n.SetString(spl[0], 10)

	n2 := new(big.Int)
	n2, _ = n2.SetString(spl[1], 10)

	fmt.Print(n.Add(n, n2))
	// fmt.Print(big.NewInt(0).Add(n, n2))
}
