package main

import "fmt"

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
}

// Constructor creates a new LRU cache with given capacity
func Constructor[K comparable, V any](capacity int) LRUCache[K, V] {
	head := &Node[K, V]{}
	tail := &Node[K, V]{}
	head.next = tail
	tail.prev = head

	return LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
		head:     head,
		tail:     tail,
	}
}

// addNode adds a node right after the head
// Before: head <-> A <-> B <-> ... <-> tail
// After:  head <-> node <-> A <-> B <-> ... <-> tail
func (lru *LRUCache[K, V]) addNode(node *Node[K, V]) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode removes an existing node from the linked list
func (lru *LRUCache[K, V]) removeNode(node *Node[K, V]) {
	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
}

// moveToHead moves certain node to the head
func (lru *LRUCache[K, V]) moveToHead(node *Node[K, V]) {
	lru.removeNode(node)
	lru.addNode(node)
}

// popTail removes the last node (least recently used)
func (lru *LRUCache[K, V]) popTail() *Node[K, V] {
	lastNode := lru.tail.prev
	lru.removeNode(lastNode)
	return lastNode
}

// Get retrieves value by key, returns -1 if not found
func (lru *LRUCache[K, V]) Get(key K) (V, bool) {
	node, exists := lru.cache[key]
	if !exists {
		var zero V
		return zero, false
	}

	lru.moveToHead(node)
	return node.value, true
}

// Put adds or updates a key-value pair
func (lru *LRUCache[K, V]) Put(key K, value V) {
	node, exists := lru.cache[key]

	if exists {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &Node[K, V]{
		key:   key,
		value: value,
	}
	if len(lru.cache) >= lru.capacity {
		tail := lru.popTail()
		delete(lru.cache, tail.key)
	}

	lru.cache[key] = newNode
	lru.addNode(newNode)
}

// Display shows current cache state (for debugging)
func (lru *LRUCache[K, V]) Display() {
	fmt.Printf("Cache (capacity: %d, size: %d): ", lru.capacity, len(lru.cache))
	current := lru.head.next
	for current != lru.tail {
		fmt.Printf("[%v:%v]", current.key, current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== LRU Cache with Generics Demo ===")

	// Example 1: String cache (int keys, string values)
	fmt.Println("\n--- String Cache Example ---")
	stringCache := Constructor[int, string](3)

	stringCache.Put(1, "hello")
	stringCache.Put(2, "world")
	stringCache.Put(3, "golang")
	stringCache.Display() // [3:golang] [2:world] [1:hello]

	if value, found := stringCache.Get(2); found {
		fmt.Printf("Get(2): %s\n", value) // "world"
	}
	stringCache.Display() // [2:world] [3:golang] [1:hello]

	stringCache.Put(4, "generics") // Evicts key 1
	stringCache.Display()          // [4:generics] [2:world] [3:golang]

	if _, found := stringCache.Get(1); !found {
		fmt.Println("Key 1 not found (evicted)")
	}

	// Example 2: User data cache (string keys, custom struct values)
	fmt.Println("\n--- Custom Struct Cache Example ---")
	type User struct {
		ID   int
		Name string
		Age  int
	}

	userCache := Constructor[string, User](2)

	userCache.Put("user1", User{ID: 1, Name: "Alice", Age: 25})
	userCache.Put("user2", User{ID: 2, Name: "Bob", Age: 30})
	userCache.Display() // [user2:{2 Bob 30}] [user1:{1 Alice 25}]

	if user, found := userCache.Get("user1"); found {
		fmt.Printf("Found user: %+v\n", user)
	}
	userCache.Display() // [user1:{1 Alice 25}] [user2:{2 Bob 30}]

	userCache.Put("user3", User{ID: 3, Name: "Charlie", Age: 35}) // Evicts user2
	userCache.Display()                                           // [user3:{3 Charlie 35}] [user1:{1 Alice 25}]

	// Example 3: Integer cache (traditional usage)
	fmt.Println("\n--- Integer Cache Example ---")
	intCache := Constructor[int, int](2)

	intCache.Put(10, 100)
	intCache.Put(20, 200)
	intCache.Display() // [20:200] [10:100]

	if value, found := intCache.Get(10); found {
		fmt.Printf("Get(10): %d\n", value) // 100
	}

	// Example 4: String-to-string cache
	fmt.Println("\n--- String-to-String Cache Example ---")
	dictCache := Constructor[string, string](3)

	dictCache.Put("en", "hello")
	dictCache.Put("es", "hola")
	dictCache.Put("fr", "bonjour")
	dictCache.Display() // [fr:bonjour] [es:hola] [en:hello]
}
