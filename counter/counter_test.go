package counter

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup

		for i := 0; i < wantedCount; i++ {
			wg.Add(1)

			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assert.Equal(t, wantedCount, counter.Value())
	})
}
