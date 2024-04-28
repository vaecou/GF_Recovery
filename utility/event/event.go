package event

import (
	"context"
	"sync"
)

type EventFunc func(ctx context.Context, args ...interface{})

type sEvent struct {
	sync.Mutex
	list map[string][]EventFunc // 所有事件的列表
}

var event *sEvent

// Event 事件实例
func Event() *sEvent {
	if event == nil {
		event = &sEvent{
			list: make(map[string][]EventFunc),
		}
	}
	return event
}

// Call 回调一个分组的事件
func (e *sEvent) Call(group string, ctx context.Context, args ...interface{}) {
	if events, ok := e.list[group]; ok {
		for _, f := range events {
			f(ctx, args...)
		}
	}
}
