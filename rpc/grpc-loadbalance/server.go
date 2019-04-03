package main

import (
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"grpc-loadbalance/pb"
	_ "google.golang.org/grpc/balancer/roundrobin"
	"time"
	 "github.com/coreos/etcd/clientv3/naming"
	. "google.golang.org/grpc/naming"
)

var (
	addr = flag.Int("p", 0, "port")
	etcdAddr1 = flag.String("e",":2379","etcd addr")
)

type Server struct{}

func (s *Server) GetServerAddr(context.Context, *pb.Request) (*pb.Response, error) {
	resp := new(pb.Response)
	resp.Addr = fmt.Sprintf(":%d", *addr)

	return resp, nil
}

func main() {

	flag.Parse()

	err := RegisterService("/srvs")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	pb.RegisterLoadBalanceDemoServiceServer(srv, new(Server))

	fmt.Println("start!")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *addr))

	if err != nil {
		panic(err)
	}
	err = srv.Serve(l)

	if err != nil {
		panic(err)
	}

}

func RegisterService(service string) error {
	conn, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{*etcdAddr1},
		DialTimeout:          5*time.Second,
		DialKeepAliveTime:    5*time.Second,
		DialKeepAliveTimeout: 5*time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	r := &naming.GRPCResolver{
		Client: conn,
	}

	ky,err := conn.Get(context.Background(),"/hello")

	fmt.Println(ky,err)

	err = r.Update(conn.Ctx(),service,Update{Op:Add,Addr:fmt.Sprintf(":%d",*addr)})

	fmt.Println(err)

	return err
}