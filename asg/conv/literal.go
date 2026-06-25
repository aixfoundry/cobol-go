package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func BooleanLiteral(in cobol85.IBooleanLiteralContext) *pb.BooleanLiteral {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.BooleanLiteralContext)
	return &pb.BooleanLiteral{
		Value: ctx.TRUE() != nil,
	}
}

func IntegerLiteral(in cobol85.IIntegerLiteralContext) *pb.IntegerLiteral {
	return &pb.IntegerLiteral{
		Value: in.GetText(),
	}
}

func NumericLiteral(in cobol85.INumericLiteralContext) (out *pb.NumericLiteral) {
	ctx := in.(*cobol85.NumericLiteralContext)
	out = &pb.NumericLiteral{}
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		out.Type = pb.NumericLiteral_INTEGER
		out.Value = ictx.GetText()
	} else if ictx := ctx.NUMERICLITERAL(); ictx != nil {
		out.Type = pb.NumericLiteral_FLOAT
		out.Value = ictx.GetText()
	}
	return
}

func FigurativeConstant(in cobol85.IFigurativeConstantContext) (out *pb.FigurativeConstant) {
	ctx := in.(*cobol85.FigurativeConstantContext)
	out = &pb.FigurativeConstant{}
	switch {
	case ctx.ALL() != nil:
		out.Type = pb.FigurativeConstant_ALL
	case ctx.HIGH_VALUE() != nil:
		out.Type = pb.FigurativeConstant_HIGH_VALUE
	case ctx.HIGH_VALUES() != nil:
		out.Type = pb.FigurativeConstant_HIGH_VALUES
	case ctx.LOW_VALUE() != nil:
		out.Type = pb.FigurativeConstant_LOW_VALUE
	case ctx.LOW_VALUES() != nil:
		out.Type = pb.FigurativeConstant_LOW_VALUES
	case ctx.NULL() != nil:
		out.Type = pb.FigurativeConstant_NULL
	case ctx.NULLS() != nil:
		out.Type = pb.FigurativeConstant_NULLS
	case ctx.QUOTE() != nil:
		out.Type = pb.FigurativeConstant_QUOTE
	case ctx.QUOTES() != nil:
		out.Type = pb.FigurativeConstant_QUOTES
	case ctx.SPACE() != nil:
		out.Type = pb.FigurativeConstant_SPACE
	case ctx.SPACES() != nil:
		out.Type = pb.FigurativeConstant_SPACES
	case ctx.ZERO() != nil:
		out.Type = pb.FigurativeConstant_ZERO
	case ctx.ZEROES() != nil:
		out.Type = pb.FigurativeConstant_ZEROES
	case ctx.ZEROS() != nil:
		out.Type = pb.FigurativeConstant_ZEROS
	}
	if ctx.Literal() != nil {
		out.Literal = Literal(ctx.Literal())
	}
	return
}

func Literal(in cobol85.ILiteralContext) (out *pb.Literal) {
	ctx := in.(*cobol85.LiteralContext)
	out = &pb.Literal{}
	if ictx := ctx.NumericLiteral(); ictx != nil {
		out.Type = pb.Literal_NUMERIC
		out.OneOf = &pb.Literal_NumericLiteral{
			NumericLiteral: NumericLiteral(ictx),
		}
	} else if ictx := ctx.BooleanLiteral(); ictx != nil {
		out.Type = pb.Literal_BOOLEAN
		out.OneOf = &pb.Literal_BooleanLiteral{
			BooleanLiteral: BooleanLiteral(ictx),
		}
	} else if ictx := ctx.NONNUMERICLITERAL(); ictx != nil {
		out.Type = pb.Literal_NON_NUMERIC
		out.OneOf = &pb.Literal_NonNumericLiteral{
			NonNumericLiteral: &pb.NonNumericLiteral{
				Value: ictx.GetText(),
			},
		}
	} else if ictx := ctx.FigurativeConstant(); ictx != nil {
		out.Type = pb.Literal_FIGURATIVE_CONSTANT
		out.OneOf = &pb.Literal_FigurativeConstant{
			FigurativeConstant: FigurativeConstant(ictx),
		}
	}
	return
}
