package base

import (
	"bytes"
	"fmt"
	"github.com/funnycode-org/gotty/protocol"
	"github.com/funnycode-org/gotty/protocol/registry"
	"io"
	"log"
	"net"
	"reflect"
)

type Connection struct {
	isServer bool
	con      net.Conn
	session  Session
	//sessionMap
	close                chan struct{}
	RegistryProtocolKind protocol.ProtocolDecoder
	RegistryProtocol     reflect.Type
	writer               *bytes.Buffer
	wrappedSession       Session
}

func NewConnection(con net.Conn, session Session, wrappedSession Session, isServer bool) *Connection {
	registryProtocolKind, registryProtocol, err := registry.GetProtocol()
	if err != nil {
		log.Fatalf("获取自定义的协议失败:%v", err)
	}
	var buffer = make([]byte, 1024)
	return &Connection{
		isServer:             isServer,
		writer:               bytes.NewBuffer(buffer),
		RegistryProtocolKind: registryProtocolKind,
		RegistryProtocol:     registryProtocol,
		con:                  con,
		close:                make(chan struct{}, 1),
		session:              session,
		wrappedSession:       wrappedSession,
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
	var receivedBytes []byte
	if con.isServer {
		receivedBytes = make([]byte, GottyConfig.Server.ReceivedBytesLength)
	} else {
		receivedBytes = make([]byte, GottyConfig.Client.ReceivedBytesLength)
	}
	var pktBuf bytes.Buffer
READ:
	for {
		count, err := con.con.Read(receivedBytes)
		if err != nil {
			switch err {
			case io.EOF:
				log.Printf("session %d 读取包错误", con.session.SessionId())
				goto READ
			default:
				log.Printf("receive pkgs error:%v", err)
				con.close <- struct{}{}
				break READ
			}
		}
		if count == 0 {
			continue
		}
		pktBuf.Reset()
		pktBuf.Write(receivedBytes)
		// 包有问题，清空？
		err = con.tryExtractPkgs(&pktBuf)
		if err != nil {
			receivedBytes = receivedBytes[:0]
			pktBuf.Reset()
			continue
		}
		if pktBuf.Len() < 1 {
			pktBuf.Reset()
			continue
		}
	}
}

// 使用用户的自定义的协议去解包
func (con *Connection) tryExtractPkgs(pktBuf *bytes.Buffer) (err error) {
	decodeOk, err := con.RegistryProtocolKind.Decode(pktBuf, con.writer)
	if err != nil {
		return
	}
	if decodeOk {
		//serverSession := con.session.(*server.Session)
		value := reflect.New(con.session.GetRegistryProtocol()).Elem()
		customizeProtocol, ok := value.Interface().(protocol.CustomizeProtocol)
		if !ok {
			panic(fmt.Sprintf("类型 %s 没有实现接口protocol.CustomizeProtocol", con.session.GetRegistryProtocol().Name()))
		}
		bytes := con.writer.Bytes()
		err := customizeProtocol.Decode(bytes)
		if err != nil {
			log.Printf("哦豁，解析到结果失败了:%v\n", err)
			log.Println("错误解析的字节", bytes)
		}
		con.writer.Reset()
	}
	return
}

// 发送包
func (con *Connection) SendPkgs() {
	for {
		select {
		case bytes := <-con.session.GetSendChannel():
			for i := 0; i < GottyConfig.Server.Retry; i++ {
				_, err := con.con.Write(bytes)
				if err == nil {
					break
				}
			}
			break
		}
	}
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
