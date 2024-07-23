package messages

import "encoding/json"

type IMessage interface {
	AsTextMsg() (*TextMessage, bool)
	AsFaceMsg() (*FaceMessage, bool)
}

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
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

func (msg *Message) AsTextMessage() *TextMessage {
	var textMsg TextMessage
	if err := json.Unmarshal(msg.Data, &textMsg); err != nil {
		return nil
	}
	return &textMsg
}

func (msg *Message) AsFaceMessage() *FaceMessage {
	var faceMsg FaceMessage
	if err := json.Unmarshal(msg.Data, &faceMsg); err != nil {
		return nil
	}
	return &faceMsg
}
