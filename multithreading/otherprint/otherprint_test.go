package otherprint

import (
	"fmt"
	"sync"
	"testing"
)

//输入：n = 2
//输出："0102"
//说明：三条线程异步执行，其中一个调用 zero()，另一个线程调用 even()，最后一个线程调用odd()。正确的输出为 "0102"。
//
//输入：n = 5
//输出："0102030405

func OtherPrint(n int) {
	zero := make(chan struct{}, 1)
	even := make(chan struct{})
	odd := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		for i := 1; i <= n; i++ {
			<-zero
			fmt.Print(0)
			if i%2 == 0 {
				even <- struct{}{}
			}
			if i%2 != 0 {
				odd <- struct{}{}
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= n; i += 2 {
			<-odd
			fmt.Print(i)
			zero <- struct{}{}
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i <= n; i += 2 {
			<-even
			fmt.Print(i)
			zero <- struct{}{}
		}
		wg.Done()
	}()

	zero <- struct{}{}
	wg.Wait()
	fmt.Println("done")
}

func Test_Other_Print(t *testing.T) {
	OtherPrint(2)
	OtherPrint(5)
}
