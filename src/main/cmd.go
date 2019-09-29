package main

import "fmt"
//import hw "helloworld"
import event "lbsimulator/event"
import "lbsimulator"

func main() {

	es := event.NewEventStream(600 /* req/s */)

	lb := lbsimulator.NewLoadBalancer(10)
	
	//var e *event.Event
	var waitMicros uint64

	//var currTime uint64 = 0
	
	for i := 0 ; i < 10000; i += 1 {
		_, waitMicros = es.NextEvent()

		lb.TickMulti(waitMicros)
		// 50ms, just for giggles
		lb.SubmitRequest(lbsimulator.NewRequest(50000))

		//fmt.Printf("%4d: ", i)
		fmt.Println(lb.String())
	}
	
}
