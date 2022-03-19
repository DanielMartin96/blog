package main

import "fmt"

// Run - responsible for instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main()  {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}