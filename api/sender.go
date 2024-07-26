package api

import (
	"context"
	"encoding/json"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"time"
)

func MarshalJSON(action ActionStruct) ([]byte, error) {
	return json.Marshal(&struct {
		Action string                 `json:"action"`
		Params map[string]interface{} `json:"params"`
	}{
		Action: action.GetAction(),
		Params: action.GetParams(),
	})
}

type IDo interface {
	Do(ctx context.Context) error
}

func (a *ActionStruct) Do(ctx context.Context) error {
	var err error
	var data []byte
	if data, err = MarshalJSON(*a); err != nil {
		return err
	}
	log.Debug("正在发送: " + string(data))

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, a.apiUrl, nil)
	if err != nil {
		log.Error("连接WebSocket错误:", err)
		return err
	}
	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Error("设置写超时错误:", err)
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Error("发送错误:", err)
		return err
	}

	return nil
}
