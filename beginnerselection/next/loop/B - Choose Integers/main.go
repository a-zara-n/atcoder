package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	arr := scanArrayInt()
	a, b, c := arr[0], arr[1], arr[2]
	
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
