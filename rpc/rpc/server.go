package main
import (
	"net/http"
	"net/rpc"
	"net"
	"github.com/astaxie/beego"
	"io"
)

//- 方法是导出的
//- 方法有两个参数，都是导出类型或内建类型
//- 方法的第二个参数是指针
//- 方法只有一个error接口类型的返回值
//
//func (t *T) MethodName(argType T1, replyType *T2) error
type Panda int;
func (this *Panda)Getinfo(argType int, replyType *int) error {
	beego.Info(argType)
	*replyType =1 +argType

	return nil
}

type Panda1 int;
func (this *Panda1)Getinfo1(argType int, replyType *int) error {
	beego.Info(argType)
	*replyType =1 +argType

	return nil
}



func main() {
	//注册1个页面请求
	http.HandleFunc("/panda",pandatext)
	http.HandleFunc("/panda1",pandatext1)

	//new 一个对象
	pd :=new(Panda)
	pd1 :=new(Panda1)

	//注册服务
	//Register在默认服务中注册并公布 接收服务 pd对象 的方法
	rpc.Register(pd)
	rpc.Register(pd1)
	rpc.HandleHTTP()
	//建立网络监听
	ln , err :=net.Listen("tcp","127.0.0.1:10086")
	if err != nil{
		beego.Info("网络连接失败")
	}
	beego.Info("正在监听10086")
	//service接受侦听器l上传入的HTTP连接，
	http.Serve(ln,nil)
}


//用来现实网页的web函数
func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"panda")
}

//用来现实网页的web函数
func pandatext1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"panda1")
}