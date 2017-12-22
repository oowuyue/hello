package ds

// Node ...
type Node struct {
	Value int
}

// Queue ...
type Queue struct {
	nodeList []Node
}

// Enqueue ...
func (q *Queue) Enqueue(data int) {
	node := Node{Value: data}
	q.nodeList = append(q.nodeList, node)
}

// Dequeue ...
func (q *Queue) Dequeue() Node {
	dequNode := q.nodeList[0]
	q.nodeList = q.nodeList[1:]
	return dequNode
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	if len(q.nodeList) == 0 {
		return true
	}
	return false

}

// CreateQueue ...
func CreateQueue() Queue {
	newQueue := Queue{}
	return newQueue
}
