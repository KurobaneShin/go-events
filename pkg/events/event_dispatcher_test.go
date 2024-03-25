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

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) HandleEvent(event Event) {
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
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.dispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(suite.dispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.dispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.dispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_With_SameHandler() {
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.handlers[suite.event.GetName()]))

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrAlreadyRegistered, err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
