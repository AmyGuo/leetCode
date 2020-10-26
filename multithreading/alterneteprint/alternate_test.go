package alterneteprint

import (
	"fmt"
	"sync"
	"testing"
)

//这里有两个线程被异步启动。其中一个调用 foo() 方法,输出"foo", 另一个调用 bar() 方法,输出"bar"， n=1时输出"foobar" ，n=2时输出"foobarfoobar"。

func AlternatePrint(n int) {

	start := make(chan struct{})
	fooFlag := make(chan struct{}, 1)
	barFlag := make(chan struct{}, 1)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		<-start
		for i := 0; i < n; i++ {
			fmt.Println("foo")
			barFlag <- struct{}{}
			<-fooFlag
		}
		<-barFlag
		wg.Done()
	}()

	go func() {
		<-barFlag
		for i := 0; i < n; i++ {
			fmt.Println("bar")
			fooFlag <- struct{}{}
			<-barFlag
		}
		wg.Done()
	}()

	start <- struct{}{}
	wg.Wait()
	fmt.Println("done")

}

func Test_Alternate_Print(t *testing.T) {
	//AlternatePrint(1)
	//AlternatePrint(2)
	//AlternatePrint(3)
}
