package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type FileSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.FileSection
}

func NewFileSectionVisitor(section *pb.FileSection) *FileSectionVisitor {
	return &FileSectionVisitor{
		section: section,
	}
}

func (v *FileSectionVisitor) VisitFileSection(ctx *cobol85.FileSectionContext) any {
	for _, ictx := range ctx.AllFileDescriptionEntry() {
		v.section.FileDescriptionEntries = append(v.section.FileDescriptionEntries, conv.FileDescriptionEntry(ictx))
	}
	return v.VisitChildren(ctx)
}

func (v *FileSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *FileSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
