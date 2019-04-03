package main

import (
	"fmt"
	"sync"
)

type SafeInt struct {
	sync.Mutex
	Num int
}

func main() {
	waitNum := 10 // 设置等待的个数（继续往下看）

	count := SafeInt{}

	done := make(chan bool)

	for i := 0; i < waitNum; i++ {
		go func(i int) {
			count.Lock() // 加锁，防止其它例程修改 count
			count.Num = count.Num + i
			fmt.Print(count.Num, " ")
			count.Unlock()

			done <- true
		}(i)
	}

	for i := 0; i < waitNum; i++ {
		<-done
	}
}