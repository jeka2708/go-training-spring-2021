package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	id    int
	next  *Node
}

// ItemLinkedList the linked list
type ItemLinkedList struct {
	head *Node
	size int
}

//Insert will put a value in list
//s interface{} will be the value
func (ll *ItemLinkedList) Insert(s interface{}) {
	addNode := Node{
		value: s,
		next:  nil,
	}
	if ll.head == nil {
		ll.head = &addNode
	} else {
		addNode.next = ll.head
		ll.head = &addNode

		node := ll.head
		i := 0
		for node.next != nil {
			node.id = i
			node = node.next
			i++
		}
		if node.next == nil {
			node.id = i
		}
	}
	ll.size++
}

//Deletion removes a value from the top of the list
func (ll *ItemLinkedList) Deletion() {
	ll.head = ll.head.next
	ll.size--
}

//String return string of Node
func (n Node) String() string {
	return fmt.Sprintf("id: %d val: %s\n", n.id, n.value)
}

//Display print elements of list
func (ll *ItemLinkedList) Display() {
	node := ll.head
	fmt.Println("------Start list")
	for node.next != nil {
		fmt.Printf("id: %d val: %s\n", node.id, node.value)
		node = node.next
	}
	fmt.Printf("id: %d val: %s\n", node.id, node.value)
	fmt.Println("------End list")
}

//Search find elements by id
//id element`s id
func (ll *ItemLinkedList) Search(id int) Node {
	node := ll.head
	for {
		if node.id == id {
			return *node
		}
		if node.next == nil {
			return Node{}
		}
		node = node.next
	}
}

//Delete remove elements by id
//id element`s id
func (ll *ItemLinkedList) Delete(id int) {
	node := ll.head
	for {
		if node.id == id {
			ll.head = node.next
			ll.size--
			break
		}
		if node.next.id == id {
			node.next = node.next.next
			ll.size--
			break
		}
		node = node.next
	}
}

//Sort sort elements
func (ll *ItemLinkedList) Sort() {
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
}

//getValue returns a value of the type
//only strings
func (n *Node) getValue() string {
	switch str := n.value.(type) {
	case string:
		return str
	}
	return ""
}

func main() {
	var t ItemLinkedList
	t.Insert("1")
	t.Insert("2")
	t.Insert("3")
	t.Insert("4")
	t.Insert("5")

	t.Deletion()
	t.Delete(1)
	t.Sort()
	t.Display()

}
