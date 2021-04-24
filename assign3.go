// Salih Saad ss17039@auis.edu.krd
// Alaa basim ab17177@auis.edu.krd
package main

import (
	"fmt"
	"strings"
)

var keys [676]*object
var n int
var head *object

type object struct {
	name string
	next *object
}

func main() {
	getnames()
	//print() to test the linked list insertion
	hashTable()
	deleteHashByValue()
	printHash() // to test the deletion
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
	key := 0
	for index, char := range name {
		value := int(char) - 65                           // convert to asci code starting from index 0, For example, a=0 , z=25
		if index == len(name)-2 || index == len(name)-1 { // if it is the last-second character or the last character
			if index == len(name)-2 {
				key = value * 26 // let a start from 0 to 25 and b from 26 to 51 and c from 52 to 77 and so on....
			} else {
				key += value // add the second letter value to the initial value, for example, if the last second letter is b and last letter is z then key = 1 *26 + 25 = 51
				//                                             Another Example, if the last second letter is a and the last letter is y then key = 0*26 + 24 = 24

				// This was done so 'az' would be before 'ba' because if we added their asci code value, ba would be before so we did it this way.
			}
		}

	}
	return key
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
	for i := 0; i < 676; i++ {
		c := keys[i]
		for c != nil {
			fmt.Println(c.name)
			c = c.next
		}
	}
}

func deleteHashByValue() {
	z := int('Z') - 65 // get the value of z asci code where a = 0
	z = z * 26         // let each letter have 26 keys and distinguish them by the last letter
	r := int('R') - 65 // get the value of r asci code where a = 0
	r = r * 26
	for i := z; i < z+26; i++ {
		keys[i] = nil

	}
	for i := r; i < r+26; i++ {
		keys[i] = nil
	}
}
