package visitor

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type CompilationUnitVisitor struct {
	cobol85.BaseCobol85Visitor
	program *pb.Program
	name    string
}

func NewCompilationUnitVisitor(name string, program *pb.Program) *CompilationUnitVisitor {
	return &CompilationUnitVisitor{
		program: program,
		name:    name,
	}
}

func (v *CompilationUnitVisitor) VisitCompilationUnit(ctx *cobol85.CompilationUnitContext) any {
	compilationUnit := &pb.CompilationUnit{
		Name:         v.name,
		ProgramUnits: []*pb.ProgramUnit{},
	}
	v.program.CompilationUnits = append(v.program.CompilationUnits, compilationUnit)
	vr := NewProgramUnitsVisitor(compilationUnit)
	return vr.VisitChildren(ctx)
}

func (v *CompilationUnitVisitor) VisitStartRule(ctx *cobol85.StartRuleContext) any {
	return v.VisitChildren(ctx)
}

func (v *CompilationUnitVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *CompilationUnitVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
