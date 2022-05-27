package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// getCmdParams()

	f, err := os.Open("types.ts")

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
			continue
		}

		tspath := path + ".ts"

		copy(tspath, "./files/"+tspath)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// func getCmdParams() (string, string) {
// 	if len(os.Args) != 3 {
// 		log.Fatal("Usage: go run parser.go <input_file> <output_file>")
// 	}
// 	return os.Args[1], os.Args[2]
// }
