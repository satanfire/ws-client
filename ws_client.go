/*
	auth: satanfire
	date: 2023/7/24
	desc: websocket client interaction demo
*/
package wsclient

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 客户端对象
type WsClientIns struct {
	conn     *websocket.Conn
	mtx      sync.Mutex
	exitChan chan struct{}
}

/*
	创建连接
	url: 长链接地址
	query: 请求参数
*/
func (obj *WsClientIns) CreateConn(url, query string) error {
	url = fmt.Sprintf("%s?%s", url, query)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err == nil {
		obj.conn = conn
		obj.exitChan = make(chan struct{})
	}
	return err
}

// 发送心跳包, gapS: 心跳包时间间隔
func (obj *WsClientIns) SendHeartbeat(gapS int, data []byte) {
	if obj.conn == nil {
		return
	}

	ticker := time.NewTicker(time.Second * time.Duration(gapS))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			obj.mtx.Lock()
			err := obj.conn.WriteMessage(websocket.PingMessage, data)
			obj.mtx.Unlock()
			if err != nil {
				fmt.Printf("\x1b[91m heart write:%s \x1b[0m", err.Error())
			}
		case <-obj.exitChan:
			fmt.Printf("\x1b[91m heart exit. \x1b[0m")
			return
		}
	}
}

/*
	发送文本消息

*/
func (obj *WsClientIns) SendTextMsg(data []byte) error {
	if obj.conn == nil {
		return errors.New("conn is nil")
	}

	obj.mtx.Lock()
	err := obj.conn.WriteMessage(websocket.TextMessage, data)
	obj.mtx.Unlock()
	return err
}

// 接收消息
func (obj *WsClientIns) RecvMsg() (int, []byte, error) {
	if obj.conn == nil {
		return 0, nil, errors.New("conn is nil")
	}
	return obj.conn.ReadMessage()
}

// stop heart
func (obj *WsClientIns) StopHeart() {
	obj.exitChan <- struct{}{}
}
