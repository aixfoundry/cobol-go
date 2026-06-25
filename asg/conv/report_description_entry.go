package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func ReportDescriptionEntry(in cobol85.IReportDescriptionEntryContext) (out *pb.ReportDescriptionEntry) {
	ctx := in.(*cobol85.ReportDescriptionEntryContext)
	out = &pb.ReportDescriptionEntry{
		ReportName: ReportName(ctx.ReportName()),
	}
	if ictx := ctx.ReportDescriptionGlobalClause(); ictx != nil {
		out.GlobalClause = &pb.GlobalClause{}
	}
	if ictx := ctx.ReportDescriptionPageLimitClause(); ictx != nil {
		cctx := ictx.(*cobol85.ReportDescriptionPageLimitClauseContext)
		out.PageLimitClause = &pb.PageLimitClause{
			Page: IntegerLiteral(cctx.IntegerLiteral()),
		}
	}
	if ictx := ctx.ReportDescriptionHeadingClause(); ictx != nil {
		cctx := ictx.(*cobol85.ReportDescriptionHeadingClauseContext)
		out.HeadingClause = &pb.HeadingClause{
			Heading: IntegerLiteral(cctx.IntegerLiteral()),
		}
	}
	if ictx := ctx.ReportDescriptionFirstDetailClause(); ictx != nil {
		cctx := ictx.(*cobol85.ReportDescriptionFirstDetailClauseContext)
		out.FirstDetailClause = &pb.FirstDetailClause{
			FirstDetail: IntegerLiteral(cctx.IntegerLiteral()),
		}
	}
	if ictx := ctx.ReportDescriptionLastDetailClause(); ictx != nil {
		cctx := ictx.(*cobol85.ReportDescriptionLastDetailClauseContext)
		out.LastDetailClause = &pb.LastDetailClause{
			LastDetail: IntegerLiteral(cctx.IntegerLiteral()),
		}
	}
	if ictx := ctx.ReportDescriptionFootingClause(); ictx != nil {
		cctx := ictx.(*cobol85.ReportDescriptionFootingClauseContext)
		out.FootingClause = &pb.FootingClause{
			Footing: IntegerLiteral(cctx.IntegerLiteral()),
		}
	}
	return
}
