package main

import (
	"fmt"
	"unsafe"
)
// https://blog.csdn.net/cyk2396/article/details/78893420
func main()  {
	main5()
}
// 运行以下代码思考一个问题：s1和s2是指向同一个底层数组吗？
func main2() {
	array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[:5]
	s2 := append(s1, 10)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	s2[0] = 0
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmtPrint(s1, "s1")
	fmtPrint(s2, "s2")

}

func main1() {
	//1.基于数组创建数组切片
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice = array[1:7] //array[startIndex:endIndex] 不包含endIndex
	//2.直接创建数组切片
	slice2 := make([]int, 5, 10)
	//3.直接创建并初始化数组切片
	slice3 := []int{1, 2, 3, 4, 5, 6}
	//4.基于数组切片创建数组切片
	slice5 := slice3[:4]
	fmtPrint(slice3, "slice3")
	fmtPrint(slice5, "slice5")
	//5.遍历数组切片
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
	//6.len()和cap()
	var len = len(slice2) //数组切片的长度
	var cap = cap(slice)  //数组切片的容量
	fmt.Println("len(slice2) =", len)
	fmt.Println("cap(slice) =", cap)
	//7.append() 会生成新的数组切片
	slice4 := append(slice2, 6, 7, 8)
	slice4 = append(slice4, slice3...)
	fmt.Println(slice4)
	//8.copy() 如果进行操作的两个数组切片元素个数不一致，将会按照个数较小的数组切片进行复制
	copy(slice2, slice3) //将slice3的前五个元素复制给slice2
	fmt.Println(slice2, slice3)
}

func fmtPrint(slice []int, name string) {
	fmt.Println(name + ": ", slice)
	fmt.Println("len("+name+"): ", len(slice))
	fmt.Println("cap("+name+"): ", cap(slice))
	fmt.Println(unsafe.Pointer(&slice))

}

type Slice struct {
	ptr unsafe.Pointer // Array pointer
	len int            // slice length
	cap int            // slice capacity
}

func main4() {
	array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[:5]
	s2 := append(s1, 10)
	s2[0] = 0
	// 把slice转换成自定义的 Slice struct
	slice1 := (*Slice)(unsafe.Pointer(&s1))
	fmt.Printf("ptr:%v len:%v cap:%v \n", slice1.ptr, slice1.len, slice1.cap)
	slice2 := (*Slice)(unsafe.Pointer(&s2))
	fmt.Printf("ptr:%v len:%v cap:%v \n", slice2.ptr, slice2.len, slice2.cap)
}


func main5() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := append(s1, 10)
	fmtPrint(s1, "s1")
	fmtPrint(s2, "s2")
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	s2[0] = 0
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
}