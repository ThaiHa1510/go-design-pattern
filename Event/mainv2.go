package main

import (
	"fmt"
	"reflect"
	"sync"
)

type EventName string
type Listener func(...interface{}) error
type Event map[EventName][]Listener
type EventEmiter interface {
	// AddListener is a alias for On.

	AddListener(EventName, ...Listener) error

	On(EventName, ...Listener) error

	One(EventName, Listener) error

	Emit(EventName, interface{}) error

	Clear()

	SetMaxListeners(int64) (int64, error)

	RemoveListener(EventName, Listener) bool

	Len() int
}
type EventEmitter struct {
	evtListeners Event
	mu           sync.Mutex
	MaxListeners int64
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		evtListeners: Event{},
		MaxListeners: 100,
	}
}

func (e *EventEmitter) AddListener(name EventName, listener ...Listener) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	listeners := e.evtListeners[name]
	if e.MaxListeners > 0 && len(listeners) >= int(e.MaxListeners) {
		return nil
	}
	if listeners == nil {
		listeners = []Listener{}
	}
	e.evtListeners[name] = append(listeners, listener...)
	fmt.Println("Run AddListener", listener)
	return nil
}

func (e *EventEmitter) On(name EventName, listener ...Listener) error {
	return e.AddListener(name, listener...)
}

func (e *EventEmitter) Emit(name EventName, data ...interface{}) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	listeners := e.evtListeners[name]
	if listeners == nil {
		return nil
	}
	for _, listener := range listeners {
		fmt.Println("Run Emit", listener)
		if listener != nil {
			if err := listener(data...); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *EventEmitter) RemoveListener(name EventName, listener Listener) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	listeners := e.evtListeners[name]
	if listeners == nil {
		return false
	}
	for i, l := range listeners {
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(listener).Pointer() {
			e.evtListeners[name] = append(listeners[:i], listeners[i+1:]...)
			return true
		}
	}
	return false
}

func (e *EventEmitter) Len() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	return len(e.evtListeners)
}
func (e *EventEmitter) Clear() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.evtListeners = Event{}
}

func (e *EventEmitter) SetMaxListeners(value int64) (int64, error) {
	e.MaxListeners = value
	return e.MaxListeners, nil
}

// eventLoop processes adding and removing listeners
func (e *EventEmitter) eventLoop() {

} // eventLoop
func main() {
	emitter := NewEventEmitter()

	// Create a channel for listening
	eventCh := make(chan interface{})
	defer close(eventCh)
	listener := func(event ...interface{}) error {
		fmt.Println("Received event:", event)
		return fmt.Errorf("error")
	}
	// Add listener
	emitter.On("data", listener)

	// Goroutine to listen for events
	go func() {
		for event := range eventCh {
			fmt.Println("Received event:", event)
		}
	}()

	// Emit an event
	eventData := "some data"
	emitter.Emit("data", eventData)

	// Remove listener
	//emitter.RemoveListener("data", listener)
	//emitter.Emit("data", eventData)
	// Clear all listeners
	emitter.Clear()
	emitter.Emit("data", eventData)
}
