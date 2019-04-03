package main
import (
	"google.golang.org/grpc"
	pt "go-core-tree/rpc/grpc/my_grpc_proto"
	"fmt"
	"context"
)
const (
	post = "127.0.0.1:18881"
)


/*
server
1. 监听端口
2. 创建一个grpc server
3. 注册hello service server
4. 监听grpc服务



客户端
1. 链接服务器
2. 创建 server clint
3. 调用sayhello
4. 关闭
 */
func main() {
	// 客户端连接服务器
	conn,err:=grpc.Dial(post,grpc.WithInsecure())
	if err!=nil {
		fmt.Println("连接服务器失败",err)
	}


	defer conn.Close()
	//获得grpc句柄
	c:=pt.NewHelloServerClient(conn)

	// 远程调用 SayHello接口
	//远程调用 SayHello接口
	r1, err := c.SayHello(context.Background(), &pt.HelloRequest{Name: "panda", Age:100})
	if err != nil {
		fmt.Println("cloud not get Hello server ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r1.Message)

	//远程调用 GetHelloMsg接口
	r2, err := c.GetHelloMsg(context.Background(), &pt.HelloRequest{Name: "panda", Age:100})
	if err != nil {
		fmt.Println("cloud not get hello msg ..", err)
		return
	}
	fmt.Println("HelloServer resp: ", r2.Msg)
}