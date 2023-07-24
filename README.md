# websocket客户端连接包

比较简单，测试用比较方便，直接看使用demo

```go
    // 创建ws连接
	wsObj := wsclient.WsClientIns{}

    // 创建连接
    err := wsObj.CreateConn("ws://test.com", "acckey=xxx")

	// 发送心跳包, 保持连接
	go wsObj.SendHeartbeat(30, []byte(":Ping"))

    // 发送消息
    err := wsObj.SendTextMsg("test text")

    // 接收消息
    msgType, msg, err := wsObj.RecvMsg()
```

一共就几个接口，根据业务自由组合
