# Protected

`protected` is a Go library that provides thread-safe access to values using generics and mutexes. This library is useful for ensuring that concurrent read and write operations on shared data do not result in race conditions.

## Features

- Generic support for any data type.
- Thread-safe read and write operations.
- Convenient methods for updating values safely.

## Installation

To install the library, use the following command:

```sh
go get github.com/yourusername/protected
```

## Usage

### Importing the Library

```go
import "github.com/yourusername/protected"
```

### Creating a Protected Value

You can create a protected value using the `New` function:

```go
// Initialize with an integer
protInt := protected.New(10)
```

### Getting the Value

Retrieve the value safely using the `Get` method:

```go
val := protInt.Get()
fmt.Println("Initial value:", val)
```

### Setting a New Value

Set a new value safely using the `Set` method:

```go
protInt.Set(20)
fmt.Println("Updated value:", protInt.Get())
```

### Updating the Value

Update the value using a provided function with the `Update` method:

```go
protInt.Update(func(val int) int {
    return val * 2
})
fmt.Println("Doubled value:", protInt.Get())
```

### Performing Actions with Locks

Perform actions on the value safely with `DoWithLock` and `DoWithRLock` methods:

```go
// Using DoWithLock to safely perform actions on the value
protInt.DoWithLock(func(val *int) {
    *val = *val + 10
})
fmt.Println("Incremented value:", protInt.Get())

// Using DoWithRLock to safely read the value
protInt.DoWithRLock(func(val *int) {
    fmt.Println("Read value:", *val)
})
```

## Example

Here is a complete example:

```go
package main

import (
	"fmt"
	"github.com/yourusername/protected"
)

func main() {
	// Initialize with an integer
	protInt := protected.New(10)

	// Get the value
	val := protInt.Get()
	fmt.Println("Initial value:", val)

	// Set a new value
	protInt.Set(20)
	fmt.Println("Updated value:", protInt.Get())

	// Update the value using a function
	protInt.Update(func(val int) int {
		return val * 2
	})
	fmt.Println("Doubled value:", protInt.Get())

	// Using DoWithLock to safely perform actions on the value
	protInt.DoWithLock(func(val *int) {
		*val = *val + 10
	})
	fmt.Println("Incremented value:", protInt.Get())

	// Using DoWithRLock to safely read the value
	protInt.DoWithRLock(func(val *int) {
		fmt.Println("Read value:", *val)
	})
}
```
