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
	var is_div = false
	var c int
	_, a := func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }(), scanArrayInt()
	for {
		for i, v := range a {
			a[i] = v / 2
			if v%2 != 0 {
				is_div = true
				break
			}
		}
		if is_div {
			fmt.Println(c)
			break
		}
		c++
	}
}

func scanArrayInt() []int {
	var ret = []int{}
	stdin.Scan()
	for _, s := range strings.Split(stdin.Text(), " ") {
		i, _ := strconv.Atoi(s)
		ret = append(ret, i)
	}
	return ret
}
