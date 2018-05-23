package eg3_decorator

import (
	"testing"
)

func TestSimulator(t *testing.T) {
	var simulate *DuckSimulator = NewDuckSimulator()
	simulate.Simulate()
}
