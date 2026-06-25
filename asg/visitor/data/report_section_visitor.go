package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ReportSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ReportSection
}

func NewReportSectionVisitor(section *pb.ReportSection) *ReportSectionVisitor {
	return &ReportSectionVisitor{
		section: section,
	}
}

func (v *ReportSectionVisitor) VisitReportSection(ctx *cobol85.ReportSectionContext) any {
	for _, ictx := range ctx.AllReportDescription() {
		cctx := ictx.(*cobol85.ReportDescriptionContext)
		rd := &pb.ReportDescription{
			ReportDescriptionEntry: conv.ReportDescriptionEntry(cctx.ReportDescriptionEntry()),
		}
		for _, g := range cctx.AllReportGroupDescriptionEntry() {
			rd.ReportGroupDescriptionEntries = append(rd.ReportGroupDescriptionEntries, conv.ReportGroupDescriptionEntry(g))
		}
		v.section.ReportDescriptions = append(v.section.ReportDescriptions, rd)
	}
	return v.VisitChildren(ctx)
}

func (v *ReportSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ReportSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
