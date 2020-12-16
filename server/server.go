package server

import (
	"fmt"
	"github.com/funnycode-org/gotty/base"
	"net"
)

type Server struct {
	Concurrency int // 并发个数
	connections map[string]base.Connection
}

func NewServer(concurrency int) *Server {

}

func (server *Server) Start() error {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		// 每个Client一个Goroutine
		go handleConnection(conn)
	}
	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var body [4]byte
	addr := conn.RemoteAddr()
	for {
		// 读取客户端消息
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("收到%s消息: %s\n", addr, string(body[:]))
		// 回包
		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
		fmt.Printf("发送给%s: %s\n", addr, string(body[:]))
	}
	fmt.Printf("与%s断开!\n", addr)
}
