package strs

import "fmt"

func ExampleFirstLetterToUpper() {
	fmt.Println(FirstLetterToUpper("Abc"))
	fmt.Println(FirstLetterToUpper("abc"))
	// Output:
	// Abc
	// Abc
}

func ExampleFirstLetterToLower() {
	fmt.Println(FirstLetterToLower("abc"))
	fmt.Println(FirstLetterToLower("Abc"))
	// Output:
	// abc
	// abc
}
