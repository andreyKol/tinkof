package main

import (
	"fmt"
	"unicode"
)

func main() {
	var word string

	if _, err := fmt.Scanln(&word); err != nil {
		fmt.Println("Строка пустая:", err)
		return
	}

	if len(word) < 1 || len(word) > 2e5 {
		fmt.Println("Неверная длина строки")
		return
	}

	for _, c := range word {
		if !unicode.IsLower(c) {
			fmt.Println("Строка содержит большие буквы латинского алфавита")
			return
		}
	}

	fmt.Println(WordCount(word))
}

func WordCount(s string) int {
	charCount := make(map[rune]int)
	for _, c := range s {
		charCount[c]++
	}
	wordCount := 0
	word := "sheriff"

	for {
		for _, c := range word {
			if count, ok := charCount[c]; ok && count > 0 {
				charCount[c]--
			} else {
				return wordCount
			}
		}
		wordCount++
	}
}
