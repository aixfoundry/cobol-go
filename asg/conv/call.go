package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

// CallFromIdentifier creates a Call from an Identifier parse context.
// It dispatches on the identifier's oneof variant to determine the CallType.
func CallFromIdentifier(in cobol85.IIdentifierContext) (out *pb.Call) {
	ctx := in.(*cobol85.IdentifierContext)
	out = &pb.Call{}

	if ictx := ctx.QualifiedDataName(); ictx != nil {
		cctx := ictx.(*cobol85.QualifiedDataNameContext)
		callFromQualifiedDataName(cctx, out)
	} else if ictx := ctx.TableCall(); ictx != nil {
		cctx := ictx.(*cobol85.TableCallContext)
		out.Type = pb.CallType_TABLE_CALL
		if qctx := cctx.QualifiedDataName(); qctx != nil {
			qdn := QualifiedDataName(qctx)
			out.Name = qualifiedDataNameName(qdn)
		}
	} else if ictx := ctx.FunctionCall(); ictx != nil {
		cctx := ictx.(*cobol85.FunctionCallContext)
		out.Type = pb.CallType_FUNCTION_CALL
		if fnCtx := cctx.FunctionName(); fnCtx != nil {
			out.Name = fnCtx.GetText()
		}
	} else if ictx := ctx.SpecialRegister(); ictx != nil {
		out.Type = pb.CallType_SPECIAL_REGISTER_CALL
		out.Name = ictx.(*cobol85.SpecialRegisterContext).GetText()
		out.Target = &pb.Call_SpecialRegister{
			SpecialRegister: SpecialRegister(ictx),
		}
	}

	return
}

// CallFromProcedureName creates a Call from a ProcedureName parse context.
// It classifies as PROCEDURE_CALL or SECTION_CALL based on the wrapped name type.
func CallFromProcedureName(in cobol85.IProcedureNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.ProcedureNameContext)
	out = &pb.Call{}

	if pctx := ctx.ParagraphName(); pctx != nil {
		out.Type = pb.CallType_PROCEDURE_CALL
		out.Name = pctx.GetText()
		out.Target = &pb.Call_ParagraphName{
			ParagraphName: ParagraphName(pctx),
		}
	} else if sctx := ctx.SectionName(); sctx != nil {
		out.Type = pb.CallType_SECTION_CALL
		out.Name = sctx.GetText()
		out.Target = &pb.Call_SectionName{
			SectionName: SectionName(sctx),
		}
	}

	return
}

// CallFromProgramName creates a Call from a ProgramName context (used in CALL statements).
func CallFromProgramName(in cobol85.IProgramNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.ProgramNameContext)
	out = &pb.Call{
		Type: pb.CallType_UNDEFINED_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_ProgramName{
			ProgramName: ProgramName(ctx),
		},
	}
	return
}

// CallFromFileName creates a Call from a FileName context (used in file references).
func CallFromFileName(in cobol85.IFileNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.FileNameContext)
	out = &pb.Call{
		Type: pb.CallType_FILE_CONTROL_ENTRY_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_FileName{
			FileName: FileName(ctx),
		},
	}
	return
}

// CallFromConditionName creates a Call from a ConditionName context (88-level conditions).
func CallFromConditionName(in cobol85.IConditionNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.ConditionNameContext)
	out = &pb.Call{
		Type: pb.CallType_DATA_DESCRIPTION_ENTRY_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_ConditionName{
			ConditionName: ConditionName(ctx),
		},
	}
	return
}

// CallFromDataName creates a Call from a DataName context.
func CallFromDataName(in cobol85.IDataNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.DataNameContext)
	out = &pb.Call{
		Type: pb.CallType_DATA_DESCRIPTION_ENTRY_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_DataName{
			DataName: DataName(ctx),
		},
	}
	return
}

// CallFromCdName creates a Call from a CdName context (communication description).
func CallFromCdName(in cobol85.ICdNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.CdNameContext)
	out = &pb.Call{
		Type: pb.CallType_COMMUNICATION_DESCRIPTION_ENTRY_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_CdName{
			CdName: CdName(ctx),
		},
	}
	return
}

// CallFromReportName creates a Call from a ReportName context.
func CallFromReportName(in cobol85.IReportNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.ReportNameContext)
	out = &pb.Call{
		Type: pb.CallType_REPORT_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_ReportName{
			ReportName: ReportName(ctx),
		},
	}
	return
}

// CallFromMnemonicName creates a Call from a MnemonicName context.
func CallFromMnemonicName(in cobol85.IMnemonicNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.MnemonicNameContext)
	out = &pb.Call{
		Type: pb.CallType_MNEMONIC_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_MnemonicName{
			MnemonicName: MnemonicName(ctx),
		},
	}
	return
}

// CallFromIndexName creates a Call from an IndexName context.
func CallFromIndexName(in cobol85.IIndexNameContext) (out *pb.Call) {
	ctx := in.(*cobol85.IndexNameContext)
	out = &pb.Call{
		Type: pb.CallType_INDEX_CALL,
		Name: ctx.GetText(),
		Target: &pb.Call_IndexName{
			IndexName: IndexName(ctx),
		},
	}
	return
}

// callFromQualifiedDataName populates a Call from a QualifiedDataName parse context.
// It inspects the format1 (data/condition name) or format4 (in-file) variants.
func callFromQualifiedDataName(ctx *cobol85.QualifiedDataNameContext, out *pb.Call) {
	if if1ctx := ctx.QualifiedDataNameFormat1(); if1ctx != nil {
		cf1ctx := if1ctx.(*cobol85.QualifiedDataNameFormat1Context)
		if dctx := cf1ctx.DataName(); dctx != nil {
			out.Type = pb.CallType_DATA_DESCRIPTION_ENTRY_CALL
			out.Name = dctx.GetText()
			out.Target = &pb.Call_DataName{
				DataName: DataName(dctx),
			}
		} else if cctx := cf1ctx.ConditionName(); cctx != nil {
			out.Type = pb.CallType_DATA_DESCRIPTION_ENTRY_CALL
			out.Name = cctx.GetText()
			out.Target = &pb.Call_ConditionName{
				ConditionName: ConditionName(cctx),
			}
		}
	} else if if4ctx := ctx.QualifiedDataNameFormat4(); if4ctx != nil {
		cf4ctx := if4ctx.(*cobol85.QualifiedDataNameFormat4Context)
		if fctx := cf4ctx.InFile(); fctx != nil {
			cctx := fctx.(*cobol85.InFileContext)
			out.Type = pb.CallType_FILE_CONTROL_ENTRY_CALL
			if fnCtx := cctx.FileName(); fnCtx != nil {
				out.Name = fnCtx.GetText()
			}
		}
	}
}

// qualifiedDataNameName extracts the text name from a QualifiedDataName proto for diagnostic purposes.
func qualifiedDataNameName(qdn *pb.QualifiedDataName) string {
	if f1 := qdn.GetF1(); f1 != nil {
		if d := f1.GetDataName(); d != nil {
			return d.GetCobolWord().GetValue()
		}
		if c := f1.GetConditionName(); c != nil {
			return c.GetCobolWord().GetValue()
		}
	}
	return ""
}
