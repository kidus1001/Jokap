package main

import (
	"fmt"
	"jokeapp/routers"
)

func main() {
	r := routers.RouterSetUp()

	fmt.Println("Server running at port 8080")
	fmt.Println("Here are the routes:")
	fmt.Println("/ - Home page")
	fmt.Println("/joke/:id - Get a specific joke")
	fmt.Println("/jokes - Get all jokes")
	r.Run(":8080")
}
