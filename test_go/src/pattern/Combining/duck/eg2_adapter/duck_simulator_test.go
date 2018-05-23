package eg2_adapter

import (
	"testing"
)

func TestSimulator(t *testing.T) {
	var simulate *DuckSimulator = NewDuckSimulator()
	simulate.Simulate()
}
