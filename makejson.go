package main

import (
	"encoding/json"
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]


	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')
	address = address[:len(address)-1]

	data := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
