package lfu_cache

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type Item struct {
	Key              string
	Value            string
	UseFreq          atomic.Int32
	FreqListPosition *list.Element
}

/*
Пример как выглядят структуры хранения данных для наглядности.

Items:
	key: "a" → &Item{
		Key:     "a",
		Value:   "Apple",
		UseFreq: 3,
		FreqListPosition: &list.Element{...}, // указатель(!) на элемент в списке
	}

FreqList:
	FreqList = {
		1: [item1, item2],
		2: [item3],
		3: [item4, item5],
	}
*/

type LFU struct {
	capacity int32
	minFreq  int32
	Items    sync.Map // key : *Item
	FreqList sync.Map // freq (int32) → *list.List
}

func NewLFU(capacity int32) *LFU {
	cache := &LFU{
		capacity: capacity,
		Items:    sync.Map{},
		FreqList: sync.Map{},
	}

	return cache
}

func (c *LFU) incrementFrequency(item *Item) {
	oldFreq := item.UseFreq.Load()
	newFreq := item.UseFreq.Add(1)

	if gotFreqList, exist := c.FreqList.Load(oldFreq); exist {
		oldList := gotFreqList.(*list.List)
		oldList.Remove(item.FreqListPosition)

		if gotFreqList.(*list.List).Len() == 0 {
			c.FreqList.Delete(oldFreq)

			if c.minFreq == oldFreq {
				atomic.AddInt32(&c.minFreq, 1)
			}
		}
	}

	var newList *list.List

	if gotList, exist := c.FreqList.Load(newFreq); !exist {
		newList = list.New()
		c.FreqList.Store(newFreq, newList)
	} else {
		newList = gotList.(*list.List)
	}

	item.FreqListPosition = newList.PushFront(item)
}

func (c *LFU) Set(key, value string) {
	var freqList *list.List

	if rawItem, exist := c.Items.Load(key); exist {
		item := rawItem.(*Item)
		item.Value = value
		c.incrementFrequency(item)
		return
	}

	if c.getLenOfCache() >= c.capacity {
		c.removeElement()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}
	item.UseFreq.Store(1)

	if gotFreqList, exist := c.FreqList.Load(1); !exist {
		freqList = list.New()
		c.FreqList.Store(1, freqList)
	} else {
		freqList = gotFreqList.(*list.List)
	}

	elem := freqList.PushFront(item)

	item.FreqListPosition = elem

	c.Items.Store(key, item)

	c.minFreq = 1
}

func (c *LFU) Get(key string) *Item {
	rawItem, exist := c.Items.Load(key)
	if !exist {
		return nil
	}

	item := rawItem.(*Item)

	c.incrementFrequency(item)

	return item
}

func (c *LFU) removeElement() {
	gotRawList, exist := c.FreqList.Load(c.minFreq)
	if !exist {
		return
	}

	gotList := gotRawList.(*list.List)

	worseElement := gotList.Back()

	if worseElement == nil {
		return
	}

	gotList.Remove(worseElement)

	c.Items.Delete(worseElement.Value.(*Item).Key)

	if gotList.Len() == 0 {
		c.FreqList.Delete(c.minFreq)
	}
}

func (c *LFU) getLenOfCache() int32 {
	var length int32

	c.Items.Range(func(_, _ any) bool {
		length++
		return true
	})

	return length
}
