package document

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/preprocessor"
)

func GetTextWithHiddenTokens(ctx antlr.Tree, cts *antlr.CommonTokenStream) string {
	listener := NewVisitTerminalListener(cts)
	walker := antlr.NewParseTreeWalker()
	walker.Walk(listener, ctx)
	return listener.GetText()
}

func GetHiddenTokensToLeft(cts *antlr.CommonTokenStream, pos int) (out string) {
	tokens := cts.GetHiddenTokensToLeft(pos, antlr.TokenHiddenChannel)
	for _, token := range tokens {
		out += token.GetText()
	}
	return
}

type VisitTerminalListener struct {
	preprocessor.BaseCobol85PreprocessorListener
	cts      *antlr.CommonTokenStream
	notFirst bool
	text     string
}

func NewVisitTerminalListener(cts *antlr.CommonTokenStream) *VisitTerminalListener {
	return &VisitTerminalListener{
		cts:      cts,
		notFirst: false,
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *VisitTerminalListener) VisitTerminal(node antlr.TerminalNode) {
	if s.notFirst {
		pos := node.GetSourceInterval().Start
		s.text += GetHiddenTokensToLeft(s.cts, pos)
	}
	if node.GetSymbol().GetTokenType() != antlr.TokenEOF {
		s.text += node.GetText()
	}
	s.notFirst = true
}

func (s *VisitTerminalListener) GetText() string {
	return s.text
}
