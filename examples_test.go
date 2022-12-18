package textcase_test

import (
	"fmt"

	"github.com/golang-cz/textcase"
)

func ExampleCamelCase() {
	fmt.Println(textcase.CamelCase("Hello World!"))
	// Output: helloWorld
}

func ExamplePascalCase() {
	fmt.Println(textcase.PascalCase("Hello World!"))
	// Output: HelloWorld
}

func ExampleSnakeCase() {
	fmt.Println(textcase.SnakeCase("Hello World!"))
	// Output: hello_world
}

func ExampleKebabCase() {
	fmt.Println(textcase.KebabCase("Hello World!"))
	// Output: hello-world
}
