package calculator

// Add returns sum of two numbers
func Add(a, b int) int {
	return a + b
}

// Divide - Bug: no zero division check (SonarQube will flag)
func Divide(a, b int) int {
	return a / b
}

// unusedFunction - dead code (SonarQube will flag)
func unusedFunction() string {
	return "I am never called"
}