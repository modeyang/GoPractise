package algorithms

import "fmt"
import "github.com/willf/bitset"

func generate(ch chan<- int) {
	for i:=2; ;i++ {
		ch <- i
	}
}

func filter(src, dest chan int, prime int) {
	for i := range src {
		if i % prime != 0 {
			dest <- i
		}
	}
}

func Prime(n int) {
	ch := make(chan int)
	go generate(ch)
	for i:=2; i<=n; i++ {
		prime := <- ch
		fmt.Println(prime)
		out := make(chan int)
		go filter(ch, out, prime)
		out = ch
	}
}

func sieveOfEratosthenes(N int) (primes []int) {
	// use bitset store prime bool flag
	b := bitset.New(uint(N))
	for i := 2; i < N; i++ {
		if b.Test(uint(i)) {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b.Set(uint(k))
		}
	}
	return
}

func Prime2(N int) {
	primes := sieveOfEratosthenes(N)
	fmt.Println(primes)
}

