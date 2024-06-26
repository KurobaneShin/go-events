package events

import (
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

var ErrAlreadyRegistered = errors.New("handler already registered")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (eventDispatcher *EventDispatcher) Dispatch(event Event) error {
	if handlers, ok := eventDispatcher.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		wg.Add(len(handlers))
		for _, handler := range handlers {
			go handler.HandleEvent(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (eventDispatcher *EventDispatcher) Register(eventName string, handler EventHandler) error {
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for _, registeredEvent := range eventDispatcher.handlers[eventName] {
			if registeredEvent == handler {
				return ErrAlreadyRegistered
			}
		}
	}

	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], handler)

	return nil
}

func (eventDispatcher *EventDispatcher) Unregister(eventName string, handler EventHandler) error {
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for i, registeredEvent := range eventDispatcher.handlers[eventName] {
			if registeredEvent == handler {
				eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName][:i], eventDispatcher.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return nil
}

func (eventDispatcher *EventDispatcher) Clear() {
	eventDispatcher.handlers = make(map[string][]EventHandler)
}

func (eventDispatcher *EventDispatcher) Has(eventName string, handler EventHandler) bool {
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for _, registeredHandler := range eventDispatcher.handlers[eventName] {
			if registeredHandler == handler {
				return true
			}
		}
	}

	return false
}
