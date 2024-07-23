package events

type RequestEvent struct {
	BaseEvent
	RequestType string `json:"request_type"`
	UserId      int64  `json:"user_id"`
	Comment     string `json:"comment"`
	Flag        string `json:"flag"`
}

type FriendRequestEvent struct {
	RequestEvent
}

type GroupRequestEvent struct {
	RequestEvent
	SubType string `json:"sub_type"`
	GroupId int64  `json:"group_id"`
	UserId  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
}
