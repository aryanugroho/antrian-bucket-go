package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("how many slot of lokets you want to create?")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("an error occurred while reading line, errors %s", err.Error())
	}

	inputs := strings.Split(strings.Trim(input, "\n"), " ")
	fmt.Println(inputs)
}
