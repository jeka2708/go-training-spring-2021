package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	id    int
	next  *Node
}

// ItemLinkedList the linked list of Items
type ItemLinkedList struct {
	head *Node
	size int
}

func (ll *ItemLinkedList) Insert(s interface{}) {
	addNode := Node{
		value: s,
		next:  nil,
	}
	if ll.head == nil {
		ll.head = &addNode
		fmt.Println(ll.head)
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

func (ll *ItemLinkedList) Deletion() {
	ll.head = ll.head.next
	ll.size--
}

func (n Node) String() string {
	return fmt.Sprintf("id: %d val: %s\n", n.id, n.value)
}

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
func (ll *ItemLinkedList) Delete(id int) {
	node := ll.head
	for {
		if node.id == id {
			ll.head = node.next
			break
		}
		if node.next.id == id {
			node.next = node.next.next
			break
		}
		node = node.next
	}
}

//don`t work
func (ll *ItemLinkedList) Sort() {
	for i := 1; i < ll.size; i++ {
		node := ll.head
		for j := 0; j < ll.size-i; j++ {
			if node.getValue() > node.next.getValue() {
				if node.next != nil {
					temp := node.next
					node.next = node.next.next
					node.next.next = temp

				}
			}
		}
	}
}

//don`t work
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

	//
	//t.Deletion()
	//t.Delete(2)
	t.Sort()
	t.Display()

}
