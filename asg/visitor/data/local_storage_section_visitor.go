package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type LocalStorageSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.LocalStorageSection
}

func NewLocalStorageSectionVisitor(section *pb.LocalStorageSection) *LocalStorageSectionVisitor {
	return &LocalStorageSectionVisitor{
		section: section,
	}
}

func (v *LocalStorageSectionVisitor) VisitLocalStorageSection(ctx *cobol85.LocalStorageSectionContext) any {
	for _, ictx := range ctx.AllDataDescriptionEntry() {
		v.section.DataDescriptionEntries = append(v.section.DataDescriptionEntries, conv.DataDescriptionEntry(ictx))
	}
	if ictx := ctx.LocalName(); ictx != nil {
		v.section.LocalName = conv.LocalName(ictx)
	}
	return v.VisitChildren(ctx)
}

func (v *LocalStorageSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *LocalStorageSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
