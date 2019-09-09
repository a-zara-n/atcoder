package main

import "testing"

var (
	testN   = []int{125, 9999999999}
	lenN    = []int{3, 10}
	testAns = []int{176, 12656242944}
)

func TestDFS(t *testing.T) {
	t.Log("[*] これからテストを開始します")
	for i := 0; i < 2; i++ {
		sum := dfs(pow(10, lenN[i]-1), 0, testN[i])
		if sum == testAns[i] {
			t.Logf("[*] %v回目の回答はACです", i)
		} else {
			t.Errorf("[*] %v回目の回答はWAです", i)
			t.Errorf("[*] 回答は sum = %v", sum)
		}
	}
}
