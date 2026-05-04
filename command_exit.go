package main

import "os"

func callbackExit(cfg *config) error{
	println("Goodbye!")
	os.Exit(0)
	return nil
}