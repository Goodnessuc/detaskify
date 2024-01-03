package main

import (
	"errors"
	"fmt"
)

func Run() error {
	fmt.Println("starting up the application...")
	return errors.New("")

}

func main() {
	if err := Run(); err != nil {
		fmt.Println("error starting up the application")
	}
}
