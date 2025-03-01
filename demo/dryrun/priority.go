package main

import (
	"container/heap"
	"fmt"
	"math/rand/v2"

	"github.com/elin-croft/go-utils/framework/utils"
)

func main() {
	pd := utils.NewPriorityQueue()
	pd.SetCompareFunc(func(i, j int) bool {
		return pd.Queue[i].Value > pd.Queue[j].Value || (pd.Queue[i].Value == pd.Queue[j].Value && pd.Queue[i].Key > pd.Queue[j].Key)
	})
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear:":  8,
	}
	for k, v := range items {
		heap.Push(pd, &utils.QueueItem{Key: k, Value: v})
	}
	heap.Push(pd, &utils.QueueItem{Key: "orange", Value: int(rand.IntN(10))})
	for pd.Len() > 0 {
		item := heap.Pop(pd).(*utils.QueueItem)
		fmt.Printf("item: %v\n", item)
	}
}
