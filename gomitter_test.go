package gomitter

import (
	"testing"
)

func TestCorrectEventData(t *testing.T) {
	g := &Gomitter{}
	stringData := "stringData"
	intData := 5
	payload := Payload{stringData, intData}

	cb := func(p Payload) {
		if p.StringData != stringData {
			t.Error("String data does not match")
		}

		if p.IntData != intData {
			t.Error("Int data does not match")
		}
	}

	err := g.On("foo", cb)

	if err != nil {
		t.Error(err)
	}

	g.Emit("foo", payload)
}

func TestDuplicateEvent(t *testing.T) {
	g := &Gomitter{}
	stringData := "stringData"
	intData := 5
	payload := Payload{stringData, intData}

	cb := func(p Payload) {
		if p.StringData != stringData {
			t.Error("String data does not match")
		}

		if p.IntData != intData {
			t.Error("Int data does not match")
		}
	}

	err := g.On("foo", cb)

	if err != nil {
		t.Error(err)
	}

	err2 := g.On("foo", cb)

	if err2 == nil {
		t.Error("An error should have occured")
	}

	g.Emit("foo", payload)
}
