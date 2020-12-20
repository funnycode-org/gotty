package server

import (
	"fmt"
	"github.com/funnycode-org/gotty/base"
	"github.com/funnycode-org/gotty/server/listener"
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

		session := NewSession(server.FindListener())
		session.l.OnOpen(session)
		connection := NewConnection(conn, session)
		err = server.workPool.AddConnection(connection)
		if err != nil {
			fmt.Println("添加连接任务出现错误:", err)
		}
	}
	return nil
}

func (server *Server) FindListener() (serverListener listener.Listener) {
	listenerName := base.GottyConfig.Server.ListenerName
	var err error
	if len(listenerName) <= 0 {
		listenerName = listener.DefaultClientListenerName
	}
	serverListener, err = listener.FindListener(listenerName)()
	if err != nil {
		serverListener, _ = listener.FindListener(listener.DefaultClientListenerName)()
	}
	return
}
