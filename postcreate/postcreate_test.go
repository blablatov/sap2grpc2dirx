package postcreate

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		ColumnOne string
	}{
		{""},
		{"\n"},
		{"\t"},
		{"Data for test"},
		{"Yes, no"},
		{"true, or, false"},
	}

	var prevColumnOne string
	for _, test := range tests {
		if test.ColumnOne != prevColumnOne {
			fmt.Printf("\n%s\n", test.ColumnOne)
			prevColumnOne = test.ColumnOne
		}
	}
}
