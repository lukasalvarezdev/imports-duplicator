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
		os.RemoveAll("./out")
		os.Remove(test[1] + ".ts")
	}
}

type test struct {
	srcPath    string
	importPath string
	outputPath string
}

func TestGetRealPath(t *testing.T) {
	tests := []test{
		{"./test/index.ts", "./src/basic", "./test/src/basic"},
		{"./test/dir1/index.ts", "../src/basic", "./test/src/basic"},
		{"./test/dir1/dir2/dir3/index.ts", "../../../src/basic", "./test/src/basic"},
		{"./test/dir1/dir2/dir3/index.ts", "../../../out/with.multiple.dots", "./test/out/with.multiple.dots"},
	}

	for _, test := range tests {
		realPath := getRealPath(test.srcPath, test.importPath)

		if realPath != test.outputPath {
			t.Errorf("Expected %s, got %s", test.outputPath, realPath)
		}
	}
}

func TestGetFileName(t *testing.T) {
	tests := []test{
		{"./test/some.file", "", "some.file.ts"},
		{`./test/dir1/some'file`, "", `some'file.ts`},
		{"./test/dir1/dir2/dir3/index", "", "index.ts"},
	}

	for _, test := range tests {
		fileName := getFileName(test.srcPath)

		if fileName != test.outputPath {
			t.Errorf("Expected %s, got %s", test.outputPath, fileName)
		}
	}
}
