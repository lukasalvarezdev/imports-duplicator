package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

func updatePaths(paths []string, srcPath string, dstPath string, fileToReplacePath string) {
	input, err := ioutil.ReadFile(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		fileName := getFileName(path)

		// remove file extension
		spl := strings.Split(fileName, ".")
		spl = spl[:len(spl)-1]
		fileName = strings.Join(spl, ".")

		replaceFromSingleQuote := []byte("from " + "'" + path + "'")
		replaceFromDoubleQuote := []byte("from " + `"` + path + `"`)
		replaceToSingleQuote := []byte("from " + "'" + dstPath + fileName + "'")
		replaceToDoubleQuote := []byte("from " + `"` + dstPath + fileName + `"`)

		input = bytes.Replace(input, replaceFromSingleQuote, replaceToSingleQuote, -1)
		input = bytes.Replace(input, replaceFromDoubleQuote, replaceToDoubleQuote, -1)
	}

	if err = ioutil.WriteFile(fileToReplacePath, input, 0666); err != nil {
		log.Fatal(err)
	}
}
