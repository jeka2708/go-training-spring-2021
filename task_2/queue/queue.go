package queue

import (
	"fmt"
	"reflect"
)

//Node this is a Queue element.
type Node struct {
	next  *Node
	Value interface{}
}

//Queue is a data structure that.
type Queue struct {
	Head *Node
	tail *Node
	Size int
	len  int
}

//NewQueue - Creates a new queue
func NewQueue(size int, val ...interface{}) (Queue, error) {
	var q = Queue{
		Size: size,
	}
	for _, v := range val {
		e := q.Enqueue(v)
		if e != nil {
			return Queue{}, e
		}
	}
	return q, nil
}

//Enqueue - Add an element to the end of the queue.
func (q *Queue) Enqueue(val interface{}) error {
	if q.Size == q.len {
		return fmt.Errorf("queue is full")
	}
	if !q.IsEmpty() {
		currentType := reflect.TypeOf(q.Head.Value)
		addType := reflect.TypeOf(val)
		if currentType != addType {
			return fmt.Errorf("types don`t equals %s and %s", currentType, addType)
		}
	}

	addNode := &Node{
		next:  q.tail,
		Value: val,
	}
	q.tail = addNode
	element := q.tail
	for element.next != nil {
		element = element.next
	}
	q.Head = element
	q.len++
	return nil
}

//Dequeue - Remove an element from the front of the queue.
func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		return fmt.Errorf("queue is empty")
	}
	element := q.tail
	for i := 0; i < q.len-2; i++ {
		element = element.next
	}
	q.Head = element
	q.Head.next = nil
	q.len--
	return nil
}

//IsEmpty - Check if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

//IsFull - Check if the queue is full.
func (q *Queue) IsFull() bool {
	return q.Size <= q.len
}

//Peek - Get the value of the front of the queue without removing it.
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.Head.Value, nil
}

//Sort - method sorts data.
func (q *Queue) Sort() error {
	if q.IsEmpty() {
		return fmt.Errorf("the queue is empty")
	}
	if q.len == 1 {
		return fmt.Errorf("the queue contains one element")
	}
	for i := 0; i < q.len; i++ {
		element := q.tail
		for element.next != nil {
			if fmt.Sprint(element.Value) > fmt.Sprint(element.next.Value) {
				element.Value, element.next.Value = element.next.Value, element.Value
			}
			element = element.next
		}
	}
	return nil
}
