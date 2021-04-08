package main

import (
	"fmt"
	"github.com/jeka2708/go-training-spring-2021/task_2/queue"
)

func main() {
	var q, _ = queue.NewQueue(2, "1", "4", "2", "4")
	for i := 0; i < q.Size; i++ {
		fmt.Println(q.Head.Value)
	}
	//var t, e = linked_list.NewList("2", "1", "4","2","4")
	//for i:=0; i<1000;i++{
	//	t.Insert(string(i))
	//}
	////
	////
	////t.Display()
	//
	//if e != nil{
	//	fmt.Println(e)
	//}
	//t.Sort()
	////t.Display()
}
