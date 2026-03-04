package main

import (
	"fmt"
	"strings"
)

var sentence string = "hello world hello hello world cat yes no no cat cat cat"

func main() {
	frecuency := make(map[string]int)
	words := strings.Fields(sentence)
	for _, v := range words {
		frecuency[v]++
	}
	var word string
	fmt.Print("Word for check: ")
	fmt.Scan(&word)
	if val, ok := frecuency[word]; ok {
		fmt.Printf("'%s' appears %d times in the sentence.\n", word, val)
	} else {
		fmt.Printf("'%s' does not appear in the sentence.\n", word)
	}
}
