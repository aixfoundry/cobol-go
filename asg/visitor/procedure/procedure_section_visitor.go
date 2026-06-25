package procedure

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ProcedureSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ProcedureSection
}

func NewProcedureSectionVisitor(section *pb.ProcedureSection) *ProcedureSectionVisitor {
	return &ProcedureSectionVisitor{
		section: section,
	}
}
func (v *ProcedureSectionVisitor) VisitProcedureSectionHeader(ctx *cobol85.ProcedureSectionHeaderContext) any {
	v.section.ProcedureSectionHeader = &pb.ProcedureSectionHeader{}
	vr := NewProcedureSectionHeaderVisitor(v.section.ProcedureSectionHeader)
	return vr.Visit(ctx)
}

func (v *ProcedureSectionVisitor) VisitParagraphs(ctx *cobol85.ParagraphsContext) any {
	v.section.Paragraphs = &pb.Paragraphs{}
	vr := NewParagraphsVisitor(v.section.Paragraphs)
	return vr.Visit(ctx)
}

func (v *ProcedureSectionVisitor) VisitProcedureSection(ctx *cobol85.ProcedureSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProcedureSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProcedureSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
