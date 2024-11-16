package main

import (
	"fmt"
	"json-splitter/utils"
)

func main() {
	var userInput string

	fmt.Scan(&userInput)

	if utils.JsonChecking(userInput) != nil {
		return
	}

}
