package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a cache; for this, a simple integer mapped to a book.
// https://golang.org/pkg/cmd/go/internal/cache/
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	// Note: Concurrency is multiple tasks that need to be done, with no order preference.
	// Note: Parallelism is concurrency, but the facility to complete tasks in parallel (at the same time).

	// Create a wait group in order to allow functions to complete for each iteration.
	// Create a mutex for memory to resolve race conditions.
	// Create a reference to a waitgroup & mutex, as we're going to be passing this around functions.

	waitGroup := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	// In this example we need to control the output that is displayed on the command line.
	// Here we are using channels, it is to demonstrate how one could use them.
	// There is a better way to achieve this, by making query to db conditional on query to cache.
	cacheChannel := make(chan Book)
	dbChannel := make(chan Book)

	// We're doing some random queries.
	for i := 0; i < 10; i++ {
		// Random + 1 to ensure that we do not get 0, as our books start at 1.
		id := rnd.Intn(10) + 1
		// We're adding two jobs to wait group, in this case its queryCache and queryDatabase.
		waitGroup.Add(2)
		// Create an anonymous function so that we can run a query on the cache in parralel to the db.
		// We pass the id in as an arguement to avoid memory mapping issues.
		// We pass in cacheChannel as send only! This is denoted by arrow being on the right side of the chan.
		go func(id int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex, cacheChannel chan<- Book) {
			// Assign results of queryCache to b, ok; if ok then do the thing.
			if b, ok := queryCache(id, mutex); ok {
				cacheChannel <- b
			}
			waitGroup.Done()
		}(id, waitGroup, mutex, cacheChannel)
		go func(id int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex, dbChannel chan<- Book) {
			// If we don't find it cached, then we have to query the database.
			if b, ok := queryDatabase(id, mutex); ok {
				dbChannel <- b
			}
			waitGroup.Done()
		}(id, waitGroup, mutex, dbChannel)

		// We now pass the channels as a reciever, so arrow on the left.
		// Note: All channels need a send and a reciever.
		// Considering the above comment, we place the recieve inside of the loop as we need to recieve...
		// ten lots of messages because ten are sent, if we don't all hell will break loose.
		go func(cacheChannel, dbChannel <-chan Book) {
			select {
			case b := <-cacheChannel:
				fmt.Println("Found result in the cache:")
				fmt.Println(b)
				// Because we always hit the database, we need to consume the message.
				// Therefore we just dump it here, otherwise we will get duplicates on the output.
				// As the below case would be triggered as it will see the message.
				// As I mentioned in the comments earlier, there would be better ways of doing this...
				// However, it is a really good demonstration of channels.
				<-dbChannel
			case b := <-dbChannel:
				fmt.Println("Fetched result from database:")
				fmt.Println(b)
			}

		}(cacheChannel, dbChannel)
		// Allow stuff to be put in cache. I don't like this... Should I find another solution?
		time.Sleep(150 * time.Millisecond)
	}
	waitGroup.Wait()
}

func queryCache(id int, mutex *sync.RWMutex) (Book, bool) {
	// As we will not be writing to the cache here with use a Read Lock.
	// However, these should be used sparringly and used only when needed.
	// As they're more complex and expensive.
	mutex.RLock()
	b, ok := cache[id]
	mutex.RUnlock()
	return b, ok
}

func queryDatabase(id int, mutex *sync.RWMutex) (Book, bool) {
	// We use sleep to simulate a database query.
	time.Sleep(100 * time.Millisecond)
	// We actually are searching the books in book.go.
	for _, b := range books {
		if b.ID == id {
			// Lock while we write.
			mutex.Lock()
			// If we get a hit then we can add this to the cache.
			cache[id] = b
			mutex.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
