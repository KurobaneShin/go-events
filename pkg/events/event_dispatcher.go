package events

import "errors"

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

ErrAlreadyRegistered := errors.New("Handler already registered")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (eventDispatcher *EventDispatcher) Register(eventName string, handler EventHandler) error {
	if _, ok := eventDispatcher.handlers[eventName]; !ok {
		for _, registeredEvent := range eventDispatcher.handlers[eventName] {
			if registeredEvent == handler {
				return ErrAlreadyRegistered
			}
		}
	}

	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], handler)

	return nil
}
