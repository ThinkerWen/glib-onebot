package onebot

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/ThinkerWen/glib-onebot/events"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

type Bot struct {
	API    string
	events map[events.EventName][]events.EventCallbackFunc
	lock   sync.RWMutex
	client *websocket.Conn
	botQQ  *int64
	done   chan struct{}
}

// NewBot 使用指定的 API URL 创建新的 Bot 实例。
func NewBot(api string) (*Bot, error) {
	log.SetLevel(log.DebugLevel)
	return &Bot{
		API:    api,
		events: make(map[events.EventName][]events.EventCallbackFunc),
		lock:   sync.RWMutex{},
		done:   nil,
	}, nil
}

// On 为特定事件注册回调函数。
func (bot *Bot) On(event events.EventName, callback events.EventCallbackFunc) {
	bot.lock.Lock()
	defer bot.lock.Unlock()
	bot.events[event] = append(bot.events[event], callback)
}

// ListenAndWait 连接到 WebSocket 服务器并侦听事件。
func (bot *Bot) ListenAndWait(ctx context.Context) error {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		select {
		case <-interrupt:
			log.Info("用户关闭程序")
			cancel()
		case <-ctx.Done():
		}
	}()

	bot.done = make(chan struct{})
	var err error
	bot.client, _, err = websocket.DefaultDialer.DialContext(ctx, bot.API, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = bot.client.Close(); err != nil {
			log.Error("关闭连接错误:", err)
		}
	}()
	log.Info("连接成功")

	go bot.readMessages(ctx)

	<-bot.done
	return nil
}

// readMessages 持续从 WebSocket 连接读取消息。
func (bot *Bot) readMessages(ctx context.Context) {
	defer close(bot.done)

	for {
		var message []byte
		log.Debug("正在等待读取消息...")
		_, message, err := bot.client.ReadMessage()
		if err != nil {
			log.Error("ReadMessage 错误:", err)
			return
		}
		log.Debug("收到消息: " + string(message))

		select {
		case <-ctx.Done():
			log.Info("上下文已取消，停止消息读取")
			return
		default:
		}

		event, name, err := events.New(message)
		if err != nil {
			log.Error("事件创建错误:", err)
			continue
		}

		bot.executeCallbacks(ctx, name, event)
	}
}

// executeCallbacks 为给定事件执行已注册的回调。
func (bot *Bot) executeCallbacks(ctx context.Context, eventName events.EventName, event events.IEvent) {
	bot.lock.RLock()
	callbacks := bot.events[eventName]
	bot.lock.RUnlock()

	log.Debug("正在执行事件的回调: " + eventName)
	for _, callback := range callbacks {
		go callback(ctx, event)
	}
}
