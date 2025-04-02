package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
	// } else {
	// 	fmt.Println(err)
	// }
}
func main() {
	var err *MyError  // value - nil, type - *MyError
	errorHandler(err) // Ожидал - Error: nil - Вывелось Error: <nil> Почему? <nil> так всешда выводятся значения c nil
	err = &MyError{}  // value - пустая стуркура type - *MyError
	errorHandler(err) // Ожидал - Error: *MyError{} - Вывелось Error: MyError!, из-за такого, что у err есть реализация метода в стрингу

	//fmt.Println(nil) проверял строку 21
}
