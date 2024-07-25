package messages

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
