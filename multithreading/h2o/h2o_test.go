package h2o

import (
	"fmt"
	"sync"
	"testing"
)

//现在有两种线程，氧 oxygen 和氢 hydrogen，你的目标是组织这两种线程来产生水分子。
//
//存在一个屏障（barrier）使得每个线程必须等候直到一个完整水分子能够被产生出来。
//
//氢和氧线程会被分别给予 releaseHydrogen 和 releaseOxygen 方法来允许它们突破屏障。
//
//这些线程应该三三成组突破屏障并能立即组合产生一个水分子。
//
//你必须保证产生一个水分子所需线程的结合必须发生在下一个水分子产生之前。
//
//换句话说:
//
//如果一个氧线程到达屏障时没有氢线程到达，它必须等候直到两个氢线程到达。
//如果一个氢线程到达屏障时没有其它线程到达，它必须等候直到一个氧线程和另一个氢线程到达。
//书写满足这些限制条件的氢、氧线程同步代码。
//
//
//
//示例 1:
//
//输入: "HOH"
//输出: "HHO"
//解释: "HOH" 和 "OHH" 依然都是有效解。
//示例 2:
//
//输入: "OOHHHH"
//输出: "HHOHHO"
//解释: "HOHHHO", "OHHHHO", "HHOHOH", "HOHHOH", "OHHHOH", "HHOOHH", "HOHOHH" 和 "OHHOHH" 依然都是有效解。
//
//
//提示：
//
//输入字符串的总长将会是 3n, 1 ≤ n ≤ 50；
//输入字符串中的 “H” 总数将会是 2n 。
//输入字符串中的 “O” 总数将会是 n 。

type H2O struct {
	wg    *sync.WaitGroup
	hFlag chan struct{}
	oFlag chan struct{}
	hNum  int
	oNum  int
}

//产生H
func (h *H2O) hydrogen() {
	for h.hNum > 0 {
		fmt.Print("H")
		h.hNum--
		if h.hNum%2 == 0 {
			<-h.oFlag
			<-h.hFlag
		}
	}
	h.wg.Done()
}

//产生O
func (h *H2O) oxygen() {
	for h.oNum > 0 {
		h.oFlag <- struct{}{}
		fmt.Print("O")
		h.oNum--
		h.hFlag <- struct{}{}
	}

	h.wg.Done()
}

//释放一个水分子
func (h *H2O) h2oFree() {
	h.wg.Add(2)
	go h.hydrogen()
	go h.oxygen()
	h.wg.Wait()
}

func newInstance(str string) *H2O {
	h := new(H2O)
	for _, v := range []byte(str) {
		if string(v) == "H" {
			h.hNum++
		} else {
			h.oNum++
		}
	}

	h.hFlag = make(chan struct{})
	h.oFlag = make(chan struct{})

	h.wg = &sync.WaitGroup{}
	return h
}

func CreateH2O(str string) {
	if len(str)%3 != 0 || len(str) < 1 || len(str) > 50 {
		return
	}
	h := newInstance(str)
	h.h2oFree()
	fmt.Println("done")
}

func Test_Create_H2O(t *testing.T) {
	CreateH2O("HOH")
	CreateH2O("OOHHHH")
	CreateH2O("OOOHHHHHH")
}
