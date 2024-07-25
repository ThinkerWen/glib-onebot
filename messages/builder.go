package messages

func NewTextMsg(text string) Message {
	msg := new(TextMessage)
	msg.Data.Text = text
	return Message{
		Type: "text",
		Data: msg.Data,
	}
}
