package main

import (
	"io"
	"log"
	"os"
)

func copy(src string, dst string, nBytesChan chan int64, copyErrorChan chan error) {
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

func createDir(dir string) {
	//Create a folder/directory at a full qualified path
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
