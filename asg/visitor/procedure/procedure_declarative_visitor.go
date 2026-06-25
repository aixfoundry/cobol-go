package procedure

import (
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
	"github.com/antlr4-go/antlr/v4"
)

type ProcedureDeclarativeVisitor struct {
	cobol85.BaseCobol85Visitor
	declarative *pb.Declarative
}

func NewProcedureDeclarativeVisitor(declarative *pb.Declarative) *ProcedureDeclarativeVisitor {
	return &ProcedureDeclarativeVisitor{
		declarative: declarative,
	}
}

func (v *ProcedureDeclarativeVisitor) VisitUseStatement(ctx *cobol85.UseStatementContext) any {
	useStatement := &pb.UseStatement{}
	if iuse := ctx.UseAfterClause(); iuse != nil {
		clause := &pb.UseAfterClause{}
		cuse := iuse.(*cobol85.UseAfterClauseContext)
		if ion := cuse.UseAfterOn(); ion != nil {
			con := ion.(*cobol85.UseAfterOnContext)
			switch {
			case con.INPUT() != nil:
				clause.On = &pb.UseAfterClause_Type_{
					Type: pb.UseAfterClause_INPUT,
				}
			case con.OUTPUT() != nil:
				clause.On = &pb.UseAfterClause_Type_{
					Type: pb.UseAfterClause_OUTPUT,
				}
			case con.I_O() != nil:
				clause.On = &pb.UseAfterClause_Type_{
					Type: pb.UseAfterClause_I_O,
				}
			case con.EXTEND() != nil:
				clause.On = &pb.UseAfterClause_Type_{
					Type: pb.UseAfterClause_EXTEND,
				}
			case len(con.AllFileName()) != 0:
				filenames := []*pb.FileName{}
				for _, iFile := range con.AllFileName() {
					filenames = append(filenames, conv.FileName(iFile))
				}
				clause.On = &pb.UseAfterClause_FileNames_{
					FileNames: &pb.UseAfterClause_FileNames{
						FileNames: filenames,
					},
				}
			}
		}
		useStatement.Use = &pb.UseStatement_UseAfterClause{
			UseAfterClause: clause,
		}
	} else if iuse := ctx.UseDebugClause(); iuse != nil {
		ons := []*pb.UseDebugOn{}
		cuse := iuse.(*cobol85.UseDebugClauseContext)
		for _, ion := range cuse.AllUseDebugOn() {
			on := &pb.UseDebugOn{}
			con := ion.(*cobol85.UseDebugOnContext)
			switch {
			case con.ALL() != nil:
				if con.PROCEDURES() != nil {
					on.On = &pb.UseDebugOn_AllProcedures{}
				} else if con.REFERENCES() != nil {
					on.On = &pb.UseDebugOn_Identifier{
						Identifier: conv.Identifier(con.Identifier()),
					}
				}
			case con.ProcedureName() != nil:
				on.On = &pb.UseDebugOn_ProcedureName{
					ProcedureName: conv.ProcedureName(con.ProcedureName()),
				}
			case con.FileName() != nil:
				on.On = &pb.UseDebugOn_FileName{
					FileName: conv.FileName(con.FileName()),
				}
			}
			ons = append(ons, on)
		}
		useStatement.Use = &pb.UseStatement_UseDebugClause{
			UseDebugClause: &pb.UseDebugClause{
				Ons: ons,
			},
		}
	}
	return v.VisitChildren(ctx)
}

func (v *ProcedureDeclarativeVisitor) VisitParagraphs(ctx *cobol85.ParagraphsContext) any {
	v.declarative.Paragraphs = &pb.Paragraphs{}
	vr := NewParagraphsVisitor(v.declarative.Paragraphs)
	return vr.Visit(ctx)
}

func (v *ProcedureDeclarativeVisitor) VisitProcedureSectionHeader(ctx *cobol85.ProcedureSectionHeaderContext) any {
	header := &pb.ProcedureSectionHeader{}
	v.declarative.ProcedureSectionHeader = header
	vr := NewProcedureSectionHeaderVisitor(header)
	return vr.Visit(ctx)
}

func (v *ProcedureDeclarativeVisitor) VisitProcedureDeclarative(ctx *cobol85.ProcedureDeclarativeContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDeclarativeVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProcedureDeclarativeVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
