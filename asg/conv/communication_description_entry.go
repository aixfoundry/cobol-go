package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func CommunicationDescriptionEntry(in cobol85.ICommunicationDescriptionEntryContext) (out *pb.CommunicationDescriptionEntry) {
	ctx := in.(*cobol85.CommunicationDescriptionEntryContext)
	out = &pb.CommunicationDescriptionEntry{}
	if ictx := ctx.CommunicationDescriptionEntryFormat1(); ictx != nil {
		cctx := ictx.(*cobol85.CommunicationDescriptionEntryFormat1Context)
		out.CdName = CdName(cctx.CdName())
		input := &pb.CommunicationDescriptionEntry_Input{}

		for _, v := range cctx.AllDataDescName() {
			input.DataDescName = DataDescName(v)
		}
		for _, v := range cctx.AllSymbolicQueueClause() {
			vCtx := v.(*cobol85.SymbolicQueueClauseContext)
			input.SymbolicQueueClause = &pb.SymbolicQueueClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllSymbolicSubQueueClause() {
			vCtx := v.(*cobol85.SymbolicSubQueueClauseContext)
			input.SymbolicSubQueueClause = &pb.SymbolicSubQueueClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllMessageDateClause() {
			vCtx := v.(*cobol85.MessageDateClauseContext)
			input.MessageDateClause = &pb.MessageDateClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllMessageTimeClause() {
			vCtx := v.(*cobol85.MessageTimeClauseContext)
			input.MessageTimeClause = &pb.MessageTimeClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllSymbolicSourceClause() {
			vCtx := v.(*cobol85.SymbolicSourceClauseContext)
			input.SymbolicSourceClause = &pb.SymbolicSourceClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllTextLengthClause() {
			vCtx := v.(*cobol85.TextLengthClauseContext)
			input.TextLengthClause = &pb.TextLengthClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllEndKeyClause() {
			vCtx := v.(*cobol85.EndKeyClauseContext)
			input.EndKeyClause = &pb.EndKeyClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllStatusKeyClause() {
			vCtx := v.(*cobol85.StatusKeyClauseContext)
			input.StatusKeyClause = &pb.StatusKeyClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllMessageCountClause() {
			vCtx := v.(*cobol85.MessageCountClauseContext)
			input.MessageCountClause = &pb.MessageCountClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}

		out.OneOf = &pb.CommunicationDescriptionEntry_Input_{
			Input: input,
		}
	} else if ictx := ctx.CommunicationDescriptionEntryFormat2(); ictx != nil {
		cctx := ictx.(*cobol85.CommunicationDescriptionEntryFormat2Context)
		out.CdName = CdName(cctx.CdName())
		output := &pb.CommunicationDescriptionEntry_Output{}
		for _, v := range cctx.AllDestinationCountClause() {
			vCtx := v.(*cobol85.DestinationCountClauseContext)
			output.DestinationCountClause = &pb.DestinationCountClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllTextLengthClause() {
			vCtx := v.(*cobol85.TextLengthClauseContext)
			output.TextLengthClause = &pb.TextLengthClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllStatusKeyClause() {
			vCtx := v.(*cobol85.StatusKeyClauseContext)
			output.StatusKeyClause = &pb.StatusKeyClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllDestinationTableClause() {
			vCtx := v.(*cobol85.DestinationTableClauseContext)
			names := []*pb.IndexName{}
			for _, iN := range vCtx.AllIndexName() {
				names = append(names, IndexName(iN))
			}
			output.DestinationTableClause = &pb.DestinationTableClause{
				Times:   IntegerLiteral(vCtx.IntegerLiteral()),
				Indexes: names,
			}
		}
		for _, v := range cctx.AllErrorKeyClause() {
			vCtx := v.(*cobol85.ErrorKeyClauseContext)
			output.ErrorKeyClause = &pb.ErrorKeyClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllSymbolicDestinationClause() {
			vCtx := v.(*cobol85.SymbolicDestinationClauseContext)
			output.SymbolicDestinationClause = &pb.SymbolicDestinationClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		out.OneOf = &pb.CommunicationDescriptionEntry_Output_{
			Output: output,
		}
	} else if ictx := ctx.CommunicationDescriptionEntryFormat3(); ictx != nil {
		cctx := ictx.(*cobol85.CommunicationDescriptionEntryFormat3Context)
		out.CdName = CdName(cctx.CdName())
		io := &pb.CommunicationDescriptionEntry_Io{}
		for _, v := range cctx.AllMessageDateClause() {
			vCtx := v.(*cobol85.MessageDateClauseContext)
			io.MessageDateClause = &pb.MessageDateClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllMessageTimeClause() {
			vCtx := v.(*cobol85.MessageTimeClauseContext)
			io.MessageTimeClause = &pb.MessageTimeClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllSymbolicTerminalClause() {
			vCtx := v.(*cobol85.SymbolicTerminalClauseContext)
			io.SymbolicTerminalClause = &pb.SymbolicTerminalClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllTextLengthClause() {
			vCtx := v.(*cobol85.TextLengthClauseContext)
			io.TextLengthClause = &pb.TextLengthClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllEndKeyClause() {
			vCtx := v.(*cobol85.EndKeyClauseContext)
			io.EndKeyClause = &pb.EndKeyClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}
		for _, v := range cctx.AllStatusKeyClause() {
			vCtx := v.(*cobol85.StatusKeyClauseContext)
			io.StatusKeyClause = &pb.StatusKeyClause{
				DataDescName: DataDescName(vCtx.DataDescName()),
			}
		}

		out.OneOf = &pb.CommunicationDescriptionEntry_Io_{
			Io: io,
		}
	}
	return
}
