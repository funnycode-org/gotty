package server

import (
	"github.com/funnycode-org/gotty/base"
	"net"
	"strconv"
)

type Server struct {
	concurrency uint // 并发个数
	workPool    *WorkPool
}

func NewServer() *Server {
	server := &Server{
		concurrency: base.GottyConfig.Server.Concurrency,
		workPool:    newWorkPool(base.GottyConfig.Server.Concurrency, base.GottyConfig.Server.SessionNumPerConnection),
		//connections: make(map[string]base.Connection, 13),
	}
	return server
}

func (server *Server) Start() error {
	var port uint = 8080
	if base.GottyConfig.Server.Port > 0 {
		port = base.GottyConfig.Server.Port
	}
	ln, err := net.Listen("tcp", strconv.Itoa(int(port)))
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		// 每个Client一个Goroutine
		server.workPool.AddConnection(base.NewConnection(conn, base.GottyConfig.Server.SessionNumPerConnection))
	}
	return nil
}
