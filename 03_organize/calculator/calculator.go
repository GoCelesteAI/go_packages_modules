package calculator

// Add returns the sum of two numbers
func Add(a, b int) int {
  return a + b
}

// Subtract returns the difference
func Subtract(a, b int) int {
  return a - b
}

// Multiply returns the product
func Multiply(a, b int) int {
  return a * b
}

// Divide returns the quotient
func Divide(a, b int) int {
  if b == 0 {
    return 0
  }
  return a / b
}
