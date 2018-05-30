package eg2

import "testing"

func TestAll(t *testing.T) {
	subject := NewSubject()

	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")

	subject.RegisterObserver(reader1)
	subject.RegisterObserver(reader2)
	subject.RegisterObserver(reader3)

	subject.UpdateContext("observer mode")
}
