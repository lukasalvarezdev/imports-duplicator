package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func copy(src string, dst string, nBytesChan chan int64, copyErrorChan chan error) {
	createDirIfNotExists(dst)

	sourceFileStat, srcErr := os.Stat(src)
	if srcErr != nil {
		log.Fatal(srcErr)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Fatalf("%s is not a regular file", src)
	}

	source, openErr := os.Open(src)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer source.Close()

	destination, createErr := os.Create(dst)
	if createErr != nil {
		log.Fatal(createErr)
	}
	defer destination.Close()
	nBytes, copyErr := io.Copy(destination, source)

	wg.Done()
	nBytesChan <- nBytes
	copyErrorChan <- copyErr
}

func createDirIfNotExists(path string) {
	file := getFileName(path)
	// remove .ts from fileName
	file = file[:len(file)-3]
	folder := strings.Replace(path, file, "", -1)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.ModePerm)
	}
}
