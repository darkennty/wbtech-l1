package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueWordsMap := make(map[string]struct{})
	uniqueWordsSlice := make([]string, 0) // Сохраняем порядок слов

	for _, word := range words {
		if _, ok := uniqueWordsMap[word]; !ok {
			uniqueWordsMap[word] = struct{}{}
			uniqueWordsSlice = append(uniqueWordsSlice, word)
		}
	}

	fmt.Println("New word set: ", uniqueWordsSlice)
}
