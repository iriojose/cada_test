package main

import "github.com/iriojose/cada_test/initializers"

func Init() {
	initializers.LoadEnv()
	initializers.NewServer(":8000")
}

func main() {
	Init()
}
