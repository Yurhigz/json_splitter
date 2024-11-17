package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var userInput string

	// fmt.Scan(&userInput)

	// if utils.JsonChecking(userInput) != nil {
	// 	return
	// }

	// Open our jsonFile
	jsonFile, err := os.Open("D:/Mes_documents/Programmation/json_splitter/test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	r := bufio.NewReader(jsonFile)

	// fmt.Println(r)

	data := make([]byte, 1000)

	_, err = r.Read(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
