package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func ArithmeticExpression(in cobol85.IArithmeticExpressionContext) (out *pb.ArithmeticExpression) {
	ctx := in.(*cobol85.ArithmeticExpressionContext)
	out = &pb.ArithmeticExpression{}
	if ictx := ctx.MultDivs(); ictx != nil {
		out.MultDivs = MultDivs(ictx)
	}
	for _, v := range ctx.AllPlusMinus() {
		out.PlusMinus = append(out.PlusMinus, PlusMinus(v))
	}
	return
}

func PlusMinus(in cobol85.IPlusMinusContext) (out *pb.PlusMinus) {
	ctx := in.(*cobol85.PlusMinusContext)
	out = &pb.PlusMinus{}
	if ctx.MINUSCHAR() != nil {
		out.Operator = pb.PlusMinus_MINUS
	} else {
		out.Operator = pb.PlusMinus_PLUS
	}
	if ctx.MultDivs() != nil {
		out.MultDivs = MultDivs(ctx.MultDivs())
	}
	return
}

func MultDivs(in cobol85.IMultDivsContext) (out *pb.MultDivs) {
	ctx := in.(*cobol85.MultDivsContext)
	out = &pb.MultDivs{
		Powers:  &pb.Powers{},
		MultDiv: []*pb.MultDiv{},
	}
	if ictx := ctx.Powers(); ictx != nil {
		out.Powers = Powers(ictx)
	}
	for _, v := range ctx.AllMultDiv() {
		out.MultDiv = append(out.MultDiv, MultDiv(v))
	}
	return
}

func MultDiv(in cobol85.IMultDivContext) (out *pb.MultDiv) {
	ctx := in.(*cobol85.MultDivContext)
	out = &pb.MultDiv{
		Powers: &pb.Powers{},
	}
	if ctx.SLASHCHAR() != nil {
		out.Operator = pb.MultDiv_DIV
	} else {
		out.Operator = pb.MultDiv_MULT
	}
	if ictx := ctx.Powers(); ictx != nil {
		out.Powers = Powers(ictx)
	}
	return
}

func Powers(in cobol85.IPowersContext) (out *pb.Powers) {
	ctx := in.(*cobol85.PowersContext)
	out = &pb.Powers{}
	for _, v := range ctx.AllPower() {
		out.Powers = append(out.Powers, Power(v))
	}
	if ctx.Basis() != nil {
		out.Basis = Basis(ctx.Basis())
	}
	return
}

func Power(in cobol85.IPowerContext) (out *pb.Power) {
	ctx := in.(*cobol85.PowerContext)
	out = &pb.Power{}
	if ctx.DOUBLEASTERISKCHAR() != nil {
		out.HasPowerOperator = true
	}
	if ctx.Basis() != nil {
		out.Basis = Basis(ctx.Basis())
	}
	return
}

func Basis(in cobol85.IBasisContext) (out *pb.Basis) {
	ctx := in.(*cobol85.BasisContext)
	out = &pb.Basis{}
	if ctx.ArithmeticExpression() != nil {
		out.OneOf = &pb.Basis_ArithmeticExpression{
			ArithmeticExpression: ArithmeticExpression(ctx.ArithmeticExpression()),
		}
	} else if ctx.Identifier() != nil {
		out.OneOf = &pb.Basis_Identifier{
			Identifier: Identifier(ctx.Identifier()),
		}
	} else if ctx.Literal() != nil {
		out.OneOf = &pb.Basis_Literal{
			Literal: Literal(ctx.Literal()),
		}
	}
	return
}
