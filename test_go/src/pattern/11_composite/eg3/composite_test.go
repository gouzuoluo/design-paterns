package eg3

import "testing"

func TestAll(t *testing.T) {

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	c0 := NewComponent(CompositeNode, "c3")
	c1 := NewComponent(CompositeNode, "c1")
	c1.AddChild(c0)
	c1.AddChild(l1)

	c2 := NewComponent(CompositeNode, "c2")
	c2.AddChild(l2)
	c2.AddChild(l3)

	root := NewComponent(CompositeNode, "root")
	root.AddChild(c1)
	root.AddChild(c2)

	root.Print("")
}
