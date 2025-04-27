package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func countWords(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			word = strings.ToLower(word)
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordCount, nil
}

func getTopWords(wordCount map[string]int) []string {
	var wordPairs []struct {
		Word  string
		Count int
	}

	for word, count := range wordCount {
		wordPairs = append(wordPairs, struct {
			Word  string
			Count int
		}{Word: word, Count: count})
	}

	sort.Slice(wordPairs, func(i, j int) bool {
		return wordPairs[i].Count > wordPairs[j].Count
	})

	var topWords []string
	for i := 0; i < 5 && i < len(wordPairs); i++ {
		topWords = append(topWords, fmt.Sprintf("%s: %d", wordPairs[i].Word, wordPairs[i].Count))
	}

	return topWords
}

func main() {
	filename := "sample.txt"

	wordCount, err := countWords(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	topWords := getTopWords(wordCount)

	fmt.Println("Top 5 most frequent words:")
	for _, word := range topWords {
		fmt.Println(word)
	}
}
