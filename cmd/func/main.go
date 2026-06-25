package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/document"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
)

func main() {
	f, err := os.OpenFile("./func.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	buff, err := io.ReadAll(os.Stdin)
	if err != nil {
		f.WriteString(err.Error() + "\n")
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	processed := string(buff)
	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	listener := NewFuncListener(cts)
	antlr.ParseTreeWalkerDefault.Walk(listener, cpp.StartRule())

	output := listener.GetParagraphs()
	buff, err = json.Marshal(output)
	if err != nil {
		f.WriteString(err.Error() + "\n")
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	n, err := fmt.Fprintf(os.Stdout, "%s", string(buff))
	if err != nil {
		f.WriteString(err.Error() + "\n")
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	f.WriteString(fmt.Sprintf("%d - %d - %d\n", len(processed), len(buff), n))
}

type Tok struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Context struct {
	buffer string
	tokens []*Tok
}

func NewContext() *Context {
	return &Context{
		buffer: "",
	}
}

func (ctx *Context) Read() string {
	return ctx.buffer
}

func (ctx *Context) Write(s string) {
	ctx.buffer += s
}

func (ctx *Context) GetTokens() []*Tok {
	return ctx.tokens
}

func (ctx *Context) AppendTokens(tokens []*Tok) {
	ctx.tokens = append(ctx.tokens, tokens...)
}

type Paragraph struct {
	Name   string `json:"name"`
	Text   string `json:"text"`
	Tokens []*Tok `json:"tokens"`
}

type FuncListener struct {
	cobol85.BaseCobol85Listener
	cts        *antlr.CommonTokenStream
	contexts   []*Context
	paragraphs []*Paragraph
}

func NewFuncListener(cts *antlr.CommonTokenStream) *FuncListener {
	return &FuncListener{
		contexts:   []*Context{NewContext()},
		cts:        cts,
		paragraphs: []*Paragraph{},
	}
}

func (s *FuncListener) push() {
	s.contexts = append([]*Context{NewContext()}, s.contexts...)
}

func (s *FuncListener) pop() {
	if len(s.contexts) > 0 {
		s.contexts = s.contexts[1:]
	}
}

func (s *FuncListener) peek() *Context {
	if len(s.contexts) == 0 {
		return nil
	}
	return s.contexts[0]
}

func (s *FuncListener) Context() *Context {
	return s.peek()
}

func (s *FuncListener) GetText() string {
	return s.Context().Read()
}

func (s *FuncListener) GetParagraphs() []*Paragraph {
	return s.paragraphs
}

func (s *FuncListener) EnterParagraph(ctx *cobol85.ParagraphContext) {
	s.push()
}

func (s *FuncListener) ExitParagraph(ctx *cobol85.ParagraphContext) {
	ctxt := s.Context()
	paragraph := &Paragraph{
		Name:   ctx.ParagraphName().GetText(),
		Text:   ctxt.Read(),
		Tokens: ctxt.GetTokens(),
	}
	s.paragraphs = append(s.paragraphs, paragraph)
	s.pop()
}

// VisitTerminal is called when a terminal node is visited.
func (s *FuncListener) VisitTerminal(node antlr.TerminalNode) {
	pos := node.GetSourceInterval().Start
	tok := node.GetSymbol()
	tt := tok.GetTokenType()
	if tt != cobol85.Cobol85ParserCOMMENTENTRYLINE &&
		tt != cobol85.Cobol85ParserCOMMENTLINE &&
		tt != cobol85.Cobol85ParserEOF {
		text := strings.TrimSpace(tok.GetText())
		if tt == cobol85.Cobol85ParserDOT_FS ||
			tt == cobol85.Cobol85ParserDOT {
			s.Context().AppendTokens([]*Tok{
				{
					Type:  "DOT",
					Value: text,
				},
				{
					Type:  "NEWLINE",
					Value: "\n",
				},
			})
		} else if tt == cobol85.Cobol85ParserNONNUMERICLITERAL {
			s.Context().AppendTokens([]*Tok{
				{
					Type:  "NONNUMERICLITERAL",
					Value: text,
				},
			})
		} else {
			s.Context().AppendTokens([]*Tok{
				{
					Type:  "RULE",
					Value: text,
				},
			})
		}
	}

	left := document.GetHiddenTokensToLeft(s.cts, pos)
	s.Context().Write(left)
	if node.GetSymbol().GetTokenType() != antlr.TokenEOF {
		s.Context().Write(node.GetText())
	}
}
