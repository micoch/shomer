package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("your code watchman")
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "There is less than two indecies")
		os.Exit(1)
	}
	dir := os.Args[1]
	path := filepath.Join(dir, "package-lock.json")
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is a problem with the path", err)
		os.Exit(1)
	}
	fmt.Println(len(content))
}
