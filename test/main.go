package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/ThinkerWen/glib-onebot/api"
	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/ThinkerWen/glib-onebot/messages"
)

func main() {
	bot, err := onebot.NewBot("ws://127.0.0.1:3001")
	if err != nil {
		return
	}
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		message := events.EventParser(event).AsGroupMessage()
		for _, msg := range message.Message {
			_ = msg.Type
			_ = messages.MsgParser(msg).AsText().Text
		}
		_ = message.MessageId
		msg1 := messages.NewTextMsg("我爱你")
		msg2 := messages.NewFaceMsg("66")
		msg3 := messages.NewTextMsg("群主")
		go api.Build(bot.API).SendGroupMsg(message.GroupId, false, msg1, msg2, msg3).Do(ctx)
		if resp, err := api.Build(bot.API).GetFriendList().DoWithResponse(ctx); err == nil {
			if data, err := json.Marshal(resp); err == nil {
				fmt.Println(string(data))
			}
		}
	})
	err = bot.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
