package main

import (
	"flag"
	"github.com/labstack/gommon/log"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-loadbalance/pb"
	"time"
)

var (
	etcdAddr   = flag.String("etcd", ":2379", "etcd addr")
	serverAddr = flag.String("p", ":0", "server addr")
)

func main() {
	flag.Parse()
	e, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{*etcdAddr},
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    5 * time.Second,
		DialKeepAliveTimeout: 5 * time.Second,
	})

	b := grpc.RoundRobin(&etcdnaming.GRPCResolver{Client: e})

	conn, err := grpc.Dial("/srvs",
		grpc.WithBalancer(b),
		grpc.WithInsecure(),
		//grpc.WithBlock(),
	)

	if err != nil {
		panic(err)
	}

	cli := pb.NewLoadBalanceDemoServiceClient(conn)

	for {

		resp, err := cli.GetServerAddr(context.Background(), &pb.Request{})
		log.Info(resp,err)
		time.Sleep(1 * time.Second)
	}
}

//func RegisterService(service string) error {
//	conn, err := clientv3.New(clientv3.Config{
//		Endpoints:            []string{*etcdAddr},
//		DialTimeout:          5*time.Second,
//		DialKeepAliveTime:    5*time.Second,
//		DialKeepAliveTimeout: 5*time.Second,
//	})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	r := &etcdnaming.GRPCResolver{
//		Client: conn,
//	}
//
//	return r.Update(conn.Ctx(),service,naming.Update{Op:naming.Add,Addr:*serverAddr})
//}
