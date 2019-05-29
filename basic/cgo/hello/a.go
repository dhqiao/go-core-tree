package main

import (
	"fmt"
)

/*
typedef long LONG;
typedef unsigned long DWORD;
typedef long long LONGLONG;

typedef union _LARGE_INTEGER {
    struct {
        DWORD LowPart;
        LONG HighPart;
    };
    struct {
        DWORD LowPart;
        LONG HighPart;
    } u;
    LONGLONG QuadPart;
} LARGE_INTEGER, *PLARGE_INTEGER;

void Show(LARGE_INTEGER li)
{
	li.u.LowPart = 1;
	li.u.HighPart = 4;
}
*/
import "C"

func main() {
	var li C.LARGE_INTEGER // 等价于： var li [8]byte
	var b [8]byte = li     // 正确，因为[8]byte和C.LARGE_INTEGER相同
	C.Show(b)              // 参数类型为LARGE_INTEGER，可以接收[8]byte
	li[0] = 75
	fmt.Println(li) // [75 0 0 0 0 0 0 0]
	li[4] = 23
	Test(li) // 参数类型为[8]byte，可以接收C.LARGE_INTEGER
}

func Test(b [8]byte) {
	fmt.Println(b)
}
