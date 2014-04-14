package books

import "fmt"

func ExampleBook() {
	potter := Book{
		Title: "Harry Potter",
		Author: &Author{
			First: "JK",
			Last:  "Rowling",
		},
	}
	fmt.Println(potter.String())
	// Harry Potter by JK ROWLING

	cool := Book{
		Title: "The Cool Book",
	}
	fmt.Println(cool.String())
	// The Cool Book by unknown author

	// Output:
	// Harry Potter by JK ROWLING
	// The Cool Book by unknown author
}
