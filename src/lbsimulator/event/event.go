package lbsimulator

import "math/rand"
//import "fmt"

const microsPerSecond float64 = 1000000

type EventStream struct {
	impl *esImpl
}

type esImpl struct {
	eventRand *rand.Rand
	rate float64
}

type Event struct {}

func NewEventStream(eventsPerSecond uint32) *EventStream {
	eventSource := rand.NewSource(1000)
	eventRand := rand.New(eventSource)

	rate := float64(eventsPerSecond) / microsPerSecond
	impl := esImpl{ eventRand, rate }
	return &EventStream{ &impl }
}

func (es *EventStream) NextEvent() (*Event, uint64) {
	impl := es.impl
	waitMicros := uint64(impl.eventRand.ExpFloat64() / impl.rate)
	return &Event{}, waitMicros
}

