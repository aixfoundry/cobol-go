package main

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"strings"

	"github.com/aixfoundry/cobol-go/document"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/options"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	f, err := os.OpenFile("./dobf.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
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

	listener := NewDobfListener(cts)
	antlr.ParseTreeWalkerDefault.Walk(listener, cpp.StartRule())
	vars := map[string]string{}
	maps.Copy(vars, listener.GetVars())
	maps.Copy(vars, listener.GetFuncs())
	output := map[string]any{
		"text":   listener.GetText(),
		"vars":   vars,
		"tokens": listener.GetTokens(),
	}
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

type DobfListener struct {
	cobol85.BaseCobol85Listener
	contexts []*Context
	cts      *antlr.CommonTokenStream
	opts     *options.Options
	vars     map[string]int
	varIndex int

	funcs     map[string]int
	funcIndex int
}

func NewDobfListener(cts *antlr.CommonTokenStream, opts ...options.Option) *DobfListener {
	o := options.NewOptions()
	for _, opt := range opts {
		opt.Apply(o)
	}
	return &DobfListener{
		contexts: []*Context{NewContext()},
		opts:     o,
		cts:      cts,
		vars:     map[string]int{},
		funcs:    map[string]int{},
	}
}

func (s *DobfListener) push() {
	s.contexts = append([]*Context{NewContext()}, s.contexts...)
}

func (s *DobfListener) pop() {
	if len(s.contexts) > 0 {
		s.contexts = s.contexts[1:]
	}
}

func (s *DobfListener) peek() *Context {
	if len(s.contexts) == 0 {
		return nil
	}
	return s.contexts[0]
}

func (s *DobfListener) Context() *Context {
	return s.peek()
}

func (s *DobfListener) GetText() string {
	return s.Context().Read()
}

func (s *DobfListener) GetVars() (out map[string]string) {
	out = map[string]string{}
	for k, v := range s.vars {
		out[fmt.Sprintf("VAR_%d", v)] = k
	}
	return
}

func (s *DobfListener) GetFuncs() (out map[string]string) {
	out = map[string]string{}
	for k, v := range s.funcs {
		out[fmt.Sprintf("FUNC_%d", v)] = k
	}
	return
}

func (s *DobfListener) GetTokens() []*Tok {
	return s.Context().GetTokens()
}

// EnterParagraph is called when production paragraph is entered.
func (s *DobfListener) EnterParagraphName(ctx *cobol85.ParagraphNameContext) {
	s.push()
}

// ExitParagraph is called when production paragraph is exited.
func (s *DobfListener) ExitParagraphName(ctx *cobol85.ParagraphNameContext) {
	popped := s.Context().Read()
	s.pop()
	tree := ctx.GetParent()
	oldFunc := ctx.GetText()
	newFunc := ""
	switch tree.(type) {
	case cobol85.IParagraphContext:
		i, ok := s.funcs[oldFunc]
		if !ok {
			s.funcIndex++
			s.funcs[oldFunc] = s.funcIndex
			newFunc = fmt.Sprintf("FUNC_%d", s.funcIndex)
		} else {
			newFunc = fmt.Sprintf("FUNC_%d", i)
		}
	case cobol85.IProcedureNameContext:
		i, ok := s.funcs[oldFunc]
		if !ok {
			s.funcIndex++
			s.funcs[oldFunc] = s.funcIndex
			newFunc = fmt.Sprintf("FUNC_%d", s.funcIndex)
		} else {
			newFunc = fmt.Sprintf("FUNC_%d", i)
		}
	case cobol85.IQualifiedDataNameFormat2Context:
		i, ok := s.funcs[oldFunc]
		if !ok {
			s.funcIndex++
			s.funcs[oldFunc] = s.funcIndex
			newFunc = fmt.Sprintf("FUNC_%d", s.funcIndex)
		} else {
			newFunc = fmt.Sprintf("FUNC_%d", i)
		}
	default:
		newFunc = oldFunc
	}

	content := strings.ReplaceAll(popped, oldFunc, newFunc)
	s.Context().AppendTokens([]*Tok{
		{
			Type:  "RULE",
			Value: newFunc,
		},
	})
	s.Context().Write(content)
}

// EnterDataName is called when production dataName is entered.
func (s *DobfListener) EnterDataName(ctx *cobol85.DataNameContext) {
	s.push()
}

// ExitDataName is called when production dataName is exited.
func (s *DobfListener) ExitDataName(ctx *cobol85.DataNameContext) {
	popped := s.Context().Read()
	s.pop()
	tree := ctx.GetParent()
	oldVariable := ctx.GetText()
	newVariable := ""
	switch tree.(type) {
	case cobol85.IDataDescriptionEntryFormat1Context:
		i, ok := s.vars[oldVariable]
		if !ok {
			s.varIndex++
			s.vars[oldVariable] = s.varIndex
			newVariable = fmt.Sprintf("VAR_%d", s.varIndex)
		} else {
			newVariable = fmt.Sprintf("VAR_%d", i)
		}
	case cobol85.IDataDescriptionEntryFormat2Context:
		i, ok := s.vars[oldVariable]
		if !ok {
			s.varIndex++
			s.vars[oldVariable] = s.varIndex
			newVariable = fmt.Sprintf("VAR_%d", s.varIndex)
		} else {
			newVariable = fmt.Sprintf("VAR_%d", i)
		}
	case cobol85.IQualifiedDataNameFormat1Context:
		i, ok := s.vars[oldVariable]
		if ok {
			newVariable = fmt.Sprintf("VAR_%d", i)
		} else {
			newVariable = oldVariable
		}
	default:
		newVariable = oldVariable
	}

	content := strings.ReplaceAll(popped, oldVariable, newVariable)
	s.Context().AppendTokens([]*Tok{
		{
			Type:  "RULE",
			Value: newVariable,
		},
	})
	s.Context().Write(content)
}

// VisitTerminal is called when a terminal node is visited.
func (s *DobfListener) VisitTerminal(node antlr.TerminalNode) {
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
