package line

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/aixfoundry/cobol-go/constant"
	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
)

type LinkedLine struct {
	Line
	origin                *Line
	prev                  *LinkedLine
	next                  *LinkedLine
	isTriggerCommentEntry bool
	isInCommentEntry      bool
}

func NewLinkedLine(r io.Reader, opts ...options.Option) (ll *LinkedLine) {
	o := options.NewOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	f := o.Format
	dialect := o.Dialect

	scan := bufio.NewScanner(r)
	var last *LinkedLine
	no := 0
	for scan.Scan() {
		l := NewLine(scan.Text(), no, f, dialect)
		if l == nil {
			continue
		}
		tmp := &LinkedLine{
			Line:   *l,
			origin: l,
			prev:   last,
		}
		if last == nil {
			ll = tmp
		} else {
			last.next = tmp
		}
		last = tmp
		no++
	}
	ll = normalizesLines(ll)
	return
}

func Combine(ll *LinkedLine) (code string) {
	source := ll
	for {
		if source == nil {
			break
		}
		if source.Type != CONTINUATION {
			if source.No > 0 {
				code += "\n"
			}
			code += LinePrefix(source.Format)
			code += source.Indicator
		}
		code += source.Content()
		source = source.next
	}
	return
}

func LinePrefix(f format.Format) (ret string) {
	if f != format.TANDEM && f != format.FREE {
		ret += strings.Repeat(constant.CHAR_WHITESPACE, 6)
	}
	return ret
}

func normalizesLines(source *LinkedLine) (target *LinkedLine) {
	target = source
	for {
		if source == nil {
			break
		}
		source.processIndicator()
		source.processInlineComment()
		source.processCommentEntry()
		source = source.next
	}
	source = target
	return
}

var (
	triggersStart = []*regexp.Regexp{
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.AUTHOR, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.INSTALLATION, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.DATE_WRITTEN, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.DATE_COMPILED, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.SECURITY, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.REMARKS, constant.CHAR_DOT)),
	}
	triggersEnd = []*regexp.Regexp{
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.PROGRAM_ID, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.AUTHOR, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.INSTALLATION, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.DATE_WRITTEN, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.DATE_COMPILED, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.SECURITY, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*.*`, constant.ENVIRONMENT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.DATA, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.PROCEDURE, constant.CHAR_DOT)),
		regexp.MustCompile(fmt.Sprintf(`^%s\s*\%s`, constant.REMARKS, constant.CHAR_DOT)),
	}
	triggersGroup = regexp.MustCompile(fmt.Sprintf(
		`([ \\t]*)(%s|%s|%s|%s|%s|%s)(.+)`,
		fmt.Sprintf(`%s\s*\%s`, constant.AUTHOR, constant.CHAR_DOT),
		fmt.Sprintf(`%s\s*\%s`, constant.INSTALLATION, constant.CHAR_DOT),
		fmt.Sprintf(`%s\s*\%s`, constant.DATE_WRITTEN, constant.CHAR_DOT),
		fmt.Sprintf(`%s\s*\%s`, constant.DATE_COMPILED, constant.CHAR_DOT),
		fmt.Sprintf(`%s\s*\%s`, constant.SECURITY, constant.CHAR_DOT),
		fmt.Sprintf(`%s\s*\%s`, constant.REMARKS, constant.CHAR_DOT),
	))

	// Pre-compiled regexes for hot paths
	reInlineComment     = regexp.MustCompile(`\*>[^ ]`)
	reDoubleQuote       = regexp.MustCompile(`"([^"]|""|'')*"`)
	reSingleQuote       = regexp.MustCompile(`'([^']|''|"")*'`)
	reLeadingSpace      = regexp.MustCompile(`^\s+`)
	reTrailingSpace     = regexp.MustCompile(`\s+$`)
)

func isTriggerStart(ctt string) bool {
	upper := strings.ToUpper(strings.TrimSpace(ctt))
	for _, v := range triggersStart {
		if v.MatchString(upper) {
			return true
		}
	}
	return false
}

func isTriggerEnd(ctt string) bool {
	upper := strings.ToUpper(ctt)
	for _, v := range triggersEnd {
		if v.MatchString(upper) {
			return true
		}
	}
	return false
}

func (ll *LinkedLine) processCommentEntry() {
	//
	// If the Compiler directive SOURCEFORMAT is specified as or defaulted to FIXED,
	// the comment-entry can be contained on one or more lines but is restricted to
	// area B of those lines; the next line commencing in area A begins the next
	// non-comment entry.
	//
	if ll.Format.MultiLine() {
		ll.processMultiLineCommentEntry()
	} else {
		ll.processSingleLineCommentEntry()
	}
}

func (ll *LinkedLine) processMultiLineCommentEntry() {
	isTriggerCommentEntry := isTriggerStart(ll.Content())
	if isTriggerCommentEntry {
		ll.escapeCommentEntry()
	} else if ll.prev != nil && (ll.prev.isTriggerCommentEntry || ll.prev.isInCommentEntry) {
		isContentAEmpty := strings.TrimSpace(ll.ContentA) == ""
		isInOsvsCommentEntry := ll.Dialect == format.OSVS && !isTriggerEnd(ll.Content())
		ll.isInCommentEntry = ll.Type == COMMENT || isContentAEmpty || isInOsvsCommentEntry
		if ll.isInCommentEntry {
			ll.Indicator = constant.COMMENT_ENTRY_TAG + constant.CHAR_WHITESPACE
		}
	}
	ll.isTriggerCommentEntry = isTriggerCommentEntry
}

func (ll *LinkedLine) processSingleLineCommentEntry() {
	isTriggerCommentEntry := isTriggerStart(ll.Content())
	if isTriggerCommentEntry {
		ll.escapeCommentEntry()
	}
}

func (ll *LinkedLine) escapeCommentEntry() {
	ctt := ll.Content()
	groups := triggersGroup.FindAllStringSubmatch(ctt, -1)
	if len(groups) != 0 {
		group := groups[0]
		if len(group) == 4 {
			ws := group[1]
			trigger := group[2]
			commentEntry := group[3]
			ctt = ws + trigger + constant.CHAR_WHITESPACE + constant.COMMENT_ENTRY_TAG + commentEntry
			ll.SetContent(ctt)
		}
	}
}

func (ll *LinkedLine) processInlineComment() {
	content := ll.Content()
	if reInlineComment.MatchString(content) {
		content = strings.ReplaceAll(content, constant.COMMENT_TAG, constant.COMMENT_TAG+constant.CHAR_WHITESPACE)
		ll.SetContent(content)
	}
}

func (ll *LinkedLine) processIndicator() {
	content := ll.ContentTrimSuffix()
	trimedContent := TrimPrefix(content)
	switch ll.Type {
	case DEBUG:
		ll.Indicator = constant.CHAR_WHITESPACE
		ll.SetContent(content)
	case CONTINUATION:
		if len(content) == 0 {
			ll.Indicator = constant.CHAR_WHITESPACE
			ll.SetContent(constant.EMPTY_STRING)
		} else if ll.prev != nil {
			prev := ll.prev
			prevContent := prev.Content()
			prevOriginContent := prev.origin.Content()
			if strings.HasSuffix(prevOriginContent, "\"") ||
				strings.HasSuffix(prevOriginContent, "'") {
				/*
					上一行以引号结尾，当前行以引号开始
						"xxxxxxxxxxxx"
					-   "xxxxxxxxxxxx"

					上一行以引号结尾，当前行不以引号开始
						"xxxxxxxxxxxx"
					-   xxxxxxxxxxxx"
				*/
				if strings.HasPrefix(trimedContent, "\"") ||
					strings.HasPrefix(trimedContent, "'") {
					trimedContent = trimedContent[1:]
				}

				ll.Indicator = constant.CHAR_WHITESPACE
				ll.SetContent(trimedContent)
			} else if prev.IsEndingWithOpenLiteral() {
				/*
					上一行以开放引号结束，当前行以引号开始
						"xxxxxxxxxxxx
					-   "xxxxxxxxxxxx"

					上一行以开放引号结束，当前行不以引号开始
						"xxxxxxxxxxxx
					-   xxxxxxxxxxxx"
				*/
				if strings.HasPrefix(trimedContent, "\"") ||
					strings.HasPrefix(trimedContent, "'") {
					trimedContent = trimedContent[1:]
					ll.Indicator = constant.CHAR_WHITESPACE
					ll.SetContent(trimedContent)
				} else {
					ll.Indicator = constant.CHAR_WHITESPACE
					ll.SetContent(content)
				}
			} else if strings.HasSuffix(prevContent, "\"") ||
				strings.HasSuffix(prevContent, "'") {
				trimedContent = constant.CHAR_WHITESPACE + trimedContent
				ll.Indicator = constant.CHAR_WHITESPACE
				ll.SetContent(trimedContent)
			} else {
				ll.Indicator = constant.CHAR_WHITESPACE
				ll.SetContent(trimedContent)
			}
		} else {
			ll.Indicator = constant.CHAR_WHITESPACE
			ll.SetContent(trimedContent)
		}
	case COMMENT:
		// [*]    =>    [*> ]
		ll.Indicator = constant.COMMENT_TAG + constant.CHAR_WHITESPACE
		ll.SetContent(content)
	case COMPILER_DIRECTIVE:
		ll.Indicator = constant.CHAR_WHITESPACE
		ll.SetContent(constant.EMPTY_STRING)
	case NORMAL:
		ll.Indicator = constant.CHAR_WHITESPACE
		ll.SetContent(content)
	default:
		ll.Indicator = constant.CHAR_WHITESPACE
		ll.SetContent(content)
	}
}

func (ll *LinkedLine) IsNextContinuation() bool {
	if ll.next == nil {
		return false
	}
	if ll.next.Type == CONTINUATION {
		return true
	}
	return false
}

// IsEndingWithOpenLiteral 判断是否是开放字符串
//
//	example:
//
//		MOVE "11111111111111111111
//		22222222222222222222
//		33333333333333333333" TO DATA
func (ll *LinkedLine) IsEndingWithOpenLiteral() bool {
	content := reDoubleQuote.ReplaceAllString(ll.origin.Content(), constant.EMPTY_STRING)
	content = reSingleQuote.ReplaceAllString(content, constant.EMPTY_STRING)
	return strings.Contains(content, constant.CHAR_DOUBLE_QUOTE) ||
		strings.Contains(content, constant.CHAR_SINGLE_QUOTE)
}

func TrimPrefix(ctt string) string {
	return reLeadingSpace.ReplaceAllString(ctt, constant.EMPTY_STRING)
}

func (ll *LinkedLine) ContentTrimPrefix() string {
	return TrimPrefix(ll.origin.Content())
}

func TrimSuffix(ctt string) string {
	return reTrailingSpace.ReplaceAllString(ctt, constant.EMPTY_STRING)
}

func (ll *LinkedLine) ContentTrimSuffix() string {
	content := TrimSuffix(ll.origin.Content())
	// 当末尾是逗号或者分号时补一个空格
	if len(content) != 0 {
		c := content[len(content)-1]
		switch string(c) {
		case constant.CHAR_COMMA:
			content += constant.CHAR_WHITESPACE
		case constant.CHAR_SEMICOLON:
			content += constant.CHAR_WHITESPACE
		}
	}

	// When this line is continued and ends with an open literal,
	// use the original untrimmed content to preserve the literal.
	if ll.IsNextContinuation() && ll.IsEndingWithOpenLiteral() {
		content = ll.origin.Content()
	}
	return content
}
