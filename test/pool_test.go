package test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/qyqx233/gtool/lib/queue"
)

func TestChan2(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	select {
	case ch <- 2:
		t.Log("sended")
	default:
		t.Log("send failed")
		break
	}
	select {
	case ch <- 3:
		t.Log("sended")
	default:
		t.Log("send failed")
		break
	}
here:
	fmt.Print("so\n")
}

func countArray(arr []uint64) uint64 { return 0 }

func TestChan1(t *testing.T) {
	// ch := make(chan struct{}, 10)
	const n = 100
	const size = 10
	var counter uint64 = 0
	var arr = make([]uint64, size)
	f := func(i int) {
		var x uint64
		for {
			var ii = atomic.AddUint64(&counter, 1)
			if ii > n {
				break
			}
			x += ii
		}
		arr[i] = x
		t.Logf("%d go die\n", i)
	}
	for j := 0; j < size; j++ {
		go f(j)
	}
	// for x := 0; x < 10; x++ {
	// 	time.Sleep(time.Second)
	// }
	println(countArray(arr))
}

func TestSpmc(t *testing.T) {
	deque := queue.NewSpmcDequeue(64)
	deque.PushHead(1)
	deque.PushHead(2)
	t.Log(deque.PopTail())
	// for i := 0; i < 70; i++ {
	// 	if deque.PushHead(i) == false {
	// 		break
	// 	}
	// }
	// for i := 0; i < 10; i++ {
	// 	v, b := deque.PopTail()
	// 	noUse(v, b)
	// }

	// t.Log(deque)
}

func noUse(v ...interface{}) {}

func Benchmark(b *testing.B) {
	deque := queue.NewSpmcDequeue(64)
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			if deque.PushHead(i) == false {
				break
			}
		}
		for i := 0; i < 10; i++ {
			v, b := deque.PopTail()
			noUse(v, b)
			// t.Log(v.(int), b)
		}
	}
}
