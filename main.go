package main

import (
	"fmt"
	"bufio"
	"lem-in/parse"
	"os"
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
	fmt.Println(parse.ParseFile(text))
}
