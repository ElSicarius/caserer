# caserer

A simple golang tool to help you generate wordlists based on case types.


## Usage
```bash
Usage of ./caserer:
  -e string
        Comma-separated list of file extensions to preserve during conversion (default "php,js,jsp,do,aspx")
  -f string
        File containing HTTP resources
  -l string
        Path to a dictionary file for language matching
  -t string
        Case conversion type ('snake' or 'camel') (default "snake")

```

## Example
```bash
cat /opt/lists/web/jsfiles.txt | caserer -t snake | tee /opt/lists/web/jsfilesSnake.txt
cat /opt/lists/web/phpassetnoteBIG.txt | caserer -t camel | tee /opt/lists/web/phpassetnoteBIGCamel.php
```