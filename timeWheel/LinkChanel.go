package core

import "goEasyChat/job"

type LinkChanel struct {
	Head *Node
	Last *Node
}

type Node struct {
	Value *job.InfoJob
	Next  *Node
}

func (linkChanel *LinkChanel) Push(value *job.InfoJob) {
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

func (linkChanel *LinkChanel) Pop() (*job.InfoJob, bool) {
	if linkChanel.Head == nil {
		return nil, false
	} else {
		value := linkChanel.Head.Value
		linkChanel.Head = linkChanel.Head.Next
		return value, true
	}
}
