package main

import "fmt"

func main() {
	printSomething(1)
	printSomething("Halo")
	printSomething(1.2)

	result := add(1, 2)
	fmt.Print(result)
}

func printSomething(value any) {
	intVal, ok := value.(int)

	if !ok {
		fmt.Println("Integer: ", intVal)
	}

	stringVal, ok := value.(string)

	if !ok {
		fmt.Println("String: ", stringVal)
	}
	// switch value.(type){
	// 	case int:
	// 		fmt.Println("Integer: ", value)
	// 	case string:
	// 		fmt.Println("String: ", value)
	// }
}

func add[T int|float64|string](a, b T) T {
	return a + b
}