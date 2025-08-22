# LRU Cache Implementation in Go

A simple, efficient, and type-safe LRU (Least Recently Used) cache implementation in Go using generics.

## Features

- **Generic Types**: Supports any comparable key type and any value type
- **O(1) Operations**: Constant time complexity for Get and Put operations
- **Thread-Safe Ready**: Easy to extend with mutex for concurrent usage
- **Memory Efficient**: Automatic eviction of least recently used items
- **Type Safe**: Compile-time type checking with Go generics

## Installation

This is a standalone Go module. Simply copy the code into your project or save it as a `.go` file.

```bash
# Run the example
go run lru_cache.go
```

## Quick Start

```go
package main

import "fmt"

func main() {
    // Create a string cache with capacity of 3
    cache := Constructor[int, string](3)

    // Add items
    cache.Put(1, "hello")
    cache.Put(2, "world")
    cache.Put(3, "golang")

    // Get item (moves it to front)
    if value, found := cache.Get(2); found {
        fmt.Println(value) // "world"
    }

    // Add another item (evicts least recently used)
    cache.Put(4, "generics")

    // Check cache size
    fmt.Printf("Cache size: %d\n", cache.Size())
}
```

## API Reference

### Constructor

#### `Constructor[K, V](capacity int) *LRUCache[K, V]`

Creates a new LRU cache with the specified capacity.

**Type Parameters:**

- `K comparable` - Key type (must be comparable)
- `V any` - Value type (can be any type)

**Parameters:**

- `capacity` - Maximum number of items the cache can hold

**Returns:**

- Pointer to a new LRUCache instance

### Methods

#### `Get(key K) (V, bool)`

Retrieves a value by key and marks it as recently used.

**Parameters:**

- `key` - The key to look up

**Returns:**

- `value` - The associated value (zero value if not found)
- `found` - Boolean indicating if the key was found

**Time Complexity:** O(1)

#### `Put(key K, value V)`

Adds a new key-value pair or updates an existing one. If the cache is at capacity, evicts the least recently used item.

**Parameters:**

- `key` - The key to store
- `value` - The value to associate with the key

**Time Complexity:** O(1)

#### `Display()`

Prints the current state of the cache (for debugging purposes).

## Usage Examples

### String Cache

```go
// Integer keys, string values
stringCache := Constructor[int, string](3)
stringCache.Put(1, "hello")
stringCache.Put(2, "world")

if value, found := stringCache.Get(1); found {
    fmt.Println(value) // "hello"
}
```

### Custom Struct Cache

```go
type User struct {
    ID   int
    Name string
    Age  int
}

// String keys, User struct values
userCache := Constructor[string, User](10)
userCache.Put("user1", User{ID: 1, Name: "Alice", Age: 25})

if user, found := userCache.Get("user1"); found {
    fmt.Printf("User: %+v\n", user)
}
```

### Dictionary Cache

```go
// String keys and values
dictCache := Constructor[string, string](100)
dictCache.Put("en", "hello")
dictCache.Put("es", "hola")
dictCache.Put("fr", "bonjour")
```

### Traditional Integer Cache

```go
// Integer keys and values
intCache := Constructor[int, int](5)
intCache.Put(1, 100)
intCache.Put(2, 200)
```

## How It Works

The LRU cache uses two main data structures:

1. **Hash Map**: Provides O(1) key lookup
2. **Doubly Linked List**: Maintains order of usage

### Internal Structure

```
Head (dummy) ↔ [Most Recent] ↔ [...] ↔ [Least Recent] ↔ Tail (dummy)
```

### Operations

- **Get**: Move accessed node to head (most recent position)
- **Put**: Add new node at head, evict from tail if at capacity
- **Eviction**: Remove least recently used node from tail

## Performance

- **Time Complexity**: O(1) for all operations
- **Space Complexity**: O(capacity)

## Limitations

- **Not Thread-Safe**: Requires external synchronization for concurrent access
- **Generic Constraints**: Keys must be comparable, values can be any type
- **Memory**: Holds references to all cached items until evicted

## Making It Thread-Safe

To use in concurrent environments, wrap operations with a mutex:

```go
type SafeLRUCache[K comparable, V any] struct {
    cache *LRUCache[K, V]
    mutex sync.RWMutex
}

func (s *SafeLRUCache[K, V]) Get(key K) (V, bool) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return s.cache.Get(key)
}

func (s *SafeLRUCache[K, V]) Put(key K, value V) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.cache.Put(key, value)
}
```

## Requirements

- Go 1.18+ (for generics support)

## License

This implementation is provided as-is for educational and practical use.

## Contributing

Feel free to submit issues, suggestions, or improvements to make this LRU cache implementation even better!
