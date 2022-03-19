package main

import (
	"context"
	"fmt"

	"github.com/DanielMartin96/blog/internal/db"
)

// Run - responsible for instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		fmt.Println(err, "hello")
		return err
	}
	fmt.Println("successfully connected and pinged database")

	return nil
}

func main()  {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}