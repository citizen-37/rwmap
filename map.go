package rwmap

import "sync"

type RWMap[K comparable, V any] struct {
	mx   *sync.RWMutex
	list map[K]V
}

func New[K comparable, V any]() *RWMap[K, V] {
	return &RWMap[K, V]{
		mx:   &sync.RWMutex{},
		list: make(map[K]V),
	}
}

func (c *RWMap[K, V]) Set(key K, val V) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.list[key] = val
}

func (c *RWMap[K, V]) Unset(key K) {
	c.mx.Lock()
	defer c.mx.Unlock()

	delete(c.list, key)
}

func (c *RWMap[K, V]) Get(key K) V {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.list[key]
}

func (c *RWMap[K, V]) List() []V {
	c.mx.RLock()
	defer c.mx.RUnlock()

	res := make([]V, 0, len(c.list))
	for _, v := range c.list {
		res = append(res, v)
	}

	return res
}
