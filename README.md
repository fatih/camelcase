# Camelcase [![GoDoc](https://godoc.org/github.com/fatih/camelcase?status.png)](http://godoc.org/github.com/fatih/camelcase) [![Build Status](https://travis-ci.org/fatih/camelcase.png)](https://travis-ci.org/fatih/camelcase)

Camelcase is a micro package to split the words of a camelcase type string into
a slice of words. Use it to convert a camelcase into any type of word.

## Install

```bash
go get github.com/fatih/camelcase
```

## Examples

```go
splitted := camelcase.Split("GolangPackage")

splitted[0] => "Golang"
splitted[1] => "Package"
```

Supported cases:

```
lowercase =>       ["lowercase"]
Class =>           ["Class"]
MyClass =>         ["My", "Class"]
MyC =>             ["My", "C"]
HTML =>            ["HTML"]
PDFLoader =>       ["PDF", "Loader"]
AString =>         ["A", "String"]
SimpleXMLParser => ["Simple", "XML", "Parser"]
vimRPCPlugin =>    ["vim", "RPC", "Plugin"]
GL11Version =>     ["GL", "11", "Version"]
99Bottles =>       ["99", "Bottles"]
May5 =>            ["May", "5"]
BFG9000 =>         ["BFG", "9000"]
```
