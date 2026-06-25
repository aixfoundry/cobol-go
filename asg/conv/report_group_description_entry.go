package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func ReportGroupDescriptionEntry(in cobol85.IReportGroupDescriptionEntryContext) (out *pb.ReportGroupDescriptionEntry) {
	ctx := in.(*cobol85.ReportGroupDescriptionEntryContext)
	out = &pb.ReportGroupDescriptionEntry{}
	if iformat := ctx.ReportGroupDescriptionEntryFormat1(); iformat != nil {
		cformat := iformat.(*cobol85.ReportGroupDescriptionEntryFormat1Context)
		out.DataName = DataName(cformat.DataName())
		out.IntegerLiteral = IntegerLiteral(cformat.IntegerLiteral())
		vertial := &pb.ReportGroupDescriptionEntry_Vertical{}
		if iclause := cformat.ReportGroupLineNumberClause(); iclause != nil {
			cclause := iclause.(*cobol85.ReportGroupLineNumberClauseContext)
			lineNumberClause := &pb.LineNumberClause{}
			if inp := cclause.ReportGroupLineNumberNextPage(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupLineNumberNextPageContext)
				lineNumberClause.OneOf = &pb.LineNumberClause_NextPage{
					NextPage: IntegerLiteral(cnp.IntegerLiteral()),
				}
			} else if inp := cclause.ReportGroupLineNumberPlus(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupLineNumberPlusContext)
				lineNumberClause.OneOf = &pb.LineNumberClause_Plus{
					Plus: IntegerLiteral(cnp.IntegerLiteral()),
				}
			}
			vertial.LineNumberClause = lineNumberClause
		} else if iclause := cformat.ReportGroupNextGroupClause(); iclause != nil {
			cclause := iclause.(*cobol85.ReportGroupNextGroupClauseContext)
			next_group_clause := &pb.NextGroupClause{}
			if inp := cclause.ReportGroupNextGroupNextPage(); inp != nil {
				next_group_clause.OneOf = &pb.NextGroupClause_NextPage{
					NextPage: true,
				}
			} else if inp := cclause.IntegerLiteral(); inp != nil {
				next_group_clause.OneOf = &pb.NextGroupClause_NextGroup{
					NextGroup: IntegerLiteral(inp),
				}
			} else if inp := cclause.ReportGroupNextGroupPlus(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupNextGroupPlusContext)
				next_group_clause.OneOf = &pb.NextGroupClause_Plus{
					Plus: IntegerLiteral(cnp.IntegerLiteral()),
				}
			}
			vertial.NextGroupClause = next_group_clause
		} else if iclause := cformat.ReportGroupTypeClause(); iclause != nil {
			cclause := iclause.(*cobol85.ReportGroupTypeClauseContext)
			typeClause := &pb.TypeClause{}
			if ich := cclause.ReportGroupTypeControlHeading(); ich != nil {
				cch := ich.(*cobol85.ReportGroupTypeControlHeadingContext)
				ch := &pb.TypeClause_ControlHeading{}
				if cch.FINAL() != nil {
					ch.OneOf = &pb.TypeClause_ControlHeading_Final{}
				} else if cch.DataName() != nil {
					ch.OneOf = &pb.TypeClause_ControlHeading_DataName{}
				}
				typeClause.Type = &pb.TypeClause_ControlHeading_{
					ControlHeading: ch,
				}
			} else if cclause.ReportGroupTypePageHeading() != nil {
				typeClause.Type = &pb.TypeClause_PageHeading_{
					PageHeading: &pb.TypeClause_PageHeading{},
				}
			} else if cclause.ReportGroupTypeReportHeading() != nil {
				typeClause.Type = &pb.TypeClause_ReportHeading_{
					ReportHeading: &pb.TypeClause_ReportHeading{},
				}
			} else if ich := cclause.ReportGroupTypeControlFooting(); ich != nil {
				cch := ich.(*cobol85.ReportGroupTypeControlFootingContext)
				ch := &pb.TypeClause_ControlFooting{}
				if cch.FINAL() != nil {
					ch.OneOf = &pb.TypeClause_ControlFooting_Final{}
				} else if cch.DataName() != nil {
					ch.OneOf = &pb.TypeClause_ControlFooting_DataName{}
				}
				typeClause.Type = &pb.TypeClause_ControlFooting_{
					ControlFooting: ch,
				}
			} else if cclause.ReportGroupTypeDetail() != nil {
				typeClause.Type = &pb.TypeClause_Detail_{
					Detail: &pb.TypeClause_Detail{},
				}
			} else if cclause.ReportGroupTypePageFooting() != nil {
				typeClause.Type = &pb.TypeClause_PageFooting_{
					PageFooting: &pb.TypeClause_PageFooting{},
				}
			} else if cclause.ReportGroupTypeReportFooting() != nil {
				typeClause.Type = &pb.TypeClause_ReportFooting_{
					ReportFooting: &pb.TypeClause_ReportFooting{},
				}
			}
			vertial.TypeClause = typeClause
		} else if iclause := cformat.ReportGroupUsageClause(); iclause != nil {
			vertial.UsageClause = &pb.UsageClause{}
		}
		out.OneOf = &pb.ReportGroupDescriptionEntry_Vertical_{
			Vertical: vertial,
		}
	} else if iformat := ctx.ReportGroupDescriptionEntryFormat2(); iformat != nil {
		cformat := iformat.(*cobol85.ReportGroupDescriptionEntryFormat2Context)
		out.DataName = DataName(cformat.DataName())
		out.IntegerLiteral = IntegerLiteral(cformat.IntegerLiteral())
		single := &pb.ReportGroupDescriptionEntry_Single{}
		if iclause := cformat.ReportGroupLineNumberClause(); iclause != nil {
			cclause := iclause.(*cobol85.ReportGroupLineNumberClauseContext)
			lineNumberClause := &pb.LineNumberClause{}
			if inp := cclause.ReportGroupLineNumberNextPage(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupLineNumberNextPageContext)
				lineNumberClause.OneOf = &pb.LineNumberClause_NextPage{
					NextPage: IntegerLiteral(cnp.IntegerLiteral()),
				}
			} else if inp := cclause.ReportGroupLineNumberPlus(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupLineNumberPlusContext)
				lineNumberClause.OneOf = &pb.LineNumberClause_Plus{
					Plus: IntegerLiteral(cnp.IntegerLiteral()),
				}
			}
			single.LineNumberClause = lineNumberClause
		} else if iclause := cformat.ReportGroupUsageClause(); iclause != nil {
			single.UsageClause = &pb.UsageClause{}
		}
		out.OneOf = &pb.ReportGroupDescriptionEntry_Single_{
			Single: single,
		}
	} else if iformat := ctx.ReportGroupDescriptionEntryFormat3(); iformat != nil {
		cformat := iformat.(*cobol85.ReportGroupDescriptionEntryFormat3Context)
		out.DataName = DataName(cformat.DataName())
		out.IntegerLiteral = IntegerLiteral(cformat.IntegerLiteral())
		printable := &pb.ReportGroupDescriptionEntry_Printable{}
		for _, iv := range cformat.AllReportGroupPictureClause() {
			cv := iv.(*cobol85.ReportGroupPictureClauseContext)
			printable.PictureClause = &pb.PictureClause{
				PictureString: PictureString(cv.PictureString()),
			}
		}
		for range cformat.AllReportGroupUsageClause() {
			printable.UsageClause = &pb.UsageClause{}
		}
		for _, iv := range cformat.AllReportGroupSignClause() {
			cv := iv.(*cobol85.ReportGroupSignClauseContext)
			var sign pb.SignClause_Type
			switch {
			case cv.TRAILING() != nil:
				sign = pb.SignClause_TRAILING
			case cv.LEADING() != nil:
				sign = pb.SignClause_LEADING
			}
			printable.SignClause = &pb.SignClause{
				Type: sign,
			}
		}
		for _, iv := range cformat.AllReportGroupJustifiedClause() {
			cv := iv.(*cobol85.ReportGroupJustifiedClauseContext)
			printable.JustifiedClause = &pb.JustifiedClause{
				Right: cv.RIGHT() != nil,
			}
		}
		for range cformat.AllReportGroupBlankWhenZeroClause() {
			printable.BlankWhenZeroClause = &pb.BlankWhenZeroClause{}
		}
		for _, iclause := range cformat.AllReportGroupLineNumberClause() {
			cclause := iclause.(*cobol85.ReportGroupLineNumberClauseContext)
			lineNumberClause := &pb.LineNumberClause{}
			if inp := cclause.ReportGroupLineNumberNextPage(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupLineNumberNextPageContext)
				lineNumberClause.OneOf = &pb.LineNumberClause_NextPage{
					NextPage: IntegerLiteral(cnp.IntegerLiteral()),
				}
			} else if inp := cclause.ReportGroupLineNumberPlus(); inp != nil {
				cnp := inp.(*cobol85.ReportGroupLineNumberPlusContext)
				lineNumberClause.OneOf = &pb.LineNumberClause_Plus{
					Plus: IntegerLiteral(cnp.IntegerLiteral()),
				}
			}
			printable.LineNumberClause = lineNumberClause
		}
		for _, iv := range cformat.AllReportGroupColumnNumberClause() {
			cv := iv.(*cobol85.ReportGroupColumnNumberClauseContext)
			printable.ColumnNumberClause = &pb.ColumnNumberClause{
				ColumnNumber: IntegerLiteral(cv.IntegerLiteral()),
			}
		}
		for range cformat.AllReportGroupIndicateClause() {
			printable.IndicateClause = &pb.IndicateClause{}
		}
		for _, iv := range cformat.AllReportGroupSourceClause() {
			cv := iv.(*cobol85.ReportGroupSourceClauseContext)
			printable.SourceClause = &pb.SourceClause{
				Source: Identifier(cv.Identifier()),
			}
		}
		for _, iv := range cformat.AllReportGroupValueClause() {
			cv := iv.(*cobol85.ReportGroupValueClauseContext)
			printable.ValueClause = &pb.ValueClause{
				Value: Literal(cv.Literal()),
			}
		}
		for _, iv := range cformat.AllReportGroupSumClause() {
			cv := iv.(*cobol85.ReportGroupSumClauseContext)
			sumClause := &pb.SumClause{}
			for _, is := range cv.AllIdentifier() {
				sumClause.Sums = append(sumClause.Sums, Identifier(is))
			}
			for _, is := range cv.AllDataName() {
				sumClause.Upons = append(sumClause.Upons, DataName(is))
			}
			printable.SumClause = sumClause
		}
		for _, iv := range cformat.AllReportGroupResetClause() {
			cv := iv.(*cobol85.ReportGroupResetClauseContext)
			resetClause := &pb.ResetClause{}
			if cv.FINAL() != nil {
				resetClause.OneOf = &pb.ResetClause_Final{
					Final: true,
				}
			} else if cv.DataName() != nil {
				resetClause.OneOf = &pb.ResetClause_DataName{
					DataName: DataName(cv.DataName()),
				}
			}
			printable.ResetClause = resetClause
		}
		out.OneOf = &pb.ReportGroupDescriptionEntry_Printable_{
			Printable: printable,
		}
	}
	return
}
