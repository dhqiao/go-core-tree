package main

import (
	"strconv"
	"fmt"
	"time"
	"sync"
)

const N  = 1000 // fatal error: concurrent map read and map write


// N太小时不会（比如10），因机器而异
// fatal error: concurrent map read and map write
func mapDemo1() {
	m := make(map[string]int)

	go func() {
		for i := 0; i < N; i++ {
			m[strconv.Itoa(i)] = i // write
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			fmt.Println(i, m[strconv.Itoa(i)]) // read
		}
	}()

	time.Sleep(time.Second * 5)
}

// 可以看到Store就是写，Load就是读。
func mapDemo3() {
	var m sync.Map

	go func() {
		for i := 0; i < N; i++ {
			m.Store(strconv.Itoa(i), i) // 写
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			v, _ := m.Load(strconv.Itoa(i)) // 读
			fmt.Println(i, v)
		}
	}()

	time.Sleep(time.Second * 5)
}




type Cmap struct {
	m map[string]int
	//lock *sync.Mutex
	lock *sync.RWMutex
}

func (c *Cmap) Get(key string) int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.m[key]
}

func (c *Cmap) Set(key string, val int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.m[key] = val
}

func mapDemo2() {
	m := make(map[string]int)
	//lock := new(sync.Mutex)
	lock := new(sync.RWMutex)
	cm := Cmap{
		m:    m,
		lock: lock,
	}

	go func() {
		for i := 0; i < N; i++ {
			cm.Set(strconv.Itoa(i), i)
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			fmt.Println(i, cm.Get(strconv.Itoa(i)))
		}
	}()

	time.Sleep(time.Second * 5)
}


// https://blog.csdn.net/butterfly5211314/article/details/83316687
func main()  {
	mapDemo5()
}

//可以看到Store就是写，Load就是读。
// 仔细看下面的例子，就会理解，这个LoadOrStore方法其实是先取，
//如果有key的话就返回，没有再存储，其实就是优先取，没有再存储；
//和缓存的做法很类似：先从缓存中取，没有就从数据库中读，读回来存入缓存，下次就可以从缓存中取到了。
func mapDemo5()  {
	var m sync.Map

	actual, loaded := m.LoadOrStore("k1", "v1")
	fmt.Println(actual, loaded) // v1 false
	fmt.Println(m.Load("k1"))

	actual, loaded = m.LoadOrStore("k1", "v11")
	fmt.Println(actual, loaded) // v1 true
	fmt.Println(m.Load("k1"))

	actual, loaded = m.LoadOrStore("k1", "v2")
	fmt.Println(actual, loaded) // // v1 true
	fmt.Println(m.Load("k1"))

	actual, loaded = m.LoadOrStore("k2", "v22")
	fmt.Println(actual, loaded) // v2 false
	fmt.Println(m.Load("k2"))

}