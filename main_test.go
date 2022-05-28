package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	makeCopy("types.ts", "./files/out/", "new.ts")

	os.Remove("./new.ts")
	os.RemoveAll("./files/out/")

	// we need the folder out again
	os.Mkdir("./files/out/", os.ModePerm)
}
