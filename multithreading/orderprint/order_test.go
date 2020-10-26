package orderprint

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type operation func(string)

//func withPrint2() operation {
//	return func(s string) {
//		fmt.Println(s)
//	}
//}

func withPrint(s string) {
	fmt.Println(s)
}

func first(f chan struct{}) {
	fmt.Println("first print")
	f <- struct{}{}
}

func second(f chan struct{}) {
	fmt.Println("second print")
	f <- struct{}{}
}

func third(f chan struct{}) {
	fmt.Println("third print")
	f <- struct{}{}
}

func OrderPrint1() {
	flag := make(chan struct{})
	count := 3
	go first(flag)
	go second(flag)
	go third(flag)

	for range flag {
		count--
		if count == 0 {
			close(flag)
		}
	}
	fmt.Println("main print")
}

func generate(cptr *uint32, total int, op operation, ch chan struct{}) map[int]func(uint32) {
	var m = make(map[int]func(order uint32), total)
	for i := 1; i <= total; i++ {
		m[i] = func(order uint32) {
			for {
				if !atomic.CompareAndSwapUint32(cptr, order, order) {
					time.Sleep(500 * time.Millisecond)
				} else {
					break
				}
			}
			op(fmt.Sprintf("this is goroutine %d", order))
			ch <- struct{}{}
			atomic.AddUint32(cptr, 1)
		}
	}
	return m
}

func OrderPrint2() {
	var count uint32 = 1 //协程开始的编号
	total := 5
	ch := make(chan struct{}) //信号量，保证协程在主程之前完成
	m := generate(&count, total, withPrint, ch)
	//m := generate(&count, total, withPrint2(), ch)

	for index, fn := range m {
		go fn(uint32(index))
	}

	for range ch {
		total--
		if total == 0 {
			close(ch)
		}
	}
}

func OrderPrint3() {
	a := make(chan struct{})
	b := make(chan struct{})
	c := make(chan struct{})

	go A(a, b)
	go B(b, c)
	go C(c)

	close(a)
	select {
	case <-time.After(time.Second):
		break
	}
	fmt.Println("main print")
}

func A(a, b chan struct{}) {
	<-a
	fmt.Println("this is goroutine A")
	close(b)
}
func B(b, c chan struct{}) {
	<-b
	fmt.Println("this is goroutine B")
	close(c)
}

func C(c chan struct{}) {
	<-c
	fmt.Println("this is goroutine C")
}

//保证协程在主程之前完成
func Test_Order_Print1(t *testing.T) {
	OrderPrint1()
}

//保证协程串行运行
func Test_order_Print2(t *testing.T) {
	OrderPrint2()
}

//保证协程串行运行
func Test_order_Print3(t *testing.T) {
	OrderPrint3()
}
