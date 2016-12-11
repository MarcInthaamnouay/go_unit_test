package main

import (
	"fmt"
	"modules/product"
	"start"
)

func main() {
	userOne := start.CreateUser("marc@gmail.com", "marc", "intha", 18)
	fmt.Println(userOne)

	// Create a second user (pointer)
	userTwo := &start.UserStruct{
		Email:     "chachou@gmail.com",
		Firstname: "cc",
		Name:      "pp",
		Age:       18,
	}

	// Testing the product for test purpose...
	var boeing = product.CreateProduct("boeing 777-300ER", 1, userOne)
	var chocker = product.CreateProduct("chocker", 1, userTwo)

	fmt.Println(boeing.Name)
	fmt.Println(chocker.Name)

	fmt.Println("owner %s", boeing.Owner.Firstname)
	fmt.Println("ower of chocker", chocker.Owner.Email)

	fmt.Println("product valid", chocker.IsValid())

}
