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
	n, k := arr[0], arr[1]
	stdin.Scan()
	s := []rune(stdin.Text())
	for i := 0; i < k; i++ {
		tmp := s
		for l := 0; l < n; l++ {
			for r := l; r < n; r++ {
				for j := 0; j < l-r; j++ {
					if tmp[l+j] == 'L' {
						tmp[l+j] = 'R'
					} else {
						tmp[l+j] = 'L'
					}
					tmp, b, c, d := tmp[:l+j], tmp[l+j], tmp[l+j+1:r-j], tmp[r-j:]
					tmp = append(tmp, c...)
					tmp = append(tmp, b)
					tmp = append(tmp, d...)
				}
			}
		}
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
