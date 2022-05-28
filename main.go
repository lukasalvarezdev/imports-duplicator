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
	start := time.Now()
	srcPath, dstPath, fileToReplacePath := getCmdParams()
	paths := make([]string, 0)

	fmt.Println("Program started...")

	f, err := os.Open(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Printf("Scanning file %s...\n", srcPath)

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
		fileName := getFileName(path)

		wg.Add(1)
		var nBytes int64
		var copyErr error

		go copy(importPath, dstPath+"/"+fileName, make(chan int64), make(chan error))

		if copyErr != nil {
			log.Fatal(copyErr)
		}

		fmt.Printf("Copied %d bytes from %s to %s\n", nBytes, importPath, dstPath+fileName)
		paths = append(paths, path)
	}

	fmt.Printf("Updating %s imports in %s...\n", srcPath, fileToReplacePath)
	updatePaths(paths, srcPath, dstPath, fileToReplacePath)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Program finished.")

	elapsed := time.Since(start)
	log.Printf("Execution time was %s", elapsed)
	wg.Wait()
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
