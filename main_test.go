package main

import (
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1+2", args{1, 2}, 3},
		{"2+3", args{2, 3}, 5},
		{"3+4", args{3, 4}, 7},
		{"4+5", args{4, 5}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	n := rand.Intn(10000)
	if simpleSum(n) != Sum(n) {
		t.Errorf("Sum(%d) = %d, want %d", n, Sum(n), simpleSum(n))
	}
}
func simpleSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}
