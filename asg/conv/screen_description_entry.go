package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func ScreenDescriptionEntry(in cobol85.IScreenDescriptionEntryContext) (out *pb.ScreenDescriptionEntry) {
	ctx := in.(*cobol85.ScreenDescriptionEntryContext)
	out = &pb.ScreenDescriptionEntry{}
	if ctx.FILLER() != nil {
		out.IntegerLiteral = &pb.ScreenDescriptionEntry_Filler{}
	} else if ctx.ScreenName() != nil {
		out.IntegerLiteral = &pb.ScreenDescriptionEntry_ScreenName{
			ScreenName: ScreenName(ctx.ScreenName()),
		}
	}
	for _, iv := range ctx.AllScreenDescriptionBlankClause() {
		out.BlankClause = &pb.BlankClause{}
		cv := iv.(*cobol85.ScreenDescriptionBlankClauseContext)
		if cv.SCREEN() != nil {
			out.BlankClause.Type = pb.BlankClause_SCREEN
		} else if cv.LINE() != nil {
			out.BlankClause.Type = pb.BlankClause_LINE
		}
	}

	for _, iv := range ctx.AllScreenDescriptionBellClause() {
		out.BellClause = &pb.BellClause{}
		cv := iv.(*cobol85.ScreenDescriptionBellClauseContext)
		if cv.BELL() != nil {
			out.BellClause.Type = pb.BellClause_BELL
		} else if cv.BEEP() != nil {
			out.BellClause.Type = pb.BellClause_BEEP
		}
	}

	for range ctx.AllScreenDescriptionBlinkClause() {
		out.BlinkClause = &pb.BlinkClause{}
	}

	for _, iv := range ctx.AllScreenDescriptionEraseClause() {
		out.EraseClause = &pb.EraseClause{}
		cv := iv.(*cobol85.ScreenDescriptionEraseClauseContext)
		if cv.EOL() != nil {
			out.EraseClause.Type = pb.EraseClause_EOL
		} else if cv.EOS() != nil {
			out.EraseClause.Type = pb.EraseClause_EOS
		}
	}

	for _, iv := range ctx.AllScreenDescriptionLightClause() {
		out.LightClause = &pb.LightClause{}
		cv := iv.(*cobol85.ScreenDescriptionLightClauseContext)
		if cv.HIGHLIGHT() != nil {
			out.LightClause.Type = pb.LightClause_HIGHLIGHT
		} else if cv.LOWLIGHT() != nil {
			out.LightClause.Type = pb.LightClause_LOWLIGHT
		}
	}

	for _, iv := range ctx.AllScreenDescriptionGridClause() {
		out.GridClause = &pb.GridClause{}
		cv := iv.(*cobol85.ScreenDescriptionGridClauseContext)
		if cv.GRID() != nil {
			out.GridClause.Type = pb.GridClause_GRID
		} else if cv.LEFTLINE() != nil {
			out.GridClause.Type = pb.GridClause_LEFTLINE
		} else if cv.OVERLINE() != nil {
			out.GridClause.Type = pb.GridClause_OVERLINE
		}
	}

	for range ctx.AllScreenDescriptionReverseVideoClause() {
		out.ReverseVideoClause = &pb.ReverseVideoClause{}
	}

	for range ctx.AllScreenDescriptionUnderlineClause() {
		out.UnderlineClause = &pb.UnderlineClause{}
	}

	for _, iv := range ctx.AllScreenDescriptionSizeClause() {
		out.SizeClause = &pb.SizeClause{}
		cv := iv.(*cobol85.ScreenDescriptionSizeClauseContext)
		if cv.Identifier() != nil {
			out.SizeClause.Size = &pb.SizeClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.IntegerLiteral() != nil {
			out.SizeClause.Size = &pb.SizeClause_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(cv.IntegerLiteral()),
			}
		}
	}

	for _, iv := range ctx.AllScreenDescriptionLineClause() {
		out.LineClause = &pb.LineClause{}
		cv := iv.(*cobol85.ScreenDescriptionLineClauseContext)
		if cv.PLUS() != nil {
			out.LineClause.Type = pb.LineClause_PLUS
		} else if cv.PLUSCHAR() != nil {
			out.LineClause.Type = pb.LineClause_PLUS
		} else if cv.MINUSCHAR() != nil {
			out.LineClause.Type = pb.LineClause_MINUS
		}

		if cv.Identifier() != nil {
			out.LineClause.Line = &pb.LineClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.IntegerLiteral() != nil {
			out.LineClause.Line = &pb.LineClause_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(cv.IntegerLiteral()),
			}
		}
	}

	for _, iv := range ctx.AllScreenDescriptionColumnClause() {
		out.ColumnClause = &pb.ColumnClause{}
		cv := iv.(*cobol85.ScreenDescriptionColumnClauseContext)
		if cv.PLUS() != nil {
			out.ColumnClause.Type = pb.ColumnClause_PLUS
		} else if cv.PLUSCHAR() != nil {
			out.ColumnClause.Type = pb.ColumnClause_PLUS
		} else if cv.MINUSCHAR() != nil {
			out.ColumnClause.Type = pb.ColumnClause_MINUS
		}

		if cv.Identifier() != nil {
			out.ColumnClause.Column = &pb.ColumnClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.IntegerLiteral() != nil {
			out.ColumnClause.Column = &pb.ColumnClause_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(cv.IntegerLiteral()),
			}
		}
	}

	for _, iv := range ctx.AllScreenDescriptionForegroundColorClause() {
		out.ForegroundColorClause = &pb.ForegroundColorClause{}
		cv := iv.(*cobol85.ScreenDescriptionForegroundColorClauseContext)
		if cv.Identifier() != nil {
			out.ForegroundColorClause.ForegroundColor = &pb.ForegroundColorClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.IntegerLiteral() != nil {
			out.ForegroundColorClause.ForegroundColor = &pb.ForegroundColorClause_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(cv.IntegerLiteral()),
			}
		}
	}

	for _, iv := range ctx.AllScreenDescriptionBackgroundColorClause() {
		out.BackgroundColorClause = &pb.BackgroundColorClause{}
		cv := iv.(*cobol85.ScreenDescriptionBackgroundColorClauseContext)
		if cv.Identifier() != nil {
			out.BackgroundColorClause.BackgroundColor = &pb.BackgroundColorClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.IntegerLiteral() != nil {
			out.BackgroundColorClause.BackgroundColor = &pb.BackgroundColorClause_IntegerLiteral{
				IntegerLiteral: IntegerLiteral(cv.IntegerLiteral()),
			}
		}
	}

	for _, iv := range ctx.AllScreenDescriptionControlClause() {
		out.ControlClause = &pb.ControlClause{}
		cv := iv.(*cobol85.ScreenDescriptionControlClauseContext)
		if cv.Identifier() != nil {
			out.ControlClause.Control = Identifier(cv.Identifier())
		}
	}

	for _, iv := range ctx.AllScreenDescriptionValueClause() {
		out.ValueClause = &pb.ValueClause{}
		cv := iv.(*cobol85.ScreenDescriptionValueClauseContext)
		if cv.Literal() != nil {
			out.ValueClause.Value = Literal(cv.Literal())
		}
	}

	for _, iv := range ctx.AllScreenDescriptionPictureClause() {
		out.PictureClause = &pb.PictureClause{}
		cv := iv.(*cobol85.ScreenDescriptionPictureClauseContext)
		out.PictureClause.PictureString = PictureString(cv.PictureString())
	}

	for _, iv := range ctx.AllScreenDescriptionFromClause() {
		out.FromClause = &pb.FromClause{}
		cv := iv.(*cobol85.ScreenDescriptionFromClauseContext)
		if cv.Identifier() != nil {
			out.FromClause.From = &pb.FromClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			out.FromClause.From = &pb.FromClause_Literal{
				Literal: Literal(cv.Literal()),
			}
		}

		if it := cv.ScreenDescriptionToClause(); it != nil {
			ct := it.(*cobol85.ScreenDescriptionToClauseContext)
			out.FromClause.To = Identifier(ct.Identifier())
		}
	}

	for _, iv := range ctx.AllScreenDescriptionUsingClause() {
		out.UsingClause = &pb.UsingClause{}
		cv := iv.(*cobol85.ScreenDescriptionUsingClauseContext)
		if cv.Identifier() != nil {
			out.UsingClause.Using = Identifier(cv.Identifier())
		}
	}

	for range ctx.AllScreenDescriptionUsageClause() {
		out.UsageClause = &pb.UsageClause{}
	}

	for range ctx.AllScreenDescriptionBlankWhenZeroClause() {
		out.BlankWhenZeroClause = &pb.BlankWhenZeroClause{}
	}

	for _, iv := range ctx.AllScreenDescriptionJustifiedClause() {
		cv := iv.(*cobol85.ScreenDescriptionJustifiedClauseContext)
		out.JustifiedClause = &pb.JustifiedClause{
			Right: cv.RIGHT() != nil,
		}
	}

	for _, iv := range ctx.AllScreenDescriptionSignClause() {
		out.SignClause = &pb.SignClause{}
		cv := iv.(*cobol85.ScreenDescriptionSignClauseContext)
		if cv.LEADING() != nil {
			out.SignClause.Type = pb.SignClause_LEADING
		} else if cv.TRAILING() != nil {
			out.SignClause.Type = pb.SignClause_TRAILING
		}
	}

	for _, iv := range ctx.AllScreenDescriptionAutoClause() {
		out.AutoClause = &pb.AutoClause{}
		cv := iv.(*cobol85.ScreenDescriptionAutoClauseContext)
		if cv.AUTO() != nil {
			out.AutoClause.Type = pb.AutoClause_AUTO
		} else if cv.AUTO_SKIP() != nil {
			out.AutoClause.Type = pb.AutoClause_AUTO_SKIP
		}
	}

	for _, iv := range ctx.AllScreenDescriptionSecureClause() {
		out.SecureClause = &pb.SecureClause{}
		cv := iv.(*cobol85.ScreenDescriptionSecureClauseContext)
		if cv.SECURE() != nil {
			out.SecureClause.Type = pb.SecureClause_SECURE
		} else if cv.NO_ECHO() != nil {
			out.SecureClause.Type = pb.SecureClause_NO_ECHO
		}
	}

	for _, iv := range ctx.AllScreenDescriptionRequiredClause() {
		out.RequiredClause = &pb.RequiredClause{}
		cv := iv.(*cobol85.ScreenDescriptionRequiredClauseContext)
		if cv.REQUIRED() != nil {
			out.RequiredClause.Type = pb.RequiredClause_REQUIRED
		} else if cv.EMPTY_CHECK() != nil {
			out.RequiredClause.Type = pb.RequiredClause_EMPTY_CHECK
		}
	}

	for _, iv := range ctx.AllScreenDescriptionPromptClause() {
		out.PromptClause = &pb.PromptClause{}
		cv := iv.(*cobol85.ScreenDescriptionPromptClauseContext)
		if cv.Identifier() != nil {
			out.PromptClause.Prompt = &pb.PromptClause_Identifier{
				Identifier: Identifier(cv.Identifier()),
			}
		} else if cv.Literal() != nil {
			out.PromptClause.Prompt = &pb.PromptClause_Literal{
				Literal: Literal(cv.Literal()),
			}
		}

		if ic := cv.ScreenDescriptionPromptOccursClause(); ic != nil {
			cc := ic.(*cobol85.ScreenDescriptionPromptOccursClauseContext)
			out.PromptClause.Times = IntegerLiteral(cc.IntegerLiteral())
		}
	}

	for _, iv := range ctx.AllScreenDescriptionFullClause() {
		out.FullClause = &pb.FullClause{}
		cv := iv.(*cobol85.ScreenDescriptionFullClauseContext)
		if cv.FULL() != nil {
			out.FullClause.Type = pb.FullClause_FULL
		} else if cv.LENGTH_CHECK() != nil {
			out.FullClause.Type = pb.FullClause_LENGTH_CHECK
		}
	}

	for range ctx.AllScreenDescriptionZeroFillClause() {
		out.ZeroFillClause = &pb.ZeroFillClause{}
	}
	return
}
