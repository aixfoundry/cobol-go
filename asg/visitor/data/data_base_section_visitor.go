package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type DataBaseSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.DataBaseSection
}

func NewDataBaseSectionVisitor(section *pb.DataBaseSection) *DataBaseSectionVisitor {
	return &DataBaseSectionVisitor{
		section: section,
	}
}

func (v *DataBaseSectionVisitor) VisitDataBaseSection(ctx *cobol85.DataBaseSectionContext) any {
	for _, ictx := range ctx.AllDataBaseSectionEntry() {
		cctx := ictx.(*cobol85.DataBaseSectionEntryContext)
		entry := &pb.DataBaseSectionEntry{
			IntegerLiteral: conv.IntegerLiteral(cctx.IntegerLiteral()),
		}
		for _, literal := range cctx.AllLiteral() {
			if entry.Literal == nil {
				entry.Literal = conv.Literal(literal)
			} else {
				entry.Invoke = conv.Literal(literal)
			}
		}
		v.section.DataBaseSectionEntries = append(v.section.DataBaseSectionEntries, entry)
	}
	return v.VisitChildren(ctx)
}

func (v *DataBaseSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *DataBaseSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
