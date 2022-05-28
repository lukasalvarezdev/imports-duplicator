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
	srcPath, dstPath, fileToReplacePath := getCmdParams()
	makeCopy(srcPath, dstPath, fileToReplacePath)
}

func makeCopy(srcPath string, dstPath string, fileToReplacePath string) {
	start := time.Now()
	paths := make([]string, 0)
	fmt.Println("Program started...")
	dstPath = fixDstPath(dstPath)

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
		go copy(realPath, dstPath+fileName, nBytesChnl, copyErrChnl)

		nBytes, copyErr := <-nBytesChnl, <-copyErrChnl

		if copyErr != nil {
			log.Fatal(copyErrChnl, " Error copying file")
		}

		paths = append(paths, realPath)
		fmt.Printf("Copied %v bytes from %s to %s\n", nBytes, realPath, dstPath+fileName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err, "Error while reading file")
	}

	// fmt.Printf("Updating %s imports in %s...\n", srcPath, fileToReplacePath)
	// updatePaths(paths, srcPath, dstPath, fileToReplacePath)

	elapsed := time.Since(start)
	log.Printf("Execution time was %s", elapsed)

	wg.Wait()
}

func fixDstPath(dstPath string) string {
	if !strings.HasSuffix(dstPath, "/") {
		dstPath = dstPath + "/"
	}

	return dstPath
}

func getFileName(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1] + ".ts"
}

func getImportPathWithExtension(path string) string {
	return path + ".ts"
}

func getPath(split []string) (string, error) {
	path := split[len(split)-1]
	path = strings.Replace(path, " ", "", -1)

	if path == "" {
		return "", fmt.Errorf("path is empty")
	}

	if strings.Contains(path, ";") {
		path = path[:len(path)-1]
	}

	// remove last ' or "" from path
	path = path[:len(path)-1]

	// remove first ' or "" from path
	path = path[1:]

	return path, nil
}

func getCmdParams() (string, string, string) {
	if len(os.Args) != 4 {
		log.Fatal("Usage: go run . <source_path> <destination_path> <file_to_replace_path>")
	}
	return os.Args[1], os.Args[2], os.Args[3]
}

func getRealPath(src, importPath string) (realPath string) {
	split := strings.Split(importPath, "/")
	fileSplit := strings.Split(src, "/")

	for _, path := range split {
		if path == ".." {
			last := fileSplit[len(fileSplit)-1]

			if strings.Contains(last, ".ts") {
				fileSplit = fileSplit[:len(fileSplit)-2]
				continue
			}

			fileSplit = fileSplit[:len(fileSplit)-1]
			continue
		}

		if path == "." {
			last := fileSplit[len(fileSplit)-1]

			if strings.Contains(last, ".ts") {
				fileSplit = fileSplit[:len(fileSplit)-1]
				continue
			}

			continue
		}

		fileSplit = append(fileSplit, path)
	}

	return strings.Join(fileSplit, "/")
}
