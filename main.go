package main

import (
	"fmt"

	apigin "example.com/api-with-go/apiGin"
)

func main() {

	// fmt.Println("Api using Mux")
	// apimux.ApiMUX()

	fmt.Println("Api using Gin")
	apigin.ApiGIN()

}
