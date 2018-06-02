package slaves

import (
	"testing"
)

func TestServe_SlavePool(t *testing.T) {
	ch := make(chan int, 20)
	sp := NewPool(func(obj interface{}) {
		ch <- obj.(int)
	})

	rounds := 100000

	go func() {
		for i := 0; i < rounds; i++ {
			sp.W <- i
		}
	}()

	i := 0
	for i < rounds {
		select {
		case <-ch:
			i++
		}
	}
	sp.Close()
}

func BenchmarkSlavePool(b *testing.B) {
	ch := make(chan int, b.N)

	sp := NewPool(func(obj interface{}) {
		ch <- obj.(int)
	})

	go func() {
		for i := 0; i < b.N; i++ {
			sp.W <- i
		}
	}()

	i := 0
	for i < b.N {
		select {
		case <-ch:
			i++
		}
	}
	close(ch)
	sp.Close()
}
