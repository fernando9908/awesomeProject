package main

import (
	"awesomeProject/hello/kitex_gen/api"
	"awesomeProject/hello/kitex_gen/api/hello"
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func main() {
	//c, err := hello.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	req := &api.Request{Message: "my request"}
	//	resp, err := c.Echo(context.Background(), req)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(resp)
	//	time.Sleep(time.Second)
	//}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	// 使用newClient也行
	c := hello.MustNewClient("hello", client.WithResolver(r))
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := c.Echo(ctx, &api.Request{Message: "hello"})
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
