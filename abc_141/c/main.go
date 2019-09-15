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
	n, k, q := arr[0], arr[1], arr[2]
	as := make([]int, q)
	for i := 0; i < q; i++ {
		as[i] = func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	}
	if k > q {
		for i := 0; i < n; i++ {
			fmt.Println("Yes")
		}
		return
	}
	points := map[int]int{}
	for i := 0; i < q; i++ {
		points[as[i]-1]++
	}
	for i := 0; i < n; i++ {
		if q-k < points[i] {
			fmt.Println("Yes")
			continue
		}
		fmt.Println("No")
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
