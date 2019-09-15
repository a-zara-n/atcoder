package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() {
	stdin.Scan()
	s := stdin.Text()
	flag := true
	k := map[rune]bool{'R': true, 'U': true, 'D': true}
	g := map[rune]bool{'L': true, 'U': true, 'D': true}
	q := map[int]map[rune]bool{1: g, 0: k}
	for i := 0; i < len(s); i++ {
		if !q[i%2][rune(s[i])] {
			flag = false
		}
	}
	fmt.Println(map[bool]string{true: "Yes", false: "No"}[flag])
}
