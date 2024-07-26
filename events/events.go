package events

import (
	"context"
	"encoding/json"
	"errors"
)

type EventName string
type EventCallbackFunc func(ctx context.Context, event IEvent)

const (
	OnPrivateMessageEvent EventName = "friend"
	OnGroupMessageEvent   EventName = "group"
	OnFriendRequestEvent  EventName = "friend"
	OnGroupRequestEvent   EventName = "group"
	OnGroupUploadEvent    EventName = "group_upload"
	OnGroupAdminEvent     EventName = "group_admin"
	OnGroupDecreaseEvent  EventName = "group_decrease"
	OnGroupIncreaseEvent  EventName = "group_increase"
	OnGroupBanEvent       EventName = "group_ban"
	OnFriendAddEvent      EventName = "friend_add"
	OnFriendRecallEvent   EventName = "friend_recall"
	OnNotifyEvent         EventName = "notify"
	OnLifecycleEvent      EventName = "lifecycle"
	OnHeartbeatEvent      EventName = "heartbeat"
)

// IEvent 定义了所有事件类型都需要实现的接口
type IEvent interface {
	GetTime() int64
	GetSelfId() int64
	GetEventName() EventName
}

// BaseEvent 基础事件结构体
type BaseEvent struct {
	Time     int64  `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}

// New 解析 JSON 数据并返回适当的 IEvent 实现
func New(data []byte) (IEvent, EventName, error) {
	var base BaseEvent
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, "", err
	}

	unmarshalEvent := func(name EventName, event IEvent) (IEvent, EventName, error) {
		if err := json.Unmarshal(data, event); err != nil {
			return nil, "", err
		}
		return event, name, nil
	}

	eventTypeMap := map[string]map[EventName]func() IEvent{
		"message": {
			OnPrivateMessageEvent: func() IEvent { return &PrivateMessageEvent{} },
			OnGroupMessageEvent:   func() IEvent { return &GroupMessageEvent{} },
		},
		"request": {
			OnFriendRequestEvent: func() IEvent { return &FriendRequestEvent{} },
			OnGroupRequestEvent:  func() IEvent { return &GroupRequestEvent{} },
		},
		"notice": {
			OnGroupUploadEvent:   func() IEvent { return &GroupUploadEvent{} },
			OnGroupAdminEvent:    func() IEvent { return &GroupAdminEvent{} },
			OnGroupDecreaseEvent: func() IEvent { return &GroupDecreaseEvent{} },
			OnGroupIncreaseEvent: func() IEvent { return &GroupIncreaseEvent{} },
			OnGroupBanEvent:      func() IEvent { return &GroupBanEvent{} },
			OnFriendAddEvent:     func() IEvent { return &FriendAddEvent{} },
			OnFriendRecallEvent:  func() IEvent { return &FriendRecallEvent{} },
			OnNotifyEvent:        func() IEvent { return &NotifyEvent{} },
		},
		"meta_event": {
			OnLifecycleEvent: func() IEvent { return &LifecycleEvent{} },
			OnHeartbeatEvent: func() IEvent { return &HeartbeatEvent{} },
		},
	}

	// 获取事件类型和子类型的映射
	if subEventMap, ok := eventTypeMap[base.PostType]; ok {
		var subtype string
		var subtypes struct {
			MessageType   string `json:"message_type"`
			RequestType   string `json:"sub_type"`
			NoticeType    string `json:"notice_type"`
			MetaEventType string `json:"meta_event_type"`
		}
		if err := json.Unmarshal(data, &subtypes); err != nil {
			return nil, "", err
		}
		switch base.PostType {
		case "message":
			subtype = subtypes.MessageType
		case "request":
			subtype = subtypes.RequestType
		case "notice":
			subtype = subtypes.NoticeType
		case "meta_event":
			subtype = subtypes.MetaEventType
		}

		if eventFunc, ok := subEventMap[EventName(subtype)]; ok {
			return unmarshalEvent(EventName(subtype), eventFunc())
		}
		return nil, "", errors.New("未知的子事件类型")
	}
	return nil, "", errors.New("未知的事件类型")
}

func (b *BaseEvent) GetTime() int64 {
	return b.Time
}

func (b *BaseEvent) GetSelfId() int64 {
	return b.SelfId
}

func (b *BaseEvent) GetEventName() EventName {
	return EventName(b.PostType)
}
