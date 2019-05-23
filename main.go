package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	arrayA := [10]int{}
	fmt.Println(arrayA)
	aa := make(map[int]string)
	fmt.Println(aa)
	bb := new(map[int]string)
	fmt.Println(bb)
	arrayB := make([]int, 10)
	fmt.Println(arrayB)
	var i = 0
	changei(i)
	fmt.Println(i)
	ii := new(int)
	fmt.Println(*ii)
	changeii(ii)
	fmt.Println(*ii)

}

func changei(i int)  {
	i = 1000
}
func changeii(i *int)  {
	*i = 1000
}
