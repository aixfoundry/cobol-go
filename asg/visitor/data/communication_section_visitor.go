package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type CommunicationSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.CommunicationSection
}

func NewCommunicationSectionVisitor(section *pb.CommunicationSection) *CommunicationSectionVisitor {
	return &CommunicationSectionVisitor{
		section: section,
	}
}

func (v *CommunicationSectionVisitor) VisitCommunicationSection(ctx *cobol85.CommunicationSectionContext) any {
	for _, ictx := range ctx.AllDataDescriptionEntry() {
		v.section.DataDescriptionEntries = append(v.section.DataDescriptionEntries, conv.DataDescriptionEntry(ictx))
	}
	for _, ictx := range ctx.AllCommunicationDescriptionEntry() {
		v.section.CommunicationDescriptionEntries = append(v.section.CommunicationDescriptionEntries, conv.CommunicationDescriptionEntry(ictx))
	}
	return v.VisitChildren(ctx)
}

func (v *CommunicationSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *CommunicationSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
