package utils

import (
	"errors"
	"testing"
)

func TestFileUtil_ContentToList(t *testing.T) {
	cases := []struct {
		testName       string
		ioUtilReadFile ReadFile
		expected       int
	}{
		{
			testName: "1. Positive case",
			ioUtilReadFile: func(filename string) ([]byte, error) {
				return []byte("900 1000\n1000 1100"), nil
			},
			expected: 2,
		},
		{
			testName: "2. Positive case",
			ioUtilReadFile: func(filename string) ([]byte, error) {
				return []byte(""), errors.New("error read file")
			},
			expected: 0,
		},
	}

	for _, c := range cases {
		t.Logf("Currently testing: %s\n", c.testName)
		util := NewFileUtil(c.ioUtilReadFile)
		result, _ := util.ContentToList("list.txt")
		if c.expected != len(result) {
			t.Errorf("Expected: %v, Got: %v\n", c.expected, len(result))
		}
	}
}
