package main

import (
	"fmt"
	"sync"
	"time"
)
//  https://my.oschina.net/90design/blog/1814499
var m *sync.RWMutex
var wg sync.WaitGroup

func main() {
	m = new(sync.RWMutex)
	wg.Add(2)
	go write(1)
	go write(4)
	go write(5)

	time.Sleep(1 * time.Second)
	go read(2)
	go read(3);
	wg.Wait()
	time.Sleep(6*time.Second)
}
func write(i int) {
	fmt.Println(i, "写开始.")
	m.Lock()
	fmt.Println(i, "正在写入中......")
	time.Sleep(3 * time.Second)
	m.Unlock()
	fmt.Println(i, "写入结束.")
	wg.Done()
}
func read(i int) {
	fmt.Println(i, "读开始.")
	m.RLock()
	fmt.Println(i, "正在读取中......")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	fmt.Println(i, "读取结束.")
	wg.Done()
}

func read2(i int) {
	fmt.Println(i, "读开始.")
	m.RLock()
	fmt.Println(i, "正在读取中......")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	fmt.Println(i, "读取结束.")
	wg.Done()
}