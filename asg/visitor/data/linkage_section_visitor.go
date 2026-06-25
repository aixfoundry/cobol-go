package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type LinkageSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.LinkageSection
}

func NewLinkageSectionVisitor(section *pb.LinkageSection) *LinkageSectionVisitor {
	return &LinkageSectionVisitor{
		section: section,
	}
}

func (v *LinkageSectionVisitor) VisitLinkageSection(ctx *cobol85.LinkageSectionContext) any {
	for _, ictx := range ctx.AllDataDescriptionEntry() {
		v.section.DataDescriptionEntries = append(v.section.DataDescriptionEntries, conv.DataDescriptionEntry(ictx))
	}
	return v.VisitChildren(ctx)
}

func (v *LinkageSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *LinkageSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
