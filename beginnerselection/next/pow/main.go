package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	n := func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	fmt.Println(factorial(n))
}
func permutation(n int, k int) int {
	v := 1
	w := int(math.Pow(10, 9)) + 7
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v = v * (n - i) % w
		}
	} else if k > n {
		v = 0
	}
	return v
}
func factorial(n int) int {
	return permutation(n, n-1)
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
