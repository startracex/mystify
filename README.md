---
author: "[STARTRACEX](https://github.com/startracex/mystify)"
date: 2022-11-01
update: 2023-07-30
summary: A simple site builder
---

# Mystify

## Getting Started

### URL as Path

The following rules will be followed by default

- "/,/zh,/ru" will return the readme of the root directory "./README.md,./README. zh.md,./README.ru.md"
- The document will be matched with the prefix "/docs". The path before it (such as en or zh) and the path after it (the
  relative file path) will be mapped to the local "./posts/...", for example (./posts/ru/a/b.md will be matched
  by/ru/docs/a/b)
- If the file as the path does not exist after splicing, the path will be spliced with index.md. If the file still does
  not exist, the path will be spliced with "/index.md". If the file does not exist, all documents with default values
  will be returned
- The metadata of the document can be the yaml content between three dashes and the newline. If the metadata does not
  exist, each line of the text will be used as the metadata, or the default value will be maintained
- The. yml document under each language name directory and the. yml document within the directory path will become the
  sidebar directory when using this language
- Resource paths that do not begin with "/public" will be considered as paths relative to document files

### Use yaml

You can use yaml in the document, or use the content separated by lines as the document metadata, as shown in the first
line of the document.
Yaml parses the following contents: "author, summary, date, update, title" (string type), "author" (string array type)
**Note** The author will be overwritten by the existing authors. The non-existent title will be the first level title of
the document. The author can use the markdown syntax. Please use (") to include

### Out-of-the-box

Start the service by running the executable

Set the port number and domain name by carrying the following flags

```shell
-p <port number>
-d <domain name>
-dev <Set gin mode to debug>

Example
./main \ ./main.exe -p 8080 -d localhost --dev

```

Access these pages after running

[Home](http:localhost:8080/)
[Documentation](http:localhost:8080/en/docs/)

Change the language to add '/zh' or '/ru' or '/en' at the beginning of the path (default)

Basic styles and scripts are ready in the folder
