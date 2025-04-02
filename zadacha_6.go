package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println(generatePassword(10))
}

func generatePassword(length int) string {

	//random := rand.New(rand.NewSource(42)) - для тестирование будет выдвать одинаковое значение
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var strBuilder strings.Builder
	for i := 0; i < length; i++ { //начиная с версии 1.20 rand будет выдавать разные сиды
		strBuilder.WriteRune(chars[rand.Intn(len(chars))]) //тестирование - strBuilder.WriteRune(chars[random.Intn(len(chars))])
	}
	return (strBuilder.String())
}

//да можно угадать последовательность так как программа генерирует число из
// random := rand.New(rand.NewSource(98)) это использовать для тестирования

// 2.1 можно угадать если знать seed
//2.2 для тестирования описал комментарии в коде
