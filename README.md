# caserer

A simple golang tool to help you generate wordlists based on case types.

I know there are other tools that can do this, but they where not satisfying my needs, so i made my own.


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

## Cases examples
### Without language dictionnary
Base wordlist:
```bash
getUser.php
get_info.php
getresources.php
SetResources.php
get_People.php
```

```bash
╰─➤  cat ../../examples/wordlist.txt | caserer -t camel
getUser.php
getInfo.php
getresources.php
SetResources.php
getPeople.php
```

```bash
╰─➤  cat ../../examples/wordlist.txt | caserer -t snake
get_user.php
get_info.php
getresources.php
set_resources.php
get_People.php
get_people.php
```

### Prefix/Suffix adding

Base wordlist:
```bash
getUser.php
get_info.php
getresources.php
SetResources.php
get_People.php
```
```bash
╰─➤  cat ../../examples/wordlist.txt | caserer -t snake -S _suffix                     
get_user_suffix.php
get_info_suffix.php
getresources_suffix.php
set_resources_suffix.php
get_People_suffix.php
```

```bash
cat ../../examples/wordlist.txt | caserer -t snake -P prefix_
prefix_get_user.php
prefix_get_info.php
prefix_getresources.php
prefix_set_resources.php
prefix_get_People.php
```

### With language dictionnary

Base wordlist:
```bash
getUser.php
get_info.php
getresources.php
SetResources.php
get_People.php
get.resource.php
test.getterResource.php
infophp.php
phpinfo.php
testingpurpose.php
```

```bash
caserer -f ../../examples/wordlist.txt -t snake -l /opt/lists/web/englishwords.txt                                    
get_user.php
get_info.php
get_resources.php
set_resources.php
get_People.php
get.resource.php
test.getter_resource.php
info_php.php
phpinfo.php
testing_purpose.php
```
    
```bash
caserer -f ../../examples/wordlist.txt -t camel -l /opt/lists/web/englishwords.txt
getUser.php
getInfo.php
getResources.php
SetResources.php
getPeople.php
get.resource.php
test.getterResource.php
infoPhp.php
phpInfo.php
testingPurpose.php
```

**limits**:
- The language dictionnary is not perfect and can't handle all cases
- The matching algorith is simple for speed purpose, i'm not doing any kind of fuzzy matching