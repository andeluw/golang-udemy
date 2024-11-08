package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func newProduct(title string, id string, price float64) Product {
	return Product{
		title: title,
		id:    id,
		price: price,
	}
}

// func main(){
// 	// 1
// 	hobbies := [3]string{"Code", "Basket", "Game"}
// 	fmt.Println(hobbies)

// 	// 2
// 	fmt.Println(hobbies[0])
// 	fmt.Println(hobbies[1:])

// 	//3
// 	featuredHobbies := hobbies[:2]
// 	fmt.Println(featuredHobbies)
// 	// featuredHobbies2 := hobbies[0:2]
// 	// fmt.Println(featuredHobbies2)

// 	// 4
// 	featuredHobbies = featuredHobbies[1:3]
// 	fmt.Println(featuredHobbies)

// 	// 5
// 	goals := []string{"Backend", "Frontend"}
// 	fmt.Println(goals)

// 	// 6
// 	goals[1] = "UI UX"
// 	goals = append(goals, "DevOps")
// 	fmt.Println(goals)

// 	// 7
// 	products := []Product{
// 		// newProduct("Bola", "1", 2000),
// 		// newProduct("Meja", "2", 300000),
// 		{"Bola", "1", 2000},
// 		{"Meja", "2", 300000},
// 	}
// 	products = append(products,
// 		newProduct("Mouse", "3", 200000))
// 	fmt.Println(products)

// 	newProducts := []Product{
// 		{"Pensil", "4", 5000},
// 		{"Laptop", "5", 100000},
// 	}
// 	products = append(products, newProducts...)
// 	fmt.Println(products)
// }

// func main() {
// 	prices := []float64{10.99}
// 	prices = append(prices, 5.99)
// 	fmt.Println(prices)
// }

// func main() {
// 	var productNames [4]string
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	productNames[1] = "A Carpet"

// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }

func main() {
	userNames := make([]string, 3, 5)

	userNames[0] = "Julie"
	fmt.Println(userNames)

	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}
}