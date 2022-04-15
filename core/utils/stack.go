package utils

import "sync"

type element struct {
	data interface{}
	next *element
}

type Stack struct {
	lock *sync.Mutex
	head *element
	Size int
}

func (stk *Stack) Push(data interface{}) {
	stk.lock.Lock()

	element := new(element)
	element.data = data
	temp := stk.head
	element.next = temp
	stk.head = element
	stk.Size++

	stk.lock.Unlock()
}

func (stk *Stack) Pop() interface{} {
	if stk.head == nil {
		return nil
	}
	stk.lock.Lock()
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	stk.lock.Unlock()

	return r
}

func (stk *Stack) Head() interface{} {
	return stk.head.data
}

func (stk *Stack) IsEmpty() bool {
	return stk.head == nil
}

func NewStack() *Stack {
	stk := new(Stack)
	stk.lock = &sync.Mutex{}

	return stk
}
