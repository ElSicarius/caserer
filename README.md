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

## Install

I'm providing a precompiled binary for linux, you can download it from the release page.

If you want to compile it yourself, you can use the following command:
```bash
git clone git@github.com:ElSicarius/caserer.git
cd caserer/cmd/caserer
go build .
cp caserer /usr/local/bin
caserer -h
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
### Uniformizing the wordlist
Base wordlist:
```bash
getUser.php
GetUser.php
Get_User.php
get_info.php
getresources.php
SetResources.php
get_People.php
get.resource.php
test.getterResource.php
infophp.php
phpinfo.php
testingpurpose.php
testing-purpose.php
```
Result 
```bash
╰─➤  caserer -f ../../examples/wordlist.txt -t snake -u                                    
get_user.php
get_user.php
get_user.php
get_info.php
getresources.php
set_resources.php
get_people.php
get.resource.php
test.getter_resource.php
infophp.php
phpinfo.php
testingpurpose.php
testing_purpose.php
```
/!\ You mush pipe it to sort -uV to remove duplicates

```bash
╰─➤  caserer -f ../../examples/wordlist.txt -t snake -u | sort -uV
get.resource.php
getresources.php
get_info.php
get_people.php
get_user.php
infophp.php
phpinfo.php
set_resources.php
testingpurpose.php
testing_purpose.php
test.getter_resource.php
```

Also, this will affect (uniformize) prefixes and suffixes output:
```bash
╰─➤  caserer -f ../../examples/wordlist.txt -t snake -u -P ajax_ | sort -uV
ajax_get.resource.php
ajax_getresources.php
ajax_get_info.php
ajax_get_people.php
ajax_get_user.php
ajax_infophp.php
ajax_phpinfo.php
ajax_set_resources.php
ajax_testingpurpose.php
ajax_testing_purpose.php
ajax_test.getter_resource.php
```

```bash
caserer -f ../../examples/wordlist.txt -t camel -u -P ajax_ | sort -uV
ajaxGet.resource.php
ajaxGetInfo.php
ajaxGetPeople.php
ajaxGetUser.php
ajaxGetresources.php
ajaxInfophp.php
ajaxPhpinfo.php
ajaxSetResources.php
ajaxTest.getterResource.php
ajaxTestingPurpose.php
ajaxTestingpurpose.php
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