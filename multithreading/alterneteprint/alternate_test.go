package alterneteprint

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//这里有两个线程被异步启动。其中一个调用 foo() 方法,输出"foo", 另一个调用 bar() 方法,输出"bar"， n=1时输出"foobar" ，n=2时输出"foobarfoobar"。

func AlternatePrint1(n int) {

	start := make(chan struct{})
	flag := make(chan struct{})

	go func() {
		for i := 0; i < n; i++ {
			<-start
			fmt.Print("foo")
			flag <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			<-flag
			fmt.Print("bar")
		}
	}()

	close(start)
	select {
	case <-time.After(time.Second):
		break
	}
	fmt.Println("done")
}

func AlternatePrint2(n int) {
	start := make(chan struct{})
	flag := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 0; i < n; i++ {
			<-start
			fmt.Print("foo")
			flag <- struct{}{}
			time.Sleep(time.Millisecond * 5)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < n; i++ {
			<-flag
			fmt.Print("bar")
		}
		wg.Done()
	}()

	close(start)
	wg.Wait()
	fmt.Println("done")
}

func AlternatePrint3(n int) {
	start := make(chan int)
	x := make(chan int)
	y := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		<-start
		for i := 0; i < n; i++ {
			fmt.Print("foo")
			x <- 1
			<-y
		}
		x <- 1
		wg.Done()
	}()

	go func() {
		<-x
		for i := 0; i < n; i++ {
			fmt.Print("bar")
			y <- 1
			<-x
		}
		wg.Done()
	}()

	start <- 1

	wg.Wait()
	fmt.Println("done")
}

func Test_Alternate_Print1(t *testing.T) {
	//AlternatePrint1(1)
	AlternatePrint1(2)
	//AlternatePrint1(3)
}

func Test_Alternate_Print2(t *testing.T) {
	//AlternatePrint2(1)
	AlternatePrint2(2)
	//AlternatePrint2(3)
}

func Test_Alternate_Print3(t *testing.T) {
	//AlternatePrint3(1)
	AlternatePrint3(2)
	//AlternatePrint3(3)
}
