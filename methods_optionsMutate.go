package cache

import "time"

func (c *Cache) SetMaxValues(n int) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.mutated = true
	c.options.MaxValues = n
}

func (c *Cache) SetExpiryDuration(d time.Duration) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.mutated = true
	c.options.ExpiryDuration = d
}

func (c *Cache) SetCleanDuration(d time.Duration) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.mutated = true
	c.options.CleanDuration = d
}

func (c *Cache) SetCleanMaxValuesPerSweep(n int) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.mutated = true
	c.options.CleanMaxValuesPerSweep = n
}

func (c *Cache) SetPreDeletionFunc(p CallbackFunc) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.mutated = true
	c.options.PreDeletionFunc = p
}

func (c *Cache) SetPostDeletionFunc(p CallbackFunc) {
	defer c.mutex.Unlock()
	c.mutex.Lock()

	c.mutated = true
	c.options.PostDeletionFunc = p
}