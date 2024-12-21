package main

import "fmt"

func main() {

	// Time to practice what you learned!
	var hobbies = []string{"Exercise", "Coding", "PlayStation"}

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	fmt.Println(hobbies[0:2])
	updatedArray := hobbies[:2]

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	fmt.Println(updatedArray[1:3])

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Learning Go", "Learning Docker"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Learning Redis"
	goals = append(goals, "Learning Docker")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		title string
		id    int
		price float64
	}

	productArray := []Product{{
		title: "Apple",
		id:    1,
		price: 2.99,
	}, {
		title: "Orange",
		id:    2,
		price: 2.50,
	}}

	fmt.Println(productArray)

	// newItem := Product{
	// 	title: "Pineapple",
	// 	id:    3,
	// 	price: 10.99,
	// }

	// productArray = append(productArray, newItem)

	productArray = append(productArray, Product{
		title: "Pineapple",
		id:    3,
		price: 10.99,
	})
	fmt.Println(productArray)
}
