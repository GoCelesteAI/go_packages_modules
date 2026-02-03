# Go Packages & Modules

Demo code from **Go Tutorial Lesson 21: Packages & Modules**

Watch the full tutorial: [YouTube Link]

## Topics Covered

1. **Module Basics** - `go mod init`, go.mod file structure
2. **Creating Packages** - Package directories, public vs private (capitalization)
3. **Importing Packages** - Local package imports, using exported functions
4. **Multi-Package Projects** - Organizing code across multiple packages

## Project Structure

```
├── 01_module/           # Basic module creation
│   ├── go.mod
│   └── main.go
├── 02_package/          # Creating and importing packages
│   ├── go.mod
│   ├── greetings/
│   │   └── greetings.go
│   └── main.go
└── 03_organize/         # Multi-package project
    ├── go.mod
    ├── calculator/
    │   └── calculator.go
    ├── stringutil/
    │   └── stringutil.go
    └── main.go
```

## Running the Examples

```bash
# Module basics
cd 01_module
go run .

# Package creation and import
cd 02_package
go run .

# Multi-package project
cd 03_organize
go run .
```

## Key Concepts

### Public vs Private
- **Uppercase** first letter = Public (exported)
- **Lowercase** first letter = Private (unexported)

```go
func Welcome() string  // Public - can be used from other packages
func formatName() string  // Private - only within same package
```

### Module Path
The module path in `go.mod` determines the import path:
```go
module myproject

import "myproject/greetings"
```

## Channel

Subscribe for more Go tutorials: [GoCeleste AI](https://youtube.com/@GoCelesteAI)
