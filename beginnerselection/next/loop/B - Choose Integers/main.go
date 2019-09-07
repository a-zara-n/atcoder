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
	a, b, c := arr[0], arr[1], arr[2]
	for i := 1; i <= b; i++ {
		d := a * i % b
		if d == c {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
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
