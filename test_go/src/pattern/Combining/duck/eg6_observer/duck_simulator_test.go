package eg6_observer

import (
	"testing"
)

func TestSimulator(t *testing.T) {
	var simulate *DuckSimulator = NewDuckSimulator()
	var duckFactory AbstractDuckFactory = new(CountingDuckFactory)
	simulate.Simulate(duckFactory)
}
