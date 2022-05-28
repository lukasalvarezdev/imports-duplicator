package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	makeCopy("./test/index.ts", "./test/out/", "./test/parsed.ts")

	os.Remove("./test/parsed.ts")
	os.RemoveAll("./test/out")

	// we need the folder out again
	os.Mkdir("./test/out", os.ModePerm)
}
