package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./data/test_data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	wordsMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			//cleaning the text and lowering case
			word = strings.ToLower(strings.Trim(word, ",.?!~@$%^&*()_-+={}[]"))
			wordsMap[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	//structure to work with data, holding together word and its count
	type frequencyOfWords struct {
		Word  string
		Count int
	}

	//creates slice of struct that can be sorted
	var frequencies []frequencyOfWords
	for word, count := range wordsMap {
		frequencies = append(frequencies, frequencyOfWords{word, count})
	}

	//sorts the words in descending order
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	//displays the result, here's cases included: if the there are no top 5 words
	//it will show the top N words
	if len(frequencies) == 0 {
		fmt.Println("This file does not contain repeated words or empty.")
	} else if len(frequencies) >= 5 {
		fmt.Println("Top 5 most used words:")
		for i := 0; i < 5 && i < len(frequencies); i++ {
			fmt.Printf("%s: %d\n", frequencies[i].Word, frequencies[i].Count)
		}
	} else {
		fmt.Printf("Top %d most used words:\n", len(frequencies))
		for i := 0; i < len(frequencies); i++ {
			fmt.Printf("%s: %d\n", frequencies[i].Word, frequencies[i].Count)
		}
	}

}
