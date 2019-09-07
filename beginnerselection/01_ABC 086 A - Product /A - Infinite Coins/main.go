package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	sum := func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	ichien := func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	sum = sum % 500
	var o = "No"
	if sum <= ichien {
		o = "Yes"
	}
	fmt.Println(o)
}
