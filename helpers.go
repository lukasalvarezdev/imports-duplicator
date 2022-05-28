package main

import (
	"fmt"
	"strings"
)

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
