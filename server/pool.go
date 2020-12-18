package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"
)

const (
	defaultTimeout = 300
)

type WorkPool struct {
	pools   chan chan *Connection
	workNum uint
	timeout uint // Millisecond unit
}

func (wp *WorkPool) AddConnection(con *Connection) error {
	if wp.timeout == 0 {
		wp.timeout = defaultTimeout
	}

	connectionChannelContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(wp.timeout))
	defer cancel()
	now := time.Now()
	var err error
	select {
	case connectionChannel := <-wp.pools:
		timeout, _ := context.WithTimeout(connectionChannelContext, time.Millisecond*time.Duration(wp.timeout)-time.Now().Sub(now))
		select {
		case <-timeout.Done():
			log.Println("任务入队超时")
			break
		case connectionChannel <- con:
			log.Println("添加了一个连接任务")
			break
		}
		break
	case <-connectionChannelContext.Done():
		err = errors.New("处理连接超时")
		log.Println(err)
		break
	}
	return err
}

type WorkConnection struct {
	connectionChannel chan *Connection
	sessionNum        uint
}

func (wc *WorkConnection) AcceptConnection(pools chan chan *Connection) {
	for {
		pools <- wc.connectionChannel
		select {
		case con := <-wc.connectionChannel:
			go con.Do()
			break
		}
	}
}

func newWorkPool(workNum, sessionNum uint) *WorkPool {
	return &WorkPool{
		workNum: workNum,
		pools: func() (pools chan chan *Connection) {
			var err error
			if pools, err = initPools(workNum, sessionNum); err == nil {
				return pools
			}
			panic(fmt.Sprintf("工作协程池初始化失败:%v", err))
		}(),
	}
}

func initPools(workNum, sessionNum uint) (chan chan *Connection, error) {
	numCpu := uint(runtime.NumCPU())
	if workNum == 0 {
		workNum = numCpu
	}
	var workChannels = make(chan chan *Connection, workNum)
	for ; workNum > 0; workNum-- {
		wc := &WorkConnection{
			connectionChannel: make(chan *Connection, sessionNum),
			sessionNum:        sessionNum,
		}
		go wc.AcceptConnection(workChannels)
	}
	return workChannels, nil
}
