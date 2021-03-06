# quip2md [![GoDoc](https://godoc.org/github.com/coxley/quip2md?status.svg)](https://godoc.org/github.com/coxley/quip2md) [![Build Status](https://travis-ci.org/coxley/quip2md.svg?branch=master)](https://travis-ci.org/coxley/quip2md)

Quip is a cloud documentation suite which provides a web API for interacting
with it. While the UI has export options for different formats, it all happens
client-side. The API only returns HTML which is heavily customized to work with
their styling. Unfortunately this means no markdown converter knows how to
act when processing.

There's an [existing issue](https://github.com/quip/quip-api/issues/8) on the
[quip-api](https://github.com/quip/quip-api/) repo but it doesn't seem to be a
priority. That's where `quip2md` comes in!

## Supported Features

* Headers 1-3
* Bold
* Italics
* Underline [*]
* Strikethrough [*]
* Inline Code
* Hyperlinks
* Unordered Lists
* Ordered Lists
* Checked Lists w/ check marks maintained
* Nested lists w/ a-z and roman numerals on unordered lists
* Quote Blocks w/ maintained text styling
* Code Blocks - including adjacent ones

### Caveats [*]
* Not every Markdown renderer supports underline the same way. We use `_this_`
* Not every Markdown renderer supports strikethrough. We use `~~this~~`
* Highlighted text will be converted with no styling applied

## Download

```shell
go get -u github.com/coxley/quip2md/...
```

## Usage

`quip2md` can be used as both a library in your Go program or from the CLI.
This is done to enable any project regardless of implementation language to
take advantage

If you're using the [go-quip](https://github.com/mduvall/go-quip) library,
usage might look like the following:

```golang
package main

import (
	"fmt"

	"github.com/coxley/quip2md"
	quip "github.com/mduvall/go-quip"
)

const quip_token = "YOUR_API_TOKEN"
const quip_ref_doc = "EXAMPLE_DOCUMENT_THREAD_ID"

func main() {
	q := quip.NewClient(quip_token)
	thread := q.GetThread(quip_ref_doc)
	md, _ := quip2md.QuipToMarkdown(thread.Html)
	fmt.Println(md)

}
```


From the CLI:

```
$ quip2md --help
NAME:
   quip2md - Convert Quip HTML to Markdown

USAGE:
   quip2md [file to convert]
   cat [file to convert] | quip2md
   custom_program | quip2md

```
