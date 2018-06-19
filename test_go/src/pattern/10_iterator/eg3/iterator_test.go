package eg3

import "testing"

func TestAll(t *testing.T) {
	var aggregate Aggregate = NewNumbers(1, 50)

	IteratorPrint(aggregate.Iterator())
}
