package main

import (
	"bufio"
	"fmt"
	"os"

	parsing "lem-in/parse"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Usage: main <filename>")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		println("Error opening file: ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + "\n"
		if err := scanner.Err(); err != nil {
			println("Error reading file: ", err)
			return
		}
	}
	Graph := parsing.NewGraph()
	fmt.Println(Graph.ParseFile(text))
}
