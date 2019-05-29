package main

/*
#include <stdint.h>

static int32_t add(int32_t a, int32_t b) {
    return a + b;
}
*/
import "C"
import "fmt"

func main() {
	var a, b int32 = 1, 2
	var c int32 = int32(C.add(C.int32_t(a), C.int32_t(b)))
	fmt.Println(c) // 3
}