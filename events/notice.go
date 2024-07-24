package events

type NoticeEvent struct {
	BaseEvent
	NoticeType string `json:"notice_type"`
}

type GroupUploadEvent struct {
	NoticeEvent
	GroupId int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
	File    File  `json:"file"`
}
type File struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}

type GroupAdminEvent struct {
	NoticeEvent
	SubType string `json:"sub_type"`
	GroupId int64  `json:"group_id"`
	UserId  int64  `json:"user_id"`
}

type GroupDecreaseEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupId    int64  `json:"group_id"`
	OperatorId int64  `json:"operator_id"`
	UserId     int64  `json:"user_id"`
}

type GroupIncreaseEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupId    int64  `json:"group_id"`
	OperatorId int64  `json:"operator_id"`
	UserId     int64  `json:"user_id"`
}

type GroupBanEvent struct {
	NoticeEvent
	SubType    string `json:"sub_type"`
	GroupId    int64  `json:"group_id"`
	OperatorId int64  `json:"operator_id"`
	UserId     int64  `json:"user_id"`
	Duration   int64  `json:"duration"`
}

type FriendAddEvent struct {
	NoticeEvent
	UserId int64 `json:"user_id"`
}

type GroupRecallEvent struct {
	NoticeEvent
	GroupId    int64 `json:"group_id"`
	UserId     int64 `json:"user_id"`
	OperatorId int64 `json:"operator_id"`
	MessageId  int64 `json:"message_id"`
}

type FriendRecallEvent struct {
	NoticeEvent
	UserId    int64 `json:"user_id"`
	MessageId int64 `json:"message_id"`
}

type NotifyEvent struct {
	NoticeEvent
	SubType   string `json:"sub_type"`
	GroupId   int64  `json:"group_id"`
	UserId    int64  `json:"user_id"`
	TargetId  int64  `json:"target_id"`
	HonorType string `json:"honor_type"`
}
