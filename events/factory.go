package events

import "encoding/json"

// EventFactory 用于解析事件
type EventFactory struct {
	data []byte
}

func EventParser(event IEvent) *EventFactory {
	if data, err := json.Marshal(event); err != nil {
		return nil
	} else {
		factory := new(EventFactory)
		factory.data = data
		return factory
	}
}

func (ef *EventFactory) AsPrivateMessage() PrivateMessageEvent {
	var e PrivateMessageEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return PrivateMessageEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupMessage() GroupMessageEvent {
	var e GroupMessageEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupMessageEvent{}
	}
	return e
}

func (ef *EventFactory) AsFriendRequest() FriendRequestEvent {
	var e FriendRequestEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return FriendRequestEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupRequest() GroupRequestEvent {
	var e GroupRequestEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupRequestEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupUpload() GroupUploadEvent {
	var e GroupUploadEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupUploadEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupAdmin() GroupAdminEvent {
	var e GroupAdminEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupAdminEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupDecrease() GroupDecreaseEvent {
	var e GroupDecreaseEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupDecreaseEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupIncrease() GroupIncreaseEvent {
	var e GroupIncreaseEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupIncreaseEvent{}
	}
	return e
}

func (ef *EventFactory) AsGroupBan() GroupBanEvent {
	var e GroupBanEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return GroupBanEvent{}
	}
	return e
}

func (ef *EventFactory) AsFriendAdd() FriendAddEvent {
	var e FriendAddEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return FriendAddEvent{}
	}
	return e
}

func (ef *EventFactory) AsFriendRecall() FriendRecallEvent {
	var e FriendRecallEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return FriendRecallEvent{}
	}
	return e
}

func (ef *EventFactory) AsNotify() NotifyEvent {
	var e NotifyEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return NotifyEvent{}
	}
	return e
}

func (ef *EventFactory) AsLifecycle() LifecycleEvent {
	var e LifecycleEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return LifecycleEvent{}
	}
	return e
}

func (ef *EventFactory) AsHeartbeat() HeartbeatEvent {
	var e HeartbeatEvent
	if err := json.Unmarshal(ef.data, &e); err != nil {
		return HeartbeatEvent{}
	}
	return e
}
