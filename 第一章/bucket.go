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
	next  *Student
}

func main() {

	var arr [11]*Student

	var stuCount = 5
	var tname string
	var tscore int
	for index := 0; index < stuCount; index++ {
		fmt.Scan(&tname, &tscore)
		if arr[tscore] == nil {
			arr[tscore] = &Student{name: tname, score: tscore}
		} else {
			arr[tscore] = &Student{name: tname, score: tscore, next: arr[tscore]}
		}
		fmt.Println(arr[tscore])
	}

	fmt.Println("sorted:")
	var tmpStu *Student
	for index := 0; index <= 10; index++ {
		if arr[index] != nil {
			tmpStu = arr[index]
			for tmpStu != nil {
				fmt.Println(tmpStu)
				tmpStu = tmpStu.next
			}
		}
	}

}
