package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ElSicarius/caserer/pkg/argparse"
	"github.com/ElSicarius/caserer/pkg/caser"
	"github.com/ElSicarius/caserer/pkg/dictionnary"
)

// Command-line flags

func main() {
	flag.Parse()

	if *argparse.DictionaryPath != "" {
		if err := dictionnary.LoadDictionary(*argparse.DictionaryPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading dictionary: %v\n", err)
			os.Exit(1)
		}
	}

	var scanner *bufio.Scanner
	if *argparse.FilePath != "" {
		file, err := os.Open(*argparse.FilePath)
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

		if *argparse.Uniform {
			line = caser.ApplyPrefixSuffix(line, *argparse.Prefix, *argparse.Suffix, caser.ParseExtensions(*argparse.Extensions))
		}

		var finalLine string
		finalLine = caser.ConvertCase(line, *argparse.CaseType, *argparse.Uniform)

		if !*argparse.Uniform {
			// Use the new function to apply prefix and suffix properly
			finalLine = caser.ApplyPrefixSuffix(finalLine, *argparse.Prefix, *argparse.Suffix, caser.ParseExtensions(*argparse.Extensions))
		}

		fmt.Println(finalLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
