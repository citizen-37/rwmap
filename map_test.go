package rwmap_test

import (
	"testing"

	"github.com/citizen-37/rwmap"
)

func TestMap(t *testing.T) {
	cache := rwmap.New[int, int]()

	t.Run("reader", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 10000; i++ {
			cache.Get(i)
		}
	})

	t.Run("writer", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 10000; i++ {
			cache.Set(i, i)
		}
	})

	t.Run("list reader", func(t *testing.T) {
		t.Parallel()
		for i := 0; i < 10000; i++ {
			cache.List()
		}
	})
}
