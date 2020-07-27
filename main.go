package main

import (
	"fmt"
	"github.com/boddumanohar/event-generator/generator"
)

func main() {

	// TODO:
	// take cli arguments for:
	// * Log file
	// * rate of events
	// * rate of errors
	fmt.Println("Generating random events")
	generator.Generate()
}
