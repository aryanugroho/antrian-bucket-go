package main

import (
	"bufio"
	"fmt"
	"github.com/aryanugroho/antrian-bucket-go/pkg/antrian"
	"os"
	"strconv"
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
	arg := inputs[0]
	isValid := antrian.Validate(arg)
	if !isValid {
		fmt.Printf("number of slot should be 1 - 5")
		return
	}

	// init antrian
	dummyQueueProcess := []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}
	slot, _ := strconv.Atoi(arg)
	antrian := antrian.Antrian{Slot: slot}
	antrian.Start(dummyQueueProcess)

	antrian.Close()
}
