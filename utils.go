package main

import (
	"fmt"
	"os"
)

func isFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a)
}
