package main

import "github.com/iriojose/cada_test/initializers"

func Init() {
	initializers.LoadEnv()
	initializers.NewServer()
}

func main() {
	Init()
}
