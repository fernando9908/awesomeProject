package main

import (
	"awesomeProject/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	//svr := api.NewServer(new(HelloImpl))
	//
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	s := hello.NewServer(new(HelloImpl), server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}))
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
