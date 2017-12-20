package main

import (
	"fmt"
)

/*
Student
*/
type Student struct {
	name  string
	score int
}

func swap(t1 *Student, t2 *Student) {
	var tmp Student
	tmp = *t1
	*t1 = *t2
	*t2 = tmp
}

func main() {

	var arr [5]Student
	var stuCount = 5

	var tname string
	var tscore int
	for index := 0; index < stuCount; index++ {
		fmt.Scan(&tname, &tscore)
		arr[index] = Student{name: tname, score: tscore}
		fmt.Println(arr[index])
	}

	for i := 1; i <= stuCount-1; i++ { // n位数组  n-1趟
		for j := 0; j < stuCount-i; j++ {
			if arr[j].score < arr[j+1].score {
				swap(&arr[j], &arr[j+1])
			}
		}
	}

	fmt.Println("sorted:")
	fmt.Println(arr)
}
