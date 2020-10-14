// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filename := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filename)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filename)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("重复行的文件名称：")
			for k := range filename[line] {
				fmt.Printf("%s\t", k)
			}
			fmt.Printf("||%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filename[input.Text()] == nil {
			filename[input.Text()] = make(map[string]bool)
		}
		filename[input.Text()][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
