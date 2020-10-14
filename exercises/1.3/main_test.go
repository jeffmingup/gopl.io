package main

import "testing"

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join()
	}
}
func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plus()
	}
}
