package main

import "fmt"

func main() {
	var foo []int
	var bar []int
	foo = append(foo, 1)
	foo = append(foo, 2)
	foo = append(foo, 3)
	bar = append(foo, 4)
	//fmt.Println(len(foo))
	//foo = foo[:4] - этим я проверил что в массиве на который ссылает foo есть элемент с индексом 3, иначе была бы паника
	//Длина foo по-прежнему 3 поэтому он не видит добавленный элемент "4" и при добавлении "5" он заменяет элемент под индексом 3 (4 на 5)
	foo = append(foo, 5)
	fmt.Println(foo, bar)
}

//Ожидал увидеть foo - 1235 bar - 1234
