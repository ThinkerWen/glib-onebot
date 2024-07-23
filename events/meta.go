package events

type MetaEvent struct {
	BaseEvent
	MetaEventType string `json:"meta_event_type"`
	SubType       string `json:"sub_type"`
}

type LifecycleEvent struct {
	MetaEvent
}

type HeartbeatEvent struct {
	MetaEvent
	Status   interface{} `json:"status"`
	Interval int64       `json:"interval"`
}
