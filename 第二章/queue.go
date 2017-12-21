package main

import (
	"fmt"
)

//Node ...
type Node struct {
	data int
	next *Node
	prve *Node
}

//Linkqueue ..
type Linkqueue struct {
	head *Node
	tail *Node
	size int
}

//Dequeue ..
func (q *Linkqueue) Dequeue() *Node {
	if q.head == nil {
		return nil
	}
	var res *Node
	if q.size == 1 {
		res = q.head
		q.head = nil
		q.tail = nil
	} else {
		res = q.head
		q.head = q.head.next
		q.head.prve = nil
	}
	q.size--
	return res
}

//Enqueue ..
func (q *Linkqueue) Enqueue(value int) *Node {
	if q.tail == nil {
		q.tail = &Node{data: value}
		q.head = q.tail
		q.size++
		return q.tail
	}
	if q.size == 1 {
		q.tail = &Node{data: value, next: nil, prve: q.tail}
		q.head.next = q.tail
	} else {
		q.tail = &Node{data: value, next: nil, prve: q.tail}
		q.tail.prve.next = q.tail
	}
	q.size++
	return q.tail
}

func main() {

	q := &Linkqueue{}

	numArr := [9]int{4, 2, 9, 8, 5, 7, 1, 3, 6}
	for index := 8; index >= 0; index-- {
		q.Enqueue(numArr[index])
	}

	for q.tail != q.head {
		fmt.Printf("%d ", q.head.data)
		q.Dequeue()

		q.Enqueue(q.head.data)
		q.Dequeue()
	}
	fmt.Printf("%d ", q.head.data)

	var tmp *Node
	tmp = q.tail
	for tmp != nil {
		fmt.Printf("\r\n Linkqueue node %v", tmp)
		tmp = tmp.prve
	}
	return
}
