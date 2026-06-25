package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

func Identifier(in cobol85.IIdentifierContext) (out *pb.Identifier) {
	ctx := in.(*cobol85.IdentifierContext)
	out = &pb.Identifier{}
	if ictx := ctx.QualifiedDataName(); ictx != nil {
		out.OneOf = &pb.Identifier_QualifiedDataName{
			QualifiedDataName: QualifiedDataName(ictx),
		}
	} else if ictx := ctx.TableCall(); ictx != nil {
		out.OneOf = &pb.Identifier_TableCall{
			TableCall: TableCall(ictx),
		}
	} else if ictx := ctx.FunctionCall(); ictx != nil {
		out.OneOf = &pb.Identifier_FunctionCall{
			FunctionCall: FunctionCall(ictx),
		}
	} else if ictx := ctx.SpecialRegister(); ictx != nil {
		out.OneOf = &pb.Identifier_SpecialRegister{
			SpecialRegister: &pb.SpecialRegister{},
		}
	}
	return
}

func FunctionCall(in cobol85.IFunctionCallContext) (out *pb.FunctionCall) {
	ctx := in.(*cobol85.FunctionCallContext)
	args := []*pb.Argument{}
	for _, v := range ctx.AllArgument() {
		args = append(args, Argument(v))
	}
	out = &pb.FunctionCall{
		FunctionName:      FunctionName(ctx.FunctionName()),
		Arguments:         args,
		ReferenceModifier: ReferenceModifier(ctx.ReferenceModifier()),
	}
	return
}

func Argument(in cobol85.IArgumentContext) (out *pb.Argument) {
	ctx := in.(*cobol85.ArgumentContext)
	if ictx := ctx.Identifier(); ictx != nil {
		out.OneOf = &pb.Argument_Identifier{
			Identifier: Identifier(ictx),
		}
	} else if ictx := ctx.Literal(); ictx != nil {
		out.OneOf = &pb.Argument_Literal{
			Literal: Literal(ictx),
		}
	} else if ictx := ctx.QualifiedDataName(); ictx != nil {
		out.OneOf = &pb.Argument_QualifiedDataName{
			QualifiedDataName: &pb.QualifiedDataNameIntegerLiteral{
				QualifiedDataName: QualifiedDataName(ctx.QualifiedDataName()),
				IntegerLiteral:    IntegerLiteral(ctx.IntegerLiteral()),
			},
		}
	} else if ictx := ctx.IndexName(); ictx != nil {
		out.OneOf = &pb.Argument_IndexName{
			IndexName: &pb.IndexNameIntegerLiteral{
				IndexName:      IndexName(ctx.IndexName()),
				IntegerLiteral: IntegerLiteral(ctx.IntegerLiteral()),
			},
		}
	} else if ictx := ctx.ArithmeticExpression(); ictx != nil {
		out.OneOf = &pb.Argument_ArithmeticExpression{
			ArithmeticExpression: &pb.ArithmeticExpression{},
		}
	}
	return
}
