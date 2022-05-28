package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := [][]string{
		{"./test/index.ts", "my-output"},
		{"./test/dir1/index.ts", "my-output"},
		{"./test/dir1/dir2/dir3/index.ts", "my-output"},
	}

	for _, test := range tests {
		runDuplicator(test[0], test[1])
		os.Remove(test[1])
		os.RemoveAll("./out")
	}
}
