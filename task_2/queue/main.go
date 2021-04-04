package main

import "fmt"

//Enqueue Add an element to the end of the queue
func Enqueue(queue []interface{}, val interface{}) []interface{} {
	return append(queue, val)
}

//Dequeue Remove an element from the front of the queue
func Dequeue(queue []interface{}) []interface{} {
	return queue[1:]
}

//IsEmpty Check if the queue is empty
func IsEmpty(queue []interface{}) bool {
	return len(queue) == 0
}

//IsFull Check if the queue is full
func IsFull(queue []interface{}) bool {
	return cap(queue) == len(queue)
}

//Peek Get the value of the front of the queue without removing it
func Peek(queue []interface{}) {
	fmt.Println(queue[1])
}
func main() {
	var queue = make([]interface{}, 1, 1)
	queue = Enqueue(queue, 2)
	queue = Enqueue(queue, 3)
	queue = Enqueue(queue, 1)
	queue = Dequeue(queue)
	fmt.Println(IsEmpty(queue))
	fmt.Println(IsFull(queue))
	Peek(queue)
	fmt.Println(queue)
}
