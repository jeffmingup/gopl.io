// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var algo string
	flag.StringVar(&algo, "algo", "256", "加密算法")
	flag.Parse()
	var f func(data []byte) interface{}
	switch algo {
	case "384":
		fmt.Println("使用sha512.Sum384算法")
		f = func(data []byte) interface{} {
			return sha512.Sum384(data)
		}
	case "512":
		fmt.Println("使用sha512.Sum512算法")
		f = func(data []byte) interface{} {
			return sha512.Sum512(data)
		}
	default:
		fmt.Println("使用sha256.Sum256算法")
		f = func(data []byte) interface{} {
			return sha256.Sum256(data)
		}
	}
	c1 := f([]byte("x"))
	c2 := f([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
