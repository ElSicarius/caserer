package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ElSicarius/caserer/pkg/argparse"
	"github.com/ElSicarius/caserer/pkg/caser"
	"github.com/ElSicarius/caserer/pkg/dictionnary"
	"github.com/zenthangplus/goccm"
)

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

	// Define max number of concurrent goroutines
	ccm := goccm.New(60) // limits to 4 concurrent goroutines

	for scanner.Scan() {
		line := scanner.Text()
		ccm.Wait() // wait if there are already max goroutines running

		go func(text string) {
			defer ccm.Done() // mark this goroutine as completed on exit
			processedLine := processLine(text)
			fmt.Println(processedLine) // handle output; consider synchronization if order matters
		}(line)
	}

	ccm.WaitAllDone() // wait for all goroutines to complete

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

func processLine(line string) string {
	if *argparse.Uniform {
		line = caser.ApplyPrefixSuffix(line, *argparse.Prefix, *argparse.Suffix, caser.ParseExtensions(*argparse.Extensions))
	}

	finalLine := caser.ConvertCase(line, *argparse.CaseType, *argparse.Uniform)

	if !*argparse.Uniform {
		finalLine = caser.ApplyPrefixSuffix(finalLine, *argparse.Prefix, *argparse.Suffix, caser.ParseExtensions(*argparse.Extensions))
	}

	return finalLine
}
