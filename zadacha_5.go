package main

import (
	"fmt"
)

func main() {
	str := "Привет"
	//str[2] = 'e' // сторки не изменяемый тип, по индексу находятся байты а не символ

	runes := []rune(str)
	runes[2] = 'e'

	fmt.Println(string(runes))

}

//строка не изменяемый тип данных, строка это последовательность байт utf-8
