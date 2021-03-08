package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this is how dynamic arrays work under the hood,
// theyre just linked lists.
type linkedList struct {
	value   int
	pointer *linkedList
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var myList *linkedList = createList(10, nil)
	// Here some funky stuff is happening, stay with me...
	// First we have nothing to initialise, as we're not using an index to interate.
	// If myList is nil, we know we have hit the end, so we print out the last record.
	// After each iteration we set myList to the pointer of the next record.
	for ; myList != nil; myList = myList.pointer {
		fmt.Printf("%v\n", myList.value)
	}
}

// Takes a number, this is how many items you want to add.
// Takes a list, can either be and existing list or an empty one.
func createList(num int, aLinkedList *linkedList) *linkedList {
	// Let's do our due diligence at all times.
	if num == 0 {
		return nil
	}

	// If the list does not exist, we create it, with nil as the pointer to the next item.
	// As we do not know num here, but if it's 1 it will create on item, and one only one...
	// ..as the for loop will be satisfied on its first iteration.
	if aLinkedList == nil {
		aLinkedList = &linkedList{rand.Intn(100), nil}
		num--
	}

	// The below took me a bit to figure out whats going on, but I think I sussed it...
	// We create another list, so we can keep track of what the last record in the list is.
	// We then return the original aLinkedList, which is the first item with its pointer updated (if needed)
	// To point to the next record along.
	var lastRecordInList *linkedList = aLinkedList

	for i := 0; i < num; i++ {
		// Create a new record, with nil incase this is the last element, we take the address of this.
		newRecord := &linkedList{rand.Intn(100), nil}
		// We assign the new record address to the pointer in the last record of the linked list.
		lastRecordInList.pointer = newRecord
		// We now assign the last record in the list to the value at the address of newRecord...
		// ... because the lastRecordInList was updated in the previous line to point to newRecord.
		lastRecordInList = lastRecordInList.pointer
	}
	return aLinkedList
}
