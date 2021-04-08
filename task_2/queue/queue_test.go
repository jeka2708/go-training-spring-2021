package queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func initVal() Queue {
	q, _ := NewQueue(5, "0")
	_ = q.Enqueue("1")
	_ = q.Enqueue("2")
	_ = q.Enqueue("3")
	_ = q.Enqueue("4")
	return q
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	q := initVal()
	expectedQueue := []string{"0", "1", "2", "3", "4"}
	for i, expected := range expectedQueue {
		node := q.tail
		for j := 1; j < q.len-i; j++ {
			node = node.next
		}
		actual := node.Value
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}
func TestQueue_EnqueueWrongType(t *testing.T) {
	assert := assert.New(t)
	q, _ := NewQueue(2, "1")
	err := q.Enqueue(2)
	expected := fmt.Errorf("types don`t equals string and int")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))
}
func TestQueue_EnqueueIsFull(t *testing.T) {
	assert := assert.New(t)
	q := initVal()
	err := q.Enqueue("six")
	expected := fmt.Errorf("queue is full")
	actual := err
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected:", expected, actual))

}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	q := initVal()
	expectedQueue := []string{"0", "1", "2", "3"}
	for i, expected := range expectedQueue {
		element := q.tail
		for j := 1; j < q.len-i; j++ {
			element = element.next
		}
		actual := element.Value
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	q := initVal()
	actual := q.IsFull()
	assert.Equal(true, actual, fmt.Sprintf("%s and %s not equal, but expected:", "true", strconv.FormatBool(actual)))
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	q := initVal()
	actual, _ := q.Peek()
	expected := "0"
	assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
}

func TestList_Sort(t *testing.T) {
	assert := assert.New(t)
	q := initVal()
	_ = q.Sort()
	expectedList := []string{"0", "1", "2", "3", "4"}
	for i := 0; i < q.len; i++ {
		expected := expectedList[i]
		element := q.tail
		for j := 0; j < i; j++ {
			element = element.next
		}
		actual := element.Value
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}
