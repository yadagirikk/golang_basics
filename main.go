package main

import "errors"

func AddOnlyPositive(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b is negative")
	}
	return a + b, nil
}

// func Add(a int, b float64) (int, int) {
// 	result := a + int(b)
// 	return result, result + 1
// }

// func SayHelloWorld() string {
// 	return "Hello World"
// }

func main() {
	// result, _ := Add(10, 20.4)
	// result2, _ := Add(11, 20.4)
	// result3, _ := Add(12, 20.4)
	// println(result, result2, result3)

	result, err := AddOnlyPositive(10, -100)
	if err != nil {
		println("Error: ", err.Error())
	} else {
		println(result)
	}
}
