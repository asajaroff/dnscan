package main

import (
	"fmt"
)

var chars = []string{"a","b","c"}

func dictionaryCreate() {
	for i, _ := range chars {
		fmt.Printf("%s\n", chars[i])
	}
}

func createWord() string {
	// Iterate through the amount of characters defined on chars
	for n := 0; n <= len(chars); n++ {
		for i, _ := range chars {
			fmt.Println(chars[i])
		}
	}

//	for i, _ := range chars {
//		for o, _ := range chars {
//			currentWord = chars[i] + chars[o]
//			log.Println(currentWord)
//		}
//	}
	currentWord := "string"
	return currentWord
}
