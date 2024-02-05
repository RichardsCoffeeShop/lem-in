package main

import (
	"fmt"
	"os"

	"richardscull/lem-in/lib"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Invalid number of arguments")
		return
	}

	// Read file
	filePath := os.Args[1]
	lines, err := lib.ReadFile(filePath)
	if err != nil {
		fmt.Println("ERROR: File not found")
		return
	}

	// Print file content
	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println() // Empty line

	// Resolve input and fill field
	field := lib.Field{}
	err = lib.ResolveInput(lines, &field)
	if err != nil {
		fmt.Println("ERROR: invalid data format, " + err.Error())
		return
	}

	// Note: used for debugging purposes
	// lib.PrintFieldData(&field)

	// Find all shortest paths
	var shortestPaths [][]string
	visited := make(map[string]bool)
	lib.FindShortestPaths(field, "", []string{}, visited, &shortestPaths)
	shortestPaths = lib.RemoveTooLongPaths(shortestPaths)

	if len(shortestPaths) == 0 {
		fmt.Printf("ERROR: invalid data format, farm is unsolvable")
		return
	}

	// Start turn-based solving and print result
	lib.StartTurnBasedSolving(&field, shortestPaths)
}
