package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	n := func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	as, flag := make([]int, n), false
	for i := 0; i < n; i++ {
		as[i] = func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
		if as[i] == 2 {
			flag = true
		}
	}
	if !flag {
		fmt.Println(-1)
		return
	}
	route, count, index := map[int]bool{}, 0, 0
	for {
		if route[index] || count > n {
			fmt.Println(-1)
			break
		} else if index == 1 {
			fmt.Println(count)
			break
		}
		route[index] = true
		index = as[index] - 1
		count++
	}
}
