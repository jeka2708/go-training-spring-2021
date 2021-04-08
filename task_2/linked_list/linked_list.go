package linked_list

import (
	"fmt"
	"reflect"
	"time"
)

type Node struct {
	value interface{}
	next  *Node
}

// ItemLinkedList the linked_list
type ItemLinkedList struct {
	head *Node
	size int
}

func NewList(val ...interface{}) (ItemLinkedList, error) {
	var ll ItemLinkedList
	for _, v := range val {
		e := ll.Insert(v)
		if e != nil {
			return ItemLinkedList{}, e
		}
	}
	return ll, nil
}

//Insert will put a value in list
//s interface{} will be the value
func (ll *ItemLinkedList) Insert(val interface{}) error {
	addNode := Node{
		value: val,
		next:  nil,
	}
	if ll.head == nil {
		ll.head = &addNode
	} else {
		currentType := reflect.TypeOf(ll.head.value)
		addType := reflect.TypeOf(val)
		if currentType != addType {
			return fmt.Errorf("types don`t equals %s and %s", currentType, addType)
		}
		addNode.next = ll.head
		ll.head = &addNode
	}
	ll.size++
	return nil
}

//Deletion removes a value from the top of the list
func (ll *ItemLinkedList) Deletion() error {
	if ll.size > 0 {
		ll.head = ll.head.next
		ll.size--
		return nil
	}
	return fmt.Errorf("list if empty")
}

//String return string of Node
func (n Node) String() string {
	return fmt.Sprintf("val: %s\n", n.value)
}

//Display print elements of list
func (ll *ItemLinkedList) Display() {
	node := ll.head
	fmt.Println("------Start list")
	for node.next != nil {
		fmt.Printf("val: %s\n", node.value)
		node = node.next
	}
	fmt.Printf("val: %s\n", node.value)
	fmt.Println("------End list")
}

//Search find elements by id
//id element`s id
func (ll *ItemLinkedList) Search(id int) (Node, error) {
	node := ll.head
	for i := 0; i < ll.size; i++ {
		if i == id {
			return *node, nil
		}
		node = node.next
	}
	return Node{}, fmt.Errorf("id not found")
}

//Delete remove elements by id
//id element`s id
func (ll *ItemLinkedList) Delete(id int) error {
	node := ll.head
	for i := 0; i < ll.size; i++ {
		if i == id {
			ll.head = node.next
			ll.size--
			return nil
		}
		if i+1 == id && i+1 < ll.size {
			node.next = node.next.next
			ll.size--
			return nil
		}
		node = node.next
	}
	return fmt.Errorf("id not found %d", id)
}

//Sort sort elements
func (ll *ItemLinkedList) Sort() error {

	if ll.size < 2 {
		return fmt.Errorf("not enough elements")
	}
	start := time.Now()
	for i := 1; i < ll.size; i++ {
		currentNode := ll.head
		prevNode := ll.head
		for j := 0; j < ll.size-i; j++ {
			if currentNode.getValue() > currentNode.next.getValue() {
				if ll.head == currentNode {
					temp := ll.head
					ll.head = ll.head.next
					temp.next = ll.head.next
					ll.head.next = temp
					prevNode = ll.head

				} else {
					nodeNext := currentNode.next
					prevNode.next = currentNode.next
					prevNode = currentNode.next
					currentNode.next = nodeNext.next
					nodeNext.next = currentNode
				}
			}

		}

		currentNode = currentNode.next
	}
	elapsed := time.Since(start).Seconds()
	fmt.Println(elapsed)
	return nil
}

//getValue returns a value of the type
//only strings
func (n *Node) getValue() string {
	return fmt.Sprintf(n.value.(string))
}
