package events

import "errors"

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

var ErrAlreadyRegistered = errors.New("handler already registered")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
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
