// Salih Saad ss17039@auis.edu.krd
// Alaa basim ab17177@auis.edu.krd
package main

import (
	"fmt"
	"strings"
)

var keys [26]*object
var n int
var head *object

type object struct {
	name string
	next *object
}

func main() {
	fmt.Println("Hello, playground")
	getnames()
	print()
	hashTable()
	printHash()
}

func getnames() {
	fmt.Println("Enter the number of names: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var name string
		fmt.Println(" Enter name Number ", i+1, "  : ")
		fmt.Scan(&name)
		storeNames(name)
	}
}

func storeNames(nam string) {

	// create an object
	n := object{name: nam, next: nil}

	// check if it is the first name:
	if head == nil {
		head = &n
	} else {
		c := head
		for c.next != nil {
			c = c.next
		}
		c.next = &n
	}

}

func print() {
	c := head

	for c != nil {
		fmt.Println(c.name)
		c = c.next
	}
}

func hash(n *object) int {
	//https://stackoverflow.com/questions/18130859/how-can-i-iterate-over-a-string-by-runes-in-go
	name := strings.ToUpper(n.name)
	println(name)
	key := 0
	for index, char := range name {
		if index == len(name)-2 { // if it is the last-second character
			println(char, index)
			key = int(char) // convert to asci code
			break
		}

	}
	return key - 65 // because A = 65
}

func hashTable() {
	c := head

	for c != nil {
		k := hash(c)
		nextObject := c.next // get the next object
		c.next = keys[k]     // will point at nothing in the first iteration , point at the previous object in other iteration
		keys[k] = c          // let the key point at the last in object
		c = nextObject       // iterate to the next object
	}

}

func printHash() {
	for i := 0; i < 26; i++ {
		c := keys[i]
		for c != nil {
			fmt.Println(c.name)
			c = c.next
		}
	}
}
