package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	indata := []string{
		"1 3", "0 1", "32 21",
	}
	outdata := []string{
		"3", "0", "58",
	}
	for i, input := range indata {
		in := strings.Split(input, " ")
		ret := Run(in)
		out := outdata[i]
		if ret != out {
			t.Errorf("\nテストデータ \"%v\"が異なります。\nTestCase 	: %v\n正当なデータ  : %v\nReturn 		: %v", i, input, out, ret)
		}
	}
}
