package greetings

import "fmt"

// Welcome is a public function (capitalized)
// It can be called from other packages
func Welcome(name string) string {
  return fmt.Sprintf("Welcome, %s!", name)
}

// Farewell is another public function
func Farewell(name string) string {
  return fmt.Sprintf("Farewell, %s!", name)
}

// formatMessage is private (lowercase)
// It can only be used within this package
func formatMessage(msg string) string {
  return fmt.Sprintf("[GREETING] %s", msg)
}

// Version is a public constant
const Version = "1.0.0"

// maxRetries is a private constant
const maxRetries = 3
