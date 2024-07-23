package events

type NoticeEvent struct {
	BaseEvent
	NoticeType string `json:"notice_type"`
}

type GroupUploadEvent struct {
	NoticeEvent
}

type GroupAdminEvent struct {
	NoticeEvent
}
type GroupDecreaseEvent struct {
	NoticeEvent
}
type GroupIncreaseEvent struct {
	NoticeEvent
}
type GroupBanEvent struct {
	NoticeEvent
}
type FriendAddEvent struct {
	NoticeEvent
}
type GroupRecallEvent struct {
	NoticeEvent
}
type FriendRecallEvent struct {
	NoticeEvent
}
type NotifyEvent struct {
	NoticeEvent
}
