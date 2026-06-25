package visitor

import (
	"github.com/aixfoundry/cobol-go/asg/visitor/environment"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
	"github.com/antlr4-go/antlr/v4"
)

type EnvironmentDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.EnvironmentDivision
}

func NewEnvironmentDivisionVisitor(division *pb.EnvironmentDivision) *EnvironmentDivisionVisitor {
	return &EnvironmentDivisionVisitor{
		division: division,
	}
}

func (v *EnvironmentDivisionVisitor) VisitSpecialNamesParagraph(ctx *cobol85.SpecialNamesParagraphContext) any {
	v.division.SpecialNamesParagraph = &pb.SpecialNamesParagraph{}
	vr := environment.NewSpecialNamesParagraphVisitor(v.division.SpecialNamesParagraph)
	return vr.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitInputOutputSection(ctx *cobol85.InputOutputSectionContext) any {
	v.division.InputOutputSection = &pb.InputOutputSection{}
	vr := environment.NewInputOutputSectionVisitor(v.division.InputOutputSection)
	return vr.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitConfigurationSection(ctx *cobol85.ConfigurationSectionContext) any {
	v.division.ConfigurationSection = &pb.ConfigurationSection{}
	vr := environment.NewConfigurationSectionVisitor(v.division.ConfigurationSection)
	return vr.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) any {
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) VisitEnvironmentDivisionBody(ctx *cobol85.EnvironmentDivisionBodyContext) any {
	return v.VisitChildren(ctx)
}

func (v *EnvironmentDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *EnvironmentDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
