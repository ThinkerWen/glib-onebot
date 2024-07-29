# OneBot Golang SDK

## 安装 💡

```shell
go get -u github.com/ThinkerWen/glib-onebot@latest
```

## 使用方法

```go
package main

import (
	"context"
	"github.com/ThinkerWen/glib-onebot"
	"github.com/ThinkerWen/glib-onebot/api"
	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/ThinkerWen/glib-onebot/messages"
)

func main() {
	bot, err := onebot.NewBot("ws://127.0.0.1:3001")
	if err != nil {
		panic(err)
	}
	bot.On(events.OnGroupMessageEvent, func(ctx context.Context, event events.IEvent) {
		message := events.EventParser(event).AsGroupMessage()
		msg := messages.NewTextMsg("Hello OneBot")
		go api.Build(bot.API).SendGroupMsg(message.GroupId, false, msg).Do(ctx)
	})
	err = bot.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
}
```