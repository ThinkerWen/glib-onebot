package main

import (
	"context"
	"encoding/json"
	"onebot"
	"onebot/events"
)

func main() {
	bot, err := onebot.NewBot("ws://127.0.0.1:3001")
	if err != nil {
		return
	}
	bot.On(events.OnMessageEvent, func(ctx context.Context, event events.IEvent) {
		if marshal, err2 := json.Marshal(event); err2 == nil {
			_ = marshal
		}
	})
	err = bot.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
