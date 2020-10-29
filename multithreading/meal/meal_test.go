package meal

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//5 个沉默寡言的哲学家围坐在圆桌前，每人面前一盘意面。叉子放在哲学家之间的桌面上。（5 个哲学家，5 根叉子）
//
//所有的哲学家都只会在思考和进餐两种行为间交替。哲学家只有同时拿到左边和右边的叉子才能吃到面，而同一根叉子在同一时间只能被一个哲学家使用。每个哲学家吃完面后都需要把叉子放回桌面以供其他哲学家吃面。只要条件允许，哲学家可以拿起左边或者右边的叉子，但在没有同时拿到左右叉子时不能进食。
//
//假设面的数量没有限制，哲学家也能随便吃，不需要考虑吃不吃得下。
//
//设计一个进餐规则（并行算法）使得每个哲学家都不会挨饿；也就是说，在没有人知道别人什么时候想吃东西或思考的情况下，每个哲学家都可以在吃饭和思考之间一直交替下去。

//
//哲学家从 0 到 4 按 顺时针 编号。请实现函数 void wantsToEat(philosopher, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)：
//
//philosopher 哲学家的编号。
//pickLeftFork 和 pickRightFork 表示拿起左边或右边的叉子。
//eat 表示吃面。
//putLeftFork 和 putRightFork 表示放下左边或右边的叉子。

//由于哲学家不是在吃面就是在想着啥时候吃面，所以思考这个方法没有对应的回调。
//给你 5 个线程，每个都代表一个哲学家，请你使用类的同一个对象来模拟这个过程。在最后一次调用结束之前，可能会为同一个哲学家多次调用该函数。

type DiningPhilosophers struct {
	wg                     *sync.WaitGroup
	streamForks            [5]chan interface{}
	missingDoubleForkTimes int
}

func Think() {
	Random := func(max int) int {
		rand.Seed(time.Now().Unix())
		return rand.Int() % (max + 1)
	}
	<-time.After(time.Millisecond * time.Duration(Random(50)))
}

func PickLeftFork(i int) {
	fmt.Printf("%d号哲学家拿起左叉子\n", i)
}

func PickRightFork(i int) {
	fmt.Printf("%d号哲学家拿起右叉子\n", i)
}

func Eat(i int) {
	fmt.Printf("%d号哲学家开始进餐。。。。。。\n", i)
}

func PutLeftFork(i int) {
	fmt.Printf("%d号哲学家放下左叉子\n", i)
}

func PutRightFork(i int) {
	fmt.Printf("%d号哲学家放下右叉子\n\n", i)
}
func (this *DiningPhilosophers) WantEat(philosopher int, pickLeftFork func(int), pickRightFork func(int), eat func(int), putLeftFork func(int), putRightFork func(int)) {
	defer this.wg.Done()

	var leftNum = (philosopher + 4) % 5  //取得该哲学家左边的号码
	var rightNum = (philosopher + 6) % 5 //取得该哲学家右边的号码

	for {
		select {
		case this.streamForks[leftNum] <- philosopher: //尝试拿起左边筷子
			PickLeftFork(philosopher) //成功拿起左边筷子
			select {
			case this.streamForks[rightNum] <- philosopher: //尝试拿起右边筷子
				PickRightFork(philosopher)  //成功拿起又边筷子
				Eat(philosopher)            //左右边都拿到了，开始吃
				<-this.streamForks[leftNum] //吃完了，放下左边筷子
				PutLeftFork(philosopher)
				<-this.streamForks[rightNum] //吃完了，放下右边筷子
				PutRightFork(philosopher)
				return //吃饱离开
			default: //无法拿起右边筷子
				fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, rightNum)
				<-this.streamForks[leftNum] //把已经拿起来的左边筷子释放出去
				PutLeftFork(philosopher)
			}
		default: //无法拿起左边筷子
			fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, leftNum)
		}
		this.missingDoubleForkTimes++
		Think()
	}
}

func Test_Meal(t *testing.T) {
	d := new(DiningPhilosophers)
	d.wg = &sync.WaitGroup{}
	for i := range d.streamForks {
		d.streamForks[i] = make(chan interface{}, 1)
	}

	for i := range d.streamForks {
		d.wg.Add(1)
		go d.WantEat(i, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	}

	d.wg.Wait()
}

//参考链接：https://www.zhihu.com/column/c_1286635591711801344
