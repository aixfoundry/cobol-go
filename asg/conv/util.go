package conv

import (
	"strings"

	"github.com/aixfoundry/cobol-go/constant"
	"github.com/antlr4-go/antlr/v4"
)

func Symbol(name string) string {
	return strings.ToUpper(name)
}

func GetUntaggedText(nodes []antlr.TerminalNode, tags ...string) (ret string) {
	for _, node := range nodes {
		text := node.GetText()
		for _, tag := range tags {
			text = strings.ReplaceAll(text, tag, "")
		}
		ret += strings.TrimSpace(text) + constant.CHAR_WHITESPACE
	}
	ret = strings.TrimSpace(ret)
	return
}

func TreesStringTree(tree antlr.Tree, ruleNames []string, depth int) string {
	s := antlr.TreesGetNodeText(tree, ruleNames, nil)

	s = antlr.EscapeWhitespace(s, false)
	c := tree.GetChildCount()
	if c == 0 {
		return s
	}
	var res strings.Builder
	if depth > 0 {
		res.WriteString("\n")
	}
	res.WriteString(strings.Repeat("\t", depth) + "(" + s + " ")
	for k, child := range tree.GetChildren() {
		if k > 0 {
			res.WriteString(" ")
		}
		s = TreesStringTree(child, ruleNames, depth+1)
		res.WriteString(s)
	}
	res.WriteString(")")
	return res.String()
}
