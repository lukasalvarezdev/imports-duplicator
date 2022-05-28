package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	srcPath, fileToReplacePath := getCmdParams()
	runDuplicator(srcPath, fileToReplacePath)
}

func runDuplicator(srcPath string, outPutFileName string) {
	start := time.Now()
	paths := make([]string, 0)
	fmt.Println("Program started...")

	f, err := os.Open(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		txt := scanner.Text()
		// "from" to make sure that there are no multiline import issues
		res, err := regexp.Compile("from")

		if matches := res.MatchString(txt); err != nil || !matches {
			continue
		}

		// splits the line by the "from" keyword, and removes it
		split := res.Split(txt, -1)
		path, err := getPath(split)

		if err != nil {
			continue
		}

		importPath := getImportPathWithExtension(path)
		realPath := getRealPath(srcPath, importPath)
		fileName := getFileName(path)
		nBytesChnl := make(chan int64)
		copyErrChnl := make(chan error)

		wg.Add(1)
		go copy(realPath, "./out/"+fileName, nBytesChnl, copyErrChnl)

		nBytes, copyErr := <-nBytesChnl, <-copyErrChnl

		if copyErr != nil {
			log.Fatal(copyErrChnl, " Error copying file")
		}

		// remove file extension from import path
		importPath = strings.Replace(importPath, ".ts", "", -1)
		paths = append(paths, importPath)
		fmt.Printf("Copied %v bytes from %s to %s\n", nBytes, realPath, "./out/"+fileName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Error while reading file")
	}

	fmt.Printf("Updating %s imports in %s...\n", srcPath, outPutFileName)
	updatePaths(paths, srcPath, outPutFileName)

	elapsed := time.Since(start)
	log.Printf("Execution time was %s", elapsed)

	wg.Wait()
}

func getCmdParams() (string, string) {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run . <source_path> <output_file_name>")
	}

	if strings.Contains(os.Args[2], "/") {
		log.Fatal("<output_file_name> must be a file name, not a path")
	}

	return os.Args[1], os.Args[2]
}
