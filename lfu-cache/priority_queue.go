package lfu_cache

// PriorityQueue is a slice of Item
type PriorityQueue []*Item

// Less checks if i UseFreq < j UseFreq
func (pq PriorityQueue) Less(i, j int) bool {
	iFreq := pq[i].UseFreq.Load()
	jFreq := pq[j].UseFreq.Load()

	return iFreq < jFreq
}

// Len returns length of queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Swap swapping a pair of Item in PriorityQueue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].HeapIndex = i
	pq[j].HeapIndex = j
}

// Push adding new Item to PriorityQueue
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.HeapIndex = len(*pq)
	*pq = append(*pq, item)
}

// Pop delete and return least frequently used (min priority) element
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.HeapIndex = -1
	*pq = old[0 : n-1]

	return item
}
