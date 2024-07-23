package onebot

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"onebot/events"
)

type Bot struct {
	API    string
	events map[events.EventName][]events.EventCallbackFunc
	lock   sync.RWMutex
	client *websocket.Conn
	botQQ  *int64
	done   chan struct{}
}

// NewBot creates a new Bot instance with the specified API URL.
func NewBot(api string) (*Bot, error) {
	log.SetLevel(log.DebugLevel)
	return &Bot{
		API:    api,
		events: make(map[events.EventName][]events.EventCallbackFunc),
		lock:   sync.RWMutex{},
		done:   nil,
	}, nil
}

// On registers a callback function for a specific event.
func (bot *Bot) On(event events.EventName, callback events.EventCallbackFunc) {
	bot.lock.Lock()
	defer bot.lock.Unlock()
	bot.events[event] = append(bot.events[event], callback)
}

// ListenAndWait connects to the WebSocket server and listens for events.
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

// readMessages continuously reads messages from the WebSocket connection.
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

		event, err := events.New(message)
		if err != nil {
			log.Error("事件创建错误:", err)
			continue
		}

		bot.executeCallbacks(ctx, event)
	}
}

// executeCallbacks executes registered callbacks for the given event.
func (bot *Bot) executeCallbacks(ctx context.Context, event events.IEvent) {
	bot.lock.RLock()
	callbacks := bot.events[event.GetEventName()]
	bot.lock.RUnlock()

	log.Debug("正在执行事件的回调: " + event.GetEventName())
	for _, callback := range callbacks {
		go callback(ctx, event)
	}
}
