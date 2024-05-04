package dictionnary

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

var Dictionaries = map[string]bool{}

func LoadDictionary(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, strings.ToLower(word))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Sort words by length in descending order
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	// Add sorted words to the dictionaries map
	for _, word := range words {
		Dictionaries[word] = true
	}
	// print the first 20 words
	return nil
}
