package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	EventType string
	Payload   string
	CreatedAt time.Time
}
type service struct {
	subs map[string][]chan Event
	mu   sync.RWMutex
}

func NewBusService() *service {
	return &service{
		subs: make(map[string][]chan Event),
	}
}
func (svc *service) Pub(e Event) {
	svc.mu.RLock()
	defer svc.mu.RUnlock()
	for _, ch := range svc.subs[e.EventType] {
		ch <- e
	}
}
func (svc *service) Sub(e_type string, ch chan Event) {
	svc.mu.Lock()
	defer svc.mu.Unlock()
	svc.subs[e_type] = append(svc.subs[e_type], ch)
}
func main() {
	bus := NewBusService()

	ch1 := make(chan Event)
	ch2 := make(chan Event)
	ch3 := make(chan Event)

	bus.Sub("ev1", ch1)
	bus.Sub("ev2", ch2)
	bus.Sub("ev1", ch3)

	go func() {
		for c := range ch1 {
			fmt.Println("ch1", c.EventType, c.Payload, c.CreatedAt)
		}
	}()
	go func() {
		for c := range ch2 {
			fmt.Println("ch2", c.EventType, c.Payload, c.CreatedAt)
		}
	}()
	go func() {
		for c := range ch3 {
			fmt.Println("ch3", c.EventType, c.Payload, c.CreatedAt)
		}
	}()
	bus.Pub(Event{EventType: "ev1", Payload: "This is Event 1", CreatedAt: time.Now()})
	time.Sleep(time.Second)
	bus.Pub(Event{EventType: "ev1", Payload: "This is Event 1", CreatedAt: time.Now()})
	time.Sleep(time.Second)
	bus.Pub(Event{EventType: "ev2", Payload: "This is Event 2", CreatedAt: time.Now()})
	time.Sleep(time.Second * 3)
	defer func() {
		close(ch1)
		close(ch2)
		close(ch3)
	}()
}
