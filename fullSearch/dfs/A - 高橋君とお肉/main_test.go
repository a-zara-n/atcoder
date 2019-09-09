package main

import "testing"

var (
	testN   = []int{4, 3, 1}
	testA   = [][]int{{4, 6, 7, 10}, {1, 2, 4}, {29}}
	testAns = []int{14, 4, 29}
)

func TestDFS(t *testing.T) {
	t.Log("[*] これからテストを開始します")
	for i := 0; i < 3; i++ {
		n = testN[i]
		arr = testA[i]
		sum := dfs(0, 0, 0)
		if sum == testAns[i] {
			t.Logf("[*] %v回目の回答はACです", i)
		} else {
			t.Errorf("[*] %v回目の回答はWAです", i)
			t.Errorf("[*] 回答は sum = %v", sum)
		}
	}
}
