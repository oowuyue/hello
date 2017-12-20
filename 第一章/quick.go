package main

import (
	"fmt"
)

type Student1 struct {
	name  string
	score int
}

var stus []Student1

func main() {

	var stuCount int
	fmt.Println("输入同学个数：")
	fmt.Scan(&stuCount)

	var tmpName string
	var tmpScore int
	for index := 0; index < stuCount; index++ {
		fmt.Printf("请输入第 %d 个同学名字和分数：", index)
		fmt.Scan(&tmpName, &tmpScore)
		stus = append(stus, Student1{name: tmpName, score: tmpScore})
	}
	fmt.Println(stus)
	fmt.Println("sorted")
	quick(0, stuCount-1)
	fmt.Println(stus)

	var search int
	var searchRes Student1
	fmt.Println("输入要查找的分数：")
	fmt.Scan(&search)
	searchRes = binsearch(0, stuCount-1, search)
	fmt.Println(searchRes)

	fmt.Println("输入要查找的分数：")
	fmt.Scan(&search)
	searchRes = binsearch(0, stuCount-1, search)
	fmt.Println(searchRes)

}

func quick(left int, right int) {

	if left >= right {
		return
	}

	var standard = stus[left]
	var i = left
	var j = right

	for i < j {
		for stus[j].score >= standard.score && i < j {
			j--
		}

		for stus[i].score <= standard.score && i < j {
			i++
		}
		if i < j {
			var tmp Student1
			tmp = stus[i]
			stus[i] = stus[j]
			stus[j] = tmp
		}

	}

	stus[left] = stus[i]
	stus[i] = standard

	quick(left, i-1)
	quick(i+1, right)

}

func binsearch(left int, right int, search int) Student1 {

	var result Student1
	if left > right {
		return result
	}
	var middle = (left + right) / 2

	if stus[middle].score == search {
		return stus[middle]
	}

	if stus[middle].score < search {
		return binsearch(middle+1, right, search)
	}

	if stus[middle].score > search {
		return binsearch(left, middle-1, search)
	}
	return result
}
