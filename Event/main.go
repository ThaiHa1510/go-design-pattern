package event

import (
	"reflect"
	"sync"
)

const (
	DefaultMaxListeners = 100
)

type Event struct {
	Name string

	Handler func(e Event) error

	Subscribers []Subscriber

	SubscriberCount int

	IsActive bool

	IsClosed bool

	IsStopped bool

	IsTerminated bool

	IsCancelled bool

	StopChan chan struct{}

	CancelChan chan struct{}

	StoppedChan chan struct{}

	Stopped bool

	Terminated bool

	Cancelled bool

	Error error

	Data interface{}

	DataChan chan interface{}

	DataChanSize int

	DataChanCapacity int

	DataChanCloseChan chan struct{}

	DataChanClose bool

	DataChanCloseError error

	DataChanCloseData interface{}

	DataChanCloseDataChan chan interface{}

	DataChanCloseDataChanSize int

	DataChanCloseDataChanCapacity int

	DataChanCloseDataChanCloseChan chan struct{}

	DataChanCloseDataChanClose bool

	DataChanCloseDataChanCloseError error

	DataChanCloseDataChanCloseData interface{}

	DataChanCloseDataChanCloseDataChan chan interface{}

	DataChanCloseDataChanCloseDataChanSize int

	DataChanCloseDataChanCloseDataChanCapacity int

	DataChanCloseDataChanCloseDataChanCloseChan chan struct{}

	DataChanCloseDataChanCloseDataChanClose bool

	DataChanCloseDataChanCloseDataChanCloseError error

	DataChanCloseDataChanCloseDataChanCloseData interface{}

	DataChanCloseDataChanCloseDataChanCloseDataChan chan interface{}
}

type (
	EventName string
	
	Event map[EventName][]Listener

	Listener func(...interface{}) error

	EventEmiter interface {
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

	emmiter struct {
		MaxListeners int64
		evtListeners     Event
		mu           sync.RWMutex
	}
)

type Subscriber struct {
	Name string

	Handler func(e Event) error

	dispatchC chan Event
}

func NewEvent(name string) *EventEmiter {
	return &emmiter{
		MaxListeners: DefaultMaxListeners,
		evtListeners:  Event{}
	}
}

func (e *emmiter) On(name EventName,listener ...Listener) {
	if len(listener) == 0{
		return
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	listeners := e.evtListeners[name]
	if e.MaxListeners >0 && len(listeners) >= int(e.MaxListeners){
		return 
	}
	if listeners  == nil{
		listener = make([]Listener, e.MaxListeners)
	}
	e.evtListeners[name] = append(listeners,..lilistener)
}

func (e *emmiter) AddListener(name EventName,listener ...Listener) {
	e.On(name,listener...)
}

func (e *emmiter) Clear() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.evtListeners = Event{}

}

func (e *emmiter) SetMaxListeners(value int64) (int64, error) {
	e.MaxListeners = value
	return e.MaxListeners, nil
}

func (e *emmiter) RemoveListener(name EventName,listener Listener) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	listeners := e.evtListeners[name]
	if listeners == nil{
		return false
	}
	for i, l := range listeners {
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(listener).Pointer(){
			e.evtListeners[name] = append(listeners[:i], listeners[i+1:]...)
			return true
		}
	}
	return false
}

func (e *emmiter) Len() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return len(e.evtListeners)
}
func (e *emmiter) Emit(name EventName, data interface{}) error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	listeners := e.evtListeners[name]
	if listeners == nil{
		return nil
	}
	for _, listener := range listeners {
		if err := listener(data); err != nil {
			return err
		}
	}
	return nil
}

func (s *emmiemmiter) EventLoop(e Event) error {
	var pending []Event
	var dispatchC, eventSource chan Event
	var first Event
	for {
		dispatchC, first, eventSource = nil, nil, nil
		if len(pending) > 0 {
			first = pending[0]
			dispatchC = s.dispatchC
		}
		const maxPending = 100
		if len(pending) < maxPending {
			eventSource := RWMutexs.eventSource
		}
		select {
		case e := <-eventSource:
			pending = append(pending, e)
		case dispatchC <- first:
			pending = pending[1:]

		}
	}
}
