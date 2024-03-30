package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandler interface {
	HandleEvent(event Event, wg *sync.WaitGroup)
}

type EventPublisher interface {
	Dispatch(event Event) error
	Register(eventName string, handler EventHandler) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear() error
}
