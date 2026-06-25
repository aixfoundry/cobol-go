package visitor

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/constant"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type IdentificationDivisionVisitor struct {
	cobol85.BaseCobol85Visitor
	division *pb.IdentificationDivision
}

func NewIdentificationDivisionVisitor(division *pb.IdentificationDivision) *IdentificationDivisionVisitor {
	return &IdentificationDivisionVisitor{
		division: division,
	}
}

func (v *IdentificationDivisionVisitor) VisitProgramIdParagraph(ctx *cobol85.ProgramIdParagraphContext) interface{} {
	attribute := pb.ProgramIdParagraph_COMMON
	switch {
	case ctx.COMMON() != nil:
		attribute = pb.ProgramIdParagraph_COMMON
	case ctx.INITIAL() != nil:
		attribute = pb.ProgramIdParagraph_INITIAL
	case ctx.LIBRARY() != nil:
		attribute = pb.ProgramIdParagraph_LIBRARY
	case ctx.DEFINITION() != nil:
		attribute = pb.ProgramIdParagraph_DEFINITION
	case ctx.RECURSIVE() != nil:
		attribute = pb.ProgramIdParagraph_RECURSIVE
	default:
		attribute = pb.ProgramIdParagraph_COMMON
	}
	nameCtx := ctx.ProgramName().(*cobol85.ProgramNameContext)
	programName := &pb.ProgramName{}
	if ictx := nameCtx.NONNUMERICLITERAL(); ictx != nil {
		programName.OneOf = &pb.ProgramName_NonNumericLiteral{
			NonNumericLiteral: &pb.NonNumericLiteral{
				Value: ictx.GetText(),
			},
		}
	} else if ictx := nameCtx.CobolWord(); ictx != nil {
		programName.OneOf = &pb.ProgramName_CobolWord{
			CobolWord: &pb.CobolWord{
				Value: ictx.GetText(),
			},
		}
	}
	v.division.ProgramIdParagraph = &pb.ProgramIdParagraph{
		ProgramName: programName,
		Attribute:   attribute,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitAuthorParagraph(ctx *cobol85.AuthorParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = conv.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.AuthorParagraph = &pb.AuthorParagraph{
		Author: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitInstallationParagraph(ctx *cobol85.InstallationParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = conv.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.InstallationParagraph = &pb.InstallationParagraph{
		Installation: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitDateWrittenParagraph(ctx *cobol85.DateWrittenParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = conv.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.DateWrittenParagraph = &pb.DateWrittenParagraph{
		DateWritten: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitDateCompiledParagraph(ctx *cobol85.DateCompiledParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = conv.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.DateCompiledParagraph = &pb.DateCompiledParagraph{
		DateCompiled: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitSecurityParagraph(ctx *cobol85.SecurityParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = conv.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.SecurityParagraph = &pb.SecurityParagraph{
		Security: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitRemarksParagraph(ctx *cobol85.RemarksParagraphContext) interface{} {
	text := ""
	if ictx := ctx.CommentEntry(); ictx != nil {
		cctx := ictx.(*cobol85.CommentEntryContext)
		text = conv.GetUntaggedText(cctx.AllCOMMENTENTRYLINE(), constant.COMMENT_ENTRY_TAG, constant.EXEC_END_TAG)
	}
	v.division.RemarksParagraph = &pb.RemarksParagraph{
		Remarks: text,
	}
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitIdentificationDivision(ctx *cobol85.IdentificationDivisionContext) any {
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *IdentificationDivisionVisitor) VisitIdentificationDivisionBody(ctx *cobol85.IdentificationDivisionBodyContext) any {
	return v.VisitChildren(ctx)
}

func (v *IdentificationDivisionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
