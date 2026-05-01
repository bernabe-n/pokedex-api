package main

import "os"

func callbackExit() error{
	println("Goodbye!")
	os.Exit(0)
	return nil
}