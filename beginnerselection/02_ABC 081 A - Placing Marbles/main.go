package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	stdin.Scan()
	s := stdin.Text()
	var out int
	for _, i := range s {
		if i == '1' {
			out++
		}
	}
	fmt.Println(out)
}
