package lbsimulator

//import "fmt" 
//import "math/rand"
import "strings"
import "math"
import "strconv"
//import event "lbsimulator/event"

// TODO: several load balancer strategies (random, round-robin, least requests)

type machine struct {
	requests [](*Request)
}

func (m *machine) submitRequest(r *Request) {
	m.requests = append(m.requests, r)
}

func (m *machine) numRequests() int {
	return len(m.requests)
}

func (m *machine) tickMulti(micros uint64) {
	i := 0
	requests := m.requests

	// in-place filter:
	for _, r := range requests {
		if r.remainingDuration > float64(micros) {
			r.remainingDuration -= float64(micros)
			requests[i] = r
			i += 1
		}
	}
	m.requests = requests[:i]

}

type LoadBalancer interface {
	choose() *machine
	SubmitRequest(*Request)
	Tick()
	TickMulti(uint64)
	String() string
}


type leastConnectionsLB struct {
	machines []machine
}

func (lb *leastConnectionsLB) choose() (*machine) {
	leastConnections := math.MaxInt32
	var leastConnsMachine *machine

	for idx := range lb.machines {
		m := &lb.machines[idx]
		l := len(m.requests)
		// TODO: this is supposed to do round-robin load-balancing
		// if all LBs have same number of requests active
		if l < leastConnections {
			leastConnections = l
			leastConnsMachine = m
		}
	}

	return leastConnsMachine
}	

func (lb *leastConnectionsLB) TickMulti(t uint64) {
	for idx := range lb.machines {
		lb.machines[idx].tickMulti(t)
	}
}

func (lb *leastConnectionsLB) Tick() {
	lb.TickMulti(1)
}

func NewLoadBalancer (size int) LoadBalancer {
	machines := make([](machine), 10)
	return &leastConnectionsLB{ machines }
}	

func (lb *leastConnectionsLB) SubmitRequest(r *Request) {
	m := lb.choose()
	m.submitRequest(r)
}

func (lb *leastConnectionsLB) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	machines := lb.machines
	if len(machines) > 0 {
		sb.WriteString(strconv.Itoa(machines[0].numRequests()))
		machines = machines[1:]
	}
	for _, m := range machines {
		//fmt.Println(m)
		sb.WriteString(",")
		sb.WriteString(strconv.Itoa(m.numRequests()))
	}
	sb.WriteString("]")
	return sb.String()
}

type Request struct {
	originalDuration float64
	remainingDuration float64
}

func NewRequest(duration float64) *Request {
	return &Request{ duration, duration }
}
