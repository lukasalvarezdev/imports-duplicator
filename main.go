package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	srcPath, dstPath := getCmdParams()

	f, err := os.Open(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		txt := scanner.Text()
		res, err := regexp.Compile("from")

		if matches := res.MatchString(txt); err != nil || !matches {
			continue
		}

		split := res.Split(txt, -1)
		path, err := getPath(split)

		if err != nil {
			continue
		}

		srcPath := getSrc(path)
		fileName := getFileName(path)

		copy(srcPath, dstPath+fileName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getFileName(path string) string {
	return strings.Split(path, "/")[len(strings.Split(path, "/"))-1] + ".ts"
}

func getSrc(path string) string {
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

func getCmdParams() (string, string) {
	// if len(os.Args) != 4 {
	// 	log.Fatal("Usage: go run . <source_path> <destination_path> <file_to_replace_path>")
	// }
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run . <source_path> <destination_path>")
	}
	return os.Args[1], os.Args[2]
}
