package main

import "fmt"


type Queue struct {
   data [5]int
   head int
   tail int
   size int
}

//Dequeue ..
func (q *Queue) Dequeue() int {
    if q.size == 0{
        return -1
    }

    if q.size == 1{
        q.size = 0
        return q.data[q.head]
    }
    
    var res int
    if q.head == 4{
        res = q.data[q.head]
        q.head = 0
    } else{
        res = q.data[q.head]
        q.head++
    }
    q.size--
    return res
}

//Enqueue ..
func (q *Queue) Enqueue(value int)  int {
    if q.size == 5 {
        return -1
    }

    if q.size == 0{
        q.data[q.tail] = value
        q.size = 1
        return q.tail
    }

    if q.tail == 4{
        q.tail = 0
    } else{
        q.tail++
    }
    q.data[q.tail] = value
    q.size++
    return q.tail
}

func main() {

    myque := &Queue{}
    myque.Enqueue(1)
    myque.Enqueue(2)
    myque.Enqueue(3)
    myque.Enqueue(7)
    myque.Enqueue(4)
    myque.Enqueue(5)

    myque.Dequeue()
    myque.Dequeue()

    myque.Enqueue(6)
    myque.Enqueue(8)
    myque.Dequeue()
    myque.Enqueue(9)

    fmt.Printf("head ndex %d \r\n", myque.head)
    fmt.Printf("tail ndex %d \r\n", myque.tail)
    fmt.Printf("queue size %d \r\n", myque.size)

    for i := 0; i < myque.size; i++ {
        index := myque.head + i
        if index > 4 {
            index = index -5
        }
        fmt.Print(myque.data[index])
    }
}
