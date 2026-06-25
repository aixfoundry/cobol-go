package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func FileDescriptionEntry(in cobol85.IFileDescriptionEntryContext) (out *pb.FileDescriptionEntry) {
	ctx := in.(*cobol85.FileDescriptionEntryContext)
	out = &pb.FileDescriptionEntry{
		FileName: FileName(ctx.FileName()),
	}
	for _, ictx := range ctx.AllFileDescriptionEntryClause() {
		cctx := ictx.(*cobol85.FileDescriptionEntryClauseContext)
		if iClause := cctx.ExternalClause(); iClause != nil {
			out.ExternalClause = &pb.ExternalClause{}
		} else if iClause := cctx.GlobalClause(); iClause != nil {
			out.GlobalClause = &pb.GlobalClause{}
		} else if iClause := cctx.BlockContainsClause(); iClause != nil {
			cCtx := iClause.(*cobol85.BlockContainsClauseContext)
			var unit pb.BlockContainsClause_Unit
			switch {
			case cCtx.RECORDS() != nil:
				unit = pb.BlockContainsClause_RECORDS
			case cCtx.CHARACTERS() != nil:
				unit = pb.BlockContainsClause_CHARACTERS
			}
			out.BlockContainsClause = &pb.BlockContainsClause{
				From: IntegerLiteral(cCtx.IntegerLiteral()),
				To:   IntegerLiteral(cCtx.IntegerLiteral()),
				Unit: unit,
			}
		} else if iClause := cctx.RecordContainsClause(); iClause != nil {
			cCtx := iClause.(*cobol85.RecordContainsClauseContext)
			out.RecordContainsClause = &pb.RecordContainsClause{}
			if iFormat := cCtx.RecordContainsClauseFormat1(); iFormat != nil {
				fCtx := iFormat.(*cobol85.RecordContainsClauseFormat1Context)
				out.RecordContainsClause.From = IntegerLiteral(fCtx.IntegerLiteral())
			} else if iFormat := cCtx.RecordContainsClauseFormat2(); iFormat != nil {
				fCtx := iFormat.(*cobol85.RecordContainsClauseFormat2Context)
				out.RecordContainsClause.From = IntegerLiteral(fCtx.IntegerLiteral())
				out.RecordContainsClause.Varying = true
				if iTo := fCtx.RecordContainsTo(); iTo != nil {
					toCtx := iTo.(*cobol85.RecordContainsToContext)
					out.RecordContainsClause.To = IntegerLiteral(toCtx.IntegerLiteral())
				}
				if fCtx.QualifiedDataName() != nil {
					out.RecordContainsClause.QualifiedDataName = QualifiedDataName(fCtx.QualifiedDataName())
				}
			} else if iFormat := cCtx.RecordContainsClauseFormat3(); iFormat != nil {
				fCtx := iFormat.(*cobol85.RecordContainsClauseFormat3Context)
				out.RecordContainsClause.From = IntegerLiteral(fCtx.IntegerLiteral())
				if iTo := fCtx.RecordContainsTo(); iTo != nil {
					toCtx := iTo.(*cobol85.RecordContainsToContext)
					out.RecordContainsClause.To = IntegerLiteral(toCtx.IntegerLiteral())
				}
			}
		} else if iClause := cctx.LabelRecordsClause(); iClause != nil {
			cCtx := iClause.(*cobol85.LabelRecordsClauseContext)
			out.LabelRecordsClause = &pb.LabelRecordsClause{}
			switch {
			case cCtx.OMITTED() != nil:
				out.LabelRecordsClause.OneOf = &pb.LabelRecordsClause_Type_{
					Type: pb.LabelRecordsClause_OMITTED,
				}
			case cCtx.STANDARD() != nil:
				out.LabelRecordsClause.OneOf = &pb.LabelRecordsClause_Type_{
					Type: pb.LabelRecordsClause_STANDARD,
				}
			case len(cCtx.AllDataName()) != 0:
				names := []*pb.DataName{}
				for _, ia := range cCtx.AllDataName() {
					names = append(names, DataName(ia))
				}
				out.LabelRecordsClause.OneOf = &pb.LabelRecordsClause_DataNames_{
					DataNames: &pb.LabelRecordsClause_DataNames{
						DataNames: names,
					},
				}
			}
		} else if iClause := cctx.ValueOfClause(); iClause != nil {
			cCtx := iClause.(*cobol85.ValueOfClauseContext)
			out.ValueOfClause = &pb.ValueOfClause{}
			for _, ia := range cCtx.AllValuePair() {
				vpCtx := ia.(*cobol85.ValuePairContext)
				vp := &pb.ValuePair{
					SystemName: SystemName(vpCtx.SystemName()),
				}
				if vpCtx.QualifiedDataName() != nil {
					vp.Value = &pb.ValuePair_QualifiedDataName{
						QualifiedDataName: QualifiedDataName(vpCtx.QualifiedDataName()),
					}
				} else if vpCtx.Literal() != nil {
					vp.Value = &pb.ValuePair_Literal{
						Literal: Literal(vpCtx.Literal()),
					}
				}
			}
		} else if iClause := cctx.DataRecordsClause(); iClause != nil {
			cCtx := iClause.(*cobol85.DataRecordsClauseContext)
			out.DataRecordsClause = &pb.DataRecordsClause{}
			for _, iDataName := range cCtx.AllDataName() {
				out.DataRecordsClause.DataNames = append(out.DataRecordsClause.DataNames, DataName(iDataName))
			}
		} else if iClause := cctx.LinageClause(); iClause != nil {
			cCtx := iClause.(*cobol85.LinageClauseContext)
			out.LinageClause = &pb.LinageClause{}
			nofl := &pb.DataNameOrIntegerLiteral{}
			if cCtx.DataName() != nil {
				nofl.OneOf = &pb.DataNameOrIntegerLiteral_DataName{
					DataName: DataName(cCtx.DataName()),
				}
			} else if cCtx.IntegerLiteral() != nil {
				nofl.OneOf = &pb.DataNameOrIntegerLiteral_IntegerLiteral{
					IntegerLiteral: IntegerLiteral(cCtx.IntegerLiteral()),
				}
			}
			out.LinageClause.NumberOfLines = nofl
			for _, iAt := range cCtx.AllLinageAt() {
				atCtx := iAt.(*cobol85.LinageAtContext)
				if ilat := atCtx.LinageFootingAt(); ilat != nil {
					lat := ilat.(*cobol85.LinageFootingAtContext)
					nofl := &pb.DataNameOrIntegerLiteral{}
					if lat.DataName() != nil {
						nofl.OneOf = &pb.DataNameOrIntegerLiteral_DataName{
							DataName: DataName(lat.DataName()),
						}
					} else if lat.IntegerLiteral() != nil {
						nofl.OneOf = &pb.DataNameOrIntegerLiteral_IntegerLiteral{
							IntegerLiteral: IntegerLiteral(lat.IntegerLiteral()),
						}
					}
					out.LinageClause.FootingAt = nofl
				} else if ilat := atCtx.LinageLinesAtTop(); ilat != nil {
					lat := ilat.(*cobol85.LinageLinesAtTopContext)
					nofl := &pb.DataNameOrIntegerLiteral{}
					if lat.DataName() != nil {
						nofl.OneOf = &pb.DataNameOrIntegerLiteral_DataName{
							DataName: DataName(lat.DataName()),
						}
					} else if lat.IntegerLiteral() != nil {
						nofl.OneOf = &pb.DataNameOrIntegerLiteral_IntegerLiteral{
							IntegerLiteral: IntegerLiteral(lat.IntegerLiteral()),
						}
					}
					out.LinageClause.LinesAtTop = nofl
				} else if ilat := atCtx.LinageLinesAtBottom(); ilat != nil {
					lat := ilat.(*cobol85.LinageLinesAtBottomContext)
					nofl := &pb.DataNameOrIntegerLiteral{}
					if lat.DataName() != nil {
						nofl.OneOf = &pb.DataNameOrIntegerLiteral_DataName{
							DataName: DataName(lat.DataName()),
						}
					} else if lat.IntegerLiteral() != nil {
						nofl.OneOf = &pb.DataNameOrIntegerLiteral_IntegerLiteral{
							IntegerLiteral: IntegerLiteral(lat.IntegerLiteral()),
						}
					}
					out.LinageClause.LinesAtBottom = nofl
				}
			}
		} else if iClause := cctx.CodeSetClause(); iClause != nil {
			cCtx := iClause.(*cobol85.CodeSetClauseContext)
			out.CodeSetClause = &pb.CodeSetClause{
				AlphabetName: AlphabetName(cCtx.AlphabetName()),
			}
		} else if iClause := cctx.ReportClause(); iClause != nil {
			cCtx := iClause.(*cobol85.ReportClauseContext)
			out.ReportClause = &pb.ReportClause{}
			for _, iReport := range cCtx.AllReportName() {
				out.ReportClause.ReportNames = append(out.ReportClause.ReportNames, ReportName(iReport))
			}
		} else if iClause := cctx.RecordingModeClause(); iClause != nil {
			cCtx := iClause.(*cobol85.RecordingModeClauseContext)
			msCtx := cCtx.ModeStatement().(*cobol85.ModeStatementContext)
			out.RecordingModeClause = &pb.RecordingModeClause{
				ModeStatement: &pb.ModeStatement{
					CobolWord: CobolWord(msCtx.CobolWord()),
				},
			}
		}
	}
	for _, ictx := range ctx.AllDataDescriptionEntry() {
		out.DataDescriptionEntries = append(out.DataDescriptionEntries, DataDescriptionEntry(ictx))
	}
	return
}
