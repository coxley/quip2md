# Quip to Markdown: Reference MVP

This document stands as the reference conversion for Quip-to-Markdown. Formatting and style parameters within here are expected to be supported. Should they no longer be, unit tests should fail.

Adding **more** content in a _separate_ paragraph to test more than *simple* content between headings or other style ~~things~~ objects.

# H1

## H2

### H3

### Bulleted List

* **Bold**
* *Italics*
* _Underline (Appears italic because most Markdown flavors don't render underline)_
* ~~Strikethrough~~
* `Inline Code`
* [Hyperlink](https://github.com/coxley)
* Highlighted
* Item
with
linebreaks
* Primary subject
  * Sub Element
      * Nested Sub Element

### Numbered List

1. **Bold**
2. *Italics*
3. _Underline (Appears italic because most Markdown flavors don't render underline)_
4. ~~Strikethrough~~
5. `Inline Code`
6. [Hyperlink](https://github.com/coxley)
7. Highlighted
8. Primary subject
  a. Sub Element
      i. Nested Sub Element

### Check List

[ ] **Bold**
[X] *Italics*
[X] _Underline_
[ ] ~~Strikethrough~~
[ ] `Inline Code`
[ ] [Hyperlink](https://github.com/coxley)
[ ] Highlighted


### Quotes

> “How about this?' Simmon asked me. "Which is worse, stealing a pie or killing Ambrose?"
> I gave it a moment's hard thought. "A meat pie, or a fruit pie?” 
> ― **Patrick Rothfuss,** [The Wise Man's Fear](https://www.goodreads.com/work/quotes/2502882)

### Code


```
#!/usr/bin/env python3

import asyncio

async def perodic(i: int):

    while True:
        print("A period has passed")

        await asyncio.sleep(i)
        

        
asyncio.create_task(periodic(1))

asyncio.create_task(periodic(2))
loop = asyncio.get_event_loop()

loop.run_forever()

```

---

```
package main

import "fmt"

func main() {
    fmt.Println("quip to markdown")

}
```