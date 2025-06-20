package lfu_cache

// getLenOfCache return length of Items map
func (c *LFU) getLenOfCache() int {
	var length int

	c.Items.Range(func(_, _ any) bool {
		length++
		return true
	})

	return length
}
