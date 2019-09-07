package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	stdin.Scan()
	stdin.Scan()
	r := 0
	for _, istr := range strings.Split(stdin.Text(), " ") {
		r ^= func() int { i, _ := strconv.Atoi(istr); return i }()
	}
	fmt.Println(map[bool]string{true: "Yes", false: "NO"}[r == 0])
}
