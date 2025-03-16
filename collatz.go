package main

import (
	"fmt"
	"time"
)

func collatzIterativeTraditional(n uint64) {
	for n > 1 {
		if n%2 == 1 {
			n = 3*n + 1
		} else {
			n /= 2
		}
	}
}

func collatzIterativeBitwise(n uint64) {
	for n > 1 {
		if (n & 1) == 1 {
			n = ((n << 1) | 1) + n
		} else {
			n >>= 1
		}
	}
}

func collatzRecursiveTraditional(n uint64) {
	if n <= 1 {
		return
	}
	if n%2 == 1 {
		collatzRecursiveTraditional(3*n + 1)
	} else {
		collatzRecursiveTraditional(n / 2)
	}
}

func collatzRecursiveBitwise(n uint64) {
	if n <= 1 {
		return
	}
	if (n & 1) == 1 {
		collatzRecursiveBitwise(((n << 1) | 1) + n)
	} else {
		collatzRecursiveBitwise(n >> 1)
	}
}

func timeFunc(f func(uint64), n uint64) time.Duration {
	start := time.Now()
	for i := uint64(1); i <= n; i++ {
		f(i)
	}
	return time.Since(start)
}

func main() {
	const N = 10_000_000
	recursiveTraditional := timeFunc(collatzRecursiveTraditional, N)
	recursiveBitwise := timeFunc(collatzRecursiveBitwise, N)
	iterativeTraditional := timeFunc(collatzIterativeTraditional, N)
	iterativeBitwise := timeFunc(collatzIterativeBitwise, N)
	fmt.Println("Traditional (recursive): ", recursiveTraditional)
	fmt.Println("Bitwise (recursive): ", recursiveBitwise)
	fmt.Println("Traditional (iterative): ", iterativeTraditional)
	fmt.Println("Bitwise (iterative): ", iterativeBitwise)
}
