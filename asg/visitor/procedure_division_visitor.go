package visitor

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/asg/visitor/procedure"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ProcedureDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.ProcedureDivision
}

func NewProcedureDivisionVisitor(division *pb.ProcedureDivision) *ProcedureDivisionVisitor {
	return &ProcedureDivisionVisitor{
		division: division,
	}
}

func (v *ProcedureDivisionVisitor) VisitProcedureDeclarative(ctx *cobol85.ProcedureDeclarativeContext) interface{} {
	declarative := &pb.Declarative{}
	v.division.Declaratives = append(v.division.Declaratives, declarative)
	vr := procedure.NewProcedureDeclarativeVisitor(declarative)
	return vr.Visit(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDeclaratives(ctx *cobol85.ProcedureDeclarativesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivisionGivingClause(ctx *cobol85.ProcedureDivisionGivingClauseContext) interface{} {
	gc := &pb.ProcedureDivision_GivingClause{
		DataName: conv.DataName(ctx.DataName()),
	}
	var typ pb.ProcedureDivision_GivingClause_Type
	switch {
	case ctx.GIVING() != nil:
		typ = pb.ProcedureDivision_GivingClause_GIVING
	case ctx.RETURNING() != nil:
		typ = pb.ProcedureDivision_GivingClause_RETURNING
	}
	gc.Type = typ
	v.division.GivingClause = gc
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivisionUsingClause(ctx *cobol85.ProcedureDivisionUsingClauseContext) interface{} {
	uc := &pb.ProcedureDivision_UsingClause{}
	switch {
	case ctx.USING() != nil:
		uc.Type = pb.ProcedureDivision_UsingClause_USING
	case ctx.CHAINING() != nil:
		uc.Type = pb.ProcedureDivision_UsingClause_CHAINING
	}
	for _, ictx := range ctx.AllProcedureDivisionUsingParameter() {
		cctx := ictx.(*cobol85.ProcedureDivisionUsingParameterContext)
		up := &pb.ProcedureDivision_UsingParameter{}
		if ip := cctx.ProcedureDivisionByReferencePhrase(); ip != nil {
			byReferencePhrase := &pb.ProcedureDivision_ByReferencePhrase{}
			cp := ip.(*cobol85.ProcedureDivisionByReferencePhraseContext)
			for _, ir := range cp.AllProcedureDivisionByReference() {
				reference := &pb.ProcedureDivision_ByReference{}
				cr := ir.(*cobol85.ProcedureDivisionByReferenceContext)
				switch {
				case cr.ANY() != nil:
					reference.OneOf = &pb.ProcedureDivision_ByReference_Any{
						Any: true,
					}
				case cr.FileName() != nil:
					reference.OneOf = &pb.ProcedureDivision_ByReference_FileName{
						FileName: conv.FileName(cr.FileName()),
					}
				case cr.Identifier() != nil:
					reference.OneOf = &pb.ProcedureDivision_ByReference_Identifier{
						Identifier: conv.Identifier(cr.Identifier()),
					}
				}
				byReferencePhrase.References = append(byReferencePhrase.References, reference)
			}
			up.OneOf = &pb.ProcedureDivision_UsingParameter_ByReferencePhrase{
				ByReferencePhrase: byReferencePhrase,
			}
		} else if ip := cctx.ProcedureDivisionByValuePhrase(); ip != nil {
			cp := ip.(*cobol85.ProcedureDivisionByValuePhraseContext)
			values := []*pb.ProcedureDivision_ByValue{}
			for _, iv := range cp.AllProcedureDivisionByValue() {
				cv := iv.(*cobol85.ProcedureDivisionByValueContext)
				value := &pb.ProcedureDivision_ByValue{}
				switch {
				case cv.ANY() != nil:
					value.OneOf = &pb.ProcedureDivision_ByValue_Any{
						Any: true,
					}
				case cv.Identifier() != nil:
					value.OneOf = &pb.ProcedureDivision_ByValue_Identifier{
						Identifier: conv.Identifier(cv.Identifier()),
					}
				case cv.Literal() != nil:
					value.OneOf = &pb.ProcedureDivision_ByValue_Literal{
						Literal: conv.Literal(cv.Literal()),
					}
				}
				values = append(values, value)
			}
			up.OneOf = &pb.ProcedureDivision_UsingParameter_ByValuePhrase{
				ByValuePhrase: &pb.ProcedureDivision_ByValuePhrase{
					Values: values,
				},
			}
		}
		uc.UsingParameters = append(uc.UsingParameters, up)
	}
	v.division.UsingClause = uc
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitParagraphs(ctx *cobol85.ParagraphsContext) interface{} {
	v.division.Paragraphs = &pb.Paragraphs{}
	vr := procedure.NewParagraphsVisitor(v.division.Paragraphs)
	return vr.Visit(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureSection(ctx *cobol85.ProcedureSectionContext) interface{} {
	section := &pb.ProcedureSection{}
	v.division.ProcedureSections = append(v.division.ProcedureSections, section)
	vr := procedure.NewProcedureSectionVisitor(section)
	return vr.Visit(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivisionBody(ctx *cobol85.ProcedureDivisionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) VisitProcedureDivision(ctx *cobol85.ProcedureDivisionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ProcedureDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProcedureDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
