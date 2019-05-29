package main
import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	)

// https://studygolang.com/articles/20519
// https://segmentfault.com/a/1190000019222661?utm_source=tag-newest
func main() {    // 开启pprof，监听请求
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
