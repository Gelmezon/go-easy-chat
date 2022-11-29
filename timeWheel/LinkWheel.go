package core

import (
	"errors"
	"fmt"
	"time"
)

type LinkWheel struct {
	Head     *ChanelNode
	Last     *ChanelNode
	NodeList []*ChanelNode
}

type ChanelNode struct {
	Value *LinkChanel
	Next  *ChanelNode
}

func (linkWheel *LinkWheel) PushNode(value *LinkChanel) {
	if linkWheel.Head == nil {
		node := ChanelNode{Value: value}
		linkWheel.Head = &node
		linkWheel.Last = &node
	} else {
		node := ChanelNode{Value: value, Next: linkWheel.Head}
		linkWheel.Last.Next = &node
		linkWheel.Last = &node
	}
}

func (linkWheel *LinkWheel) PushValueByIndex(index int, value int) error {
	if index > len(linkWheel.NodeList) {
		return errors.New("未找到index")
	}
	linkWheel.NodeList[index].Value.Push(value)
	return nil
}

func (linkWheel *LinkWheel) Move() {
	if linkWheel.Head == nil {
		return
	} else {
		value := linkWheel.Head.Value
		linkWheel.Head = linkWheel.Head.Next
		for {
			if value, ok := value.Pop(); ok {
				go fmt.Println("执行任务 ", value, " 现在事件", time.Now().Format("2006-01-02 15:04:05"))
			} else {
				break
			}
		}
	}
}
