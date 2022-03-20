package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		Id     string
		Status string
	}{
		{"", "\n"},
		{".", ""},
		{"\t", "one\ttwo\tthree\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes"},
	}

	var prevId string
	for _, test := range tests {
		if test.Id != prevId {
			fmt.Printf("\n%s\n", test.Id)
			prevId = test.Id
		}
	}

	var prevStatus string
	for _, test := range tests {
		if test.Status != prevStatus {
			fmt.Printf("\n%s\n", test.Status)
			prevStatus = test.Status
		}
	}
}
