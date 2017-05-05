package round_robin

type RoundRobin struct {
	Index int
	Keys  map[int]interface{}
	_Prev string
}
