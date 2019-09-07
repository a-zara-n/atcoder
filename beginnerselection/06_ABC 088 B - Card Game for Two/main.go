package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	n, a := func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }(), scanArrayInt()
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	m := map[int]int{0: 0, 1: 0}
	for i := 0; i < n; i++ {
		m[i%2] += a[i]
	}
	fmt.Println(m[0] - m[1])
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
