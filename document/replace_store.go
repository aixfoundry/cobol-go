package document

import (
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/preprocessor"
)

var (
	rePseudoTextStart = regexp.MustCompile("^==")
	rePseudoTextEnd   = regexp.MustCompile("==$")
)

type ReplaceStore struct {
	replaceable preprocessor.IReplaceableContext
	replacement preprocessor.IReplacementContext
}

func NewReplaceStore(replaceable preprocessor.IReplaceableContext, replacement preprocessor.IReplacementContext) *ReplaceStore {
	return &ReplaceStore{
		replaceable: replaceable,
		replacement: replacement,
	}
}

func (s *ReplaceStore) GetRegex(replaceable string) (ret string) {
	if replaceable != "" {
		parts := strings.Fields(replaceable)
		regexParts := []string{}
		for _, v := range parts {
			regexParts = append(regexParts, "\\Q"+v+"\\E")
		}
		sep := "[\\r\\n\\s]+"
		ret = strings.Join(regexParts, sep)
	}
	return
}

func (s *ReplaceStore) Replace(new string, cts *antlr.CommonTokenStream) (ret string) {
	replaceableText := s.getReplaceableText(s.replaceable, cts)
	replacementText := s.getReplacementText(s.replacement, cts)
	if replaceableText != "" && replacementText != "" {
		regexString := s.GetRegex(replaceableText)
		ret = regexp.MustCompile(regexString).ReplaceAllString(new, replacementText)
	} else {
		ret = new
	}
	return
}

func (s *ReplaceStore) getReplaceableText(ictx preprocessor.IReplaceableContext, cts *antlr.CommonTokenStream) (ret string) {
	ctx, ok := ictx.(*preprocessor.ReplaceableContext)
	if ok {
		if cctx := ctx.PseudoText(); cctx != nil {
			ret = s.extractPseudoText(cctx, cts)
		} else if cctx := ctx.CharDataLine(); cctx != nil {
			ret = GetTextWithHiddenTokens(cctx, cts)
		} else if cctx := ctx.CobolWord(); cctx != nil {
			ret = cctx.GetText()
		} else if cctx := ctx.Literal(); cctx != nil {
			ret = cctx.GetText()
		}
	}
	return
}

func (s *ReplaceStore) getReplacementText(ictx preprocessor.IReplacementContext, cts *antlr.CommonTokenStream) (ret string) {
	ctx, ok := ictx.(*preprocessor.ReplacementContext)
	if ok {
		if cctx := ctx.PseudoText(); cctx != nil {
			ret = s.extractPseudoText(cctx, cts)
		} else if cctx := ctx.CharDataLine(); cctx != nil {
			ret = GetTextWithHiddenTokens(cctx, cts)
		} else if cctx := ctx.CobolWord(); cctx != nil {
			ret = cctx.GetText()
		} else if cctx := ctx.Literal(); cctx != nil {
			ret = cctx.GetText()
		}
	}
	return
}

func (s *ReplaceStore) extractPseudoText(ctx preprocessor.IPseudoTextContext, cts *antlr.CommonTokenStream) string {
	pseudoText := GetTextWithHiddenTokens(ctx, cts)

	pseudoText = rePseudoTextStart.ReplaceAllString(pseudoText, "")
	pseudoText = rePseudoTextEnd.ReplaceAllString(pseudoText, "")
	pseudoText = strings.TrimSpace(pseudoText)
	return pseudoText
}

type ReplaceStores []*ReplaceStore

func (ss ReplaceStores) Len() int {
	return len(ss)
}

func (ss ReplaceStores) Less(i, j int) bool {
	iText := ss[i].replaceable.GetText()
	jText := ss[j].replaceable.GetText()
	return len(iText) > len(jText)
}

func (ss ReplaceStores) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
