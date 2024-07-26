package messages

import (
	"encoding/json"
	"github.com/charmbracelet/log"
)

type IMessage interface {
}

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TextMessage struct {
	Message
	Data struct {
		Text string `json:"text"`
	}
}

type FaceMessage struct {
	Message
	Data struct {
		Id string `json:"id"`
	}
}

func (m Message) AsTextMsg() TextMessage {
	var msg TextMessage
	msg.Type = m.Type

	data, err := json.Marshal(m.Data)
	if err != nil {
		log.Error("Marshal错误:", err)
		return msg
	}

	if err = json.Unmarshal(data, &msg.Data); err != nil {
		log.Error("Unmarshal错误:", err)
		return msg
	}

	return msg
}

func (m Message) AsFaceMsg() FaceMessage {
	var msg FaceMessage
	msg.Type = m.Type

	data, err := json.Marshal(m.Data)
	if err != nil {
		log.Error("Marshal错误:", err)
		return msg
	}

	if err = json.Unmarshal(data, &msg.Data); err != nil {
		log.Error("Unmarshal错误:", err)
		return msg
	}

	return msg
}
