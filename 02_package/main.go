package main

import (
  "fmt"

  "myapp/greetings"
)

func main() {
  fmt.Println("=== Using Our Package ===")

  // Call public functions
  msg1 := greetings.Welcome("Gopher")
  fmt.Println(msg1)

  msg2 := greetings.Farewell("World")
  fmt.Println(msg2)

  // Access public constant
  fmt.Println("Greetings version:", greetings.Version)

  // This would NOT work - private function:
  // greetings.formatMessage("test") // Error!

  // This would NOT work - private constant:
  // fmt.Println(greetings.maxRetries) // Error!
}
