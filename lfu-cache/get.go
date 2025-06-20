package lfu_cache

import (
	"container/heap"
	"fmt"
)

// Get returns Item.Value
func (c *LFU) Get(key string) string {
	rawItem, exist := c.Items.Load(key)
	if !exist {
		fmt.Println("item doesnt exist")
		return ""
	}

	item := rawItem.(*Item)
	item.UseFreq.Add(1)

	heap.Fix(&c.Queue, item.HeapIndex)

	return item.Value
}
