package environment

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type SpecialNamesParagraphVisitor struct {
	cobol85.BaseCobol85Visitor
	paragraph *pb.SpecialNamesParagraph
}

func NewSpecialNamesParagraphVisitor(paragraph *pb.SpecialNamesParagraph) *SpecialNamesParagraphVisitor {
	return &SpecialNamesParagraphVisitor{
		paragraph: paragraph,
	}
}

func (v *SpecialNamesParagraphVisitor) VisitSpecialNamesParagraph(ctx *cobol85.SpecialNamesParagraphContext) any {
	for _, iv := range ctx.AllSpecialNameClause() {
		vctx := iv.(*cobol85.SpecialNameClauseContext)
		if ictx := vctx.ChannelClause(); ictx != nil {
			cctx := ictx.(*cobol85.ChannelClauseContext)
			v.paragraph.ChannelClause = &pb.ChannelClause{
				Channel:      conv.IntegerLiteral(cctx.IntegerLiteral()),
				MnemonicName: conv.MnemonicName(cctx.MnemonicName()),
			}
		} else if ictx := vctx.OdtClause(); ictx != nil {
			cctx := ictx.(*cobol85.OdtClauseContext)
			v.paragraph.OdtClause = &pb.OdtClause{
				MnemonicName: conv.MnemonicName(cctx.MnemonicName()),
			}
		} else if ictx := vctx.AlphabetClause(); ictx != nil {
			cctx := ictx.(*cobol85.AlphabetClauseContext)
			alphabetClause := &pb.AlphabetClause{}
			if if1 := cctx.AlphabetClauseFormat1(); if1 != nil {
				f1ctx := if1.(*cobol85.AlphabetClauseFormat1Context)
				alphabetClause.AlphabetName = conv.AlphabetName(f1ctx.AlphabetName())
				alphanumeric := &pb.AlphabetClause_Alphanumeric{}
				switch {
				case f1ctx.EBCDIC() != nil:
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_Type_{
						Type: pb.AlphabetClause_Alphanumeric_EBCDIC,
					}
				case f1ctx.ASCII() != nil:
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_Type_{
						Type: pb.AlphabetClause_Alphanumeric_ASCII,
					}
				case f1ctx.STANDARD_1() != nil:
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_Type_{
						Type: pb.AlphabetClause_Alphanumeric_STANDARD_1,
					}
				case f1ctx.STANDARD_2() != nil:
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_Type_{
						Type: pb.AlphabetClause_Alphanumeric_STANDARD_2,
					}
				case f1ctx.NATIVE() != nil:
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_Type_{
						Type: pb.AlphabetClause_Alphanumeric_NATIVE,
					}
				case f1ctx.CobolWord() != nil:
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_CobolWord{
						CobolWord: conv.CobolWord(f1ctx.CobolWord()),
					}
				case len(f1ctx.AllAlphabetLiterals()) != 0:
					als := &pb.AlphabetClause_AlphabetLiterals{}
					for _, ial := range f1ctx.AllAlphabetLiterals() {
						alCtx := ial.(*cobol85.AlphabetLiteralsContext)
						al := &pb.AlphabetClause_AlphabetLiteral{
							Literal: conv.Literal(alCtx.Literal()),
						}
						if iThrough := alCtx.AlphabetThrough(); iThrough != nil {
							throughCtx := iThrough.(*cobol85.AlphabetThroughContext)
							al.Through = conv.Literal(throughCtx.Literal())
						}
						for _, iv := range alCtx.AllAlphabetAlso() {
							vCtx := iv.(*cobol85.AlphabetAlsoContext)
							alsos := []*pb.Literal{}
							for _, iAlso := range vCtx.AllLiteral() {
								alsos = append(alsos, conv.Literal(iAlso))
							}
							al.Alsos = append(al.Alsos, &pb.AlphabetClause_Also{
								Alsos: alsos,
							})
						}
						als.Values = append(als.Values, al)
					}
					alphanumeric.OneOf = &pb.AlphabetClause_Alphanumeric_Values{
						Values: als,
					}
				}
				alphabetClause.OneOf = &pb.AlphabetClause_Alphanumeric_{
					Alphanumeric: alphanumeric,
				}
			} else if if2 := cctx.AlphabetClauseFormat2(); if2 != nil {
				f2ctx := if2.(*cobol85.AlphabetClauseFormat2Context)
				var typ pb.AlphabetClause_National_Type
				var literal *pb.Literal
				switch {
				case f2ctx.NATIVE() != nil:
					typ = pb.AlphabetClause_National_NATIVE
				case f2ctx.CCSVERSION() != nil:
					typ = pb.AlphabetClause_National_CCSVERSION
					literal = conv.Literal(f2ctx.Literal())
				}
				alphabetClause.OneOf = &pb.AlphabetClause_National_{
					National: &pb.AlphabetClause_National{
						Type:       typ,
						CcsVersion: literal,
					},
				}
			}
		} else if ictx := vctx.ClassClause(); ictx != nil {
			cctx := ictx.(*cobol85.ClassClauseContext)
			var typ pb.ClassClause_Type
			switch {
			case cctx.ALPHANUMERIC() != nil:
				typ = pb.ClassClause_ALPHANUMERIC
			case cctx.NATIONAL() != nil:
				typ = pb.ClassClause_NATIONAL
			}
			ccts := []*pb.ClassClauseThrough{}
			for _, icts := range cctx.AllClassClauseThrough() {
				ctsCtx := icts.(*cobol85.ClassClauseThroughContext)
				cct := &pb.ClassClauseThrough{}
				if iFrom := ctsCtx.ClassClauseFrom(); iFrom != nil {
					fromCtx := iFrom.(*cobol85.ClassClauseFromContext)
					if fromCtx.Identifier() != nil {
						cct.From = &pb.ClassClauseThrough_IdentifierFrom{
							IdentifierFrom: conv.Identifier(fromCtx.Identifier()),
						}
					} else if fromCtx.Literal() != nil {
						cct.From = &pb.ClassClauseThrough_LiteralFrom{
							LiteralFrom: conv.Literal(fromCtx.Literal()),
						}
					}
				}

				if iTo := ctsCtx.ClassClauseTo(); iTo != nil {
					toCtx := iTo.(*cobol85.ClassClauseToContext)
					if toCtx.Identifier() != nil {
						cct.To = &pb.ClassClauseThrough_IdentifierTo{
							IdentifierTo: conv.Identifier(toCtx.Identifier()),
						}
					} else if toCtx.Literal() != nil {
						cct.To = &pb.ClassClauseThrough_LiteralTo{
							LiteralTo: conv.Literal(toCtx.Literal()),
						}
					}
				}
				ccts = append(ccts, cct)
			}
			v.paragraph.ClassClause = &pb.ClassClause{
				ClassName:     conv.ClassName(cctx.ClassName()),
				Type:          typ,
				ClassThroughs: ccts,
			}
		} else if ictx := vctx.CurrencySignClause(); ictx != nil {
			cctx := ictx.(*cobol85.CurrencySignClauseContext)
			clause := &pb.CurrencySignClause{}
			for _, iLiteral := range cctx.AllLiteral() {
				if clause.CurrencyLiteral == nil {
					clause.CurrencyLiteral = conv.Literal(iLiteral)
				} else {
					clause.PictureSymbolLiteral = conv.Literal(iLiteral)
				}
			}
			v.paragraph.CurrencySignClause = clause
		} else if ictx := vctx.SymbolicCharactersClause(); ictx != nil {
			cctx := ictx.(*cobol85.SymbolicCharactersClauseContext)
			var typ pb.SymbolicCharactersClause_Type
			switch {
			case cctx.ALPHANUMERIC() != nil:
				typ = pb.SymbolicCharactersClause_ALPHANUMERIC
			case cctx.NATIONAL() != nil:
				typ = pb.SymbolicCharactersClause_NATIONAL
			}
			v.paragraph.SymbolicCharactersClause = &pb.SymbolicCharactersClause{
				Type: typ,
			}
		} else if ictx := vctx.EnvironmentSwitchNameClause(); ictx != nil {
			// TODO: cctx := ictx.(*cobol85.EnvironmentSwitchNameClauseContext)
		} else if ictx := vctx.DefaultDisplaySignClause(); ictx != nil {
			cctx := ictx.(*cobol85.DefaultDisplaySignClauseContext)
			var typ pb.DefaultDisplaySignClause_Type
			switch {
			case cctx.LEADING() != nil:
				typ = pb.DefaultDisplaySignClause_LEADING
			case cctx.TRAILING() != nil:
				typ = pb.DefaultDisplaySignClause_TRAILING
			}
			v.paragraph.DefaultDisplaySignClause = &pb.DefaultDisplaySignClause{
				Type: typ,
			}
		} else if ictx := vctx.DefaultComputationalSignClause(); ictx != nil {
			// TODO: cctx := ictx.(*cobol85.DefaultComputationalSignClauseContext)
		} else if ictx := vctx.ReserveNetworkClause(); ictx != nil {
			// cctx := ictx.(*cobol85.ReserveNetworkClauseContext)
			v.paragraph.ReserveNetworkClause = &pb.ReserveNetworkClause{}
		}
	}
	return v.VisitChildren(ctx)
}

func (v *SpecialNamesParagraphVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *SpecialNamesParagraphVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
