package visitor

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ProgramUnitVisitor struct {
	cobol85.BaseCobol85Visitor
	programUnit *pb.ProgramUnit
}

func NewProgramUnitVisitor(programUnit *pb.ProgramUnit) *ProgramUnitVisitor {
	return &ProgramUnitVisitor{
		programUnit: programUnit,
	}
}

func (v *ProgramUnitVisitor) VisitDataDivision(ctx *cobol85.DataDivisionContext) any {
	v.programUnit.DataDivision = &pb.DataDivision{}
	vr := NewDataDivisionVisitor(v.programUnit.DataDivision)
	return vr.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitEnvironmentDivision(ctx *cobol85.EnvironmentDivisionContext) any {
	v.programUnit.EnvironmentDivision = &pb.EnvironmentDivision{}
	vr := NewEnvironmentDivisionVisitor(v.programUnit.EnvironmentDivision)
	return vr.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitIdentificationDivision(ctx *cobol85.IdentificationDivisionContext) any {
	v.programUnit.IdentificationDivision = &pb.IdentificationDivision{}
	vr := NewIdentificationDivisionVisitor(v.programUnit.IdentificationDivision)
	return vr.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitProcedureDivision(ctx *cobol85.ProcedureDivisionContext) any {
	v.programUnit.ProcedureDivision = &pb.ProcedureDivision{}
	vr := NewProcedureDivisionVisitor(v.programUnit.ProcedureDivision)
	return vr.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) VisitProgramUnit(ctx *cobol85.ProgramUnitContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProgramUnitVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProgramUnitVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
