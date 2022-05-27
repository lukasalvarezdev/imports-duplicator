package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 1. read cmd params: <source_path> <destination_path> <file_to_replace_path>
// 1. read file params[0]
// 2. split string, only imports, with "" or ''
// 3. get path
// 4. copy files to ./files/
// 5. replace main file to import the types from ./files/

func main() {
	srcPath, dstPath, fileToReplacePath := getCmdParams()

	f, err := os.Open(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		txt := scanner.Text()
		split := strings.Split(txt, "'")
		path := split[1]

		if path == "" {
			fmt.Println(fileToReplacePath)
			continue
		}

		tspath := path + ".ts"

		copy(tspath, dstPath+tspath)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCmdParams() (string, string, string) {
	if len(os.Args) != 4 {
		log.Fatal("Usage: go run . <source_path> <destination_path> <file_to_replace_path>")
	}
	return os.Args[1], os.Args[2], os.Args[3]
}
