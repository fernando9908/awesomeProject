package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	// 监听端口
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		// 获取请求，如果成功则获得一个连接connection
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		// 处理连接，启动一个goroutine(子线程)处理该连接
		go process(client)
	}
}

func process(conn net.Conn) {
	// 函数退出时关闭该连接
	defer conn.Close()
	// 创建一个带缓冲的流
	reader := bufio.NewReader(conn)
	for {
		// 每次读一个字节b，看起来低效，但由于reader带缓冲一次不仅只读一个
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		// 将读到的字节写入
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
