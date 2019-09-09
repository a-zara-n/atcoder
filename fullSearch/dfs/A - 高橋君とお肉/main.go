package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = bufio.NewScanner(os.Stdin)
var (
	n   int
	arr []int
)

func main() {
	n = func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = func() int { stdin.Scan(); i, _ := strconv.Atoi(stdin.Text()); return i }()
	}
	sum := dfs(0, 0, 0)
	fmt.Println(sum)
}
func dfs(i int, conroA, conroB int) int {
	if i == n {
		//全ての探索終了時に結果を返す
		max := map[bool]int{true: conroA, false: conroB}[conroA > conroB]
		return max
	}
	//コンロBへの加算を確認する枝
	sum1 := dfs(i+1, conroA, conroB+arr[i])
	//コンロAへの加算を確認する枝
	sum2 := dfs(i+1, conroA+arr[i], conroB)
	//最短時間を求めるので最短を返す
	if sum1 < sum2 {
		return sum1
	}
	return sum2
}
