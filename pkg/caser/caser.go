package caser

import (
	"strings"

	"github.com/ElSicarius/caserer/pkg/argparse"
	"github.com/ElSicarius/caserer/pkg/dictionnary"
)

func ConvertCase(s, caseType string) string {
	var result strings.Builder
	extensionSet := ParseExtensions(*argparse.Extensions) // Parse the extensions from the flag
	words := SplitIntoWords(s)

	if *argparse.DictionaryPath != "" {
		var newWords []string
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			longestMatch := ""
			matchIndex := -1
			// don't process extensions
			if extensionSet[lowerWord] {
				newWords = append(newWords, word)
				continue
			}

			// Find the longest dictionary match within the word
			for dictWord := range dictionnary.Dictionaries {
				if index := strings.Index(lowerWord, dictWord); index != -1 {
					// Ensure the match is the longest and properly bounded (check word boundaries)
					if len(dictWord) > len(longestMatch) {
						longestMatch = dictWord
						matchIndex = index
					}
				}
			}

			// If a longest match is found, split and add the segments to newWords
			if matchIndex != -1 {
				if matchIndex > 0 {
					newWords = append(newWords, word[:matchIndex])
				}
				newWords = append(newWords, longestMatch)
				if matchIndex+len(longestMatch) < len(word) {
					newWords = append(newWords, word[matchIndex+len(longestMatch):])
				}
			} else {
				// No match found, add the original word
				newWords = append(newWords, word)
			}
		}
		words = newWords
	}

	// Use the appropriate case conversion based on the caseType flag
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
