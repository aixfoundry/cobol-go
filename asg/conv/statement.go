package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func Statement(in cobol85.IStatementContext) (out *pb.Statement) {
	ctx := in.(*cobol85.StatementContext)
	out = &pb.Statement{}
	if t := ctx.AcceptStatement(); t != nil {
		out.OneOf = &pb.Statement_AcceptStatement{
			AcceptStatement: AcceptStatement(t),
		}
	} else if t := ctx.AddStatement(); t != nil {
		out.OneOf = &pb.Statement_AddStatement{
			AddStatement: AddStatement(t),
		}
	} else if t := ctx.CallStatement(); t != nil {
		out.OneOf = &pb.Statement_CallStatement{
			CallStatement: CallStatement(t),
		}
	} else if t := ctx.CancelStatement(); t != nil {
		out.OneOf = &pb.Statement_CancelStatement{
			CancelStatement: CancelStatement(t),
		}
	} else if t := ctx.CloseStatement(); t != nil {
		out.OneOf = &pb.Statement_CloseStatement{
			CloseStatement: CloseStatement(t),
		}
	} else if t := ctx.ComputeStatement(); t != nil {
		out.OneOf = &pb.Statement_ComputeStatement{
			ComputeStatement: ComputeStatement(t),
		}
	} else if t := ctx.ContinueStatement(); t != nil {
		out.OneOf = &pb.Statement_ContinueStatement{
			ContinueStatement: ContinueStatement(t),
		}
	} else if t := ctx.DeleteStatement(); t != nil {
		out.OneOf = &pb.Statement_DeleteStatement{
			DeleteStatement: DeleteStatement(t),
		}
	} else if t := ctx.DisableStatement(); t != nil {
		out.OneOf = &pb.Statement_DisableStatement{
			DisableStatement: DisableStatement(t),
		}
	} else if t := ctx.DisplayStatement(); t != nil {
		out.OneOf = &pb.Statement_DisplayStatement{
			DisplayStatement: DisplayStatement(t),
		}
	} else if t := ctx.DivideStatement(); t != nil {
		out.OneOf = &pb.Statement_DivideStatement{
			DivideStatement: DivideStatement(t),
		}
	} else if t := ctx.EnableStatement(); t != nil {
		out.OneOf = &pb.Statement_EnableStatement{
			EnableStatement: EnableStatement(t),
		}
	} else if t := ctx.EntryStatement(); t != nil {
		out.OneOf = &pb.Statement_EntryStatement{
			EntryStatement: EntryStatement(t),
		}
	} else if t := ctx.EvaluateStatement(); t != nil {
		out.OneOf = &pb.Statement_EvaluateStatement{
			EvaluateStatement: EvaluateStatement(t),
		}
	} else if t := ctx.ExhibitStatement(); t != nil {
		out.OneOf = &pb.Statement_ExhibitStatement{
			ExhibitStatement: ExhibitStatement(t),
		}
	} else if t := ctx.ExecCicsStatement(); t != nil {
		out.OneOf = &pb.Statement_ExecCicsStatement{
			ExecCicsStatement: ExecCicsStatement(t),
		}
	} else if t := ctx.ExecSqlStatement(); t != nil {
		out.OneOf = &pb.Statement_ExecSqlStatement{
			ExecSqlStatement: ExecSqlStatement(t),
		}
	} else if t := ctx.ExecSqlImsStatement(); t != nil {
		out.OneOf = &pb.Statement_ExecSqlImsStatement{
			ExecSqlImsStatement: ExecSqlImsStatement(t),
		}
	} else if t := ctx.ExitStatement(); t != nil {
		out.OneOf = &pb.Statement_ExitStatement{
			ExitStatement: ExitStatement(t),
		}
	} else if t := ctx.GenerateStatement(); t != nil {
		out.OneOf = &pb.Statement_GenerateStatement{
			GenerateStatement: GenerateStatement(t),
		}
	} else if t := ctx.GobackStatement(); t != nil {
		out.OneOf = &pb.Statement_GobackStatement{
			GobackStatement: GobackStatement(t),
		}
	} else if t := ctx.GoToStatement(); t != nil {
		out.OneOf = &pb.Statement_GoToStatement{
			GoToStatement: GoToStatement(t),
		}
	} else if t := ctx.IfStatement(); t != nil {
		out.OneOf = &pb.Statement_IfStatement{
			IfStatement: IfStatement(t),
		}
	} else if t := ctx.InitializeStatement(); t != nil {
		out.OneOf = &pb.Statement_InitializeStatement{
			InitializeStatement: InitializeStatement(t),
		}
	} else if t := ctx.InitiateStatement(); t != nil {
		out.OneOf = &pb.Statement_InitiateStatement{
			InitiateStatement: InitiateStatement(t),
		}
	} else if t := ctx.InspectStatement(); t != nil {
		out.OneOf = &pb.Statement_InspectStatement{
			InspectStatement: InspectStatement(t),
		}
	} else if t := ctx.MergeStatement(); t != nil {
		out.OneOf = &pb.Statement_MergeStatement{
			MergeStatement: MergeStatement(t),
		}
	} else if t := ctx.MoveStatement(); t != nil {
		out.OneOf = &pb.Statement_MoveStatement{
			MoveStatement: MoveStatement(t),
		}
	} else if t := ctx.MultiplyStatement(); t != nil {
		out.OneOf = &pb.Statement_MultiplyStatement{
			MultiplyStatement: MultiplyStatement(t),
		}
	} else if t := ctx.NextSentenceStatement(); t != nil {
		out.OneOf = &pb.Statement_NextSentenceStatement{
			NextSentenceStatement: NextSentenceStatement(t),
		}
	} else if t := ctx.OpenStatement(); t != nil {
		out.OneOf = &pb.Statement_OpenStatement{
			OpenStatement: OpenStatement(t),
		}
	} else if t := ctx.PerformStatement(); t != nil {
		out.OneOf = &pb.Statement_PerformStatement{
			PerformStatement: PerformStatement(t),
		}
	} else if t := ctx.PurgeStatement(); t != nil {
		out.OneOf = &pb.Statement_PurgeStatement{
			PurgeStatement: PurgeStatement(t),
		}
	} else if t := ctx.ReadStatement(); t != nil {
		out.OneOf = &pb.Statement_ReadStatement{
			ReadStatement: ReadStatement(t),
		}
	} else if t := ctx.ReceiveStatement(); t != nil {
		out.OneOf = &pb.Statement_ReceiveStatement{
			ReceiveStatement: ReceiveStatement(t),
		}
	} else if t := ctx.ReleaseStatement(); t != nil {
		out.OneOf = &pb.Statement_ReleaseStatement{
			ReleaseStatement: ReleaseStatement(t),
		}
	} else if t := ctx.ReturnStatement(); t != nil {
		out.OneOf = &pb.Statement_ReturnStatement{
			ReturnStatement: ReturnStatement(t),
		}
	} else if t := ctx.RewriteStatement(); t != nil {
		out.OneOf = &pb.Statement_RewriteStatement{
			RewriteStatement: RewriteStatement(t),
		}
	} else if t := ctx.SearchStatement(); t != nil {
		out.OneOf = &pb.Statement_SearchStatement{
			SearchStatement: SearchStatement(t),
		}
	} else if t := ctx.SendStatement(); t != nil {
		out.OneOf = &pb.Statement_SendStatement{
			SendStatement: SendStatement(t),
		}
	} else if t := ctx.SetStatement(); t != nil {
		out.OneOf = &pb.Statement_SetStatement{
			SetStatement: SetStatement(t),
		}
	} else if t := ctx.SortStatement(); t != nil {
		out.OneOf = &pb.Statement_SortStatement{
			SortStatement: SortStatement(t),
		}
	} else if t := ctx.StartStatement(); t != nil {
		out.OneOf = &pb.Statement_StartStatement{
			StartStatement: StartStatement(t),
		}
	} else if t := ctx.StopStatement(); t != nil {
		out.OneOf = &pb.Statement_StopStatement{
			StopStatement: StopStatement(t),
		}
	} else if t := ctx.StringStatement(); t != nil {
		out.OneOf = &pb.Statement_StringStatement{
			StringStatement: StringStatement(t),
		}
	} else if t := ctx.SubtractStatement(); t != nil {
		out.OneOf = &pb.Statement_SubtractStatement{
			SubtractStatement: SubtractStatement(t),
		}
	} else if t := ctx.TerminateStatement(); t != nil {
		out.OneOf = &pb.Statement_TerminateStatement{
			TerminateStatement: TerminateStatement(t),
		}
	} else if t := ctx.UnstringStatement(); t != nil {
		out.OneOf = &pb.Statement_UnstringStatement{
			UnstringStatement: UnstringStatement(t),
		}
	} else if t := ctx.WriteStatement(); t != nil {
		out.OneOf = &pb.Statement_WriteStatement{
			WriteStatement: WriteStatement(t),
		}
	}
	return
}

func AcceptStatement(in cobol85.IAcceptStatementContext) (out *pb.AcceptStatement) {
	ctx := in.(*cobol85.AcceptStatementContext)
	out = &pb.AcceptStatement{
		Identifier: Identifier(ctx.Identifier()),
	}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}

	if ictx := ctx.AcceptFromDateStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AcceptFromDateStatementContext)

		typ := pb.AcceptStatement_FromDate_DATE
		switch {
		case cctx.DATE() != nil:
			typ = pb.AcceptStatement_FromDate_DATE
		case cctx.DAY() != nil:
			typ = pb.AcceptStatement_FromDate_DAY
		case cctx.DAY_OF_WEEK() != nil:
			typ = pb.AcceptStatement_FromDate_DAY_OF_WEEK
		case cctx.TIME() != nil:
			typ = pb.AcceptStatement_FromDate_TIME
		case cctx.TIMER() != nil:
			typ = pb.AcceptStatement_FromDate_TIMER
		case cctx.TODAYS_DATE() != nil:
			typ = pb.AcceptStatement_FromDate_TODAYS_DATE
		case cctx.TODAYS_NAME() != nil:
			typ = pb.AcceptStatement_FromDate_TODAYS_NAME
		case cctx.YEAR() != nil:
			typ = pb.AcceptStatement_FromDate_YEAR
		case cctx.YYYYMMDD() != nil:
			typ = pb.AcceptStatement_FromDate_YYYYMMDD
		case cctx.YYYYDDD() != nil:
			typ = pb.AcceptStatement_FromDate_YYYYDDD
		}

		out.OneOf = &pb.AcceptStatement_FromDate_{
			FromDate: &pb.AcceptStatement_FromDate{
				Type: typ,
			},
		}
	} else if ictx := ctx.AcceptFromEscapeKeyStatement(); ictx != nil {
		out.OneOf = &pb.AcceptStatement_FromEscapeKey_{
			FromEscapeKey: &pb.AcceptStatement_FromEscapeKey{},
		}
	} else if ictx := ctx.AcceptFromMnemonicStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AcceptFromMnemonicStatementContext)
		out.OneOf = &pb.AcceptStatement_FromMnemonic_{
			FromMnemonic: &pb.AcceptStatement_FromMnemonic{
				MnemonicName: MnemonicName(cctx.MnemonicName()),
			},
		}
	} else if ictx := ctx.AcceptMessageCountStatement(); ictx != nil {
		out.OneOf = &pb.AcceptStatement_MessageCount_{
			MessageCount: &pb.AcceptStatement_MessageCount{},
		}
	}
	return
}

func AddStatement(in cobol85.IAddStatementContext) (out *pb.AddStatement) {
	ctx := in.(*cobol85.AddStatementContext)
	out = &pb.AddStatement{}
	if ctx.OnSizeErrorPhrase() != nil {
		out.OnSizeErrorPhrase = OnSizeErrorPhrase(ctx.OnSizeErrorPhrase())
	}
	if ctx.NotOnSizeErrorPhrase() != nil {
		out.NotOnSizeErrorPhrase = NotOnSizeErrorPhrase(ctx.NotOnSizeErrorPhrase())
	}

	if ictx := ctx.AddToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AddToStatementContext)
		addTo := &pb.AddStatement_To{}
		for _, v := range cctx.AllAddFrom() {
			from := &pb.AddStatement_AddFrom{}
			iv := v.(*cobol85.AddFromContext)
			if iv.Identifier() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Identifier{
					Identifier: Identifier(iv.Identifier()),
				}
			} else if iv.Literal() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Literal{
					Literal: Literal(iv.Literal()),
				}
			}
			addTo.Froms = append(addTo.Froms, from)
		}
		for _, v := range cctx.AllAddTo() {
			iv := v.(*cobol85.AddToContext)
			if iv.Identifier() != nil {
				addTo.Tos = append(addTo.Tos, Identifier(iv.Identifier()))
			}
		}
		out.OneOf = &pb.AddStatement_To_{
			To: addTo,
		}
	} else if ictx := ctx.AddToGivingStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AddToGivingStatementContext)
		addTo := &pb.AddStatement_ToGiving{}
		for _, v := range cctx.AllAddFrom() {
			iv := v.(*cobol85.AddFromContext)
			from := &pb.AddStatement_AddFrom{}
			if iv.Identifier() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Identifier{
					Identifier: Identifier(iv.Identifier()),
				}
			} else if iv.Literal() != nil {
				from.OneOf = &pb.AddStatement_AddFrom_Literal{
					Literal: Literal(iv.Literal()),
				}
			}
			addTo.Froms = append(addTo.Froms, from)
		}
		for _, v := range cctx.AllAddToGiving() {
			iv := v.(*cobol85.AddToGivingContext)
			to := &pb.AddStatement_AddToGiving{}
			if iv.Identifier() != nil {
				to.OneOf = &pb.AddStatement_AddToGiving_Identifier{
					Identifier: Identifier(iv.Identifier()),
				}
			} else if iv.Literal() != nil {
				to.OneOf = &pb.AddStatement_AddToGiving_Literal{
					Literal: Literal(iv.Literal()),
				}
			}
			addTo.Tos = append(addTo.Tos, to)
		}

		for _, v := range cctx.AllAddGiving() {
			iv := v.(*cobol85.AddGivingContext)
			if iv.Identifier() != nil {
				addTo.Givings = append(addTo.Givings, Identifier(iv.Identifier()))
			}
		}
		out.OneOf = &pb.AddStatement_ToGiving_{
			ToGiving: addTo,
		}
	} else if ictx := ctx.AddCorrespondingStatement(); ictx != nil {
		cctx := ictx.(*cobol85.AddCorrespondingStatementContext)
		addTo := &pb.AddStatement_Corresponding{}
		if cctx.Identifier() != nil {
			addTo.Corresponding = Identifier(cctx.Identifier())
		}
		if v := cctx.AddTo(); v != nil {
			iv := v.(*cobol85.AddToContext)
			addTo.To = Identifier(iv.Identifier())
		}
		out.OneOf = &pb.AddStatement_Corresponding_{
			Corresponding: addTo,
		}
	}
	return
}

func AlterStatement(in cobol85.IAlterStatementContext) (out *pb.AlterStatement) {
	ctx := in.(*cobol85.AlterStatementContext)
	out = &pb.AlterStatement{}
	for _, v := range ctx.AllAlterProceedTo() {
		iv := v.(*cobol85.AlterProceedToContext)
		names := iv.AllProcedureName()
		if len(names) == 2 {
			out.ProceedTos = append(out.ProceedTos, &pb.AlterStatement_ProceedTo{
				From: ProcedureName(names[0]),
				To:   ProcedureName(names[1]),
			})
		}
	}
	return
}

func CallStatement(in cobol85.ICallStatementContext) (out *pb.CallStatement) {
	ctx := in.(*cobol85.CallStatementContext)
	out = &pb.CallStatement{}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}
	if ctx.OnOverflowPhrase() != nil {
		out.OnOverflowPhrase = OnOverflowPhrase(ctx.OnOverflowPhrase())
	}
	if ictx := ctx.CallUsingPhrase(); ictx != nil {
		cctx := ictx.(*cobol85.CallUsingPhraseContext)
		out.UsingPhrase = &pb.CallStatement_UsingPhrase{}
		for _, v := range cctx.AllCallUsingParameter() {
			usingParameter := &pb.CallStatement_UsingParameter{}
			iv := v.(*cobol85.CallUsingParameterContext)
			if iiv := iv.CallByContentPhrase(); iiv != nil {
				byContentPhrase := &pb.CallStatement_ByContentPhrase{}
				ciiv := iiv.(*cobol85.CallByContentPhraseContext)
				for _, bc := range ciiv.AllCallByContent() {
					content := &pb.CallStatement_ByContent{}
					cbc := bc.(*cobol85.CallByContentContext)
					if cbc.Literal() != nil {
						content.OneOf = &pb.CallStatement_ByContent_Literal{
							Literal: Literal(cbc.Literal()),
						}
					} else if cbc.Identifier() != nil {
						content.OneOf = &pb.CallStatement_ByContent_Identifier{
							Identifier: Identifier(cbc.Identifier()),
						}
					} else if cbc.OMITTED() != nil {
						content.OneOf = &pb.CallStatement_ByContent_Omitted{
							Omitted: true,
						}
					}
					byContentPhrase.Contents = append(byContentPhrase.Contents, content)
				}
				usingParameter.OneOf = &pb.CallStatement_UsingParameter_ByContentPhrase{
					ByContentPhrase: byContentPhrase,
				}
			} else if iiv := iv.CallByValuePhrase(); iiv != nil {
				byValuePhrase := &pb.CallStatement_ByValuePhrase{}
				ciiv := iiv.(*cobol85.CallByValuePhraseContext)
				for _, bv := range ciiv.AllCallByValue() {
					ibv := &pb.CallStatement_ByValue{}
					cbv := bv.(*cobol85.CallByValueContext)
					if cbv.Literal() != nil {
						ibv.OneOf = &pb.CallStatement_ByValue_Literal{
							Literal: Literal(cbv.Literal()),
						}
					} else if cbv.Identifier() != nil {
						ibv.OneOf = &pb.CallStatement_ByValue_Identifier{
							Identifier: Identifier(cbv.Identifier()),
						}
					}
					byValuePhrase.Values = append(byValuePhrase.Values, ibv)
				}
				usingParameter.OneOf = &pb.CallStatement_UsingParameter_ByValuePhrase{
					ByValuePhrase: byValuePhrase,
				}
			} else if iiv := iv.CallByReferencePhrase(); iiv != nil {
				byReferencePhrase := &pb.CallStatement_ByReferencePhrase{}
				ciiv := iiv.(*cobol85.CallByReferencePhraseContext)
				for _, bv := range ciiv.AllCallByReference() {
					ibv := &pb.CallStatement_ByReference{}
					cbv := bv.(*cobol85.CallByReferenceContext)
					if cbv.Literal() != nil {
						ibv.OneOf = &pb.CallStatement_ByReference_Literal{
							Literal: Literal(cbv.Literal()),
						}
					} else if cbv.Identifier() != nil {
						ibv.OneOf = &pb.CallStatement_ByReference_Identifier{
							Identifier: Identifier(cbv.Identifier()),
						}
					} else if cbv.FileName() != nil {
						ibv.OneOf = &pb.CallStatement_ByReference_FileName{
							FileName: FileName(cbv.FileName()),
						}
					} else if cbv.OMITTED() != nil {
						ibv.OneOf = &pb.CallStatement_ByReference_Omitted{
							Omitted: true,
						}
					}
					byReferencePhrase.Refs = append(byReferencePhrase.Refs, ibv)
				}
				usingParameter.OneOf = &pb.CallStatement_UsingParameter_ByReferencePhrase{
					ByReferencePhrase: byReferencePhrase,
				}
			}
			out.UsingPhrase.Parameters = append(out.UsingPhrase.Parameters, usingParameter)
		}
	} else if ictx := ctx.CallGivingPhrase(); ictx != nil {
		cctx := ictx.(*cobol85.CallGivingPhraseContext)
		out.GivingPhrase = &pb.CallStatement_GivingPhrase{
			Identifier: Identifier(cctx.Identifier()),
		}
	}
	return
}
func CancelStatement(in cobol85.ICancelStatementContext) (out *pb.CancelStatement) {
	ctx := in.(*cobol85.CancelStatementContext)
	out = &pb.CancelStatement{}
	for _, v := range ctx.AllCancelCall() {
		call := &pb.CancelCall{}
		cc := v.(*cobol85.CancelCallContext)
		if cc.LibraryName() != nil {
			call.OneOf = &pb.CancelCall_LibraryName{
				LibraryName: LibraryName(cc.LibraryName()),
			}
		} else if cc.Literal() != nil {
			call.OneOf = &pb.CancelCall_Literal{
				Literal: Literal(cc.Literal()),
			}
		} else if cc.Identifier() != nil {
			call.OneOf = &pb.CancelCall_Identifier{
				Identifier: Identifier(cc.Identifier()),
			}
		}
		out.Calls = append(out.Calls, call)
	}
	return
}
func CloseStatement(in cobol85.ICloseStatementContext) (out *pb.CloseStatement) {
	ctx := in.(*cobol85.CloseStatementContext)
	out = &pb.CloseStatement{}
	for _, v := range ctx.AllCloseFile() {
		cc := v.(*cobol85.CloseFileContext)
		cf := &pb.CloseFile{}
		if cc.FileName() != nil {
			cf.FileName = FileName(cc.FileName())
		}

		if cc.CloseRelativeStatement() != nil {
			cf.Statement = &pb.CloseFile_CloseRelativeStatement{
				CloseRelativeStatement: &pb.CloseRelativeStatement{},
			}
		} else if ic := cc.ClosePortFileIOStatement(); ic != nil {
			icc := ic.(*cobol85.ClosePortFileIOStatementContext)
			closePortFileIoStatement := &pb.ClosePortFileIOStatement{
				Wait: icc.WAIT() != nil,
			}
			for _, vv := range icc.AllClosePortFileIOUsing() {
				using := &pb.ClosePortFileIOUsing{}
				ivv := vv.(*cobol85.ClosePortFileIOUsingContext)
				if idata := ivv.ClosePortFileIOUsingAssociatedData(); idata != nil {
					data := idata.(*cobol85.ClosePortFileIOUsingAssociatedDataContext)
					associatedData := &pb.ClosePortFileIOUsing_AssociatedData{}
					if data.IntegerLiteral() != nil {
						associatedData.OneOf = &pb.ClosePortFileIOUsing_AssociatedData_IntegerLiteral{
							IntegerLiteral: IntegerLiteral(data.IntegerLiteral()),
						}
					} else if data.Identifier() != nil {
						associatedData.OneOf = &pb.ClosePortFileIOUsing_AssociatedData_Identifier{
							Identifier: Identifier(data.Identifier()),
						}
					}
					using.OneOf = &pb.ClosePortFileIOUsing_AssociatedData_{
						AssociatedData: associatedData,
					}
				} else if idata := ivv.ClosePortFileIOUsingAssociatedDataLength(); idata != nil {
					data := idata.(*cobol85.ClosePortFileIOUsingAssociatedDataLengthContext)
					associatedDataLength := &pb.ClosePortFileIOUsing_AssociatedDataLength{}
					if data.IntegerLiteral() != nil {
						associatedDataLength.OneOf = &pb.ClosePortFileIOUsing_AssociatedDataLength_IntegerLiteral{
							IntegerLiteral: IntegerLiteral(data.IntegerLiteral()),
						}
					} else if data.Identifier() != nil {
						associatedDataLength.OneOf = &pb.ClosePortFileIOUsing_AssociatedDataLength_Identifier{
							Identifier: Identifier(data.Identifier()),
						}
					}
					using.OneOf = &pb.ClosePortFileIOUsing_AssociatedDataLength_{
						AssociatedDataLength: associatedDataLength,
					}
				} else if idata := ivv.ClosePortFileIOUsingCloseDisposition(); idata != nil {
					using.OneOf = &pb.ClosePortFileIOUsing_CloseDisposition_{
						CloseDisposition: &pb.ClosePortFileIOUsing_CloseDisposition{},
					}
				}
				closePortFileIoStatement.Usings = append(closePortFileIoStatement.Usings, using)
			}
			cf.Statement = &pb.CloseFile_ClosePortFileIoStatement{
				ClosePortFileIoStatement: closePortFileIoStatement,
			}
		} else if ic := cc.CloseReelUnitStatement(); ic != nil {
			cf.Statement = &pb.CloseFile_CloseReelUnitStatement{
				CloseReelUnitStatement: &pb.CloseReelUnitStatement{},
			}
		}
		out.CloseFiles = append(out.CloseFiles, cf)
	}
	return
}
func ComputeStatement(in cobol85.IComputeStatementContext) (out *pb.ComputeStatement) {
	ctx := in.(*cobol85.ComputeStatementContext)
	out = &pb.ComputeStatement{}
	if ctx.OnSizeErrorPhrase() != nil {
		out.OnSizeErrorPhrase = OnSizeErrorPhrase(ctx.OnSizeErrorPhrase())
	}
	if ctx.NotOnSizeErrorPhrase() != nil {
		out.NotOnSizeErrorPhrase = NotOnSizeErrorPhrase(ctx.NotOnSizeErrorPhrase())
	}
	for _, v := range ctx.AllComputeStore() {
		iv := v.(*cobol85.ComputeStoreContext)
		store := &pb.ComputeStore{
			Identifier: Identifier(iv.Identifier()),
		}
		out.Stores = append(out.Stores, store)
	}
	if ctx.ArithmeticExpression() != nil {
		out.ArithmeticExpression = ArithmeticExpression(ctx.ArithmeticExpression())
	}
	return
}
func ContinueStatement(in cobol85.IContinueStatementContext) (out *pb.ContinueStatement) {
	out = &pb.ContinueStatement{}
	return
}
func DeleteStatement(in cobol85.IDeleteStatementContext) (out *pb.DeleteStatement) {
	ctx := in.(*cobol85.DeleteStatementContext)
	out = &pb.DeleteStatement{}
	if ctx.FileName() != nil {
		out.FileName = FileName(ctx.FileName())
	}
	if ctx.InvalidKeyPhrase() != nil {
		out.InvalidKeyPhrase = InvalidKeyPhrase(ctx.InvalidKeyPhrase())
	}

	if ctx.NotInvalidKeyPhrase() != nil {
		out.NotInvalidKeyPhrase = NotInvalidKeyPhrase(ctx.NotInvalidKeyPhrase())
	}
	return
}
func DisableStatement(in cobol85.IDisableStatementContext) (out *pb.DisableStatement) {
	ctx := in.(*cobol85.DisableStatementContext)
	out = &pb.DisableStatement{}
	if ctx.CdName() != nil {
		out.CdName = CdName(ctx.CdName())
	}
	if ctx.Literal() != nil {
		out.Key = &pb.DisableStatement_Literal{
			Literal: Literal(ctx.Literal()),
		}
	} else if ctx.Identifier() != nil {
		out.Key = &pb.DisableStatement_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	}
	return
}

func DisplayStatement(in cobol85.IDisplayStatementContext) (out *pb.DisplayStatement) {
	ctx := in.(*cobol85.DisplayStatementContext)
	out = &pb.DisplayStatement{}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}
	for _, v := range ctx.AllDisplayOperand() {
		iv := v.(*cobol85.DisplayOperandContext)
		operand := &pb.DisplayOperand{}
		if iv.Identifier() != nil {
			operand.OneOf = &pb.DisplayOperand_Identifier{
				Identifier: Identifier(iv.Identifier()),
			}
		} else if iv.Literal() != nil {
			operand.OneOf = &pb.DisplayOperand_Literal{
				Literal: Literal(iv.Literal()),
			}
		}
		out.Operands = append(out.Operands, operand)
	}

	if ictx := ctx.DisplayAt(); ictx != nil {
		cctx := ictx.(*cobol85.DisplayAtContext)
		at := &pb.DisplayAt{}
		if cctx.Identifier() != nil {
			at.OneOf = &pb.DisplayAt_Identifier{
				Identifier: Identifier(cctx.Identifier()),
			}
		} else if cctx.Literal() != nil {
			at.OneOf = &pb.DisplayAt_Literal{
				Literal: Literal(cctx.Literal()),
			}
		}
		out.At = at
	}

	if ictx := ctx.DisplayUpon(); ictx != nil {
		cctx := ictx.(*cobol85.DisplayUponContext)
		upon := &pb.DisplayUpon{}
		if cctx.MnemonicName() != nil {
			upon.OneOf = &pb.DisplayUpon_MnemonicName{
				MnemonicName: MnemonicName(cctx.MnemonicName()),
			}
		} else if cctx.EnvironmentName() != nil {
			upon.OneOf = &pb.DisplayUpon_EnvironmentName{
				EnvironmentName: EnvironmentName(cctx.EnvironmentName()),
			}
		}
		out.Upon = upon
	}

	if ictx := ctx.DisplayWith(); ictx != nil {
		out.With = &pb.DisplayWith{}
	}
	return
}
func DivideStatement(in cobol85.IDivideStatementContext) (out *pb.DivideStatement) {
	ctx := in.(*cobol85.DivideStatementContext)
	out = &pb.DivideStatement{}
	if ctx.OnSizeErrorPhrase() != nil {
		out.OnSizeErrorPhrase = OnSizeErrorPhrase(ctx.OnSizeErrorPhrase())
	}
	if ctx.NotOnSizeErrorPhrase() != nil {
		out.NotOnSizeErrorPhrase = NotOnSizeErrorPhrase(ctx.NotOnSizeErrorPhrase())
	}
	if ctx.Identifier() != nil {
		out.Divide = &pb.DivideStatement_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	} else if ctx.Literal() != nil {
		out.Divide = &pb.DivideStatement_Literal{
			Literal: Literal(ctx.Literal()),
		}
	}

	if ictx := ctx.DivideIntoStatement(); ictx != nil {
		cctx := ictx.(*cobol85.DivideIntoStatementContext)
		into := &pb.DivideStatement_IntoStatement{}
		for _, v := range cctx.AllDivideInto() {
			iv := v.(*cobol85.DivideIntoContext)
			into.Intos = append(into.Intos, &pb.DivideStatement_Into{
				Identifier: Identifier(iv.Identifier()),
			})
		}
		out.Statement = &pb.DivideStatement_IntoStatement_{
			IntoStatement: into,
		}
	} else if ictx := ctx.DivideIntoGivingStatement(); ictx != nil {
		cctx := ictx.(*cobol85.DivideIntoGivingStatementContext)
		giving := &pb.DivideStatement_IntoGivingStatement{}
		if cctx.Identifier() != nil {
			giving.Into = &pb.DivideStatement_IntoGivingStatement_Identifier{
				Identifier: Identifier(cctx.Identifier()),
			}
		} else if cctx.Literal() != nil {
			giving.Into = &pb.DivideStatement_IntoGivingStatement_Literal{
				Literal: Literal(cctx.Literal()),
			}
		}

		if idgp := cctx.DivideGivingPhrase(); idgp != nil {
			givingPhrase := &pb.DivideStatement_GivingPhrase{}
			cdgp := idgp.(*cobol85.DivideGivingPhraseContext)
			for _, v := range cdgp.AllDivideGiving() {
				cv := v.(*cobol85.DivideGivingContext)
				givingPhrase.Givings = append(givingPhrase.Givings, &pb.DivideStatement_Giving{
					Identifier: Identifier(cv.Identifier()),
				})
			}
			giving.GivingPhrase = givingPhrase
		}
		out.Statement = &pb.DivideStatement_IntoGivingStatement_{
			IntoGivingStatement: giving,
		}
	} else if ictx := ctx.DivideByGivingStatement(); ictx != nil {
		cctx := ictx.(*cobol85.DivideByGivingStatementContext)
		giving := &pb.DivideStatement_ByGivingStatement{}
		if cctx.Identifier() != nil {
			giving.By = &pb.DivideStatement_ByGivingStatement_Identifier{
				Identifier: Identifier(cctx.Identifier()),
			}
		} else if cctx.Literal() != nil {
			giving.By = &pb.DivideStatement_ByGivingStatement_Literal{
				Literal: Literal(cctx.Literal()),
			}
		}

		if idgp := cctx.DivideGivingPhrase(); idgp != nil {
			givingPhrase := &pb.DivideStatement_GivingPhrase{}
			cdgp := idgp.(*cobol85.DivideGivingPhraseContext)
			for _, v := range cdgp.AllDivideGiving() {
				cv := v.(*cobol85.DivideGivingContext)
				givingPhrase.Givings = append(givingPhrase.Givings, &pb.DivideStatement_Giving{
					Identifier: Identifier(cv.Identifier()),
				})
			}
			giving.GivingPhrase = givingPhrase
		}
		out.Statement = &pb.DivideStatement_ByGivingStatement_{
			ByGivingStatement: giving,
		}
	}

	if ictx := ctx.DivideRemainder(); ictx != nil {
		cctx := ictx.(*cobol85.DivideRemainderContext)
		out.Remainder = &pb.DivideStatement_Remainder{
			Identifier: Identifier(cctx.Identifier()),
		}
	}
	return
}
func EnableStatement(in cobol85.IEnableStatementContext) (out *pb.EnableStatement) {
	ctx := in.(*cobol85.EnableStatementContext)
	out = &pb.EnableStatement{}

	if ctx.CdName() != nil {
		out.CdName = CdName(ctx.CdName())
	}

	if ctx.Identifier() != nil {
		out.Key = &pb.EnableStatement_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	}

	if ctx.Literal() != nil {
		out.Key = &pb.EnableStatement_Literal{
			Literal: Literal(ctx.Literal()),
		}
	}
	return
}
func EntryStatement(in cobol85.IEntryStatementContext) (out *pb.EntryStatement) {
	ctx := in.(*cobol85.EntryStatementContext)
	out = &pb.EntryStatement{}

	if ctx.Literal() != nil {
		out.Entry = Literal(ctx.Literal())
	}

	for _, v := range ctx.AllIdentifier() {
		out.Usings = append(out.Usings, Identifier(v))
	}
	return
}

func EvaluateSelect(in cobol85.IEvaluateSelectContext) (out *pb.EvaluateStatement_Select) {
	ctx := in.(*cobol85.EvaluateSelectContext)
	out = &pb.EvaluateStatement_Select{}
	if ctx.Identifier() != nil {
		out.OneOf = &pb.EvaluateStatement_Select_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	} else if ctx.Literal() != nil {
		out.OneOf = &pb.EvaluateStatement_Select_Literal{
			Literal: Literal(ctx.Literal()),
		}
	} else if ctx.Condition() != nil {
		out.OneOf = &pb.EvaluateStatement_Select_Condition{
			Condition: Condition(ctx.Condition()),
		}
	} else if ctx.ArithmeticExpression() != nil {
		out.OneOf = &pb.EvaluateStatement_Select_ArithmeticExpression{
			ArithmeticExpression: ArithmeticExpression(ctx.ArithmeticExpression()),
		}
	}
	return
}

func EvaluateValue(in cobol85.IEvaluateValueContext) (out *pb.EvaluateStatement_Value) {
	ctx := in.(*cobol85.EvaluateValueContext)
	out = &pb.EvaluateStatement_Value{}
	if ctx.Literal() != nil {
		out = &pb.EvaluateStatement_Value{
			OneOf: &pb.EvaluateStatement_Value_Literal{
				Literal: Literal(ctx.Literal()),
			},
		}
	} else if ctx.Identifier() != nil {
		out = &pb.EvaluateStatement_Value{
			OneOf: &pb.EvaluateStatement_Value_Identifier{
				Identifier: Identifier(ctx.Identifier()),
			},
		}
	} else if ctx.ArithmeticExpression() != nil {
		out = &pb.EvaluateStatement_Value{
			OneOf: &pb.EvaluateStatement_Value_ArithmeticExpression{
				ArithmeticExpression: ArithmeticExpression(ctx.ArithmeticExpression()),
			},
		}
	}
	return
}

func EvaluateCondition(in cobol85.IEvaluateConditionContext) (out *pb.EvaluateStatement_Condition) {
	ctx := in.(*cobol85.EvaluateConditionContext)
	out = &pb.EvaluateStatement_Condition{}
	if ctx.ANY() != nil {
		out.OneOf = &pb.EvaluateStatement_Condition_Any{
			Any: true,
		}
	} else if ctx.Condition() != nil {
		out.OneOf = &pb.EvaluateStatement_Condition_Condition{
			Condition: Condition(ctx.Condition()),
		}
	} else if ctx.BooleanLiteral() != nil {
		out.OneOf = &pb.EvaluateStatement_Condition_BooleanLiteral{
			BooleanLiteral: BooleanLiteral(ctx.BooleanLiteral()),
		}
	} else if ictx := ctx.EvaluateValue(); ictx != nil {
		valueThrough := &pb.EvaluateStatement_ValueThrough{
			Value: EvaluateValue(ictx),
		}

		if ictx := ctx.EvaluateThrough(); ictx != nil {
			cctx := ictx.(*cobol85.EvaluateThroughContext)
			through := &pb.EvaluateStatement_Through{}
			if cctx.EvaluateValue() != nil {
				through.Value = EvaluateValue(cctx.EvaluateValue())
			}
			valueThrough.Through = through
		}
		out.OneOf = &pb.EvaluateStatement_Condition_ValueThrough{
			ValueThrough: valueThrough,
		}
	}
	return
}

func EvaluateStatement(in cobol85.IEvaluateStatementContext) (out *pb.EvaluateStatement) {
	ctx := in.(*cobol85.EvaluateStatementContext)
	out = &pb.EvaluateStatement{}

	if ictx := ctx.EvaluateSelect(); ictx != nil {
		out.Select = EvaluateSelect(ictx)
	}

	for _, iv := range ctx.AllEvaluateAlsoSelect() {
		cv := iv.(*cobol85.EvaluateAlsoSelectContext)
		alsoSelect := &pb.EvaluateStatement_AlsoSelect{
			Select: EvaluateSelect(cv.EvaluateSelect()),
		}
		out.AlsoSelects = append(out.AlsoSelects, alsoSelect)
	}

	for _, iv := range ctx.AllEvaluateWhenPhrase() {
		cv := iv.(*cobol85.EvaluateWhenPhraseContext)
		when := &pb.EvaluateStatement_WhenPhrase{}
		for _, ivv := range cv.AllStatement() {
			when.Statements = append(when.Statements, Statement(ivv))
		}
		for _, ivv := range cv.AllEvaluateWhen() {
			cvv := ivv.(*cobol85.EvaluateWhenContext)
			w := &pb.EvaluateStatement_When{}
			if cvv.EvaluateCondition() != nil {
				w.Condition = EvaluateCondition(cvv.EvaluateCondition())
			}
			for _, ivvv := range cvv.AllEvaluateAlsoCondition() {
				cvvv := ivvv.(*cobol85.EvaluateAlsoConditionContext)
				if cvvv.EvaluateCondition() != nil {
					w.Alsos = append(w.Alsos, &pb.EvaluateStatement_AlsoCondition{
						Condition: EvaluateCondition(cvvv.EvaluateCondition()),
					})
				}
			}
			when.Whens = append(when.Whens, w)
		}
		out.WhenPhrases = append(out.WhenPhrases, when)
	}

	if ictx := ctx.EvaluateWhenOther(); ictx != nil {
		cctx := ictx.(*cobol85.EvaluateWhenOtherContext)
		whenOther := &pb.EvaluateStatement_WhenOther{}
		for _, iv := range cctx.AllStatement() {
			whenOther.Statements = append(whenOther.Statements, Statement(iv))
		}
		out.WhenOther = whenOther
	}
	return
}
func ExhibitStatement(in cobol85.IExhibitStatementContext) (out *pb.ExhibitStatement) {
	ctx := in.(*cobol85.ExhibitStatementContext)
	out = &pb.ExhibitStatement{}

	for _, iv := range ctx.AllExhibitOperand() {
		cv := iv.(*cobol85.ExhibitOperandContext)
		operand := &pb.ExhibitStatement_Operand{}
		if cv.Identifier() != nil {
			operand.OneOf = &pb.ExhibitStatement_Operand_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			operand.OneOf = &pb.ExhibitStatement_Operand_Literal{
				Literal: Literal(cv.Literal()),
			}
		}
		out.Operands = append(out.Operands, operand)
	}
	return
}
func ExecCicsStatement(in cobol85.IExecCicsStatementContext) (out *pb.ExecCicsStatement) {
	out = &pb.ExecCicsStatement{}
	return
}
func ExecSqlStatement(in cobol85.IExecSqlStatementContext) (out *pb.ExecSqlStatement) {
	out = &pb.ExecSqlStatement{}
	return
}
func ExecSqlImsStatement(in cobol85.IExecSqlImsStatementContext) (out *pb.ExecSqlImsStatement) {
	out = &pb.ExecSqlImsStatement{}
	return
}
func ExitStatement(in cobol85.IExitStatementContext) (out *pb.ExitStatement) {
	out = &pb.ExitStatement{}
	return
}
func GenerateStatement(in cobol85.IGenerateStatementContext) (out *pb.GenerateStatement) {
	ctx := in.(*cobol85.GenerateStatementContext)
	out = &pb.GenerateStatement{}
	if ctx.ReportName() != nil {
		out.ReportName = ReportName(ctx.ReportName())
	}
	return
}
func GobackStatement(in cobol85.IGobackStatementContext) (out *pb.GobackStatement) {
	out = &pb.GobackStatement{}
	return
}
func GoToStatement(in cobol85.IGoToStatementContext) (out *pb.GoToStatement) {
	ctx := in.(*cobol85.GoToStatementContext)
	out = &pb.GoToStatement{}
	if ictx := ctx.GoToStatementSimple(); ictx != nil {
		cctx := ictx.(*cobol85.GoToStatementSimpleContext)
		simple := &pb.GoToStatement_SimpleStatement{}
		if cctx.ProcedureName() != nil {
			simple.ProcedureName = ProcedureName(cctx.ProcedureName())
		}
		out.Statement = &pb.GoToStatement_Simple{
			Simple: simple,
		}
	} else if ictx := ctx.GoToDependingOnStatement(); ictx != nil {
		cctx := ictx.(*cobol85.GoToDependingOnStatementContext)
		depending := &pb.GoToStatement_DependingOn{}
		for _, iv := range cctx.AllProcedureName() {
			depending.ProcedureNames = append(depending.ProcedureNames, ProcedureName(iv))
		}
		if cctx.Identifier() != nil {
			depending.DependingOn = Identifier(cctx.Identifier())
		}
		if cctx.MORE_LABELS() != nil {
			out.Statement = &pb.GoToStatement_DependingOn_{
				DependingOn: &pb.GoToStatement_DependingOnStatement{
					OneOf: &pb.GoToStatement_DependingOnStatement_MoreLabels{
						MoreLabels: true,
					},
				},
			}
		} else {
			out.Statement = &pb.GoToStatement_DependingOn_{
				DependingOn: &pb.GoToStatement_DependingOnStatement{
					OneOf: &pb.GoToStatement_DependingOnStatement_DependingOn{
						DependingOn: depending,
					},
				},
			}
		}
	}
	return
}
func IfStatement(in cobol85.IIfStatementContext) (out *pb.IfStatement) {
	ctx := in.(*cobol85.IfStatementContext)
	out = &pb.IfStatement{}
	if ctx.Condition() != nil {
		out.Condition = Condition(ctx.Condition())
	}
	if ictx := ctx.IfThen(); ictx != nil {
		cctx := ictx.(*cobol85.IfThenContext)
		out.Then = &pb.IfStatement_Then{}
		if cctx.NEXT() != nil || cctx.SENTENCE() != nil {
			out.Then.NextSentence = true
		} else {
			for _, iv := range cctx.AllStatement() {
				out.Then.Statements = append(out.Then.Statements, Statement(iv))
			}
		}
	}

	if ictx := ctx.IfElse(); ictx != nil {
		cctx := ictx.(*cobol85.IfElseContext)
		out.Else = &pb.IfStatement_Else{}
		if cctx.NEXT() != nil || cctx.SENTENCE() != nil {
			out.Else.NextSentence = true
		} else {
			for _, iv := range cctx.AllStatement() {
				out.Else.Statements = append(out.Else.Statements, Statement(iv))
			}
		}
	}
	return
}
func InitializeStatement(in cobol85.IInitializeStatementContext) (out *pb.InitializeStatement) {
	ctx := in.(*cobol85.InitializeStatementContext)
	out = &pb.InitializeStatement{}
	for _, iv := range ctx.AllIdentifier() {
		out.Identifiers = append(out.Identifiers, Identifier(iv))
	}
	if ictx := ctx.InitializeReplacingPhrase(); ictx != nil {
		cctx := ictx.(*cobol85.InitializeReplacingPhraseContext)
		phrase := &pb.InitializeStatement_ReplacingPhrase{}
		for _, iv := range cctx.AllInitializeReplacingBy() {
			cv := iv.(*cobol85.InitializeReplacingByContext)
			by := &pb.InitializeStatement_ReplacingBy{}
			if cv.ALPHABETIC() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_ALPHABETIC
			} else if cv.ALPHANUMERIC() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_ALPHANUMERIC
			} else if cv.ALPHANUMERIC_EDITED() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_ALPHANUMERIC_EDITED
			} else if cv.NATIONAL() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_NATIONAL
			} else if cv.NATIONAL_EDITED() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_NATIONAL_EDITED
			} else if cv.NUMERIC() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_NUMERIC
			} else if cv.NUMERIC_EDITED() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_NUMERIC_EDITED
			} else if cv.DBCS() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_DBCS
			} else if cv.EGCS() != nil {
				by.Type = pb.InitializeStatement_ReplacingBy_EGCS
			}

			if cv.Identifier() != nil {
				by.By = &pb.InitializeStatement_ReplacingBy_Identifier{
					Identifier: Identifier(cv.Identifier()),
				}
			} else if cv.Literal() != nil {
				by.By = &pb.InitializeStatement_ReplacingBy_Literal{
					Literal: Literal(cv.Literal()),
				}
			}
			phrase.Bys = append(phrase.Bys, by)
		}
		out.ReplacingPhrase = phrase
	}
	return
}
func InitiateStatement(in cobol85.IInitiateStatementContext) (out *pb.InitiateStatement) {
	ctx := in.(*cobol85.InitiateStatementContext)
	out = &pb.InitiateStatement{}
	for _, iv := range ctx.AllReportName() {
		out.ReportNames = append(out.ReportNames, ReportName(iv))
	}
	return
}

func CharactersLeadings(inspectCharacters []cobol85.IInspectCharactersContext, allLeadings []cobol85.IInspectAllLeadingsContext) (out []*pb.InspectStatement_CharactersLeadings) {
	out = []*pb.InspectStatement_CharactersLeadings{}
	if len(inspectCharacters) == len(allLeadings) {
		for k, ivv := range inspectCharacters {
			cvv := ivv.(*cobol85.InspectCharactersContext)
			characters := &pb.InspectStatement_Characters{}
			for _, ivvv := range cvv.AllInspectBeforeAfter() {
				cvvv := ivvv.(*cobol85.InspectBeforeAfterContext)
				beforeAfter := &pb.InspectStatement_BeforeAfter{}
				if cvvv.Literal() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Literal{
						Literal: Literal(cvvv.Literal()),
					}
				} else if cvvv.Identifier() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Identifier{
						Identifier: Identifier(cvvv.Identifier()),
					}
				}
				characters.BeforeAfters = append(characters.BeforeAfters, beforeAfter)
			}
			ikv := allLeadings[k]
			ckv := ikv.(*cobol85.InspectAllLeadingsContext)
			allLeadings := &pb.InspectStatement_AllLeadings{}
			for _, ivvv := range ckv.AllInspectAllLeading() {
				cvvv := ivvv.(*cobol85.InspectAllLeadingContext)
				allLeading := &pb.InspectStatement_AllLeading{}
				beforeAfter := &pb.InspectStatement_BeforeAfter{}
				if cvvv.Literal() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Literal{
						Literal: Literal(cvvv.Literal()),
					}
				} else if cvvv.Identifier() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Identifier{
						Identifier: Identifier(cvvv.Identifier()),
					}
				}
				allLeading.BeforeAfters = append(allLeading.BeforeAfters, beforeAfter)

				if cvvv.Identifier() != nil {
					allLeading.Value = &pb.InspectStatement_AllLeading_Identifier{
						Identifier: Identifier(cvvv.Identifier()),
					}
				} else if cvvv.Literal() != nil {
					allLeading.Value = &pb.InspectStatement_AllLeading_Literal{
						Literal: Literal(cvvv.Literal()),
					}
				}
				allLeadings.AllLeadings = append(allLeadings.AllLeadings, allLeading)
			}
			out = append(out, &pb.InspectStatement_CharactersLeadings{
				Characters:  characters,
				AllLeadings: allLeadings,
			})
		}
	}
	return
}

func InspectFor(in cobol85.IInspectForContext) (out *pb.InspectStatement_For) {
	ctx := in.(*cobol85.InspectForContext)
	out = &pb.InspectStatement_For{}
	if ctx.Identifier() != nil {
		out.Identifier = Identifier(ctx.Identifier())
	}
	out.CharactersLeadings = CharactersLeadings(ctx.AllInspectCharacters(), ctx.AllInspectAllLeadings())
	return
}

func InspectReplacingPhrase(in cobol85.IInspectReplacingPhraseContext) (out *pb.InspectStatement_ReplacingPhrase) {
	out = &pb.InspectStatement_ReplacingPhrase{}
	ctx := in.(*cobol85.InspectReplacingPhraseContext)

	characters := ctx.AllInspectReplacingCharacters()
	for k, ivv := range ctx.AllInspectReplacingAllLeadings() {
		cvv := ivv.(*cobol85.InspectReplacingAllLeadingsContext)
		replacingAllLeadings := &pb.InspectStatement_ReplacingAllLeadings{}
		for _, ivvv := range cvv.AllInspectReplacingAllLeading() {
			cvvv := ivvv.(*cobol85.InspectReplacingAllLeadingContext)
			replacingAllLeading := &pb.InspectStatement_ReplacingAllLeading{}
			for _, ivvvv := range cvvv.AllInspectBeforeAfter() {
				cvvvv := ivvvv.(*cobol85.InspectBeforeAfterContext)
				beforeAfter := &pb.InspectStatement_BeforeAfter{}
				if cvvvv.Identifier() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Identifier{
						Identifier: Identifier(cvvvv.Identifier()),
					}
				} else if cvvvv.Literal() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Literal{
						Literal: Literal(cvvvv.Literal()),
					}
				}
				replacingAllLeading.BeforeAfters = append(replacingAllLeading.BeforeAfters, beforeAfter)
			}
			if cvvv.Identifier() != nil {
				replacingAllLeading.Value = &pb.InspectStatement_ReplacingAllLeading_Identifier{
					Identifier: Identifier(cvvv.Identifier()),
				}
			} else if cvvv.Literal() != nil {
				replacingAllLeading.Value = &pb.InspectStatement_ReplacingAllLeading_Literal{
					Literal: Literal(cvvv.Literal()),
				}
			}

			if i := cvvv.InspectBy(); i != nil {
				by := &pb.InspectStatement_By{}
				ci := i.(*cobol85.InspectByContext)
				if ci.Identifier() == nil {
					by.OneOf = &pb.InspectStatement_By_Identifier{
						Identifier: Identifier(ci.Identifier()),
					}
				} else if ci.Literal() != nil {
					by.OneOf = &pb.InspectStatement_By_Literal{
						Literal: Literal(ci.Literal()),
					}
				}
				replacingAllLeading.By = by
			}
			replacingAllLeadings.ReplacingAllLeadings = append(replacingAllLeadings.ReplacingAllLeadings, replacingAllLeading)
		}

		replacingCharacters := &pb.InspectStatement_ReplacingCharacters{}
		if len(characters) > k {
			cvv := characters[k].(*cobol85.InspectReplacingCharactersContext)
			for _, ivvvv := range cvv.AllInspectBeforeAfter() {
				cvvv := ivvvv.(*cobol85.InspectBeforeAfterContext)
				beforeAfter := &pb.InspectStatement_BeforeAfter{}
				if cvvv.Identifier() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Identifier{
						Identifier: Identifier(cvvv.Identifier()),
					}
				} else if cvvv.Literal() != nil {
					beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Literal{
						Literal: Literal(cvvv.Literal()),
					}
				}
				replacingCharacters.BeforeAfters = append(replacingCharacters.BeforeAfters, beforeAfter)
			}
			if i := cvv.InspectBy(); i != nil {
				by := &pb.InspectStatement_By{}
				ci := i.(*cobol85.InspectByContext)
				if ci.Identifier() == nil {
					by.OneOf = &pb.InspectStatement_By_Identifier{
						Identifier: Identifier(ci.Identifier()),
					}
				} else if ci.Literal() != nil {
					by.OneOf = &pb.InspectStatement_By_Literal{
						Literal: Literal(ci.Literal()),
					}
				}
				replacingCharacters.By = by
			}
		}
		replacingCharactersLeading := &pb.InspectStatement_ReplacingCharactersLeadings{
			ReplacingAllLeadings: replacingAllLeadings,
			ReplacingCharacters:  replacingCharacters,
		}
		out.ReplacingCharactersLeadings = append(out.ReplacingCharactersLeadings, replacingCharactersLeading)
	}
	return
}

func InspectStatement(in cobol85.IInspectStatementContext) (out *pb.InspectStatement) {
	ctx := in.(*cobol85.InspectStatementContext)
	out = &pb.InspectStatement{}
	if ctx.Identifier() != nil {
		out.Identifier = Identifier(ctx.Identifier())
	}
	if ictx := ctx.InspectTallyingPhrase(); ictx != nil {
		cctx := ictx.(*cobol85.InspectTallyingPhraseContext)
		phrase := &pb.InspectStatement_TallyingPhrase{}
		for _, iv := range cctx.AllInspectFor() {
			phrase.Fors = append(phrase.Fors, InspectFor(iv))
		}
		out.Phrase = &pb.InspectStatement_TallyingPhrase_{
			TallyingPhrase: phrase,
		}
	} else if ictx := ctx.InspectTallyingReplacingPhrase(); ictx != nil {
		cctx := ictx.(*cobol85.InspectTallyingReplacingPhraseContext)
		phrase := &pb.InspectStatement_TallyingReplacingPhrase{}
		for _, iv := range cctx.AllInspectFor() {
			phrase.Fors = append(phrase.Fors, InspectFor(iv))
		}
		for _, iv := range cctx.AllInspectReplacingPhrase() {
			replacingPhrase := InspectReplacingPhrase(iv)
			phrase.ReplacingPhrases = append(phrase.ReplacingPhrases, replacingPhrase)
		}
		out.Phrase = &pb.InspectStatement_TallyingReplacingPhrase_{
			TallyingReplacingPhrase: phrase,
		}
	} else if ictx := ctx.InspectReplacingPhrase(); ictx != nil {
		out.Phrase = &pb.InspectStatement_ReplacingPhrase_{
			ReplacingPhrase: InspectReplacingPhrase(ictx),
		}
	} else if ictx := ctx.InspectConvertingPhrase(); ictx != nil {
		cctx := ictx.(*cobol85.InspectConvertingPhraseContext)
		phrase := &pb.InspectStatement_ConvertingPhrase{}
		for _, ivvvv := range cctx.AllInspectBeforeAfter() {
			cvvvv := ivvvv.(*cobol85.InspectBeforeAfterContext)
			beforeAfter := &pb.InspectStatement_BeforeAfter{}
			if cvvvv.Identifier() != nil {
				beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Identifier{
					Identifier: Identifier(cvvvv.Identifier()),
				}
			} else if cvvvv.Literal() != nil {
				beforeAfter.OneOf = &pb.InspectStatement_BeforeAfter_Literal{
					Literal: Literal(cvvvv.Literal()),
				}
			}
			phrase.BeforeAfters = append(phrase.BeforeAfters, beforeAfter)
		}
		if iv := cctx.InspectTo(); iv != nil {
			cv := iv.(*cobol85.InspectToContext)
			to := &pb.InspectStatement_To{}
			if cv.Identifier() != nil {
				to.OneOf = &pb.InspectStatement_To_Identifier{
					Identifier: Identifier(cv.Identifier()),
				}
			} else if cv.Literal() != nil {
				to.OneOf = &pb.InspectStatement_To_Literal{
					Literal: Literal(cv.Literal()),
				}
			}
			phrase.To = to
		}

		if cctx.Identifier() != nil {
			phrase.Value = &pb.InspectStatement_ConvertingPhrase_Identifier{
				Identifier: Identifier(cctx.Identifier()),
			}
		} else if cctx.Literal() != nil {
			phrase.Value = &pb.InspectStatement_ConvertingPhrase_Literal{
				Literal: Literal(cctx.Literal()),
			}
		}
		out.Phrase = &pb.InspectStatement_ConvertingPhrase_{
			ConvertingPhrase: phrase,
		}
	}
	return
}
func MergeStatement(in cobol85.IMergeStatementContext) (out *pb.MergeStatement) {
	ctx := in.(*cobol85.MergeStatementContext)
	out = &pb.MergeStatement{}
	if ctx.FileName() != nil {
		out.FileName = FileName(ctx.FileName())
	}
	for _, iv := range ctx.AllMergeOnKeyClause() {
		clause := &pb.MergeStatement_OnKeyClause{}
		cv := iv.(*cobol85.MergeOnKeyClauseContext)
		for _, ivv := range cv.AllQualifiedDataName() {
			clause.Keys = append(clause.Keys, QualifiedDataName(ivv))
		}
		out.OnKeyClauses = append(out.OnKeyClauses, clause)
	}
	for _, iv := range ctx.AllMergeUsing() {
		clause := &pb.MergeStatement_Using{}
		cv := iv.(*cobol85.MergeUsingContext)
		for _, ivv := range cv.AllFileName() {
			clause.FileNames = append(clause.FileNames, FileName(ivv))
		}
		out.Usings = append(out.Usings, clause)
	}

	for _, iv := range ctx.AllMergeGivingPhrase() {
		clause := &pb.MergeStatement_GivingPhrase{}
		cv := iv.(*cobol85.MergeGivingPhraseContext)
		for _, ivv := range cv.AllMergeGiving() {
			cvv := ivv.(*cobol85.MergeGivingContext)
			giving := &pb.MergeStatement_Giving{}
			if cvv.FileName() != nil {
				giving.FileName = FileName(cvv.FileName())
			}
			if cvv.LOCK() != nil {
				giving.Type = pb.MergeStatement_Giving_LOCK
			} else if cvv.SAVE() != nil {
				giving.Type = pb.MergeStatement_Giving_SAVE
			} else if cvv.NO() != nil || cvv.REWIND() != nil {
				giving.Type = pb.MergeStatement_Giving_NO_REWIND
			} else if cvv.CRUNCH() != nil {
				giving.Type = pb.MergeStatement_Giving_CRUNCH
			} else if cvv.RELEASE() != nil {
				giving.Type = pb.MergeStatement_Giving_RELEASE
			} else if cvv.WITH() != nil || cvv.REMOVE() != nil || cvv.CRUNCH() != nil {
				giving.Type = pb.MergeStatement_Giving_WITH_REMOVE_CRUNCH
			}
			clause.Givings = append(clause.Givings, giving)
		}
		out.GivingPhrases = append(out.GivingPhrases, clause)
	}

	if ictx := ctx.MergeCollatingSequencePhrase(); ictx != nil {
		cctx := ictx.(*cobol85.MergeCollatingSequencePhraseContext)
		phrase := &pb.MergeStatement_CollatingSequencePhrase{}
		for _, iv := range cctx.AllAlphabetName() {
			phrase.AlphabetNames = append(phrase.AlphabetNames, AlphabetName(iv))
		}
		if iv := cctx.MergeCollatingAlphanumeric(); iv != nil {
			alphanumeric := &pb.MergeStatement_CollatingAlphanumeric{}
			cv := iv.(*cobol85.MergeCollatingAlphanumericContext)
			if cv.AlphabetName() != nil {
				alphanumeric.AlphabetName = AlphabetName(cv.AlphabetName())
			}
			phrase.CollatingAlphanumeric = alphanumeric
		}
		if iv := cctx.MergeCollatingNational(); iv != nil {
			alphanumeric := &pb.MergeStatement_CollatingNational{}
			cv := iv.(*cobol85.MergeCollatingNationalContext)
			if cv.AlphabetName() != nil {
				alphanumeric.AlphabetName = AlphabetName(cv.AlphabetName())
			}
			phrase.CollatingNational = alphanumeric
		}
		out.CollatingSequencePhrase = phrase
	}
	if ictx := ctx.MergeOutputProcedurePhrase(); ictx != nil {
		cctx := ictx.(*cobol85.MergeOutputProcedurePhraseContext)
		phrase := &pb.MergeStatement_OutputProcedurePhrase{}
		if cctx.ProcedureName() != nil {
			phrase.ProcedureName = ProcedureName(cctx.ProcedureName())
		}
		if iv := cctx.MergeOutputThrough(); iv != nil {
			cv := iv.(*cobol85.MergeOutputThroughContext)
			phrase.OutputThrough = &pb.MergeStatement_OutputThrough{}
			if cv.ProcedureName() != nil {
				phrase.OutputThrough.ProcedureName = ProcedureName(cv.ProcedureName())
			}
		}
		out.OutputProcedurePhrase = phrase
	}

	return
}
func MoveStatement(in cobol85.IMoveStatementContext) (out *pb.MoveStatement) {
	ctx := in.(*cobol85.MoveStatementContext)
	out = &pb.MoveStatement{}
	if ictx := ctx.MoveCorrespondingToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveCorrespondingToStatementContext)
		to := &pb.MoveCorrespondingToStatement{}
		for _, v := range cctx.AllIdentifier() {
			to.To = append(to.To, Identifier(v))
		}
		if iv := cctx.MoveCorrespondingToSendingArea(); iv != nil {
			cv := iv.(*cobol85.MoveCorrespondingToSendingAreaContext)
			if cv.Identifier() != nil {
				to.SendingArea = Identifier(cv.Identifier())
			}
		}
		out.OneOf = &pb.MoveStatement_MoveCorrespondingTo{
			MoveCorrespondingTo: to,
		}
	} else if ictx := ctx.MoveToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveToStatementContext)
		to := &pb.MoveToStatement{}
		for _, v := range cctx.AllIdentifier() {
			to.To = append(to.To, Identifier(v))
		}
		if iv := cctx.MoveToSendingArea(); iv != nil {
			cv := iv.(*cobol85.MoveToSendingAreaContext)
			if cv.Literal() != nil {
				to.SendingArea = &pb.MoveToStatement_Literal{
					Literal: Literal(cv.Literal()),
				}
			} else if cv.Identifier() != nil {
				to.SendingArea = &pb.MoveToStatement_Identifier{
					Identifier: Identifier(cv.Identifier()),
				}
			}
		}
		out.OneOf = &pb.MoveStatement_MoveTo{
			MoveTo: to,
		}
	}
	return
}
func MultiplyStatement(in cobol85.IMultiplyStatementContext) (out *pb.MultiplyStatement) {
	ctx := in.(*cobol85.MultiplyStatementContext)
	out = &pb.MultiplyStatement{}
	if ctx.OnSizeErrorPhrase() != nil {
		out.OnSizeErrorPhrase = OnSizeErrorPhrase(ctx.OnSizeErrorPhrase())
	}
	if ctx.NotOnSizeErrorPhrase() != nil {
		out.NotOnSizeErrorPhrase = NotOnSizeErrorPhrase(ctx.NotOnSizeErrorPhrase())
	}

	if ctx.Identifier() != nil {
		out.Value = &pb.MultiplyStatement_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	} else if ctx.Literal() != nil {
		out.Value = &pb.MultiplyStatement_Literal{
			Literal: Literal(ctx.Literal()),
		}
	}

	if ictx := ctx.MultiplyGiving(); ictx != nil {
		giving := &pb.MultiplyStatement_Giving{}
		cctx := ictx.(*cobol85.MultiplyGivingContext)
		if iv := cctx.MultiplyGivingOperand(); iv != nil {
			cv := iv.(*cobol85.MultiplyGivingOperandContext)
			givingOperand := &pb.MultiplyStatement_GivingOperand{}
			if cv.Identifier() != nil {
				givingOperand.OneOf = &pb.MultiplyStatement_GivingOperand_Identifier{
					Identifier: Identifier(cv.Identifier()),
				}
			} else if cv.Literal() != nil {
				givingOperand.OneOf = &pb.MultiplyStatement_GivingOperand_Literal{
					Literal: Literal(cv.Literal()),
				}
			}
			giving.GivingOperand = givingOperand
		}
		for _, iv := range cctx.AllMultiplyGivingResult() {
			cv := iv.(*cobol85.MultiplyGivingResultContext)
			if cv.Identifier() != nil {
				givingResult := &pb.MultiplyStatement_GivingResult{
					Identifier: Identifier(cv.Identifier()),
				}
				giving.GivingResult = append(giving.GivingResult, givingResult)
			}

		}
		out.By = &pb.MultiplyStatement_Giving_{
			Giving: giving,
		}
	} else if ictx := ctx.MultiplyRegular(); ictx != nil {
		regular := &pb.MultiplyStatement_Regular{}
		cctx := ictx.(*cobol85.MultiplyRegularContext)
		for _, iv := range cctx.AllMultiplyRegularOperand() {
			cv := iv.(*cobol85.MultiplyRegularOperandContext)
			if cv.Identifier() != nil {
				regular.RegularOperands = append(regular.RegularOperands, Identifier(cv.Identifier()))
			}
		}
		out.By = &pb.MultiplyStatement_Regular_{
			Regular: regular,
		}
	}
	return
}
func NextSentenceStatement(in cobol85.INextSentenceStatementContext) (out *pb.NextSentenceStatement) {
	out = &pb.NextSentenceStatement{}
	return
}
func OpenStatement(in cobol85.IOpenStatementContext) (out *pb.OpenStatement) {
	ctx := in.(*cobol85.OpenStatementContext)
	out = &pb.OpenStatement{}

	for _, iv := range ctx.AllOpenIOStatement() {
		statement := &pb.OpenStatement_IOStatement{}
		cv := iv.(*cobol85.OpenIOStatementContext)
		for _, ivv := range cv.AllFileName() {
			statement.FileNames = append(statement.FileNames, FileName(ivv))
		}
		out.OneOf = &pb.OpenStatement_IoStatement{
			IoStatement: statement,
		}
	}

	for _, iv := range ctx.AllOpenInputStatement() {
		statement := &pb.OpenStatement_InputStatement{}
		cv := iv.(*cobol85.OpenInputStatementContext)
		for _, ivv := range cv.AllOpenInput() {
			cvv := ivv.(*cobol85.OpenInputContext)
			statement.Inputs = append(statement.Inputs, &pb.OpenStatement_Input{
				FileName: FileName(cvv.FileName()),
			})
		}
		out.OneOf = &pb.OpenStatement_InputStatement_{
			InputStatement: statement,
		}
	}

	for _, iv := range ctx.AllOpenOutputStatement() {
		statement := &pb.OpenStatement_OutputStatement{}
		cv := iv.(*cobol85.OpenOutputStatementContext)
		for _, ivv := range cv.AllOpenOutput() {
			cvv := ivv.(*cobol85.OpenOutputContext)
			statement.Outputs = append(statement.Outputs, &pb.OpenStatement_Output{
				FileName: FileName(cvv.FileName()),
			})
		}
		out.OneOf = &pb.OpenStatement_OutputStatement_{
			OutputStatement: statement,
		}
	}

	for _, iv := range ctx.AllOpenExtendStatement() {
		statement := &pb.OpenStatement_ExtendStatement{}
		cv := iv.(*cobol85.OpenExtendStatementContext)
		for _, ivv := range cv.AllFileName() {
			statement.FileNames = append(statement.FileNames, FileName(ivv))
		}
		out.OneOf = &pb.OpenStatement_ExtendStatement_{
			ExtendStatement: statement,
		}
	}
	return
}

func VaryingPhrase(in cobol85.IPerformVaryingPhraseContext) (out *pb.PerformStatement_VaryingPhrase) {
	out = &pb.PerformStatement_VaryingPhrase{}
	ctx := in.(*cobol85.PerformVaryingPhraseContext)
	if ctx.Identifier() != nil {
		out.Value = &pb.PerformStatement_VaryingPhrase_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	} else if ctx.Literal() != nil {
		out.Value = &pb.PerformStatement_VaryingPhrase_Literal{
			Literal: Literal(ctx.Literal()),
		}
	}

	if iv := ctx.PerformFrom(); iv != nil {
		cv := iv.(*cobol85.PerformFromContext)
		out.From = &pb.PerformStatement_From{}
		if cv.Identifier() != nil {
			out.From.OneOf = &pb.PerformStatement_From_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			out.From.OneOf = &pb.PerformStatement_From_Literal{
				Literal: Literal(cv.Literal()),
			}
		} else if cv.ArithmeticExpression() != nil {
			out.From.OneOf = &pb.PerformStatement_From_ArithmeticExpression{
				ArithmeticExpression: ArithmeticExpression(cv.ArithmeticExpression()),
			}
		}
	}

	if iv := ctx.PerformBy(); iv != nil {
		cv := iv.(*cobol85.PerformByContext)
		out.By = &pb.PerformStatement_By{}
		if cv.Identifier() != nil {
			out.By.OneOf = &pb.PerformStatement_By_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			out.By.OneOf = &pb.PerformStatement_By_Literal{
				Literal: Literal(cv.Literal()),
			}
		} else if cv.ArithmeticExpression() != nil {
			out.By.OneOf = &pb.PerformStatement_By_ArithmeticExpression{
				ArithmeticExpression: ArithmeticExpression(cv.ArithmeticExpression()),
			}
		}
	}

	if iv := ctx.PerformUntil(); iv != nil {
		cv := iv.(*cobol85.PerformUntilContext)
		out.Until = &pb.PerformStatement_Until{}
		if cv.Condition() != nil {
			out.Until.Condition = Condition(cv.Condition())
		}

		if cv.PerformTestClause() != nil {
			out.Until.TestClause = &pb.PerformStatement_TestClause{}
		}
	}
	return
}

func PerformVaryingClause(in cobol85.IPerformVaryingClauseContext) (out *pb.PerformStatement_VaryingClause) {
	out = &pb.PerformStatement_VaryingClause{}
	ctx := in.(*cobol85.PerformVaryingClauseContext)
	for _, iv := range ctx.AllPerformAfter() {
		cv := iv.(*cobol85.PerformAfterContext)
		after := &pb.PerformStatement_After{}
		if ivv := cv.PerformVaryingPhrase(); ivv != nil {
			after.VaryingPhrase = VaryingPhrase(ivv)
		}
		out.After = append(out.After, after)
	}

	if ctx.PerformVaryingPhrase() != nil {
		out.VaryingPhrase = VaryingPhrase(ctx.PerformVaryingPhrase())
	}
	return
}

func PerformType(in cobol85.IPerformTypeContext) (out *pb.PerformStatement_Type) {
	out = &pb.PerformStatement_Type{}
	cv := in.(*cobol85.PerformTypeContext)
	if ivv := cv.PerformTimes(); ivv != nil {
		cvv := ivv.(*cobol85.PerformTimesContext)
		if cvv.Identifier() != nil {
			out.OneOf = &pb.PerformStatement_Type_Times{
				Times: &pb.PerformStatement_Times{
					OneOf: &pb.PerformStatement_Times_Identifier{
						Identifier: Identifier(cvv.Identifier()),
					},
				},
			}
		} else if cvv.IntegerLiteral() != nil {
			out.OneOf = &pb.PerformStatement_Type_Times{
				Times: &pb.PerformStatement_Times{
					OneOf: &pb.PerformStatement_Times_IntegerLiteral{
						IntegerLiteral: IntegerLiteral(cvv.IntegerLiteral()),
					},
				},
			}
		}
	} else if ivv := cv.PerformUntil(); ivv != nil {
		until := &pb.PerformStatement_Until{}
		cvv := ivv.(*cobol85.PerformUntilContext)
		if cvv.Condition() != nil {
			until.Condition = Condition(cvv.Condition())
		}

		if cvv.PerformTestClause() != nil {
			until.TestClause = &pb.PerformStatement_TestClause{}
		}
		out.OneOf = &pb.PerformStatement_Type_Until{
			Until: until,
		}
	} else if ivv := cv.PerformVarying(); ivv != nil {
		varying := &pb.PerformStatement_Varying{}
		cvv := ivv.(*cobol85.PerformVaryingContext)
		if cvv.PerformTestClause() != nil {
			varying.TestClause = &pb.PerformStatement_TestClause{}
		}
		if cvv.PerformVaryingClause() != nil {
			varying.VaryingClause = PerformVaryingClause(cvv.PerformVaryingClause())
		}
		out.OneOf = &pb.PerformStatement_Type_Varying{
			Varying: varying,
		}
	}
	return
}
func PerformStatement(in cobol85.IPerformStatementContext) (out *pb.PerformStatement) {
	ctx := in.(*cobol85.PerformStatementContext)
	out = &pb.PerformStatement{}

	if ictx := ctx.PerformInlineStatement(); ictx != nil {
		statement := &pb.PerformStatement_InlineStatement{}
		cctx := ictx.(*cobol85.PerformInlineStatementContext)
		for _, iv := range cctx.AllStatement() {
			statement.Statements = append(statement.Statements, Statement(iv))
		}
		if cctx.PerformType() != nil {
			statement.Type = PerformType(cctx.PerformType())
		}
		out.OneOf = &pb.PerformStatement_InlineStatement_{
			InlineStatement: statement,
		}
	} else if ictx := ctx.PerformProcedureStatement(); ictx != nil {
		cctx := ictx.(*cobol85.PerformProcedureStatementContext)
		statement := &pb.PerformStatement_ProcedureStatement{}
		names := cctx.AllProcedureName()
		if len(names) >= 1 {
			statement.ProcedureName = ProcedureName(names[0])
		}
		if len(names) >= 2 {
			statement.Through = ProcedureName(names[1])
		}
		if cctx.PerformType() != nil {
			statement.Type = PerformType(cctx.PerformType())
		}
		out.OneOf = &pb.PerformStatement_ProcedureStatement_{
			ProcedureStatement: statement,
		}
	}
	return
}

func PurgeStatement(in cobol85.IPurgeStatementContext) (out *pb.PurgeStatement) {
	ctx := in.(*cobol85.PurgeStatementContext)
	out = &pb.PurgeStatement{}
	for _, v := range ctx.AllCdName() {
		out.CdNames = append(out.CdNames, CdName(v))
	}
	return
}

func ReadStatement(in cobol85.IReadStatementContext) (out *pb.ReadStatement) {
	ctx := in.(*cobol85.ReadStatementContext)
	out = &pb.ReadStatement{}
	if ctx.InvalidKeyPhrase() != nil {
		out.InvalidKeyPhrase = InvalidKeyPhrase(ctx.InvalidKeyPhrase())
	}
	if ctx.NotInvalidKeyPhrase() != nil {
		out.NotInvalidKeyPhrase = NotInvalidKeyPhrase(ctx.NotInvalidKeyPhrase())
	}
	if ctx.AtEndPhrase() != nil {
		out.AtEndPhrase = AtEndPhrase(ctx.AtEndPhrase())
	}
	if ctx.NotAtEndPhrase() != nil {
		out.NotAtEndPhrase = NotAtEndPhrase(ctx.NotAtEndPhrase())
	}

	if ctx.FileName() != nil {
		out.FileName = FileName(ctx.FileName())
	}

	if iv := ctx.ReadInto(); iv != nil {
		cv := iv.(*cobol85.ReadIntoContext)
		out.Into = &pb.ReadStatement_Into{
			Identifier: Identifier(cv.Identifier()),
		}
	}

	if iv := ctx.ReadWith(); iv != nil {
		out.With = &pb.ReadStatement_With{}
	}

	if iv := ctx.ReadKey(); iv != nil {
		cv := iv.(*cobol85.ReadKeyContext)
		out.Key = &pb.ReadStatement_Key{
			QualifiedDataName: QualifiedDataName(cv.QualifiedDataName()),
		}
	}
	return
}
func ReceiveStatement(in cobol85.IReceiveStatementContext) (out *pb.ReceiveStatement) {
	ctx := in.(*cobol85.ReceiveStatementContext)
	out = &pb.ReceiveStatement{}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}

	if ictx := ctx.ReceiveIntoStatement(); ictx != nil {
		statement := &pb.ReceiveStatement_IntoStatement{}
		cctx := ictx.(*cobol85.ReceiveIntoStatementContext)
		if cctx.CdName() != nil {
			statement.CdName = CdName(cctx.CdName())
		}
		if cctx.Identifier() != nil {
			statement.Identifier = Identifier(cctx.Identifier())
		}

		if iv := cctx.ReceiveNoData(); iv != nil {
			cv := iv.(*cobol85.ReceiveNoDataContext)
			statement.NoData = &pb.ReceiveStatement_NoData{}
			for _, ivv := range cv.AllStatement() {
				statement.NoData.Statements = append(statement.NoData.Statements, Statement(ivv))
			}
		}

		if iv := cctx.ReceiveWithData(); iv != nil {
			cv := iv.(*cobol85.ReceiveWithDataContext)
			statement.WithData = &pb.ReceiveStatement_WithData{}
			for _, ivv := range cv.AllStatement() {
				statement.WithData.Statements = append(statement.WithData.Statements, Statement(ivv))
			}
		}
	} else if ictx := ctx.ReceiveFromStatement(); ictx != nil {
		statement := &pb.ReceiveStatement_FromStatement{}
		cctx := ictx.(*cobol85.ReceiveFromStatementContext)
		if cctx.DataName() != nil {
			statement.DataName = DataName(cctx.DataName())
		}

		if iv := cctx.ReceiveFrom(); iv != nil {
			cv := iv.(*cobol85.ReceiveFromContext)
			statement.From = &pb.ReceiveStatement_From{}
			if cv.DataName() != nil {
				statement.From.DataName = DataName(cv.DataName())
			}
		}
		statement.Before = &pb.ReceiveStatement_Before{}
		for _, iv := range cctx.AllReceiveBefore() {
			cv := iv.(*cobol85.ReceiveBeforeContext)
			if cv.Identifier() != nil {
				statement.Before.OneOf = &pb.ReceiveStatement_Before_Identifier{
					Identifier: Identifier(cv.Identifier()),
				}
			} else if cv.NumericLiteral() != nil {
				statement.Before.OneOf = &pb.ReceiveStatement_Before_NumericLiteral{
					NumericLiteral: NumericLiteral(cv.NumericLiteral()),
				}
			}
		}

		if len(cctx.AllReceiveWith()) != 0 {
			statement.With = &pb.ReceiveStatement_With{}
		}

		statement.Thread = &pb.ReceiveStatement_Thread{}
		for _, iv := range cctx.AllReceiveThread() {
			cv := iv.(*cobol85.ReceiveThreadContext)
			if cv.DataName() != nil {
				statement.Thread.DataName = DataName(cv.DataName())
			}
		}

		statement.Size = &pb.ReceiveStatement_Size{}
		for _, iv := range cctx.AllReceiveSize() {
			cv := iv.(*cobol85.ReceiveSizeContext)
			if cv.Identifier() != nil {
				statement.Size.OneOf = &pb.ReceiveStatement_Size_Identifier{
					Identifier: Identifier(cv.Identifier()),
				}
			} else if cv.NumericLiteral() != nil {
				statement.Size.OneOf = &pb.ReceiveStatement_Size_NumericLiteral{
					NumericLiteral: NumericLiteral(cv.NumericLiteral()),
				}
			}
		}

		statement.Status = &pb.ReceiveStatement_Status{}
		for _, iv := range cctx.AllReceiveStatus() {
			cv := iv.(*cobol85.ReceiveStatusContext)
			if cv.Identifier() != nil {
				statement.Status.Identifier = Identifier(cv.Identifier())
			}
		}

	}
	return
}
func ReleaseStatement(in cobol85.IReleaseStatementContext) (out *pb.ReleaseStatement) {
	ctx := in.(*cobol85.ReleaseStatementContext)
	out = &pb.ReleaseStatement{}

	if ctx.RecordName() != nil {
		out.RecordName = RecordName(ctx.RecordName())
	}

	if ctx.QualifiedDataName() != nil {
		out.From = &pb.ReleaseStatement_From{
			QualifiedDataName: QualifiedDataName(ctx.QualifiedDataName()),
		}
	}
	return
}
func ReturnStatement(in cobol85.IReturnStatementContext) (out *pb.ReturnStatement) {
	ctx := in.(*cobol85.ReturnStatementContext)
	out = &pb.ReturnStatement{}
	if ctx.AtEndPhrase() != nil {
		out.AtEndPhrase = AtEndPhrase(ctx.AtEndPhrase())
	}
	if ctx.NotAtEndPhrase() != nil {
		out.NotAtEndPhrase = NotAtEndPhrase(ctx.NotAtEndPhrase())
	}

	if ctx.FileName() != nil {
		out.FileName = FileName(ctx.FileName())
	}

	if ictx := ctx.ReturnInto(); ictx != nil {
		cctx := ictx.(*cobol85.ReturnIntoContext)
		out.Into = &pb.ReturnStatement_Into{
			QualifiedDataName: QualifiedDataName(cctx.QualifiedDataName()),
		}
	}
	return
}
func RewriteStatement(in cobol85.IRewriteStatementContext) (out *pb.RewriteStatement) {
	ctx := in.(*cobol85.RewriteStatementContext)
	out = &pb.RewriteStatement{}
	if ctx.InvalidKeyPhrase() != nil {
		out.InvalidKeyPhrase = InvalidKeyPhrase(ctx.InvalidKeyPhrase())
	}
	if ctx.NotInvalidKeyPhrase() != nil {
		out.NotInvalidKeyPhrase = NotInvalidKeyPhrase(ctx.NotInvalidKeyPhrase())
	}

	if ctx.RecordName() != nil {
		out.RecordName = RecordName(ctx.RecordName())
	}

	if iv := ctx.RewriteFrom(); iv != nil {
		cv := iv.(*cobol85.RewriteFromContext)
		out.From = &pb.RewriteStatement_From{
			Identifier: Identifier(cv.Identifier()),
		}
	}
	return
}
func SearchStatement(in cobol85.ISearchStatementContext) (out *pb.SearchStatement) {
	ctx := in.(*cobol85.SearchStatementContext)
	out = &pb.SearchStatement{}

	if ctx.QualifiedDataName() != nil {
		out.QualifiedDataName = QualifiedDataName(ctx.QualifiedDataName())
	}

	if iv := ctx.SearchVarying(); iv != nil {
		cv := iv.(*cobol85.SearchVaryingContext)
		out.Varying = &pb.SearchStatement_Varying{
			QualifiedDataName: QualifiedDataName(cv.QualifiedDataName()),
		}
	}

	for _, iv := range ctx.AllSearchWhen() {
		cv := iv.(*cobol85.SearchWhenContext)
		when := &pb.SearchStatement_When{}
		if cv.Condition() != nil {
			when.Condition = Condition(cv.Condition())
		}
		if cv.NEXT() != nil {
			when.NextSentence = true
		}
		for _, ivv := range cv.AllStatement() {
			when.Statements = append(when.Statements, Statement(ivv))
		}
		out.When = append(out.When, when)
	}
	return
}
func SendStatement(in cobol85.ISendStatementContext) (out *pb.SendStatement) {
	ctx := in.(*cobol85.SendStatementContext)
	out = &pb.SendStatement{}
	if ctx.OnExceptionClause() != nil {
		out.OnExceptionClause = OnExceptionClause(ctx.OnExceptionClause())
	}
	if ctx.NotOnExceptionClause() != nil {
		out.NotOnExceptionClause = NotOnExceptionClause(ctx.NotOnExceptionClause())
	}

	if iv := ctx.SendStatementSync(); iv != nil {
		statement := &pb.SendStatement_SyncStatement{}
		cv := iv.(*cobol85.SendStatementSyncContext)

		if cv.Identifier() != nil {
			statement.OneOf = &pb.SendStatement_SyncStatement_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			statement.OneOf = &pb.SendStatement_SyncStatement_Literal{
				Literal: Literal(cv.Literal()),
			}
		}

		if ivv := cv.SendFromPhrase(); ivv != nil {
			statement.FromPhrase = &pb.SendStatement_FromPhrase{}
			cvv := ivv.(*cobol85.SendFromPhraseContext)
			if cvv.Identifier() != nil {
				statement.FromPhrase.Identifier = Identifier(cvv.Identifier())
			}
		}

		if ivv := cv.SendWithPhrase(); ivv != nil {
			statement.WithPhrase = &pb.SendStatement_WithPhrase{}
			cvv := ivv.(*cobol85.SendWithPhraseContext)
			if cvv.Identifier() != nil {
				statement.FromPhrase.Identifier = Identifier(cvv.Identifier())
			}
		}

		if ivv := cv.SendReplacingPhrase(); ivv != nil {
			statement.ReplacingPhrase = &pb.SendStatement_ReplacingPhrase{}
		}

		if ivv := cv.SendAdvancingPhrase(); ivv != nil {
			cvv := ivv.(*cobol85.SendAdvancingPhraseContext)
			statement.AdvancingPhrase = &pb.SendStatement_AdvancingPhrase{}

			if ivvv := cvv.SendAdvancingPage(); ivvv != nil {
				statement.AdvancingPhrase.OneOf = &pb.SendStatement_AdvancingPhrase_AdvancingPage{
					AdvancingPage: &pb.SendStatement_AdvancingPage{},
				}
			}

			if ivvv := cvv.SendAdvancingLines(); ivvv != nil {
				lines := &pb.SendStatement_AdvancingLines{}
				cvvv := ivvv.(*cobol85.SendAdvancingLinesContext)

				if cvvv.Identifier() != nil {
					lines.OneOf = &pb.SendStatement_AdvancingLines_Identifier{
						Identifier: Identifier(cvvv.Identifier()),
					}
				} else if cvvv.Literal() != nil {
					lines.OneOf = &pb.SendStatement_AdvancingLines_Literal{
						Literal: Literal(cvvv.Literal()),
					}
				}
				statement.AdvancingPhrase.OneOf = &pb.SendStatement_AdvancingPhrase_AdvancingLines{
					AdvancingLines: lines,
				}
			}

			if ivvv := cvv.SendAdvancingMnemonic(); ivvv != nil {
				cvvv := ivvv.(*cobol85.SendAdvancingMnemonicContext)
				name := &pb.SendStatement_AdvancingMnemonic{}
				if cvvv.MnemonicName() != nil {
					name.MnemonicName = MnemonicName(cvvv.MnemonicName())
				}
				statement.AdvancingPhrase.OneOf = &pb.SendStatement_AdvancingPhrase_AdvancingMnemonic{
					AdvancingMnemonic: name,
				}
			}
		}

		out.OneOf = &pb.SendStatement_SyncStatement_{
			SyncStatement: statement,
		}
	} else if iv := ctx.SendStatementAsync(); iv != nil {
		cv := iv.(*cobol85.SendStatementAsyncContext)
		statement := &pb.SendStatement_AsyncStatement{}
		if cv.Identifier() != nil {
			statement.Identifier = Identifier(cv.Identifier())
		}
	}
	return
}
func SetStatement(in cobol85.ISetStatementContext) (out *pb.SetStatement) {
	ctx := in.(*cobol85.SetStatementContext)
	out = &pb.SetStatement{}

	if iv := ctx.SetUpDownByStatement(); iv != nil {
		cv := iv.(*cobol85.SetUpDownByStatementContext)
		statement := &pb.SetStatement_UpDownByStatement{}

		for _, ivv := range cv.AllSetTo() {
			cvv := ivv.(*cobol85.SetToContext)
			to := &pb.SetStatement_To{}
			if cvv.Identifier() != nil {
				to.Identifier = Identifier(cvv.Identifier())
			}
			statement.Tos = append(statement.Tos, to)
		}

		if ivv := cv.SetByValue(); ivv != nil {
			value := &pb.SetStatement_ByValue{}
			cvv := ivv.(*cobol85.SetByValueContext)
			if cvv.Literal() != nil {
				value.OneOf = &pb.SetStatement_ByValue_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				value.OneOf = &pb.SetStatement_ByValue_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			statement.ByValue = value
		}
		out.UpDownByStatement = statement
	}

	for _, iv := range ctx.AllSetToStatement() {
		to := &pb.SetStatement_ToStatement{}
		cv := iv.(*cobol85.SetToStatementContext)
		for _, ivv := range cv.AllSetToValue() {
			value := &pb.SetStatement_ToValue{}
			cvv := ivv.(*cobol85.SetToValueContext)
			if cvv.ON() != nil {
				value.OneOf = &pb.SetStatement_ToValue_ON{
					ON: true,
				}
			} else if cvv.OFF() != nil {
				value.OneOf = &pb.SetStatement_ToValue_OFF{
					OFF: true,
				}
			} else if cvv.ENTRY() != nil {
				entry := &pb.SetStatement_Entry{}
				if cvv.Literal() != nil {
					entry.OneOf = &pb.SetStatement_Entry_Literal{
						Literal: Literal(cvv.Literal()),
					}
				} else if cvv.Identifier() != nil {
					entry.OneOf = &pb.SetStatement_Entry_Identifier{
						Identifier: Identifier(cvv.Identifier()),
					}
				}
				value.OneOf = &pb.SetStatement_ToValue_Entry{
					Entry: entry,
				}
			} else if cvv.Literal() != nil {
				value.OneOf = &pb.SetStatement_ToValue_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				value.OneOf = &pb.SetStatement_ToValue_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			to.ToValues = append(to.ToValues, value)
		}
		for _, ivv := range cv.AllSetTo() {
			value := &pb.SetStatement_To{}
			cvv := ivv.(*cobol85.SetToContext)
			if cvv.Identifier() != nil {
				value.Identifier = Identifier(cvv.Identifier())
			}
			to.Tos = append(to.Tos, value)
		}
		out.ToStatements = append(out.ToStatements, to)
	}
	return
}
func SortStatement(in cobol85.ISortStatementContext) (out *pb.SortStatement) {
	ctx := in.(*cobol85.SortStatementContext)
	out = &pb.SortStatement{}

	if ctx.FileName() != nil {
		out.FileName = FileName(ctx.FileName())
	}

	for _, iv := range ctx.AllSortOnKeyClause() {
		cv := iv.(*cobol85.SortOnKeyClauseContext)
		clause := &pb.SortStatement_OnKeyClause{}
		for _, ivv := range cv.AllQualifiedDataName() {
			clause.Keys = append(clause.Keys, QualifiedDataName(ivv))
		}
		out.OnKeyClauses = append(out.OnKeyClauses, clause)
	}

	if iv := ctx.SortDuplicatesPhrase(); iv != nil {
		out.DuplicatesPhrase = &pb.SortStatement_DuplicatesPhrase{}
	}

	if iv := ctx.SortCollatingSequencePhrase(); iv != nil {
		cv := iv.(*cobol85.SortCollatingSequencePhraseContext)
		phrase := &pb.SortStatement_CollatingSequencePhrase{}
		for _, ivv := range cv.AllAlphabetName() {
			phrase.AlphabetNames = append(phrase.AlphabetNames, AlphabetName(ivv))
		}
		if ivv := cv.SortCollatingAlphanumeric(); ivv != nil {
			phrase.CollatingAlphanumeric = &pb.SortStatement_CollatingAlphanumeric{}
			cvv := ivv.(*cobol85.SortCollatingAlphanumericContext)
			if cvv.AlphabetName() != nil {
				phrase.CollatingAlphanumeric.AlphabetName = AlphabetName(cvv.AlphabetName())
			}
		}
		if ivv := cv.SortCollatingNational(); ivv != nil {
			phrase.CollatingNational = &pb.SortStatement_CollatingNational{}
			cvv := ivv.(*cobol85.SortCollatingNationalContext)
			if cvv.AlphabetName() != nil {
				phrase.CollatingNational.AlphabetName = AlphabetName(cvv.AlphabetName())
			}
		}
		out.CollatingSequencePhrase = phrase
	}

	if iv := ctx.SortInputProcedurePhrase(); iv != nil {
		cv := iv.(*cobol85.SortInputProcedurePhraseContext)
		phrase := &pb.SortStatement_InputProcedurePhrase{}
		if cv.ProcedureName() != nil {
			phrase.ProcedureName = ProcedureName(cv.ProcedureName())
		}

		if ivv := cv.SortInputThrough(); ivv != nil {
			cvv := ivv.(*cobol85.SortInputThroughContext)
			statement := &pb.SortStatement_InputThrough{}
			if cvv.ProcedureName() != nil {
				statement.ProcedureName = ProcedureName(cvv.ProcedureName())
			}
			phrase.InputThrough = statement
		}
		out.InputProcedurePhrase = phrase
	}

	for _, iv := range ctx.AllSortUsing() {
		cv := iv.(*cobol85.SortUsingContext)
		using := &pb.SortStatement_Using{}
		for _, ivv := range cv.AllFileName() {
			using.FileNames = append(using.FileNames, FileName(ivv))
		}
		out.Usings = append(out.Usings, using)
	}

	if iv := ctx.SortOutputProcedurePhrase(); iv != nil {
		cv := iv.(*cobol85.SortOutputProcedurePhraseContext)
		phrase := &pb.SortStatement_OutputProcedurePhrase{}
		if cv.ProcedureName() != nil {
			phrase.ProcedureName = ProcedureName(cv.ProcedureName())
		}

		if ivv := cv.SortOutputThrough(); ivv != nil {
			cvv := ivv.(*cobol85.SortOutputThroughContext)
			phrase.OutputThrough = &pb.SortStatement_OutputThrough{
				ProcedureName: ProcedureName(cvv.ProcedureName()),
			}
		}
		out.OutputProcedurePhrase = phrase
	}

	for _, iv := range ctx.AllSortGivingPhrase() {
		cv := iv.(*cobol85.SortGivingPhraseContext)
		phrase := &pb.SortStatement_GivingPhrase{}
		for _, ivv := range cv.AllSortGiving() {
			cvv := ivv.(*cobol85.SortGivingContext)
			gp := &pb.SortStatement_Giving{}
			if cvv.FileName() != nil {
				gp.FileName = FileName(cvv.FileName())
			}
			if cvv.LOCK() != nil {
				gp.Type = pb.SortStatement_Giving_LOCK
			} else if cvv.SAVE() != nil {
				gp.Type = pb.SortStatement_Giving_SAVE
			} else if cvv.NO() != nil && cvv.REWIND() != nil {
				gp.Type = pb.SortStatement_Giving_NO_REWIND
			} else if cvv.CRUNCH() != nil {
				gp.Type = pb.SortStatement_Giving_CRUNCH
			} else if cvv.RELEASE() != nil {
				gp.Type = pb.SortStatement_Giving_RELEASE
			} else if cvv.WITH() != nil && cvv.REMOVE() != nil && cvv.CRUNCH() != nil {
				gp.Type = pb.SortStatement_Giving_WITH_REMOVE_CRUNCH
			}
			phrase.Givings = append(phrase.Givings, gp)
		}
		out.GivingPhrases = append(out.GivingPhrases, phrase)
	}
	return
}
func StartStatement(in cobol85.IStartStatementContext) (out *pb.StartStatement) {
	ctx := in.(*cobol85.StartStatementContext)
	out = &pb.StartStatement{}
	if ctx.InvalidKeyPhrase() != nil {
		out.InvalidKeyPhrase = InvalidKeyPhrase(ctx.InvalidKeyPhrase())
	}
	if ctx.NotInvalidKeyPhrase() != nil {
		out.NotInvalidKeyPhrase = NotInvalidKeyPhrase(ctx.NotInvalidKeyPhrase())
	}

	if ctx.FileName() != nil {
		out.FileName = FileName(ctx.FileName())
	}

	if iv := ctx.StartKey(); iv != nil {
		cv := iv.(*cobol85.StartKeyContext)
		key := &pb.StartStatement_Key{}
		if cv.QualifiedDataName() != nil {
			key.QualifiedDataName = QualifiedDataName(cv.QualifiedDataName())
		}
		if cv.EQUALCHAR() != nil {
			key.Type = pb.StartStatement_EQUAL
		} else if cv.GREATER() != nil {
			key.Type = pb.StartStatement_GREATER
			if cv.EQUAL() != nil {
				key.Type = pb.StartStatement_GREATER_OR_EQUAL
			}
		} else if cv.EQUAL() != nil {
			key.Type = pb.StartStatement_EQUAL
		} else if cv.MORETHANCHAR() != nil {
			key.Type = pb.StartStatement_GREATER
		} else if cv.MORETHANOREQUAL() != nil {
			key.Type = pb.StartStatement_GREATER_OR_EQUAL
		} else if cv.NOT() != nil {
			if cv.LESS() != nil {
				key.Type = pb.StartStatement_GREATER_OR_EQUAL
			} else if cv.LESSTHANCHAR() != nil {
				key.Type = pb.StartStatement_GREATER_OR_EQUAL
			}
		}
		out.Key = key
	}
	return
}
func StopStatement(in cobol85.IStopStatementContext) (out *pb.StopStatement) {
	ctx := in.(*cobol85.StopStatementContext)
	out = &pb.StopStatement{}

	if iv := ctx.StopStatementGiving(); iv != nil {
		cv := iv.(*cobol85.StopStatementGivingContext)
		giving := &pb.StopStatement_Giving{}
		if cv.IntegerLiteral() != nil {
			giving.OneOf = &pb.StopStatement_Giving_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(cv.IntegerLiteral()),
			}
		} else if cv.Identifier() != nil {
			giving.OneOf = &pb.StopStatement_Giving_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		}
		out.OneOf = &pb.StopStatement_Giving_{
			Giving: giving,
		}
	} else if ctx.Literal() != nil {
		out.OneOf = &pb.StopStatement_Literal{
			Literal: Literal(ctx.Literal()),
		}
	} else if ctx.RUN() != nil {
		out.OneOf = &pb.StopStatement_RUN{
			RUN: true,
		}
	}
	return
}
func StringStatement(in cobol85.IStringStatementContext) (out *pb.StringStatement) {
	ctx := in.(*cobol85.StringStatementContext)
	out = &pb.StringStatement{}
	if ctx.OnOverflowPhrase() != nil {
		out.OnOverflowPhrase = OnOverflowPhrase(ctx.OnOverflowPhrase())
	}
	if ctx.NotOnOverflowPhrase() != nil {
		out.NotOnOverflowPhrase = NotOnOverflowPhrase(ctx.NotOnOverflowPhrase())
	}
	for _, iv := range ctx.AllStringSendingPhrase() {
		cv := iv.(*cobol85.StringSendingPhraseContext)
		phrase := &pb.StringStatement_SendingPhrase{}
		for _, ivv := range cv.AllStringSending() {
			cvv := ivv.(*cobol85.StringSendingContext)
			sending := &pb.StringStatement_Sending{}
			if cvv.Identifier() != nil {
				sending.OneOf = &pb.StringStatement_Sending_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			} else if cvv.Literal() != nil {
				sending.OneOf = &pb.StringStatement_Sending_Literal{
					Literal: Literal(cvv.Literal()),
				}
			}
			phrase.Sendings = append(phrase.Sendings, sending)
		}

		if ivv := cv.StringDelimitedByPhrase(); ivv != nil {
			or := &pb.StringStatement_DelimitedByPhrase{}
			cvv := ivv.(*cobol85.StringDelimitedByPhraseContext)
			if cvv.Literal() != nil {
				or.OneOf = &pb.StringStatement_DelimitedByPhrase_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				or.OneOf = &pb.StringStatement_DelimitedByPhrase_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			phrase.DelimitedByPhrase = or
		}
		if ivv := cv.StringForPhrase(); ivv != nil {
			or := &pb.StringStatement_ForPhrase{}
			cvv := ivv.(*cobol85.StringForPhraseContext)
			if cvv.Literal() != nil {
				or.OneOf = &pb.StringStatement_ForPhrase_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				or.OneOf = &pb.StringStatement_ForPhrase_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			phrase.ForPhrase = or
		}
		out.SendingPhrases = append(out.SendingPhrases, phrase)
	}

	if iv := ctx.StringIntoPhrase(); iv != nil {
		cv := iv.(*cobol85.StringIntoPhraseContext)
		phrase := &pb.StringStatement_IntoPhrase{}
		if cv.Identifier() != nil {
			phrase.Identifier = Identifier(cv.Identifier())
		}
		out.IntoPhrase = phrase
	}

	if iv := ctx.StringWithPointerPhrase(); iv != nil {
		cv := iv.(*cobol85.StringWithPointerPhraseContext)
		phrase := &pb.StringStatement_WithPointerPhrase{}
		if cv.QualifiedDataName() != nil {
			phrase.QualifiedDataName = QualifiedDataName(cv.QualifiedDataName())
		}
		out.WithPointerPhrase = phrase
	}
	return
}
func SubtractStatement(in cobol85.ISubtractStatementContext) (out *pb.SubtractStatement) {
	ctx := in.(*cobol85.SubtractStatementContext)
	out = &pb.SubtractStatement{}
	if ctx.OnSizeErrorPhrase() != nil {
		out.OnSizeErrorPhrase = OnSizeErrorPhrase(ctx.OnSizeErrorPhrase())
	}
	if ctx.NotOnSizeErrorPhrase() != nil {
		out.NotOnSizeErrorPhrase = NotOnSizeErrorPhrase(ctx.NotOnSizeErrorPhrase())
	}

	if iv := ctx.SubtractFromStatement(); iv != nil {
		cv := iv.(*cobol85.SubtractFromStatementContext)
		statement := &pb.SubtractStatement_FromStatement{}
		for _, ivv := range cv.AllSubtractMinuend() {
			minuend := &pb.SubtractStatement_Minuend{}
			cvv := ivv.(*cobol85.SubtractMinuendContext)
			if cvv.Identifier() != nil {
				minuend.Identifier = Identifier(cvv.Identifier())
			}
			statement.Minuends = append(statement.Minuends, minuend)
		}
		for _, ivv := range cv.AllSubtractSubtrahend() {
			subtrahend := &pb.SubtractStatement_Subtrahend{}
			cvv := ivv.(*cobol85.SubtractSubtrahendContext)
			if cvv.Identifier() != nil {
				subtrahend.OneOf = &pb.SubtractStatement_Subtrahend_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			} else if cvv.Literal() != nil {
				subtrahend.OneOf = &pb.SubtractStatement_Subtrahend_Literal{
					Literal: Literal(cvv.Literal()),
				}
			}
			statement.Subtrahends = append(statement.Subtrahends, subtrahend)
		}
		out.OneOf = &pb.SubtractStatement_FromStatement_{
			FromStatement: statement,
		}
	} else if iv := ctx.SubtractFromGivingStatement(); iv != nil {
		cv := iv.(*cobol85.SubtractFromGivingStatementContext)
		statement := &pb.SubtractStatement_FromGivingStatement{}
		for _, ivv := range cv.AllSubtractGiving() {
			giving := &pb.SubtractStatement_Giving{}
			cvv := ivv.(*cobol85.SubtractGivingContext)
			if cvv.Identifier() != nil {
				giving.Identifier = Identifier(cvv.Identifier())
			}
			statement.Givings = append(statement.Givings, giving)
		}
		for _, ivv := range cv.AllSubtractSubtrahend() {
			subtrahend := &pb.SubtractStatement_Subtrahend{}
			cvv := ivv.(*cobol85.SubtractSubtrahendContext)
			if cvv.Identifier() != nil {
				subtrahend.OneOf = &pb.SubtractStatement_Subtrahend_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			} else if cvv.Literal() != nil {
				subtrahend.OneOf = &pb.SubtractStatement_Subtrahend_Literal{
					Literal: Literal(cvv.Literal()),
				}
			}
			statement.Subtrahends = append(statement.Subtrahends, subtrahend)
		}

		if ivv := cv.SubtractMinuendGiving(); ivv != nil {
			cvv := ivv.(*cobol85.SubtractMinuendGivingContext)
			giving := &pb.SubtractStatement_MinuendGiving{}
			if cvv.Literal() != nil {
				giving.OneOf = &pb.SubtractStatement_MinuendGiving_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				giving.OneOf = &pb.SubtractStatement_MinuendGiving_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			statement.MinuendGiving = giving
		}
		out.OneOf = &pb.SubtractStatement_FromGivingStatement_{
			FromGivingStatement: statement,
		}
	} else if iv := ctx.SubtractCorrespondingStatement(); iv != nil {
		cv := iv.(*cobol85.SubtractCorrespondingStatementContext)
		statement := &pb.SubtractStatement_CorrespondingStatement{}
		if cv.QualifiedDataName() != nil {
			statement.QualifiedDataName = QualifiedDataName(cv.QualifiedDataName())
		}
		if ivv := cv.SubtractMinuendCorresponding(); ivv != nil {
			cvv := ivv.(*cobol85.SubtractMinuendCorrespondingContext)
			minuend := &pb.SubtractStatement_MinuendCorresponding{}
			if cvv.QualifiedDataName() != nil {
				minuend.QualifiedDataName = QualifiedDataName(cvv.QualifiedDataName())
			}
			statement.MinuendCorresponding = minuend
		}
		out.OneOf = &pb.SubtractStatement_CorrespondingStatement_{
			CorrespondingStatement: statement,
		}
	}
	return
}
func TerminateStatement(in cobol85.ITerminateStatementContext) (out *pb.TerminateStatement) {
	ctx := in.(*cobol85.TerminateStatementContext)
	out = &pb.TerminateStatement{}

	if ctx.ReportName() != nil {
		out.ReportName = ReportName(ctx.ReportName())
	}
	return
}
func UnstringStatement(in cobol85.IUnstringStatementContext) (out *pb.UnstringStatement) {
	ctx := in.(*cobol85.UnstringStatementContext)
	out = &pb.UnstringStatement{}
	if ctx.OnOverflowPhrase() != nil {
		out.OnOverflowPhrase = OnOverflowPhrase(ctx.OnOverflowPhrase())
	}
	if ctx.NotOnOverflowPhrase() != nil {
		out.NotOnOverflowPhrase = NotOnOverflowPhrase(ctx.NotOnOverflowPhrase())
	}

	if iv := ctx.UnstringSendingPhrase(); iv != nil {
		cv := iv.(*cobol85.UnstringSendingPhraseContext)
		phrase := &pb.UnstringStatement_SendingPhrase{}
		if cv.Identifier() != nil {
			phrase.Identifier = Identifier(cv.Identifier())
		}

		if ivv := cv.UnstringDelimitedByPhrase(); ivv != nil {
			or := &pb.UnstringStatement_DelimitedByPhrase{}
			cvv := ivv.(*cobol85.UnstringDelimitedByPhraseContext)
			if cvv.Literal() != nil {
				or.OneOf = &pb.UnstringStatement_DelimitedByPhrase_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				or.OneOf = &pb.UnstringStatement_DelimitedByPhrase_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			phrase.DelimitedByPhrase = or
		}
		out.SendingPhrase = phrase
	}
	if iv := ctx.UnstringIntoPhrase(); iv != nil {
		cv := iv.(*cobol85.UnstringIntoPhraseContext)
		phrase := &pb.UnstringStatement_IntoPhrase{}
		for _, ivv := range cv.AllUnstringInto() {
			cvv := ivv.(*cobol85.UnstringIntoContext)
			into := &pb.UnstringStatement_Into{}
			if cvv.Identifier() != nil {
				into.Identifier = Identifier(cvv.Identifier())
			}

			if ivvv := cvv.UnstringDelimiterIn(); ivvv != nil {
				cvvv := ivvv.(*cobol85.UnstringDelimiterInContext)
				in := &pb.UnstringStatement_DelimiterIn{}
				if cvvv.Identifier() != nil {
					in.Identifier = Identifier(cvvv.Identifier())
				}
				into.DelimiterIn = in
			}

			if ivvv := cvv.UnstringCountIn(); ivvv != nil {
				cvvv := ivvv.(*cobol85.UnstringCountInContext)
				in := &pb.UnstringStatement_CountIn{}
				if cvvv.Identifier() != nil {
					in.Identifier = Identifier(cvvv.Identifier())
				}
				into.CountIn = in
			}
			phrase.Into = append(phrase.Into, into)
		}
		out.IntoPhrase = phrase
	}

	if iv := ctx.UnstringWithPointerPhrase(); iv != nil {
		cv := iv.(*cobol85.UnstringWithPointerPhraseContext)
		phrase := &pb.UnstringStatement_WithPointerPhrase{}
		if cv.QualifiedDataName() != nil {
			phrase.QualifiedDataName = QualifiedDataName(cv.QualifiedDataName())
		}
		out.WithPointerPhrase = phrase
	}

	if iv := ctx.UnstringTallyingPhrase(); iv != nil {
		cv := iv.(*cobol85.UnstringTallyingPhraseContext)
		phrase := &pb.UnstringStatement_TallyingPhrase{}
		if cv.QualifiedDataName() != nil {
			phrase.QualifiedDataName = QualifiedDataName(cv.QualifiedDataName())
		}
		out.TallyingPhrase = phrase
	}
	return
}
func WriteStatement(in cobol85.IWriteStatementContext) (out *pb.WriteStatement) {
	ctx := in.(*cobol85.WriteStatementContext)
	out = &pb.WriteStatement{}
	if ctx.WriteAtEndOfPagePhrase() != nil {
		out.AtEndOfPagePhrase = AtEndOfPagePhrase(ctx.WriteAtEndOfPagePhrase())
	}
	if ctx.WriteNotAtEndOfPagePhrase() != nil {
		out.NotAtEndOfPagePhrase = NotAtEndOfPagePhrase(ctx.WriteNotAtEndOfPagePhrase())
	}
	if ctx.InvalidKeyPhrase() != nil {
		out.InvalidKeyPhrase = InvalidKeyPhrase(ctx.InvalidKeyPhrase())
	}
	if ctx.NotInvalidKeyPhrase() != nil {
		out.NotInvalidKeyPhrase = NotInvalidKeyPhrase(ctx.NotInvalidKeyPhrase())
	}

	if ctx.RecordName() != nil {
		out.RecordName = RecordName(ctx.RecordName())
	}

	if iv := ctx.WriteFromPhrase(); iv != nil {
		cv := iv.(*cobol85.WriteFromPhraseContext)
		phrase := &pb.WriteStatement_FromPhrase{}
		if cv.Identifier() != nil {
			phrase.OneOf = &pb.WriteStatement_FromPhrase_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			phrase.OneOf = &pb.WriteStatement_FromPhrase_Literal{
				Literal: Literal(cv.Literal()),
			}
		}
		out.FromPhrase = phrase
	}
	if iv := ctx.WriteAdvancingPhrase(); iv != nil {
		cv := iv.(*cobol85.WriteAdvancingPhraseContext)
		phrase := &pb.WriteStatement_AdvancingPhrase{}

		if ivv := cv.WriteAdvancingPage(); ivv != nil {
			phrase.OneOf = &pb.WriteStatement_AdvancingPhrase_AdvancingPage{
				AdvancingPage: &pb.WriteStatement_AdvancingPage{},
			}
		} else if ivv := cv.WriteAdvancingLines(); ivv != nil {
			cvv := ivv.(*cobol85.WriteAdvancingLinesContext)
			adv := &pb.WriteStatement_AdvancingLines{}
			if cvv.Literal() != nil {
				adv.OneOf = &pb.WriteStatement_AdvancingLines_Literal{
					Literal: Literal(cvv.Literal()),
				}
			} else if cvv.Identifier() != nil {
				adv.OneOf = &pb.WriteStatement_AdvancingLines_Identifier{
					Identifier: Identifier(cvv.Identifier()),
				}
			}
			phrase.OneOf = &pb.WriteStatement_AdvancingPhrase_AdvancingLines{
				AdvancingLines: adv,
			}
		} else if ivv := cv.WriteAdvancingMnemonic(); ivv != nil {
			cvv := ivv.(*cobol85.WriteAdvancingMnemonicContext)
			adv := &pb.WriteStatement_AdvancingMnemonic{}
			if cvv.MnemonicName() != nil {
				adv.MnemonicName = MnemonicName(cvv.MnemonicName())
			}
			phrase.OneOf = &pb.WriteStatement_AdvancingPhrase_AdvancingMnemonic{
				AdvancingMnemonic: adv,
			}
		}
		out.AdvancingPhrase = phrase
	}
	return
}

func MoveToStatement(in cobol85.IMoveStatementContext) (out *pb.MoveStatement) {
	ctx := in.(*cobol85.MoveStatementContext)
	out = &pb.MoveStatement{}
	if ictx := ctx.MoveToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveToStatementContext)
		mts := &pb.MoveToStatement{}
		if imts := cctx.MoveToSendingArea(); imts != nil {
			cmts := imts.(*cobol85.MoveToSendingAreaContext)
			if cmts.Identifier() != nil {
				mts.SendingArea = &pb.MoveToStatement_Identifier{
					Identifier: Identifier(cmts.Identifier()),
				}
			} else if cmts.Literal() != nil {
				mts.SendingArea = &pb.MoveToStatement_Literal{
					Literal: Literal(cmts.Literal()),
				}
			}
		}
		for _, i := range cctx.AllIdentifier() {
			mts.To = append(mts.To, Identifier(i))
		}
		out.OneOf = &pb.MoveStatement_MoveTo{
			MoveTo: mts,
		}
	} else if ictx := ctx.MoveCorrespondingToStatement(); ictx != nil {
		cctx := ictx.(*cobol85.MoveCorrespondingToStatementContext)
		mts := &pb.MoveCorrespondingToStatement{}
		if imts := cctx.MoveCorrespondingToSendingArea(); imts != nil {
			cmts := imts.(*cobol85.MoveCorrespondingToSendingAreaContext)
			if cmts.Identifier() != nil {
				mts.SendingArea = Identifier(cmts.Identifier())
			}
		}
		for _, i := range cctx.AllIdentifier() {
			mts.To = append(mts.To, Identifier(i))
		}
		out.OneOf = &pb.MoveStatement_MoveCorrespondingTo{
			MoveCorrespondingTo: mts,
		}
	}
	return
}

func OnExceptionClause(in cobol85.IOnExceptionClauseContext) (out *pb.OnExceptionClause) {
	ctx := in.(*cobol85.OnExceptionClauseContext)
	out = &pb.OnExceptionClause{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotOnExceptionClause(in cobol85.INotOnExceptionClauseContext) (out *pb.NotOnExceptionClause) {
	ctx := in.(*cobol85.NotOnExceptionClauseContext)
	out = &pb.NotOnExceptionClause{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func OnSizeErrorPhrase(in cobol85.IOnSizeErrorPhraseContext) (out *pb.OnSizeErrorPhrase) {
	ctx := in.(*cobol85.OnSizeErrorPhraseContext)
	out = &pb.OnSizeErrorPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotOnSizeErrorPhrase(in cobol85.INotOnSizeErrorPhraseContext) (out *pb.NotOnSizeErrorPhrase) {
	ctx := in.(*cobol85.NotOnSizeErrorPhraseContext)
	out = &pb.NotOnSizeErrorPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func OnOverflowPhrase(in cobol85.IOnOverflowPhraseContext) (out *pb.OnOverflowPhrase) {
	ctx := in.(*cobol85.OnOverflowPhraseContext)
	out = &pb.OnOverflowPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotOnOverflowPhrase(in cobol85.INotOnOverflowPhraseContext) (out *pb.NotOnOverflowPhrase) {
	ctx := in.(*cobol85.NotOnOverflowPhraseContext)
	out = &pb.NotOnOverflowPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func InvalidKeyPhrase(in cobol85.IInvalidKeyPhraseContext) (out *pb.InvalidKeyPhrase) {
	ctx := in.(*cobol85.InvalidKeyPhraseContext)
	out = &pb.InvalidKeyPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotInvalidKeyPhrase(in cobol85.INotInvalidKeyPhraseContext) (out *pb.NotInvalidKeyPhrase) {
	ctx := in.(*cobol85.NotInvalidKeyPhraseContext)
	out = &pb.NotInvalidKeyPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func AtEndPhrase(in cobol85.IAtEndPhraseContext) (out *pb.AtEndPhrase) {
	ctx := in.(*cobol85.AtEndPhraseContext)
	out = &pb.AtEndPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotAtEndPhrase(in cobol85.INotAtEndPhraseContext) (out *pb.NotAtEndPhrase) {
	ctx := in.(*cobol85.NotAtEndPhraseContext)
	out = &pb.NotAtEndPhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func AtEndOfPagePhrase(in cobol85.IWriteAtEndOfPagePhraseContext) (out *pb.AtEndOfPagePhrase) {
	ctx := in.(*cobol85.WriteAtEndOfPagePhraseContext)
	out = &pb.AtEndOfPagePhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}

func NotAtEndOfPagePhrase(in cobol85.IWriteNotAtEndOfPagePhraseContext) (out *pb.NotAtEndOfPagePhrase) {
	ctx := in.(*cobol85.WriteNotAtEndOfPagePhraseContext)
	out = &pb.NotAtEndOfPagePhrase{
		Statements: []*pb.Statement{},
	}
	for _, v := range ctx.AllStatement() {
		out.Statements = append(out.Statements, Statement(v))
	}
	return
}
