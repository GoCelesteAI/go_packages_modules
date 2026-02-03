package main

import (
  "fmt"

  "toolkit/calculator"
  "toolkit/stringutil"
)

func main() {
  fmt.Println("=== Multi-Package Project ===")

  // Using calculator package
  fmt.Println("\n--- Calculator ---")
  fmt.Println("10 + 5 =", calculator.Add(10, 5))
  fmt.Println("10 - 5 =", calculator.Subtract(10, 5))
  fmt.Println("10 * 5 =", calculator.Multiply(10, 5))

  // Using stringutil package
  fmt.Println("\n--- String Utils ---")
  fmt.Println("Reverse 'hello':", stringutil.Reverse("hello"))
  fmt.Println("ToUpper 'hello':", stringutil.ToUpper("hello"))
}
