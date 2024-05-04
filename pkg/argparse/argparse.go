package argparse

import "flag"

var (
	Prefix         = flag.String("P", "", "Prefix to add to each line")
	Suffix         = flag.String("S", "", "Suffix to add to each line")
	FilePath       = flag.String("f", "", "File containing HTTP resources")
	CaseType       = flag.String("t", "snake", "Case conversion type ('snake' or 'camel')")
	DictionaryPath = flag.String("l", "", "Path to a dictionary file for language matching")
	Extensions     = flag.String("e", "php,js,jsp,do,aspx", "Comma-separated list of file extensions to preserve during conversion")
	Uniform        = flag.Bool("u", false, "Uniformise case conversion")
)
