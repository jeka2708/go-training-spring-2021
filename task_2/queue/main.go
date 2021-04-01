package main

import "fmt"

//type queue struct {
//	q []interface{}
//}
//
func Enqueue(queue []interface{}, val interface{}) {
	queue = append(queue, val)
}

func Dequeue(queue []interface{}) []interface{} {
	return queue[1:]
}

func IsEmpty(queue []interface{}) bool {
	return len(queue) > 0
}

func IsFull(queue []interface{}) bool {
	return cap(queue) == len(queue)
}
func Peek(queue []interface{}) {
	fmt.Println(queue[1])
}
func main() {
	var queue = make([]interface{}, 1, 1)
	fmt.Println(queue)
}
