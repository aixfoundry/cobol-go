package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ScreenSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ScreenSection
}

func NewScreenSectionVisitor(section *pb.ScreenSection) *ScreenSectionVisitor {
	return &ScreenSectionVisitor{
		section: section,
	}
}

func (v *ScreenSectionVisitor) VisitScreenSection(ctx *cobol85.ScreenSectionContext) any {
	for _, ictx := range ctx.AllScreenDescriptionEntry() {
		v.section.ScreenDescriptionEntries = append(v.section.ScreenDescriptionEntries, conv.ScreenDescriptionEntry(ictx))
	}
	return v.VisitChildren(ctx)
}

func (v *ScreenSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ScreenSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
