# Pool

## Description

`Pool` is a generic object for managing memory allocation and reuse for data types in Go. It provides efficient memory usage by minimizing the number of new allocations using a stack to index previously freed memory.

## Features

- **Insert:** Adds a new instance of type `T` to the `Pool`, returning a reference to it. If there is freed space in the stack, it reuses it instead of allocating new memory.
- **Delete:** Frees a specific instance in the `Pool` and pushes its index onto the stack for later reuse.

## Installation

1. To install the package, use go get:

    ```bash
    go get github.com/WatchJani/pool@v1.0.0
    ```

2. Run tests to ensure everything is working:

    ```bash
    go test ./...
    ```


## Usage

### Creating a New `Pool`

```go
package main

import (
	"fmt"
	"s" // Import stack package
)

func main() {
	// Create a pool for int type with an initial capacity of 10
	pool := New 

	// Iew element and get the reference
	elem1 := pool.Insert()
	*elem1 = 42
	fmt.Println(*elem1) // Output: 42

	// Insert another element
	elem2 := pool.Insert()
	*elem2 = 100
	fmt.Println(*elem2) // Output: 100

	// Delete the first element
	pool.Delete(0)

	// Insert a new element, reusing the freed space
	elem3 := pool.Insert()
	*elem3 = 55
	fmt.Println(*elem3) // Output: 55
}
```
## Functional Details

- **Initialization:** When a new `Pool` is created, memory is allocated for the initial capacity. If additional memory is needed, the `Pool` automatically expands its capacity by doubling its size.
- **Insertion and Reuse:** When inserting a new element, the `Pool` checks the stack for freed elements. If the stack is not empty, it reuses that freed space. If the stack is empty, it allocates new memory.
- **Delete:** When an element is deleted, its index is added to the stack, allowing for later reuse of that memory without a new allocation.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
