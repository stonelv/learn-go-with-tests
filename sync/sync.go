package sync

import "sync"

//Counter used to capsule the logic for counting
type Counter struct {
	value int
	mu    sync.Mutex
}

//Inc used to increment the number
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

//Value used to return the current value
func (c *Counter) Value() int {
	return c.value
}
