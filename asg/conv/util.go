package conv

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/constant"
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
	res := ""
	if depth > 0 {
		res += "\n"
	}
	res += strings.Repeat("\t", depth) + "(" + s + " "
	for k, child := range tree.GetChildren() {
		if k > 0 {
			res += " "
		}
		s = TreesStringTree(child, ruleNames, depth+1)
		res += s
	}
	res += ")"
	return res
}
