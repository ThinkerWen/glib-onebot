package events

import "onebot/messages"

type MessageEvent struct {
	BaseEvent
	MessageType string             `json:"message_type"`
	SubType     string             `json:"sub_type"`
	MessageId   int                `json:"message_id"`
	UserId      int64              `json:"user_id"`
	Message     []messages.Message `json:"message"`
	RawMessage  string             `json:"raw_message"`
	Font        int                `json:"font"`
}

type PrivateMessageEvent struct {
	MessageEvent
	Sender PrivateSender `json:"sender"`
}
type PrivateSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
}

type GroupMessageEvent struct {
	MessageEvent
	GroupId   int64       `json:"group_id"`
	Anonymous Anonymous   `json:"anonymous"`
	Sender    GroupSender `json:"sender"`
}
type Anonymous struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}
type GroupSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}
