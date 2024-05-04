package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Command-line flags
var (
	filePath       = flag.String("f", "", "File containing HTTP resources")
	caseType       = flag.String("t", "snake", "Case conversion type ('snake' or 'camel')")
	dictionaryPath = flag.String("l", "", "Path to a dictionary file for language matching")
	extensions     = flag.String("e", "php,js,jsp,do,aspx", "Comma-separated list of file extensions to preserve during conversion")
)

var dictionaries = map[string]bool{}

func main() {
	flag.Parse()

	if *dictionaryPath != "" {
		if err := loadDictionary(*dictionaryPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading dictionary: %v\n", err)
			os.Exit(1)
		}
	}

	var scanner *bufio.Scanner
	if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		line := scanner.Text()
		currentCase := detectCase(line)
		// fmt.Printf("Detected case for '%s': %s\n", line, currentCase) // Debug output
		if currentCase == *caseType {
			fmt.Println(line) // No change needed, print original line
			// if line has upper case and the caseType is snake, create an extra line with the lower case
			if *caseType == "snake" && strings.ContainsFunc(line, unicode.IsUpper) {
				fmt.Println(strings.ToLower(line))
			}
		} else {
			converted := convertCase(line, *caseType)
			// fmt.Printf("Converted '%s' to '%s'\n", line, converted) // Debug output
			fmt.Println(converted)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

func parseExtensions(extStr string) map[string]bool {
	exts := strings.Split(extStr, ",")
	extMap := make(map[string]bool)
	for _, ext := range exts {
		extMap[strings.TrimSpace(ext)] = true
	}
	return extMap
}

func loadDictionary(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		dictionaries[strings.ToLower(word)] = true
	}
	return scanner.Err()
}

func detectCase(s string) string {
	hasUnderscore := strings.Contains(s, "_")
	hasUppercase := strings.ContainsFunc(s, unicode.IsUpper)

	if hasUnderscore { // removing : && !hasUppercase {
		return "snake"
	}
	if !hasUnderscore && hasUppercase {
		return "camel"
	}
	return "unknown"
}

func convertCase(s, caseType string) string {
	var result strings.Builder
	extensionSet := parseExtensions(*extensions) // Parse the extensions from the flag
	words := splitIntoWords(s)

	if *dictionaryPath != "" {
		newWords := []string{}
		for _, word := range words {
			// try to identify extra words, split the initial word with the found words
			// append the found words to the result
			// to match the words, iter trough the dictionary and check if the word from the dictionary is in the current word
			// if it is, split the word and append the found words to the result
			matched := false
			for dictWord := range dictionaries {
				if strings.Contains(strings.ToLower(word), dictWord) {
					matched = true
					// split the word
					// append the found words to the result
					// append the rest of the word to the newWords
					// break
					// fmt.Printf("Found word: %s in %s\n", dictWord, word)
					index := strings.Index(strings.ToLower(word), dictWord)
					if index == 0 {
						newWords = append(newWords, dictWord)
						newWords = append(newWords, word[len(dictWord):])
					} else {
						newWords = append(newWords, word[:index])
						newWords = append(newWords, dictWord)
						newWords = append(newWords, word[index+len(dictWord):])
					}
					break
				}
			}
			if !matched {
				newWords = append(newWords, word)
			}
		}
		words = newWords
	}

	// Extension handling in case functions
	switch caseType {
	case "snake":
		result.WriteString(toSnakeCase(words, extensionSet))
	case "camel":
		result.WriteString(toCamelCase(words, extensionSet))
	default:
		result.WriteString(s)
	}

	return result.String()
}

func splitIntoWords(s string) []string {
	var words []string
	var word strings.Builder
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsNumber(r) {
			word.WriteRune(unicode.ToLower(r))
		} else if unicode.IsUpper(r) {
			if word.Len() > 0 {
				words = append(words, word.String())
				word.Reset()
			}
			word.WriteRune(unicode.ToLower(r))
		} else {
			if word.Len() > 0 {
				words = append(words, word.String())
				word.Reset()

			}
			words = append(words, string(unicode.ToLower(r))) // Include the punctuation as separate "word"
		}
	}
	if word.Len() > 0 {
		words = append(words, word.String())
	}
	return words
}

func toSnakeCase(words []string, extMap map[string]bool) string {
	var result strings.Builder
	for i, word := range words {
		// Append underscore only if it's not the first word and the previous word is not a dot
		// and the current word is not in the extension map
		if i > 0 && !extMap[strings.ToLower(words[i-1])] && words[i-1] != "." {
			result.WriteRune('_')
		} else {
			// remove the last underscore (2 chars back)
			curWord := result.String()

			if len(curWord) > 2 && string(curWord[len(curWord)-2]) == "_" {

				result.Reset()
				result.WriteString(curWord[:len(curWord)-2])
				result.WriteRune('.')

			}
		}
		result.WriteString(strings.ToLower(word))
	}
	return result.String()
}

func toCamelCase(words []string, extMap map[string]bool) string {
	var result strings.Builder
	for i, word := range words {
		// Check if the current word is an extension or follows a dot
		if word == "_" {
			// Skip underscores
		} else if i > 0 && words[i-1] == "." && extMap[strings.ToLower(word)] {
			result.WriteString(strings.ToLower(word)) // Keep extensions in lowercase
		} else if i > 0 && words[i-1] != "." {
			result.WriteString(strings.Title(word)) // Capitalize if not after a dot
		} else {
			result.WriteString(strings.ToLower(word)) // First word or after a dot and not an extension
		}
	}
	return result.String()
}
