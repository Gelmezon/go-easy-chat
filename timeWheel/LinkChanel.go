package core

type LinkChanel struct {
	Head *Node
	Last *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (linkChanel *LinkChanel) Push(value int) {
	if linkChanel.Head == nil {
		node := Node{Value: value}
		linkChanel.Head = &node
		linkChanel.Last = &node
	} else {
		node := Node{Value: value}
		linkChanel.Last.Next = &node
		linkChanel.Last = &node
	}
}

func (linkChanel *LinkChanel) Pop() (int, bool) {
	if linkChanel.Head == nil {
		return 0, false
	} else {
		value := linkChanel.Head.Value
		linkChanel.Head = linkChanel.Head.Next
		return value, true
	}
}
