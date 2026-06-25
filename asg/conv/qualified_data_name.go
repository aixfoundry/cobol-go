package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func Subscript(in cobol85.ISubscriptContext) (out *pb.Subscript) {
	ctx := in.(*cobol85.SubscriptContext)
	out = &pb.Subscript{}
	if ctx.ALL() != nil {
		out.OneOf = &pb.Subscript_All{
			All: true,
		}
	} else if ictx := ctx.IndexName(); ictx != nil {
		out.OneOf = &pb.Subscript_IndexName{
			IndexName: &pb.IndexNameIntegerLiteral{
				IndexName:      IndexName(ctx.IndexName()),
				IntegerLiteral: IntegerLiteral(ctx.IntegerLiteral()),
			},
		}
	} else if ictx := ctx.QualifiedDataName(); ictx != nil {
		out.OneOf = &pb.Subscript_QualifiedDataName{
			QualifiedDataName: &pb.QualifiedDataNameIntegerLiteral{
				QualifiedDataName: QualifiedDataName(ictx),
				IntegerLiteral:    IntegerLiteral(ctx.IntegerLiteral()),
			},
		}
	} else if ictx := ctx.IntegerLiteral(); ictx != nil {
		out.OneOf = &pb.Subscript_IntegerLiteral{
			IntegerLiteral: IntegerLiteral(ctx.IntegerLiteral()),
		}
	} else if ictx := ctx.ArithmeticExpression(); ictx != nil {
		out.OneOf = &pb.Subscript_ArithmeticExpression{
			ArithmeticExpression: ArithmeticExpression(ictx),
		}
	}
	return
}

func CharacterPosition(in cobol85.ICharacterPositionContext) *pb.CharacterPosition {
	ctx := in.(*cobol85.CharacterPositionContext)
	return &pb.CharacterPosition{
		ArithmeticExpression: ArithmeticExpression(ctx.ArithmeticExpression()),
	}
}

func Length(in cobol85.ILengthContext) *pb.Length {
	ctx := in.(*cobol85.LengthContext)
	return &pb.Length{
		ArithmeticExpression: ArithmeticExpression(ctx.ArithmeticExpression()),
	}
}

func ReferenceModifier(in cobol85.IReferenceModifierContext) *pb.ReferenceModifier {
	ctx := in.(*cobol85.ReferenceModifierContext)
	return &pb.ReferenceModifier{
		CharacterPosition: CharacterPosition(ctx.CharacterPosition()),
		Length:            Length(ctx.Length()),
	}
}

func TableCall(in cobol85.ITableCallContext) *pb.TableCall {
	ctx := in.(*cobol85.TableCallContext)
	subscripts := []*pb.Subscript{}
	for _, v := range ctx.AllSubscript() {
		subscripts = append(subscripts, Subscript(v))
	}
	return &pb.TableCall{
		QualifiedDataName: QualifiedDataName(ctx.QualifiedDataName()),
		Subscripts:        subscripts,
		ReferenceModifier: ReferenceModifier(ctx.ReferenceModifier()),
	}
}

func QualifiedDataName(in cobol85.IQualifiedDataNameContext) (out *pb.QualifiedDataName) {
	ctx := in.(*cobol85.QualifiedDataNameContext)
	out = &pb.QualifiedDataName{}
	if if1ctx := ctx.QualifiedDataNameFormat1(); if1ctx != nil {
		cf1ctx := if1ctx.(*cobol85.QualifiedDataNameFormat1Context)
		f1 := &pb.QualifiedDataNameFormat1{}
		if icctx := cf1ctx.DataName(); icctx != nil {
			f1.OneOf = &pb.QualifiedDataNameFormat1_DataName{
				DataName: DataName(icctx),
			}
		} else if icctx := cf1ctx.ConditionName(); icctx != nil {
			f1.OneOf = &pb.QualifiedDataNameFormat1_ConditionName{
				ConditionName: ConditionName(icctx),
			}
		}

		if icctx := cf1ctx.InFile(); icctx != nil {
			ccctx := icctx.(*cobol85.InFileContext)
			f1.InFile = &pb.InFile{
				FileName: FileName(ccctx.FileName()),
			}
		}

		for _, icctx := range cf1ctx.AllQualifiedInData() {
			ccctx := icctx.(*cobol85.QualifiedInDataContext)
			if iInData := ccctx.InData(); iInData != nil {
				inDataCtx := iInData.(*cobol85.InDataContext)
				f1.InDatas = append(f1.InDatas, &pb.InData{
					DataName: DataName(inDataCtx.DataName()),
				})
			}
			if iInTable := ccctx.InTable(); iInTable != nil {
				inTableCtx := iInTable.(*cobol85.InTableContext)
				f1.InTables = append(f1.InTables, &pb.InTable{
					TableCall: TableCall(inTableCtx.TableCall()),
				})
			}
		}
		out.OneOf = &pb.QualifiedDataName_F1{
			F1: f1,
		}
	} else if if2ctx := ctx.QualifiedDataNameFormat2(); if2ctx != nil {
		cf2ctx := if2ctx.(*cobol85.QualifiedDataNameFormat2Context)
		f2 := &pb.QualifiedDataNameFormat2{
			ParagraphName: ParagraphName(cf2ctx.ParagraphName()),
			InSection:     InSection(cf2ctx.InSection()),
		}
		out.OneOf = &pb.QualifiedDataName_F2{
			F2: f2,
		}
	} else if if3ctx := ctx.QualifiedDataNameFormat3(); if3ctx != nil {
		cf3ctx := if3ctx.(*cobol85.QualifiedDataNameFormat3Context)
		f3 := &pb.QualifiedDataNameFormat3{
			TextName:  TextName(cf3ctx.TextName()),
			InLibrary: InLibrary(cf3ctx.InLibrary()),
		}
		out.OneOf = &pb.QualifiedDataName_F3{
			F3: f3,
		}
	} else if if4ctx := ctx.QualifiedDataNameFormat4(); if4ctx != nil {
		cf4ctx := if4ctx.(*cobol85.QualifiedDataNameFormat4Context)
		f4 := &pb.QualifiedDataNameFormat4{
			InFile: InFile(cf4ctx.InFile()),
		}
		out.OneOf = &pb.QualifiedDataName_F4{
			F4: f4,
		}
	}
	return
}
