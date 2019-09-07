package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	stdin = bufio.NewScanner(os.Stdin)
	dxy   = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //{y,x}
)

func main() {
	arr := scanArrayInt()
	h, w := arr[0], arr[1]
	grids := [][]rune{}
	for i := 0; i < h; i++ {
		stdin.Scan()
		in := []rune{}
		in = append(in, []rune(stdin.Text())...)
		grids = append(grids, in)
	}

	var flag = true
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grids[y][x] == '.' {
				continue
			}
			var flag2 = false
			for _, v := range dxy {
				dx, dy := v[0], v[1]
				if 0 > x+dx || x+dx > h-1 || 0 > y+dy || y+dy > w-1 {
					continue
				}
				if grids[y+dy][x+dx] == '#' {
					flag2 = true
				}
			}
			//fmt.Println(flag2, y, x)
			if !flag2 {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
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
