package round_robin

import (
	"sync"
)

type RoundRobin struct {
	index  int
	length int
	nodes  []interface{}
	mutex  *sync.RWMutex
}

func (rr *RoundRobin) GetNode() {

}

func (rr *RoundRobin) AddNodes(node interface{}) {

}

func (rr *RoundRobin) RemoveNode(index int) {

}
