package gomitter

import (
	"sync"
	"testing"
)

func TestCorrectEventData(t *testing.T) {
	g := &Gomitter{}
	stringData := "stringData"
	intData := 5
	payload := Payload{stringData, intData}
	cb := make(chan Payload)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		select {
		case p := <-cb:
			if p.StringData != stringData {
				t.Error("String data does not match")
			}

			if p.IntData != intData {
				t.Error("Int data does not match")
			}
			wg.Done()
		}
	}()

	err := g.On("foo", cb)

	if err != nil {
		t.Error(err)
	}

	g.Emit("foo", payload)
	wg.Wait()

}

func TestDuplicateEvent(t *testing.T) {
	g := &Gomitter{}
	stringData := "stringData"
	intData := 5
	payload := Payload{stringData, intData}
	cb := make(chan Payload)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		select {
		case p := <-cb:
			if p.StringData != stringData {
				t.Error("String data does not match")
			}

			if p.IntData != intData {
				t.Error("Int data does not match")
			}
			wg.Done()
		}
	}()

	err := g.On("foo", cb)

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	err2 := g.On("foo", cb)

	if err2 == nil {
		t.Error("An error should have occured")
	}

	g.Emit("foo", payload)
	wg.Wait()
}

func TestDetatchEvent(t *testing.T) {
	g := &Gomitter{}
	stringData := "stringData"
	intData := 5
	payload := Payload{stringData, intData}
	cb := make(chan Payload)

	err := g.On("foo", cb)

	if err != nil {
		t.Error(err)
	}

	sameCb, detatchErr := g.Detatch("foo")

	if detatchErr != nil {
		t.Error(detatchErr)
	}

	if sameCb != cb {
		t.Error("Callback not the same")
	}

	emitErr := g.Emit("foo", payload)

	if emitErr == nil {
		t.Error("Should not be able to emit a detatched event")
	}
}
