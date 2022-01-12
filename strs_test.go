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

func ExampleCamelToSnake() {
	var inputs = []string{
		"XiaoMei", "xiaoMei", "HTTPStatus", "You123", "PriceP", "4sPrice", "Price4s", "goodHTTP",
		"ILoveGolangAndJSONSoMuch",
	}
	for i := range inputs {
		fmt.Println(CamelToSnake(inputs[i]))
	}
	// Output:
	// xiao_mei
	// xiao_mei
	// http_status
	// you123
	// price_p
	// 4s_price
	// price4s
	// good_http
	// i_love_golang_and_json_so_much
}

func ExampleSnakeToCamel() {
	var inputs = []string{"xiao_mei", "http_status", "you123", "price_p"}
	for i := range inputs {
		fmt.Println(SnakeToCamel(inputs[i]))
	}
	// Output:
	// XiaoMei
	// HttpStatus
	// You123
	// PriceP
}
