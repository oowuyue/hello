package main

import (
	"errors"
	"fmt"
)

//Node ...
type Node struct {
	value int
}

//Linkqueue ..
type Arrstack struct {
	data [5]Node
	top  int
}

//pop ..
func (st *Arrstack) pop() (Node, error) {
	if st.top >= 0 && st.top < 5 {
		st.top--
		return st.data[st.top+1], nil
	}
	return Node{}, errors.New("out box")
}

//push ..
func (st *Arrstack) push(value int) (Node, error) {
	if st.top >= 4 {
		return Node{}, errors.New("out box")
	}
	st.top++
	st.data[st.top] = Node{value: value}
	return st.data[st.top], nil

}

func main() {

	st := &Arrstack{}
	st.top = -1

	st.push(2)
	st.push(3)
	st.push(24)
	st.push(66)
	st.push(31)
	st.push(12)

	st.pop()
	st.pop()

	st.push(54)

	fmt.Print(st)
	return
}
