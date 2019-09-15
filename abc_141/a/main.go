package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewScanner(os.Stdin)

func main() { XXX() }

func XXX() {
	stdin.Scan()
	w := []string{"Sunny", "Cloudy", "Rainy"}
	i := map[string]int{"Sunny": 1, "Cloudy": 2, "Rainy": 0}
	s := stdin.Text()
	fmt.Println(w[i[s]])

}
