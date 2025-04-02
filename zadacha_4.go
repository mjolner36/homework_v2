package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int)
	for _, word := range []string{"hello", "world", "from", "the", "best", "language",
		"in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//Предполагаю, что выведится паника - panic
//= make(map[string]int) нужно инициализировать мапу
