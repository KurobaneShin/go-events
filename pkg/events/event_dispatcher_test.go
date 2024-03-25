package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	name    string
	payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct{}

func (h *TestEventHandler) Handle(event Event) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = TestEvent{name: "test", payload: "test"}
	suite.event2 = TestEvent{name: "test2", payload: "test2"}
	suite.handler = TestEventHandler{}
	suite.handler2 = TestEventHandler{}
	suite.dispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(suite.T(), true)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
