package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initVal() ItemLinkedList {
	q, _ := NewList("a")
	_ = q.Insert("b")
	_ = q.Insert("c")
	_ = q.Insert("d")
	return q
}

func TestItemLinkedList_Add(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	expected := 4
	actual := list.size
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected:", expected, actual))
}

func TestItemLinkedList_Search(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	expectedList := []string{"d", "c", "b", "a"}
	for i := 0; i < 4; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual.value, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual.value))
	}
}

func TestItemLinkedList_SearchNotId(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	expected := fmt.Errorf("id not found")
	_, actual := list.Search(6)
	assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
}

func TestItemLinkedList_Delete(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	id := 9
	e := list.Delete(id)
	expected := fmt.Errorf("id not found %d", id)
	actual := e
	assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
}

func TestItemLinkedList_DeleteNotId(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	list.Delete(0)
	list.Delete(1)
	expectedList := []string{"c", "a"}
	for i := 0; i < list.size; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual.value, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

//
func TestItemLinkedList_Deletion(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	list.Deletion()
	list.Deletion()
	expectedList := []string{"b", "a"}
	for i := 0; i < list.size; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual.value, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
	}
}

func TestItemLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	var list = initVal()
	_ = list.Sort()
	expectedList := []string{"a", "b", "c", "d"}
	for i := 0; i < list.size; i++ {
		expected := expectedList[i]
		actual, _ := list.Search(i)
		assert.Equal(expected, actual.value, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual.value))
	}
}
func TestItemLinkedList_SortEnoughElement(t *testing.T) {
	assert := assert.New(t)
	var list, _ = NewList("1")
	e := list.Sort()
	expected := fmt.Errorf("not enough elements")
	actual := e
	assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected:", expected, actual))
}
