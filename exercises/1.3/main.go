package main

import (
	"strings"
)

var args []string

func init() {
	for i := 0; i < 2000; i++ {
		args = append(args, "明")
	}
}
func main() {
	join()
	plus()
}
func join() {
	// var start = time.Now()
	strings.Join(args, " ")
	// secs := time.Since(start).Nanoseconds()
	// fmt.Println("join的耗时", secs)
}

func plus() {
	// var start = time.Now()
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	// secs := time.Since(start).Nanoseconds()
	// fmt.Println("加号的耗时", secs)
}
