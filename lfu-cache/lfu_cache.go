package lfu_cache

import (
	"container/heap"
	"sync"
	"sync/atomic"
)

type Item struct {
	Key       string
	Value     string
	UseFreq   atomic.Int32
	HeapIndex int
}

/*
For more understanding - this how map Items stores Item:

Items:
	key: "a" → &Item{
		Key:     "a",
		Value:   "Apple",
		UseFreq: 3,
		HeapIndex: 1,
	}
*/

type LFU struct {
	capacity int
	Items    sync.Map // key : *Item
	Queue    PriorityQueue
}

// NewLFU builds new LFU instance and init PriorityQueue in heap
func NewLFU(capacity int) *LFU {
	cache := &LFU{
		capacity: capacity,
		Items:    sync.Map{},
		Queue:    PriorityQueue{},
	}
	heap.Init(&cache.Queue)

	return cache
}
