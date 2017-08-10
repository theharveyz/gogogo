package round_robin

import (
	"errors"
	"sync"
)

type RoundRobin struct {
	sync.RWMutex // 要求所有的方法的receiver全部为pointer
	index        int
	length       int
	nodes        []interface{}
}

func (rr *RoundRobin) SeekIndex() int {
	defer rr.RUnlock()
	rr.RLock()
	return rr.index
}

func (rr *RoundRobin) GetNode() (interface{}, error) {
	defer rr.Unlock()
	rr.Lock()
	if rr.length == 0 {
		return nil, errors.New("nodes are empty")
	}
	// 当长度为1时, 则自动返回,不经过下面的计算过程
	if rr.length == 1 {
		return rr.nodes[0], nil
	}

	rr.index = (rr.index + 1) % rr.length
	// 转换下标
	i := rr.index - 1
	if i < 0 {
		i = rr.length - 1
	}
	return rr.nodes[i], nil
}

func (rr *RoundRobin) AddNode(nodes ...interface{}) {
	defer rr.Unlock()
	rr.Lock()
	rr.nodes = append(rr.nodes, nodes...)
	rr.length += len(nodes)
}

func (rr *RoundRobin) RemoveNode(index int) error {
	defer rr.Unlock()
	rr.Lock()
	// length 不能小于1
	if rr.length < 1 {
		return errors.New("cound not be remove")
	}
	if index >= rr.length {
		return errors.New("index gt nodes length")
	}
	rr.nodes = append(rr.nodes[:index], rr.nodes[(index+1):]...)
	rr.length -= 1
	return nil
}
