package utils

import "container/heap"

type QueueItem struct {
	Key   string
	Value int
}

type PriorityQueue struct {
	Queue       []*QueueItem
	compareFunc func(i, j int) bool
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		Queue: make([]*QueueItem, 0),
	}
}

var _ heap.Interface = (*PriorityQueue)(nil)

func (pq PriorityQueue) Len() int {
	return len(pq.Queue)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.compareFunc != nil {
		return pq.compareFunc(i, j)
	} else {
		return pq.Queue[i].Value < pq.Queue[j].Value
	}

}

func (pq *PriorityQueue) SetCompareFunc(fun func(i, j int) bool) {
	pq.compareFunc = fun
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.Queue[i], pq.Queue[j] = pq.Queue[j], pq.Queue[i]
}

func (pq *PriorityQueue) Push(x any) {
	pq.Queue = append(pq.Queue, x.(*QueueItem))
}

func (pq *PriorityQueue) Pop() any {
	old := pq.Queue
	n := len(old)
	item := old[n-1]
	pq.Queue = old[0 : n-1]
	return item
}
