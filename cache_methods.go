package cache

func (c *Cache) Set(key string, value interface{}) {
	c.storage[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.storage[key]
	if !ok {
		return nil, false
	}
	return value, ok
}

func (c *Cache) Delete(key string) (interface{}, bool) {
	value, ok := c.storage[key]
	if ok {
		delete(c.storage, key)
		return value, ok
	} else {
		return nil, ok
	}
}
