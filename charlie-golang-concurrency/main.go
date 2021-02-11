package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a cache; for this, a simple integer mapped to a book.
// https://golang.org/pkg/cmd/go/internal/cache/
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	// Note: Concurrency is multiple tasks that need to be done, with no order preference.
	// Note: Parallelism is concurrency, but the facility to complete tasks in parallel (at the same time).

	// We're doing some random queries.
	for i := 0; i < 10; i++ {
		// Random + 1 to ensure that we do not get 0, as our books start at 1.
		id := rnd.Intn(10) + 1
		// Create an anonymous function so that we can run a query on the cache in parralel to the db.
		// We pass the id in as an arguement to avoid memory mapping issues.
		go func(id int) {
			// Assign results of queryCache to b, ok; if ok then do the thing.
			if b, ok := queryCache(id); ok {
				fmt.Println("Found entry in cache.")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			// If we don't find it cached, then we have to query the database.
			if b, ok := queryDatabase(id); ok {
				fmt.Println("Had to hit the database.")
				fmt.Println(b)

			}
		}(id)
		// To space things out, to allow our cache time to recieve our values.
		time.Sleep(150 * time.Millisecond)
	}
	// Allow time for the previous ~20 funcs to run, because one main has exited the program finishes.
	time.Sleep(2 * time.Second)
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	// We use sleep to simulate a database query.
	time.Sleep(100 * time.Millisecond)
	// We actually are searching the books in book.go.
	for _, b := range books {
		if b.ID == id {
			// If we get a hit then we can add this to the cache.
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
