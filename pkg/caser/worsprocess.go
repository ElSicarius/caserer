package caser

import (
	"strings"
	"unicode"
)

func SplitIntoWords(s string, uniform bool) []string {
	var words []string
	var word strings.Builder
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsNumber(r) {
			if uniform {
				r = unicode.ToLower(r)
			}
			word.WriteRune(r)
		} else if unicode.IsUpper(r) {
			if word.Len() > 0 {
				words = append(words, word.String())
				word.Reset()
			}
			if uniform {
				r = unicode.ToLower(r)
			}
			word.WriteRune(r)
		} else {
			if word.Len() > 0 {
				words = append(words, word.String())
				word.Reset()
			}
			if uniform {
				r = unicode.ToLower(r)
			}
			if uniform && r == '-' {
				r = '_'
			}
			words = append(words, string(r)) // Include the punctuation as separate "word"
		}
	}
	if word.Len() > 0 {
		words = append(words, word.String())
	}
	return words
}

func ParseExtensions(extStr string) map[string]bool {
	exts := strings.Split(extStr, ",")
	extMap := make(map[string]bool)
	for _, ext := range exts {
		extMap[strings.TrimSpace(ext)] = true
	}
	return extMap
}

func DetectCase(s string, uniform bool) string {
	hasUnderscore := strings.Contains(s, "_")
	hasDash := strings.Contains(s, "-")
	hasUppercase := strings.ContainsFunc(s, unicode.IsUpper)

	if hasUnderscore || (uniform && hasDash) {
		return "snake"
	}
	if !hasUnderscore && hasUppercase {
		return "camel"
	}
	return "unknown"
}

func ApplyPrefixSuffix(line, prefix, suffix string, extMap map[string]bool) string {
	// Apply the prefix directly
	if prefix != "" {
		line = prefix + line
	}

	// Check for extensions and apply suffix before the extension
	if suffix != "" {
		line = InsertSuffixBeforeExtension(line, suffix, extMap)
	}

	return line
}

func InsertSuffixBeforeExtension(line, suffix string, extMap map[string]bool) string {
	for ext := range extMap {
		if strings.HasSuffix(line, "."+ext) {
			// Insert suffix before the extension
			return line[:len(line)-len(ext)-1] + suffix + "." + ext
		}
	}
	// If no extension matched, just append the suffix
	return line + suffix
}
