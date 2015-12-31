package utils

type Node struct {
	Value interface{}
	next  *Node
}

func NewNode() *Node {
	return new(Node)
}

func (self *Node) Next() (*Node, bool) {
	if self.next == nil {
		return nil, false
	}
	return self.next, true
}
