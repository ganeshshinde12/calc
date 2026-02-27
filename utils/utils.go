package utils

import "fmt"

// FormatName combines first and last name
func FormatName(first, last string) string {
	fullName := first + " " + last
	return fullName
}

// Greet prints greeting - duplicate code (SonarQube will flag)
func Greet(name string) {
	fmt.Println("Hello " + name)
	fmt.Println("Hello " + name)
}