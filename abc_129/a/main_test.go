package main

import (
	"strings"
	"testing"
)

func TestRunSpace(t *testing.T) {
	indata := []string{
		"1 3 4", "3 2 3",
	}
	outdata := []string{
		"4", "5",
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

/*
/**/

/*
func TestRunLine(t *testing.T) {

}

/**/
