package quip2md

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const ZERO_WIDTH_SPACE = "\u200b"

const (
	UnorderedList string = "5"
	OrderedList   string = "6"
	CheckedList   string = "7"
)

const (
	Bold          string = "b"
	Italic        string = "i"
	Underline     string = "u"
	Strikethrough string = "del"
	CodeInline    string = "code"
	Anchor        string = "a"
	Highlight     string = "annotation"

	H1         string = "h1"
	H2         string = "h2"
	H3         string = "h3"
	CodeBlock  string = "pre"
	Paragraph  string = "p"
	BlockQuote string = "blockquote"
)

// styleWrapper is a function that can take style tag tokens (non container)
// and turn into text
//
// We pass in the token itself in case a wrapper needs access
// to the attributes in the html.StartTagToken (z.Token() can only be called once)
type styleWrapper func(t html.Token, z *html.Tokenizer) string

// genericWrap is used to simplify almost every style which surround text with char(s)
func genericWrap(surround string, z *html.Tokenizer) string {
	z.Next()
	content := fmt.Sprintf("%s%s%s", surround, z.Token().Data, surround)
	z.Next()
	return content
}

func anchorWrap(t html.Token, z *html.Tokenizer) string {
	var href string
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}
	z.Next()
	linkText := z.Token().Data
	z.Next()
	return fmt.Sprintf("[%s](%s)", linkText, href)
}

var styles = map[string]styleWrapper{
	Bold:          func(t html.Token, z *html.Tokenizer) string { return genericWrap("**", z) },
	Italic:        func(t html.Token, z *html.Tokenizer) string { return genericWrap("*", z) },
	Underline:     func(t html.Token, z *html.Tokenizer) string { return genericWrap("_", z) },
	Strikethrough: func(t html.Token, z *html.Tokenizer) string { return genericWrap("~~", z) },
	CodeInline:    func(t html.Token, z *html.Tokenizer) string { return genericWrap("`", z) },
	Anchor:        anchorWrap,
	// For highlighted text, we'll leave it alone
	Highlight: func(t html.Token, z *html.Tokenizer) string { return genericWrap("", z) },
}

// QuipToMarkdown converts HTML from Quip HTML as string to sane Markdown
//
// Supported Features:
//   * Headers 1-3
//   * Bold
//   * Italics
//   * Underline [*]
//   * Strikethrough [*]
//   * Inline Code
//   * Hyperlinks
//   * Unordered Lists
//   * Ordered Lists
//   * Checked Lists w/ check marks maintained
//   * Nested lists w/ a-z and roman numerals on unordered lists
//   * Quote Blocks w/ maintained text styling
//   * Code Blocks - including adjacent ones
//
// Caveats [*]:
//   * Not every Markdown renderer supports underline the same way. We use `_this_`
//   * Not every Markdown renderer supports strikethrough. We use `~~this~~`
//   * Highlighted text will be converted with no styling applied
//
// Funny Quip HTML quirks:
//   * Everything is within a tag. Each tag is on it's own line
//   * Between paragraphs, there's a paragraph tag with 200b (unicod zero width space)
//   * <li> ends with </span><br/></li>
//   * Type of list is dictated by the class with starting value in attr
//   * Nested lists are tracked with class=parent on <li>
//   * Checked list checks are tracked by class=checked
//   * Code blocks are dictated by class=prettyprint. All content on one line in HTML
//   * For conversion purposes, almost all metadata can be ignored
func QuipToMarkdown(q string) (string, error) {
	lines := []string{}
	r := strings.NewReader(q)
	z := html.NewTokenizer(r)

	var tt html.TokenType
	var lastTag string
	for {
		// token type
		tt = z.Next()
		if tt == html.ErrorToken {
			break
		}
		t := z.Token()
		switch tt {
		case html.StartTagToken:
			tag := t.Data
			if tag == CodeBlock {
				// Some Markdown presenters will combine j.. Add a horizontal
				// rule to separate
				if lastTag == CodeBlock {
					lines = append(lines, "", "---")
				}
				lines = append(lines, toCodeBlock(z)...)
			} else if strings.HasPrefix(tag, "h") {
				lines = append(lines, toHeader(tag, z)...)
			} else if tag == Paragraph {
				lines = append(lines, toParagraph(z)...)
			} else if tag == BlockQuote {
				lines = append(lines, toBlockQuote(z)...)
			} else if tag == "div" {
				var sectionStyle string
				for _, a := range t.Attr {
					if a.Key == "data-section-style" {
						sectionStyle = a.Val
					}
				}
				lines = append(lines, toList(z, sectionStyle, 1)...)
			}
			// case html.EndTagToken:
			// case html.SelfClosingTagToken:
			lastTag = tag
		}

	}
	return strings.Join(lines, "\n"), nil
}

// toItemPrefix takes into accuont the type of list and returns the string
// placed before the <li> content (bullet, number, roman numeral, etc)
func toItemPrefix(t html.Token, cnt int, sectionStyle string, level int) (pfx string) {
	if sectionStyle == UnorderedList {
		pfx = "* "
	} else if sectionStyle == CheckedList {
		if isItemChecked(t) {
			pfx = "[X] "
		} else {
			pfx = "[ ] "
		}
	} else if sectionStyle == OrderedList && level%3 == 1 {
		// Assuming ordered list. Remainder is used so that the levels circle
		// around. Eg, level 3 goes back to 0
		pfx = fmt.Sprintf("%d. ", cnt)
	} else if sectionStyle == OrderedList && level%3 == 2 {
		// Lowercase alphabetical ordering
		letter := string('A' - 1 + cnt)
		pfx = fmt.Sprintf("%s. ", strings.ToLower(letter))
	} else if sectionStyle == OrderedList && level%3 == 0 {
		// Roman Numerals
		letter := romanEncode(cnt)
		pfx = fmt.Sprintf("%s. ", strings.ToLower(letter))
	}
	return
}

func isItemChecked(t html.Token) (answer bool) {
	for _, a := range t.Attr {
		if a.Key == "class" && a.Val == "checked" {
			answer = true
		}
	}
	return
}

// listRecursionNeeded checks if the current list item is noted as the parent
// to another nested beneath it
func listRecursionNeeded(t html.Token) (answer bool) {
	for _, a := range t.Attr {
		if a.Key == "class" && a.Val == "parent" {
			answer = true
		}
	}
	return
}

func getStartItemCount(t html.Token) (count int) {
	var err error
	for _, a := range t.Attr {
		if a.Key == "value" {
			count, err = strconv.Atoi(a.Val)
			if err != nil {
				count = 255
			}
		}
	}
	return
}

// toList begins parsing Quip HTML list and supports nesting + all list formats
//
// * Unordered, Ordered, and Checked lists are supported
// * Nested lists of each time are supported
// * Nested ordered lists cycle from: numbers, a-z, roman numerals
// * Try to manage whitespace effectively
func toList(z *html.Tokenizer, sectionStyle string, level int) []string {
	if level == 1 {
		z.Next() // <ul> (because all lists are <ul> in Quip
	}
	combinedLines := []string{}
	var itemCount int
	var recurse bool
loop:
	for {
		tt := z.Next()
		switch tt {
		case html.StartTagToken:
			t := z.Token()
			tag := t.Data
			// If list item, add prefix (*, [ ], 1., etc) to the list to
			// prepare for parsing text
			if tag == "li" {
				// We might need to parse a nested list
				recurse = listRecursionNeeded(t)
				if itemCount == 0 && level == 1 {
					itemCount = getStartItemCount(t)
				} else if itemCount == 0 && level > 1 {
					itemCount = 1
				}
				combinedLines = append(combinedLines, toItemPrefix(t, itemCount, sectionStyle, level))
				itemCount++
				continue
				// Otherwise, tag is likely within text of list item for style
			} else if styleFunc, ok := styles[tag]; ok {
				combinedLines = append(combinedLines, styleFunc(t, z))
			}
		case html.TextToken:
			line := z.Token().Data
			combinedLines = append(combinedLines, line)

			// If previously parsed <li> told us that there was another <ul>
			// beneath, call this function recursively and indent the results
			if recurse {
				recurse = false
				nested := toList(z, sectionStyle, level+1)

				// If we're at the root, trim extra whitespace before
				// indententing. Helps get rid of unnecessary newlines
				var tmp []string
				if level == 1 {
					joined := "\n" + strings.TrimSpace(strings.Join(nested, "\n"))
					nested = strings.Split(joined, "\n")
				}

				tmp = nested[:0]
				for _, v := range nested {
					// indent the nested lines
					tmp = append(tmp, strings.Repeat(" ", level*2)+v)
				}
				// Join and remove leading spaces
				combinedLines = append(combinedLines, strings.TrimLeft(strings.Join(tmp, "\n"), " "))
			}
		case html.SelfClosingTagToken:
			// Assume break, add newline
			z.Token() // <br/>
			combinedLines = append(combinedLines, "\n")
		case html.EndTagToken:
			t := z.Token()
			data := t.Data
			if data == "ul" {
				// newline after entire list
				combinedLines = append(combinedLines, "\n")
				break loop // otherwise we'd break the switch
			} else if data == "span" {
				// Every list item ends with <span><br/> -- skip to avoid too many
				// newlines
				z.Next() // skip the <br/>
			}
		}
	}
	joined := strings.Join(combinedLines, "")
	return strings.Split(joined, "\n")
}

// parseTextLike takes types that parse similarly even though they use different tags
//
// Good examples of this would be bullet points, quoted text, and paragraphs.
// They all have to handle style tags within (bold, italic, inline-code) and
// stop at their end tag
//
// It is assumed that the starting tag token was fetched already. This is safe
// because how else do you know what the closingTag would be to call here
//
// nlPrefix is a string to prefix newlines with. Useful for BlockQuotes
func parseTextLike(z *html.Tokenizer, closingTag string, nlPrefix string) (lines []string) {
	combinedLines := []string{nlPrefix}
loop:
	for {
		tt := z.Next()
		switch tt {
		case html.TextToken:
			line := z.Token().Data
			if line == ZERO_WIDTH_SPACE {
				// These are useless for our needs
				z.Next()
				return []string{}
			}
			combinedLines = append(combinedLines, line)
		case html.StartTagToken:
			t := z.Token()
			tag := t.Data
			if styleFunc, ok := styles[tag]; ok {
				combinedLines = append(combinedLines, styleFunc(t, z))
			}
		case html.SelfClosingTagToken:
			t := z.Token()
			if t.Data != "br" {
				log.Printf("Unknown Self Closer: %v", t)
			}
			combinedLines = append(combinedLines, "\n", nlPrefix)
		// When finishing block, break so we don't parse the beginning of
		// next token
		case html.EndTagToken:
			t := z.Token()
			if t.Data == closingTag {
				combinedLines = append(combinedLines, "\n")
				break loop // otherwise we'd break the switch
			}
		}
	}
	joined := strings.Join(combinedLines, "")
	return strings.Split(joined, "\n")
}

func toParagraph(z *html.Tokenizer) (lines []string) {
	return parseTextLike(z, Paragraph, "")
}

func toBlockQuote(z *html.Tokenizer) (lines []string) {
	return parseTextLike(z, BlockQuote, "> ")
}

func toHeader(hdr string, z *html.Tokenizer) (lines []string) {
	prefix := map[string]string{
		H1: "#",
		H2: "##",
		H3: "###",
	}
	z.Next()
	text := fmt.Sprintf("%s %s", prefix[hdr], z.Token().Data)
	lines = append(lines, text)
	tt := z.Next()
	if tt != html.EndTagToken {
		log.Printf("Expected header close tag, got: %v", z.Token())
		return
	}
	lines = append(lines, "")

	return
}

func toCodeBlock(z *html.Tokenizer) (lines []string) {
	// Start with newline then code fence
	lines = append(lines, "", "```")

	// Quip puts duplicate <br/> in code blocks. Use this to track and add one
	// at a time
	brCnt := 0
	// Parse until we get to the next tag
loop:
	for tt := z.Next(); tt != html.StartTagToken; tt = z.Next() {
		switch tt {
		case html.TextToken:
			lines = append(lines, z.Token().Data)
		case html.SelfClosingTagToken:
			t := z.Token()
			if t.Data != "br" {
				log.Printf("Unknown Self Closer: %v", t)
			}
			brCnt++
			if brCnt%2 == 0 {
				lines = append(lines, "")
			}
		// When finishing code block, break so we don't parse the beginning of
		// next token
		case html.EndTagToken:
			t := z.Token()
			if t.Data != CodeBlock {
				log.Printf("Expected </pre>, got: %v", t)
			}
			lines = append(lines, "```")
			break loop // otherwise we'd break the switch
		}
	}
	return
}

func romanEncode(x int) (out string) {
	if x < 0 || x == 0 {
		return "LXIX"
	}

	for x >= 1000 {
		out += "M"
		x -= 1000
	}
	if x >= 900 {
		out += "CM"
		x -= 900
	}
	for x >= 500 {
		out += "D"
		x -= 500
	}
	for x >= 100 {
		out += "C"
		x -= 100
	}
	if x >= 90 {
		out += "XC"
		x -= 90
	}
	if x >= 50 {
		out += "L"
		x -= 50
	}
	if x >= 40 {
		out += "XL"
		x -= 40
	}
	for x >= 10 {
		out += "X"
		x -= 10
	}
	if x >= 9 {
		out += "IX"
		x -= 9
	}
	if x >= 5 {
		out += "V"
		x -= 5
	}
	if x >= 4 {
		out += "IV"
		x -= 4
	}
	for x > 0 {
		out += "I"
		x -= 1
	}
	return out
}
