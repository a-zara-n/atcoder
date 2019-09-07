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
	arr := scanArrayInt()
	a, b, c, d := arr[0], arr[1], arr[2], arr[3]
	if b < c || d < a {
		fmt.Println("0")
		return
	}
	start, end :=
		map[bool]int{true: a, false: c}[a > c],
		map[bool]int{true: b, false: d}[d > b]
	fmt.Println(end - start)
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
