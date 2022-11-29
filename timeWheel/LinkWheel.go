package core

import (
	"errors"
	"goEasyChat/job"
	"reflect"
)

type LinkWheel struct {
	Head     *ChanelNode
	Last     *ChanelNode
	NodeList []*ChanelNode
	Count    int
}

type ChanelNode struct {
	Value   *LinkChanel
	Next    *ChanelNode
	IsStart bool
}

func CreateWheel(cap int) *LinkWheel {
	link := LinkWheel{}
	for i := 0; i < cap; i++ {
		chanelLink := &LinkChanel{}
		link.PushNode(chanelLink)
	}
	return &link
}

func (linkWheel *LinkWheel) PushNode(value *LinkChanel) {
	if linkWheel.Head == nil {
		node := ChanelNode{Value: value, IsStart: true}
		linkWheel.Head = &node
		linkWheel.Last = &node
		linkWheel.Count = 0
		linkWheel.NodeList = make([]*ChanelNode, 0, 100)
		linkWheel.NodeList = append(linkWheel.NodeList, &node)
	} else {
		node := ChanelNode{Value: value, Next: linkWheel.Head, IsStart: false}
		linkWheel.Last.Next = &node
		linkWheel.Last = &node
		linkWheel.NodeList = append(linkWheel.NodeList, &node)
	}
}

func (linkWheel *LinkWheel) PushValueByIndex(index int, value *job.InfoJob) error {
	if linkWheel.Count+index > len(linkWheel.NodeList) {
		return errors.New("未找到index")
	}
	linkWheel.NodeList[linkWheel.Count+index].Value.Push(value)
	return nil
}

func (linkWheel *LinkWheel) Move() {
	if linkWheel.Head == nil {
		return
	} else {
		value := linkWheel.Head.Value
		if linkWheel.Head.IsStart {
			linkWheel.Count = 0
		}
		linkWheel.Head = linkWheel.Head.Next
		for {
			if info, ok := value.Pop(); ok {
				go func() {
					value := reflect.New(job.TypeRegistry[info.JobName])
					var params []reflect.Value
					for i := range info.ParamArray {
						params = append(params, reflect.ValueOf(info.ParamArray[i]))
					}
					value.MethodByName(info.ExecuteMethod).Call(params)
				}()
			} else {
				break
			}
		}
	}
	linkWheel.Count++
}
