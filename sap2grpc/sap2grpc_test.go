package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		Id          string
		Status      string
		Name        string
		InterfaceId string
	}{
		{"", "\n", "\t", ","},
		{".", "", "null", ":"},
		{"\t", "one\ttwo\tthree\n", "\n\n", "NaN\null\n\n"},
		{"Data for test", "Number 99999 to data test", "Yes test", "Number 777777777777"},
		{"Yes, no", "No, or, yes", "Yes, no, _", "No, or, yes, _, ops"},
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

	var prevName string
	for _, test := range tests {
		if test.Name != prevName {
			fmt.Printf("\n%s\n", test.Name)
			prevName = test.Name
		}
	}

	var prevInterfaceId string
	for _, test := range tests {
		if test.InterfaceId != prevInterfaceId {
			fmt.Printf("\n%s\n", test.InterfaceId)
			prevName = test.InterfaceId
		}
	}
}
