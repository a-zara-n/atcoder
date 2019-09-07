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
	var o int
	arr := scanArrayInt()
	n, a, b := arr[0], arr[1], arr[2]
	for i := 1; i <= n; i++ {
		t, d := 0, i
		for {
			t, d = t+d%10, d/10
			if d == 0 {
				break
			}
		}
		if a <= t && t <= b {
			o += i
		}
	}
	fmt.Println(o)
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
