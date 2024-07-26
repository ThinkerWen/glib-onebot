package messages

func NewTextMsg(text string) Message {
	msg := new(TextMessage)
	msg.Data.Text = text
	return Message{
		Type: "text",
		Data: msg.Data,
	}
}

func NewFaceMsg(faceId string) Message {
	msg := new(FaceMessage)
	msg.Data.Id = faceId
	return Message{
		Type: "face",
		Data: msg.Data,
	}
}
