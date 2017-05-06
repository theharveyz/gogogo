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

func (rr *RoundRobin) GetNode() int {
	defer rr.mutex.Unlock()
	rr.mutex.Lock()
	rr.index = (rr.index + 1) % rr.length
	return rr.index
}

func (rr *RoundRobin) AddNode(nodes ...interface{}) {
	defer rr.mutex.Unlock()
	rr.mutex.Lock()
	rr.nodes = append(rr.nodes, nodes...)
}

func (rr *RoundRobin) RemoveNode(index int) error {
	defer rr.mutex.Unlock()
	rr.mutex.Lock()
	// if node := rr.nodes[index]; node {
	// 	return error
	// }
	rr.nodes = append(rr.nodes[:index], rr.nodes[(index+1):]...)
	return nil
}
