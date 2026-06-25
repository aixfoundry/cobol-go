package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func DataDescriptionEntry(in cobol85.IDataDescriptionEntryContext) (out *pb.DataDescriptionEntry) {
	ctx := in.(*cobol85.DataDescriptionEntryContext)
	out = &pb.DataDescriptionEntry{}
	if ictx := ctx.DataDescriptionEntryFormat1(); ictx != nil {
		cctx := ictx.(*cobol85.DataDescriptionEntryFormat1Context)
		f1 := &pb.DataDescriptionEntry_Format1{}
		if cctx.DataName() != nil {
			f1.DataName = DataName(cctx.DataName())
		}

		for _, v := range cctx.AllDataRedefinesClause() {
			clauseCtx := v.(*cobol85.DataRedefinesClauseContext)
			f1.DataRedefinesClause = &pb.DataRedefinesClause{
				DataName: DataName(clauseCtx.DataName()),
			}
		}
		for range cctx.AllDataIntegerStringClause() {
			f1.DataIntegerStringClause = &pb.DataIntegerStringClause{}
		}
		for _, v := range cctx.AllDataExternalClause() {
			clauseCtx := v.(*cobol85.DataExternalClauseContext)
			f1.DataExternalClause = &pb.DataExternalClause{
				Literal: Literal(clauseCtx.Literal()),
			}
		}
		for range cctx.AllDataGlobalClause() {
			f1.DataGlobalClause = &pb.DataGlobalClause{}
		}
		for range cctx.AllDataTypeDefClause() {
			f1.DataTypeDefClause = &pb.DataTypeDefClause{}
		}
		for range cctx.AllDataThreadLocalClause() {
			f1.DataThreadLocalClause = &pb.DataThreadLocalClause{}
		}
		for _, v := range cctx.AllDataPictureClause() {
			pcCtx := v.(*cobol85.DataPictureClauseContext)
			f1.DataPictureClause = &pb.DataPictureClause{
				PictureString: PictureString(pcCtx.PictureString()),
			}
		}
		for _, v := range cctx.AllDataCommonOwnLocalClause() {
			vCtx := v.(*cobol85.DataCommonOwnLocalClauseContext)
			var typ pb.DataCommonOwnLocalClause_Type
			switch {
			case vCtx.COMMON() != nil:
				typ = pb.DataCommonOwnLocalClause_COMMON
			case vCtx.OWN() != nil:
				typ = pb.DataCommonOwnLocalClause_OWN
			case vCtx.LOCAL() != nil:
				typ = pb.DataCommonOwnLocalClause_LOCAL
			}
			f1.DataCommonOwnLocalClause = &pb.DataCommonOwnLocalClause{
				Type: typ,
			}
		}
		for _, v := range cctx.AllDataTypeClause() {
			vCtx := v.(*cobol85.DataTypeClauseContext)
			switch {
			case vCtx.SHORT_DATE() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Type_{
						Type: pb.DataTypeClause_SHORT_DATE,
					},
				}
			case vCtx.LONG_DATE() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Type_{
						Type: pb.DataTypeClause_LONG_DATE,
					},
				}
			case vCtx.NUMERIC_DATE() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Type_{
						Type: pb.DataTypeClause_NUMERIC_DATE,
					},
				}
			case vCtx.NUMERIC_TIME() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Type_{
						Type: pb.DataTypeClause_NUMERIC_TIME,
					},
				}
			case vCtx.LONG_TIME() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Type_{
						Type: pb.DataTypeClause_LONG_TIME,
					},
				}
			case vCtx.CLOB() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Lob_{
						Lob: &pb.DataTypeClause_Lob{
							Mode:           pb.DataTypeClause_CLOB,
							IntegerLiteral: IntegerLiteral(vCtx.IntegerLiteral()),
						},
					},
				}
			case vCtx.BLOB() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Lob_{
						Lob: &pb.DataTypeClause_Lob{
							Mode:           pb.DataTypeClause_BLOB,
							IntegerLiteral: IntegerLiteral(vCtx.IntegerLiteral()),
						},
					},
				}
			case vCtx.DBCLOB() != nil:
				f1.DataTypeClause = &pb.DataTypeClause{
					OneOf: &pb.DataTypeClause_Lob_{
						Lob: &pb.DataTypeClause_Lob{
							Mode:           pb.DataTypeClause_DBCLOB,
							IntegerLiteral: IntegerLiteral(vCtx.IntegerLiteral()),
						},
					},
				}
			}
		}
		for _, v := range cctx.AllDataUsingClause() {
			vCtx := v.(*cobol85.DataUsingClauseContext)
			var typ pb.DataUsingClause_Type
			switch {
			case vCtx.LANGUAGE() != nil:
				typ = pb.DataUsingClause_LANGUAGE
			case vCtx.CONVENTION() != nil:
				typ = pb.DataUsingClause_CONVENTION
			}
			f1.DataUsingClause = &pb.DataUsingClause{
				Type: typ,
			}
			if vCtx.CobolWord() != nil {
				f1.DataUsingClause.OneOf = &pb.DataUsingClause_CobolWord{
					CobolWord: CobolWord(vCtx.CobolWord()),
				}
			} else if vCtx.DataName() != nil {
				f1.DataUsingClause.OneOf = &pb.DataUsingClause_DataName{
					DataName: DataName(vCtx.DataName()),
				}
			}
		}
		for _, v := range cctx.AllDataUsageClause() {
			vCtx := v.(*cobol85.DataUsageClauseContext)
			var typ pb.DataUsageClause_Type
			switch {
			case vCtx.BINARY() != nil:
				typ = pb.DataUsageClause_BINARY
				if vCtx.TRUNCATED() != nil {
					typ = pb.DataUsageClause_BINARY_TRUNCATED
				}
				if vCtx.EXTENDED() != nil {
					typ = pb.DataUsageClause_BINARY_EXTENDED
				}
			case vCtx.BIT() != nil:
				typ = pb.DataUsageClause_BIT
			case vCtx.COMP() != nil:
				typ = pb.DataUsageClause_COMP
			case vCtx.COMP_1() != nil:
				typ = pb.DataUsageClause_COMP_1
			case vCtx.COMP_2() != nil:
				typ = pb.DataUsageClause_COMP_2
			case vCtx.COMP_3() != nil:
				typ = pb.DataUsageClause_COMP_3
			case vCtx.COMP_4() != nil:
				typ = pb.DataUsageClause_COMP_4
			case vCtx.COMP_5() != nil:
				typ = pb.DataUsageClause_COMP_5
			case vCtx.COMPUTATIONAL() != nil:
				typ = pb.DataUsageClause_COMPUTATIONAL
			case vCtx.COMPUTATIONAL_1() != nil:
				typ = pb.DataUsageClause_COMPUTATIONAL_1
			case vCtx.COMPUTATIONAL_2() != nil:
				typ = pb.DataUsageClause_COMPUTATIONAL_2
			case vCtx.COMPUTATIONAL_3() != nil:
				typ = pb.DataUsageClause_COMPUTATIONAL_3
			case vCtx.COMPUTATIONAL_4() != nil:
				typ = pb.DataUsageClause_COMPUTATIONAL_4
			case vCtx.COMPUTATIONAL_5() != nil:
				typ = pb.DataUsageClause_COMPUTATIONAL_5
			case vCtx.CONTROL_POINT() != nil:
				typ = pb.DataUsageClause_CONTROL_POINT
			case vCtx.DATE() != nil:
				typ = pb.DataUsageClause_DATE
			case vCtx.DISPLAY() != nil:
				typ = pb.DataUsageClause_DISPLAY
			case vCtx.DISPLAY_1() != nil:
				typ = pb.DataUsageClause_DISPLAY_1
			case vCtx.DOUBLE() != nil:
				typ = pb.DataUsageClause_DOUBLE
			case vCtx.EVENT() != nil:
				typ = pb.DataUsageClause_EVENT
			case vCtx.FUNCTION_POINTER() != nil:
				typ = pb.DataUsageClause_FUNCTION_POINTER
			case vCtx.INDEX() != nil:
				typ = pb.DataUsageClause_INDEX
			case vCtx.KANJI() != nil:
				typ = pb.DataUsageClause_KANJI
			case vCtx.LOCK() != nil:
				typ = pb.DataUsageClause_LOCK
			case vCtx.NATIONAL() != nil:
				typ = pb.DataUsageClause_NATIONAL
			case vCtx.PACKED_DECIMAL() != nil:
				typ = pb.DataUsageClause_PACKED_DECIMAL
			case vCtx.POINTER() != nil:
				typ = pb.DataUsageClause_POINTER
			case vCtx.PROCEDURE_POINTER() != nil:
				typ = pb.DataUsageClause_PROCEDURE_POINTER
			case vCtx.REAL() != nil:
				typ = pb.DataUsageClause_REAL
			case vCtx.SQL() != nil:
				typ = pb.DataUsageClause_SQL
			case vCtx.TASK() != nil:
				typ = pb.DataUsageClause_TASK
			}
			f1.DataUsageClause = &pb.DataUsageClause{
				Type: typ,
			}
		}
		for _, v := range cctx.AllDataValueClause() {
			vCtx := v.(*cobol85.DataValueClauseContext)
			f1.DataValueClause = &pb.DataValueClause{}
			for _, vv := range vCtx.AllDataValueInterval() {
				f1.DataValueClause.DataValueInterval = append(f1.DataValueClause.DataValueInterval, DataValueInterval(vv))
			}
		}
		for _, v := range cctx.AllDataReceivedByClause() {
			vCtx := v.(*cobol85.DataReceivedByClauseContext)
			var typ pb.DataReceivedByClause_Type
			switch {
			case vCtx.CONTENT() != nil:
				typ = pb.DataReceivedByClause_CONTENT
			case vCtx.REFERENCE() != nil:
				typ = pb.DataReceivedByClause_REFERENCE
			case vCtx.REF() != nil:
				typ = pb.DataReceivedByClause_REF
			}
			f1.DataReceivedByClause = &pb.DataReceivedByClause{
				Type: typ,
			}
		}
		for _, v := range cctx.AllDataOccursClause() {
			vCtx := v.(*cobol85.DataOccursClauseContext)
			f1.DataOccursClause = &pb.DataOccursClause{
				To: IntegerLiteral(vCtx.IntegerLiteral()),
			}
			if iOD := vCtx.DataOccursDepending(); iOD != nil {
				od := iOD.(*cobol85.DataOccursDependingContext)
				f1.DataOccursClause.DependingOn = QualifiedDataName(od.QualifiedDataName())
			}
			if all := vCtx.AllDataOccursSort(); len(all) != 0 {
				for _, ia := range all {
					sortCtx := ia.(*cobol85.DataOccursSortContext)
					sort := &pb.DataOccursClause_Sort{}
					for _, iName := range sortCtx.AllQualifiedDataName() {
						sort.QualifiedDataNames = append(sort.QualifiedDataNames, QualifiedDataName(iName))
					}
					f1.DataOccursClause.Sorts = append(f1.DataOccursClause.Sorts, sort)
				}
			}
			if all := vCtx.AllDataOccursIndexed(); len(all) != 0 {
				for _, ia := range all {
					indexedCtx := ia.(*cobol85.DataOccursIndexedContext)
					indexed := &pb.DataOccursClause_Indexed{}
					for _, iIN := range indexedCtx.AllIndexName() {
						indexed.IndexNames = append(indexed.IndexNames, IndexName(iIN))
					}
					f1.DataOccursClause.Indexes = append(f1.DataOccursClause.Indexes, indexed)
				}
			}
		}
		for _, v := range cctx.AllDataSignClause() {
			vCtx := v.(*cobol85.DataSignClauseContext)
			var typ pb.DataSignClause_Type
			switch {
			case vCtx.LEADING() != nil:
				typ = pb.DataSignClause_LEADING
			case vCtx.TRAILING() != nil:
				typ = pb.DataSignClause_TRAILING
			}
			f1.DataSignClause = &pb.DataSignClause{
				Type: typ,
			}
		}
		for _, v := range cctx.AllDataSynchronizedClause() {
			vCtx := v.(*cobol85.DataSynchronizedClauseContext)
			var typ pb.DataSynchronizedClause_Type
			switch {
			case vCtx.LEFT() != nil:
				typ = pb.DataSynchronizedClause_LEFT
			case vCtx.RIGHT() != nil:
				typ = pb.DataSynchronizedClause_RIGHT
			}
			f1.DataSynchronizedClause = &pb.DataSynchronizedClause{
				Type: typ,
			}
		}
		for _, v := range cctx.AllDataJustifiedClause() {
			vCtx := v.(*cobol85.DataJustifiedClauseContext)
			var right bool
			if vCtx.RIGHT() != nil {
				right = true
			}
			f1.DataJustifiedClause = &pb.DataJustifiedClause{
				Right: right,
			}
		}
		for range cctx.AllDataBlankWhenZeroClause() {
			f1.DataBlankWhenZeroClause = &pb.DataBlankWhenZeroClause{}
		}
		for range cctx.AllDataWithLowerBoundsClause() {
			f1.DataWithLowerBoundsClause = &pb.DataWithLowerBoundsClause{}
		}
		for range cctx.AllDataAlignedClause() {
			f1.DataAlignedClause = &pb.DataAlignedClause{}
		}
		for range cctx.AllDataRecordAreaClause() {
			f1.DataRecordAreaClause = &pb.DataRecordAreaClause{}
		}
		out.OneOf = &pb.DataDescriptionEntry_F1{
			F1: f1,
		}
	} else if ictx := ctx.DataDescriptionEntryFormat2(); ictx != nil {
		cctx := ictx.(*cobol85.DataDescriptionEntryFormat2Context)
		f2 := &pb.DataDescriptionEntry_Format2{
			DataName:          DataName(cctx.DataName()),
			DataRenamesClause: &pb.DataRenamesClause{},
		}
		if iClause := cctx.DataRenamesClause(); iClause != nil {
			clauseCtx := iClause.(*cobol85.DataRenamesClauseContext)
			for _, iName := range clauseCtx.AllQualifiedDataName() {
				if f2.DataRenamesClause.From == nil {
					f2.DataRenamesClause.From = QualifiedDataName(iName)
				} else {
					f2.DataRenamesClause.To = QualifiedDataName(iName)
				}
			}
		}
		out.OneOf = &pb.DataDescriptionEntry_F2{
			F2: f2,
		}
	} else if ictx := ctx.DataDescriptionEntryFormat3(); ictx != nil {
		cctx := ictx.(*cobol85.DataDescriptionEntryFormat3Context)
		f3 := &pb.DataDescriptionEntry_Format3{
			ConditionName:   ConditionName(cctx.ConditionName()),
			DataValueClause: &pb.DataValueClause{},
		}
		idvc := cctx.DataValueClause()
		dvc := idvc.(*cobol85.DataValueClauseContext)
		for _, vv := range dvc.AllDataValueInterval() {
			f3.DataValueClause.DataValueInterval = append(f3.DataValueClause.DataValueInterval, DataValueInterval(vv))
		}
		out.OneOf = &pb.DataDescriptionEntry_F3{
			F3: f3,
		}
	} else if ictx := ctx.DataDescriptionEntryExecSql(); ictx != nil {
		// cctx := ictx.(*cobol85.DataDescriptionEntryExecSqlContext)
		out.OneOf = &pb.DataDescriptionEntry_ExecSql_{
			ExecSql: &pb.DataDescriptionEntry_ExecSql{},
		}
	}
	return
}

func DataValueInterval(in cobol85.IDataValueIntervalContext) (out *pb.DataValueInterval) {
	ctx := in.(*cobol85.DataValueIntervalContext)
	out = &pb.DataValueInterval{}
	if iFrom := ctx.DataValueIntervalFrom(); iFrom != nil {
		fromCtx := iFrom.(*cobol85.DataValueIntervalFromContext)
		if fromCtx.Literal() != nil {
			out.From = &pb.DataValueInterval_LiteralFrom{
				LiteralFrom: Literal(fromCtx.Literal()),
			}
		} else if fromCtx.CobolWord() != nil {
			out.From = &pb.DataValueInterval_CobolWordFrom{
				CobolWordFrom: CobolWord(fromCtx.CobolWord()),
			}
		}
	}
	if iTo := ctx.DataValueIntervalTo(); iTo != nil {
		toCtx := iTo.(*cobol85.DataValueIntervalToContext)
		if toCtx.Literal() != nil {
			out.To = &pb.DataValueInterval_LiteralTo{
				LiteralTo: Literal(toCtx.Literal()),
			}
		}
	}
	return
}

func PictureString(in cobol85.IPictureStringContext) (out *pb.PictureString) {
	ctx := in.(*cobol85.PictureStringContext)
	out = &pb.PictureString{}
	for _, vv := range ctx.AllPictureChars() {
		picCharCtx := vv.(*cobol85.PictureCharsContext)
		picChar := &pb.PictureChars{}
		switch {
		case picCharCtx.DOLLARCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_DOLLARCHAR,
			}
		case picCharCtx.IDENTIFIER() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_IDENTIFIER,
			}
		case picCharCtx.NUMERICLITERAL() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_NUMERICLITERAL,
			}
		case picCharCtx.SLASHCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_SLASHCHAR,
			}
		case picCharCtx.COMMACHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_COMMACHAR,
			}
		case picCharCtx.DOT() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_DOT,
			}
		case picCharCtx.COLONCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_COLONCHAR,
			}
		case picCharCtx.ASTERISKCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_ASTERISKCHAR,
			}
		case picCharCtx.DOUBLEASTERISKCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_DOUBLEASTERISKCHAR,
			}
		case picCharCtx.LPARENCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_LPARENCHAR,
			}
		case picCharCtx.RPARENCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_RPARENCHAR,
			}
		case picCharCtx.PLUSCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_PLUSCHAR,
			}
		case picCharCtx.MINUSCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_MINUSCHAR,
			}
		case picCharCtx.LESSTHANCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_LESSTHANCHAR,
			}
		case picCharCtx.MORETHANCHAR() != nil:
			picChar.OneOf = &pb.PictureChars_Type_{
				Type: pb.PictureChars_MORETHANCHAR,
			}
		case picCharCtx.IntegerLiteral() != nil:
			picChar.OneOf = &pb.PictureChars_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(picCharCtx.IntegerLiteral()),
			}
		}
		out.Chars = append(out.Chars, picChar)
	}
	return
}
