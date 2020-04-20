package main

import (
	"bufio"
	"fmt"
	"github.com/aryanugroho/antrian-bucket-go/pkg/antrian"
	"os"
	"strings"
)

func main() {
	fmt.Println("how many slot of lokets you want to create?")
	// get the input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("an error occurred while reading line, errors %s", err.Error())
		return
	}

	inputs := strings.Split(strings.Trim(input, "\n"), " ")

	// validate slot should be between 1 - 5
	isValid := antrian.Validate(inputs[0])
	if !isValid {
		fmt.Printf("number of slot should be 1 - 5")
		return
	}

	fmt.Println(inputs)
}
