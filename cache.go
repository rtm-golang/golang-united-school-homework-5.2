package cache

import "time"

type Cache struct {
	ValueMap    map[string]string
	DeadlineMap map[string]time.Time
}

func NewCache() Cache {
	return Cache{ValueMap: make(map[string]string), DeadlineMap: make(map[string]time.Time)}
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.ValueMap[key]
	if ok {
		deadline, ok := c.DeadlineMap[key]
		if ok {
			if time.Now().After(deadline) {
				delete(c.ValueMap, key)
				delete(c.DeadlineMap, key)
				return "", false
			}
		}
	}
	return value, ok
}

func (c *Cache) Put(key, value string) {
	c.ValueMap[key] = value
	_, ok := c.DeadlineMap[key]
	if ok {
		delete(c.DeadlineMap, key)
	}
}

func (c *Cache) Keys() []string {
	currentTime := time.Now()
	for k, v := range c.DeadlineMap {
		if currentTime.After(v) {
			delete(c.ValueMap, k)
			delete(c.DeadlineMap, k)
		}
	}
	valueList := make([]string, len(c.ValueMap))
	for _, v := range c.ValueMap {
		valueList = append(valueList, v)
	}
	return valueList
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.ValueMap[key] = value
	c.DeadlineMap[key] = deadline
}
