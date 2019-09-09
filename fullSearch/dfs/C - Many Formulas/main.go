package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = bufio.NewScanner(os.Stdin)
var l int

func main() {
	stdin.Scan()
	t := stdin.Text()
	l = len(t)
	s := func(t string) int { i, _ := strconv.Atoi(t); return i }(t)
	sum := dfs(pow(10, l-1), 0, s)
	fmt.Println(sum)
}

var memo = map[int]bool{}

func dfs(div, sum, s int) int {
	if div == 1 {
		return sum + s
	}
	sum1 := dfs(div/10, sum+s/div, s%div)
	sum2 := dfs(div/10, sum, s)
	return sum1 + sum2
}

func pow(num, n int) int {
	if n == 0 {
		return 1
	}
	tmp := num
	for i := 1; i < n; i++ {
		num *= tmp
	}
	return num
}
