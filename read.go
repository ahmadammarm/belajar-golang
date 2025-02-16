package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			name := Name{Fname: parts[0], Lname: parts[1]}
			names = append(names, name)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, name := range names {
		fmt.Println("First Name:", name.Fname, "Last Name:", name.Lname)
	}
}