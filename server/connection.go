package server

import (
	"bytes"
	"github.com/funnycode-org/gotty/base"
	"io"
	"log"
	"net"
)

type Connection struct {
	con     net.Conn
	session base.Session
	//sessionMap
	close chan struct{}
}

//type sessionMap struct {
//	sessionIds map[int]struct{}
//	sessions   []Session
//}

func NewConnection(con net.Conn, session base.Session) *Connection {
	return &Connection{
		con:     con,
		session: session,
		//sessionMap: sessionMap{
		//	sessionIds: make(map[int]struct{}, sessionCount),
		//	sessions:   make([]Session, sessionCount),
		//},
	}
}

// LoopReceivePkgs 接受包
func (con *Connection) LoopReceivePkgs() {
	defer func() {
		if err := recover(); err == nil {
			log.Printf("panic entered : %v", err)
			return
		}
	}()
	serverSession := con.session.(Session)
	var pktBuf bytes.Buffer
READ:
	for {
		count, err := con.con.Read(serverSession.receivedBytes)
		switch err {
		case io.EOF:
			log.Printf("session %d 读取包错误", con.session.SessionId())
			goto READ
		default:
			log.Printf("receive pkgs error:%v", err)
			con.session.Close()
			break READ
		}
		if count == 0 {
			continue
		}
		pktBuf.Reset()
		pktBuf.Write(serverSession.receivedBytes)
		// 包有问题，清空？
		err = con.tryExtractPkgs(&pktBuf)
		if err != nil {
			serverSession.receivedBytes = serverSession.receivedBytes[:0]
			pktBuf.Reset()
			continue
		}
	}

}

//
func (con *Connection) tryExtractPkgs(pktBuf *bytes.Buffer) error {
}

// 发送包
func (con *Connection) SendPkgs() {

}
func (con *Connection) Do() {
	go con.LoopReceivePkgs()
	go con.SendPkgs()
	for {
		select {
		case <-con.close:
			//todo
			break
		}
	}
}
