package lfu_cache

import "container/heap"

// Set creates or updates Item
func (c *LFU) Set(key, value string) {
	if rawItem, exist := c.Items.Load(key); exist {
		item := rawItem.(*Item)
		item.Value = value
		heap.Fix(&c.Queue, item.HeapIndex)

		return
	}

	if c.getLenOfCache() >= c.capacity {
		removedElem := heap.Pop(&c.Queue).(*Item)
		c.Items.Delete(removedElem.Key)
	}

	item := &Item{
		Key:   key,
		Value: value,
	}
	item.UseFreq.Store(0)

	heap.Push(&c.Queue, item)
	c.Items.Store(key, item)
}
