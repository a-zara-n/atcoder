package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	stdin.Scan()
	s := stdin.Text()
	//2のn乗通り組み合わせを全探索 nはsの桁数となる。
	//2のn乗は1をlen(s)-1を左シフトした値となる
	for bit := 0; bit < (1 << uint(len(s)-1)); bit++ {
		//sum はforループ時の計算結果を蓄積するための初期化
		sum, _ := strconv.Atoi(s[0:1])
		//表示の準備
		ans := s[0:1]
		for i := 0; i < len(s)-1; i++ {
			v, _ := strconv.Atoi(s[i+1 : i+2])
			if bit&(1<<uint(i)) != 0 {
				sum += v
				ans += "+"
			} else {
				sum -= v
				ans += "-"
			}
			ans += s[i+1 : i+2]
		}
		if sum == 7 {
			fmt.Println(ans + "=7")
			return
		}
	}
}
