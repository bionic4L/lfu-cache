package lfu_cache

import (
	. "lfu-cache/lfu-cache"
	"testing"
)

func TestLFUCache(t *testing.T) {
	type testCase struct {
		name       string
		capacity   int
		actions    func(lfu *LFU)
		getKey     string
		wantResult string
	}

	tests := []testCase{
		{
			name:     "Set and Get single item",
			capacity: 2,
			actions: func(c *LFU) {
				c.Set("a", "apple")
			},
			getKey:     "a",
			wantResult: "apple",
		},
		{
			name:     "Evict LFU item",
			capacity: 2,
			actions: func(c *LFU) {
				c.Set("a", "apple")  // freq = 0
				c.Set("b", "banana") // freq 0
				c.Get("a")           // freq a = 1
				c.Set("c", "cherry") // b should be evicted (freq = 0)
			},
			getKey:     "b",
			wantResult: "",
		},
		{
			name:     "Update value and frequency",
			capacity: 2,
			actions: func(c *LFU) {
				c.Set("a", "apple")
				c.Set("a", "apricot") // update value, freq must stay still
			},
			getKey:     "a",
			wantResult: "apricot",
		},
		{
			name:     "Evict when all have same frequency must remove oldest item",
			capacity: 2,
			actions: func(c *LFU) {
				c.Set("a", "apple")
				c.Set("b", "banana")
				c.Set("c", "cherry") // should evict a
			},
			getKey:     "a",
			wantResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewLFU(tt.capacity)
			tt.actions(cache)

			got := cache.Get(tt.getKey)
			if got != tt.wantResult {
				t.Errorf("Get(%q) = %q, want %q", tt.getKey, got, tt.wantResult)
			}
		})
	}
}
