package events

import (
	"context"
	"encoding/json"
	"errors"
)

type EventName string
type EventCallbackFunc func(ctx context.Context, event IEvent)

const (
	OnMessageEvent EventName = "message"
	OnRequestEvent EventName = "request"
	OnNoticeEvent  EventName = "notice"
	OnMetaEvent    EventName = "meta_event"
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
func New(data []byte) (IEvent, error) {
	var base BaseEvent
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	unmarshalEvent := func(event IEvent) (IEvent, error) {
		if err := json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, nil
	}

	eventTypeMap := map[EventName]map[string]func() IEvent{
		OnMessageEvent: {
			"friend": func() IEvent { return &PrivateMessageEvent{} },
			"group":  func() IEvent { return &GroupMessageEvent{} },
		},
		OnRequestEvent: {
			"friend": func() IEvent { return &FriendRequestEvent{} },
			"group":  func() IEvent { return &GroupRequestEvent{} },
		},
		OnNoticeEvent: {
			"group_upload":   func() IEvent { return &GroupUploadEvent{} },
			"group_admin":    func() IEvent { return &GroupAdminEvent{} },
			"group_decrease": func() IEvent { return &GroupDecreaseEvent{} },
			"group_increase": func() IEvent { return &GroupIncreaseEvent{} },
			"group_ban":      func() IEvent { return &GroupBanEvent{} },
			"friend_add":     func() IEvent { return &FriendAddEvent{} },
			"friend_recall":  func() IEvent { return &FriendRecallEvent{} },
			"notify":         func() IEvent { return &NotifyEvent{} },
		},
		OnMetaEvent: {
			"lifecycle": func() IEvent { return &LifecycleEvent{} },
			"heartbeat": func() IEvent { return &HeartbeatEvent{} },
		},
	}

	// 获取事件类型和子类型的映射
	if subEventMap, ok := eventTypeMap[base.GetEventName()]; ok {
		var subtype string
		var subtypes struct {
			MessageType   string `json:"message_type"`
			RequestType   string `json:"sub_type"`
			NoticeType    string `json:"notice_type"`
			MetaEventType string `json:"meta_event_type"`
		}
		if err := json.Unmarshal(data, &subtypes); err != nil {
			return nil, err
		}
		switch base.GetEventName() {
		case OnMessageEvent:
			subtype = subtypes.MessageType
		case OnRequestEvent:
			subtype = subtypes.RequestType
		case OnNoticeEvent:
			subtype = subtypes.NoticeType
		case OnMetaEvent:
			subtype = subtypes.MetaEventType
		}

		if eventFunc, ok := subEventMap[subtype]; ok {
			return unmarshalEvent(eventFunc())
		}
		return nil, errors.New("未知的子事件类型")
	}
	return nil, errors.New("未知的事件类型")
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

func (m *MessageEvent) AsMessageEvent() *MessageEvent {
	return m
}
