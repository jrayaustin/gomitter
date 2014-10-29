package gomitter

import (
	"errors"
)

type Payload struct {
	StringData string
	IntData    int
}

type Callback func(Payload)

type EventStore map[string]Callback

type Gomitter struct {
	Events EventStore
}

func (g *Gomitter) On(event string, cb Callback) error {

	if g.Events == nil {
		g.Events = make(map[string]Callback)
	}

	_, ok := g.Events[event]

	if !ok {
		g.Events[event] = cb
		return nil
	} else {
		return errors.New("Event already attached")
	}
}

func (g *Gomitter) Emit(event string, payload Payload) error {
	callback, ok := g.Events[event]

	if ok {
		callback(payload)
		return nil
	} else {
		return errors.New("Unknown event " + event)
	}
}
