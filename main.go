package main

import (
	"fmt"
	"my-sonar-go/calculator"
	"my-sonar-go/utils"
)

func main() {
	fmt.Println(calculator.Add(10, 5))
	fmt.Println(utils.FormatName("John", "Doe"))
}