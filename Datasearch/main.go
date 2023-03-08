package main

import (
	"datasearch/pkg/binary"
)

func main() {
	listOfWords := []string{"apple", "banana", "kiwi", "orange", "pear"}
	findWord := "banana"
	binary.BinarySearchWithSort(listOfWords, findWord)
}
