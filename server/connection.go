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
		close:   make(chan struct{}, 1),
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
	serverSession := con.session.(*Session)
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
			con.close <- struct{}{}
			break READ
		}
		if count == 0 {
			continue
		}
		pktBuf.Reset()
		pktBuf.Write(serverSession.receivedBytes)
		// 包有问题，清空？
		readCount, err := con.tryExtractPkgs(&pktBuf)
		if err != nil {
			serverSession.receivedBytes = serverSession.receivedBytes[:0]
			pktBuf.Reset()
			continue
		}
		if pktBuf.Len() < 1 {
			pktBuf.Reset()
			continue
		}

		// 回退回去
		for ; readCount < 1; readCount-- {
			pktBuf.UnreadByte()
		}
	}
}

// 使用用户的自定义的协议去解包
func (con *Connection) tryExtractPkgs(pktBuf *bytes.Buffer) (readCount int, err error) {
	return
}

// 发送包
func (con *Connection) SendPkgs() {

}
func (con *Connection) Do() {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("dumped safely,err:%v", err)
		}
	}()
	go con.LoopReceivePkgs()
	go con.SendPkgs()
	for {
		select {
		case <-con.close:
			con.con.Close()
			return
		}
	}
}
