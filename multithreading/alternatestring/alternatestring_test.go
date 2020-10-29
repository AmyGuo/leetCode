package alternatestring

import (
	"fmt"
	"sync"
	"testing"
)

//编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：
//
//如果这个数字可以被 3 整除，输出 "fizz"。
//如果这个数字可以被 5 整除，输出 "buzz"。
//如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。
//例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。
//
//假设有这么一个类：
//
//
//线程A将调用 fizz() 来判断是否能被 3 整除，如果可以，则输出 fizz。
//线程B将调用 buzz() 来判断是否能被 5 整除，如果可以，则输出 buzz。
//线程C将调用 fizzbuzz() 来判断是否同时能被 3 和 5 整除，如果可以，则输出 fizzbuzz。
//线程D将调用 number() 来实现输出既不能被 3 整除也不能被 5 整除的数字。
//

type AS struct {
	start chan struct{}
	f     chan struct{}
	b     chan struct{}
	fb    chan struct{}
	n     chan struct{}
	wg    *sync.WaitGroup
	s     []interface{}
	num   int
}

func (a *AS) fizz() {
	for i := 3; i <= a.num; i += 3 {
		if i%15 == 0 {
			continue
		}
		<-a.f
		a.s = append(a.s, "fizz")
		a.n <- struct{}{}
	}
	a.wg.Done()
}

func (a *AS) buzz() {
	for i := 5; i <= a.num; i += 5 {
		if i%15 == 0 {
			continue
		}
		<-a.b
		a.s = append(a.s, "buzz")
		a.n <- struct{}{}
	}
	a.wg.Done()
}

func (a *AS) fizzbuzz() {
	for i := 15; i <= a.num; i += 15 {
		<-a.fb
		a.s = append(a.s, "fizzbuzz")
		a.n <- struct{}{}
	}
	a.wg.Done()
}

func (a *AS) number() {
	<-a.start
	for i := 1; i <= a.num; i++ {
		if i%3 != 0 && i%5 != 0 {
			a.s = append(a.s, i)
		} else {
			if i%15 == 0 {
				a.fb <- struct{}{}
				<-a.n
			} else if i%5 == 0 {
				a.b <- struct{}{}
				<-a.n
			} else {
				a.f <- struct{}{}
				<-a.n
			}
		}
	}
	a.wg.Done()
}

func AlternateString(n int) {
	as := AS{start: make(chan struct{}), f: make(chan struct{}), b: make(chan struct{}), fb: make(chan struct{}, 1), n: make(chan struct{}), wg: &sync.WaitGroup{}, s: make([]interface{}, 0, n), num: n}
	as.wg.Add(4)
	go as.fizz()
	go as.buzz()
	go as.fizzbuzz()
	go as.number()
	close(as.start)
	as.wg.Wait()

	ret := ""
	for i, v := range as.s {
		if i < n-1 {
			ret = ret + fmt.Sprintf("%v,", v)
		}
	}
	ret = ret + fmt.Sprintf("%v", as.s[n-1])
	fmt.Println(ret)
	fmt.Println("done")
}

func Test_Alternate_String(t *testing.T) {
	AlternateString(5)
	AlternateString(15)
}
