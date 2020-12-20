package client

import (
	"errors"
	"fmt"
	"github.com/funnycode-org/gotty/base"
	"github.com/funnycode-org/gotty/client/lisener"
	"net"
	"sync"
)

type Client struct {
	Ip       string
	Port     int
	sessions map[int]*Session
	locker   sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		Ip:       base.GottyConfig.Client.ServerIp,
		Port:     base.GottyConfig.Client.Port,
		sessions: make(map[int]*Session, base.GottyConfig.Client.SessionNum),
		locker:   sync.RWMutex{},
	}
}

func (c *Client) AddNewSession(session *Session) error {
	c.locker.Lock()
	defer c.locker.Unlock()
	if _, exist := c.sessions[session.SessionId()]; exist {
		return errors.New(fmt.Sprintf("session %d 已经存在!", session.SessionId()))
	}
	c.sessions[session.SessionId()] = session
	return nil
}

func (c *Client) SelectSession() (session *WrappedSession, err error) {
	c.locker.RLock()
	defer c.locker.Unlock()
	num := len(c.sessions)
	if num > base.GottyConfig.Client.SessionNum {
		err = errors.New(fmt.Sprintf("客户端的并发数超出了最大限制个数：%d", base.GottyConfig.Client.SessionNum))
		return
	}
	for _, s := range c.sessions {
		if s.Using {
			continue
		}
		session = newWrappedSession(s)
	}

	if session != nil {
		return
	}
	newSession, err := c.newSession()
	if err != nil {
		return
	}
	session = newWrappedSession(newSession)
	return
}

func (c *Client) FindListener() (clientListener listener.Listener) {
	listenerName := base.GottyConfig.Server.ListenerName
	var err error
	if len(listenerName) <= 0 {
		listenerName = listener.DefaultClientListenerName
	}
	clientListener, err = listener.FindListener(listenerName)()
	if err != nil {
		clientListener, _ = listener.FindListener(listener.DefaultClientListenerName)()
	}
	return
}

func (c *Client) newSession() (session *Session, err error) {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Ip, c.Port))
	if err != nil {
		return
	}

	newSession := NewSession(c.FindListener())
	err = c.AddNewSession(newSession)
	if err != nil {
		return
	}
	wrappedSession := newWrappedSession(newSession)
	// 使用一个安全的session
	session.l.OnOpen(wrappedSession)
	connection := base.NewConnection(con, newSession, wrappedSession, false)
	go connection.Do()
	return
}
