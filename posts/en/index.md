# Use markdown in go

2022-09-30
2022-10-30
[STARTRACE](/)

## Getting Started

### Use directory

Write directory file (saved in ./posts/zh/.yml)
The first item is the file name with suffix removed, and the second item is the title

```yml
index: "Docs"
func: -
```

The folder does not need a title name, which is determined by the index written by. yml under it, but needs to be specified here

The first item that does not exist will be ignored

### Rendering

The go application will parse out the document and directory structure converted to markup language

javaScript will perform other functions of the page

***Code highlighting and theme colors***

Third party javascript plug-in (Highlight. js) and style parser (Less. js) are used

### Document Rules

|Property | Line | Type|
| :---:     | :----: |:---:|
|Window title | 1 | string, remove "# "|
|Center title | 1 | string, remove "# "|
|Description/Summary | 2 | Original String|
|Creation date | 3 | original string|
|Update Date | 4 | Original String|
|Author | 5 | Markup Language|
|Document |1∩[6,∞)| Markup Language|
