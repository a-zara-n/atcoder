package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
グリッド系の定石
- インプットを先に行い、探索を行う方向をあらかじめ定義する。
- 探索の際に定義した方向での座標が問題なく 範囲内であることを確認してカウントする
*/
var stdin = bufio.NewScanner(os.Stdin)
var inc = map[bool]int{true: 1, false: 0}

func main() {
	//input動作
	arr := scanArrayInt()
	h, w := arr[0], arr[1]
	grids := [][]interface{}{}
	for i := 0; i < h; i++ {
		stdin.Scan()
		in := []interface{}{}
		for _, v := range strings.Split(stdin.Text(), "") {
			in = append(in, v)
		}
		grids = append(grids, in)
	}
	//座標の定義 {x,y}
	dxy := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	//探査 for は二次元座標のxy軸
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			//その座標が数値出なかった場合continue
			if grids[i][j] == "#" {
				continue
			}
			//周辺座標の探索
			c := 0
			for _, vs := range dxy {
				dx, dy := vs[0], vs[1]
				//周辺の座標から飛び出すことのないかを確認する
				if j+dx < 0 || j+dx > w-1 || i+dy < 0 || i+dy > h-1 {
					continue
				}
				//発見したら加算する
				if grids[i+dy][j+dx] == "#" {
					c++
				}
			}
			grids[i][j] = c
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print(grids[i][j])
		}
		fmt.Print("\n")
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
