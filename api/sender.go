package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"time"
)

// 动作消息格式化方法
func marshalJSON(action ActionStruct) ([]byte, error) {
	return json.Marshal(&struct {
		Action string                 `json:"action"`
		Params map[string]interface{} `json:"params"`
	}{
		Action: action.GetAction(),
		Params: action.GetParams(),
	})
}

// IDo 执行动作抽象接口
type IDo interface {
	Do(ctx context.Context) error
	DoWithResponse(ctx context.Context) ([]byte, error)
}

// 连接 websocket
func (a *ActionStruct) connect(ctx context.Context) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, a.apiUrl, nil)
	if err != nil {
		log.Error("连接WebSocket错误:", "err", err)
		return nil, err
	}
	return conn, nil
}

// 发送动作消息
func (a *ActionStruct) sendMessage(conn *websocket.Conn, data []byte) error {
	if err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Error("设置写超时错误:", "err", err)
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Error("发送错误:", "err", err)
		return err
	}

	return nil
}

// 读取websocket内容
func (a *ActionStruct) readMessage(conn *websocket.Conn) ([]byte, error) {
	if err := conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Error("设置读超时错误:", "err", err)
		return nil, err
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Error("读取错误:", "err", err)
		return nil, err
	}

	log.Debug("收到响应: ", "message", string(message))
	return message, nil
}

// Do 执行动作
func (a *ActionStruct) Do(ctx context.Context) error {
	data, err := marshalJSON(*a)
	if err != nil {
		return err
	}

	conn, err := a.connect(ctx)
	if err != nil {
		return err
	}
	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	if err := a.sendMessage(conn, data); err != nil {
		return err
	}

	return nil
}

// DoWithResponse 执行动作（带返回值）
func (a *ActionStruct) DoWithResponse(ctx context.Context) ([]byte, error) {
	data, err := marshalJSON(*a)
	if err != nil {
		return nil, err
	}

	conn, err := a.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	if err := a.sendMessage(conn, data); err != nil {
		return nil, err
	}

	responseCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for {
		select {
		case <-responseCtx.Done():
			return nil, fmt.Errorf("等待响应超时")
		default:
			response, err := a.readMessage(conn)
			if err != nil {
				return nil, err
			}

			var responseMap map[string]interface{}
			if err := json.Unmarshal(response, &responseMap); err != nil {
				log.Error("解析响应错误:", "err", err)
				return nil, err
			}

			if status, ok := responseMap["status"].(string); ok && status == "ok" {
				return response, nil
			}
		}
	}
}
