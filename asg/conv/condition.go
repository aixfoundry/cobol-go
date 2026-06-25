package conv

import (
	"strings"

	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

// Condition converts a G4 condition tree to proto Condition
func Condition(in cobol85.IConditionContext) (out *pb.Condition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.ConditionContext)
	out = &pb.Condition{}
	if ictx := ctx.CombinableCondition(); ictx != nil {
		out.CombinableCondition = CombinableCondition(ictx)
	}
	for _, v := range ctx.AllAndOrCondition() {
		out.AndOrCondition = append(out.AndOrCondition, AndOrCondition(v))
	}
	return
}

// AndOrCondition converts AND/OR combinable or abbreviations
func AndOrCondition(in cobol85.IAndOrConditionContext) (out *pb.AndOrCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.AndOrConditionContext)
	out = &pb.AndOrCondition{}

	if ctx.AND() != nil {
		out.AndOr = pb.AndOrCondition_AND
	} else if ctx.OR() != nil {
		out.AndOr = pb.AndOrCondition_OR
	}

	if ictx := ctx.CombinableCondition(); ictx != nil {
		out.CombinableCondition = CombinableCondition(ictx)
	}

	for _, v := range ctx.AllAbbreviation() {
		out.Abbreviations = append(out.Abbreviations, Abbreviation(v))
	}
	return
}

// CombinableCondition converts optionally negated simple condition
func CombinableCondition(in cobol85.ICombinableConditionContext) (out *pb.CombinableCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.CombinableConditionContext)
	out = &pb.CombinableCondition{}
	if ctx.NOT() != nil {
		out.Not = true
	}
	if ictx := ctx.SimpleCondition(); ictx != nil {
		out.SimpleCondition = SimpleCondition(ictx)
	}
	return
}

// SimpleCondition converts a simple condition (wrapped/relation/class/name ref)
func SimpleCondition(in cobol85.ISimpleConditionContext) (out *pb.SimpleCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.SimpleConditionContext)
	out = &pb.SimpleCondition{}

	if ictx := ctx.Condition(); ictx != nil {
		out.OneOf = &pb.SimpleCondition_Condition{
			Condition: Condition(ictx),
		}
	} else if ictx := ctx.RelationCondition(); ictx != nil {
		out.OneOf = &pb.SimpleCondition_RelationCondition{
			RelationCondition: RelationCondition(ictx),
		}
	} else if ictx := ctx.ClassCondition(); ictx != nil {
		out.OneOf = &pb.SimpleCondition_ClassCondition{
			ClassCondition: ClassCondition(ictx),
		}
	} else if ictx := ctx.ConditionNameReference(); ictx != nil {
		out.OneOf = &pb.SimpleCondition_ConditionNameReference{
			ConditionNameReference: ConditionNameReference(ictx),
		}
	}
	return
}

// ClassCondition converts identifier IS NOT? (NUMERIC|ALPHABETIC|...) condition
func ClassCondition(in cobol85.IClassConditionContext) (out *pb.ClassCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.ClassConditionContext)
	out = &pb.ClassCondition{}
	if ctx.Identifier() != nil {
		out.Identifier = Identifier(ctx.Identifier())
	}
	if ctx.NUMERIC() != nil {
		out.Type = &pb.ClassCondition_Value{Value: pb.ClassCondition_NUMERIC}
	} else if ctx.ALPHABETIC() != nil {
		out.Type = &pb.ClassCondition_Value{Value: pb.ClassCondition_ALPHABETIC}
	} else if ctx.ALPHABETIC_LOWER() != nil {
		out.Type = &pb.ClassCondition_Value{Value: pb.ClassCondition_ALPHABETIC_LOWER}
	} else if ctx.ALPHABETIC_UPPER() != nil {
		out.Type = &pb.ClassCondition_Value{Value: pb.ClassCondition_ALPHABETIC_UPPER}
	} else if ctx.DBCS() != nil {
		out.Type = &pb.ClassCondition_Value{Value: pb.ClassCondition_DBCS}
	} else if ctx.KANJI() != nil {
		out.Type = &pb.ClassCondition_Value{Value: pb.ClassCondition_KANJI}
	} else if ctx.ClassName() != nil {
		out.Type = &pb.ClassCondition_ClassName{
			ClassName: ClassName(ctx.ClassName()),
		}
	}
	return
}

// ConditionNameReference converts a condition name with optional qualifiers
func ConditionNameReference(in cobol85.IConditionNameReferenceContext) (out *pb.ConditionNameReference) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.ConditionNameReferenceContext)
	out = &pb.ConditionNameReference{}
	if ictx := ctx.ConditionName(); ictx != nil {
		out.ConditionName = ConditionName(ictx)
	}

	// Check if this has inData/inFile qualifiers (means data reference style)
	allInData := ctx.AllInData()
	allMnemonic := ctx.AllInMnemonic()

	if len(allInData) > 0 || ctx.InFile() != nil {
		// InSubscript style: conditionName inData* inFile? subscript*
		inSub := &pb.ConditionNameReference_InSubscript{}
		for _, v := range allInData {
			inSub.InDatas = append(inSub.InDatas, InData(v))
		}
		if ictx := ctx.InFile(); ictx != nil {
			inSub.InFile = InFile(ictx)
		}
		for _, v := range ctx.AllConditionNameSubscriptReference() {
			subRef := &pb.ConditionNameReference_SubscriptReference{}
			cv := v.(*cobol85.ConditionNameSubscriptReferenceContext)
			for _, sctx := range cv.AllSubscript() {
				subRef.Subscripts = append(subRef.Subscripts, Subscript(sctx))
			}
			inSub.Refs = append(inSub.Refs, subRef)
		}
		out.In = &pb.ConditionNameReference_InSubscript_{
			InSubscript: inSub,
		}
	} else if len(allMnemonic) > 0 {
		// InMnemonic style: conditionName inMnemonic*
		// Use the last one
		lastMnemonic := allMnemonic[len(allMnemonic)-1]
		out.In = &pb.ConditionNameReference_InMnemonic{
			InMnemonic: InMnemonic(lastMnemonic),
		}
	}
	return
}

// RelationCondition converts a relation condition (sign/comparison/combined)
func RelationCondition(in cobol85.IRelationConditionContext) (out *pb.RelationCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.RelationConditionContext)
	out = &pb.RelationCondition{}

	if ictx := ctx.RelationSignCondition(); ictx != nil {
		out.Condition = &pb.RelationCondition_RelationSignCondition{
			RelationSignCondition: RelationSignCondition(ictx),
		}
	} else if ictx := ctx.RelationArithmeticComparison(); ictx != nil {
		out.Condition = &pb.RelationCondition_RelationArithmeticComparison{
			RelationArithmeticComparison: RelationArithmeticComparison(ictx),
		}
	} else if ictx := ctx.RelationCombinedComparison(); ictx != nil {
		out.Condition = &pb.RelationCondition_RelationCombinedComparison{
			RelationCombinedComparison: RelationCombinedComparison(ictx),
		}
	}
	return
}

// RelationSignCondition converts arithmetic expression IS NOT? (POSITIVE|NEGATIVE|ZERO)
func RelationSignCondition(in cobol85.IRelationSignConditionContext) (out *pb.RelationSignCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.RelationSignConditionContext)
	out = &pb.RelationSignCondition{}
	if ctx.ArithmeticExpression() != nil {
		out.ArithmeticExpression = ArithmeticExpression(ctx.ArithmeticExpression())
	}
	if ctx.POSITIVE() != nil {
		out.Type = pb.RelationSignCondition_POSITIVE
	} else if ctx.NEGATIVE() != nil {
		out.Type = pb.RelationSignCondition_NEGATIVE
	} else if ctx.ZERO() != nil {
		out.Type = pb.RelationSignCondition_ZERO
	}
	return
}

// RelationArithmeticComparison converts left op right
func RelationArithmeticComparison(in cobol85.IRelationArithmeticComparisonContext) (out *pb.RelationArithmeticComparison) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.RelationArithmeticComparisonContext)
	out = &pb.RelationArithmeticComparison{}
	allExprs := ctx.AllArithmeticExpression()
	if len(allExprs) >= 1 {
		out.LeftExpression = ArithmeticExpression(allExprs[0])
	}
	if len(allExprs) >= 2 {
		out.RightExpression = ArithmeticExpression(allExprs[1])
	}
	if ictx := ctx.RelationalOperator(); ictx != nil {
		out.RelationalOperator = RelationalOperator(ictx)
	}
	return
}

// RelationCombinedComparison converts expression op (combinedCondition)
func RelationCombinedComparison(in cobol85.IRelationCombinedComparisonContext) (out *pb.RelationCombinedComparison) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.RelationCombinedComparisonContext)
	out = &pb.RelationCombinedComparison{}
	if ctx.ArithmeticExpression() != nil {
		out.ArithmeticExpression = ArithmeticExpression(ctx.ArithmeticExpression())
	}
	if ictx := ctx.RelationalOperator(); ictx != nil {
		out.RelationalOperator = RelationalOperator(ictx)
	}
	if ictx := ctx.RelationCombinedCondition(); ictx != nil {
		out.RelationCombinedCondition = RelationCombinedCondition(ictx)
	}
	return
}

// RelationCombinedCondition converts expression (AND|OR expression)+
func RelationCombinedCondition(in cobol85.IRelationCombinedConditionContext) (out *pb.RelationCombinedCondition) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.RelationCombinedConditionContext)
	out = &pb.RelationCombinedCondition{}

	allExprs := ctx.AllArithmeticExpression()
	if len(allExprs) >= 1 {
		out.LeftExpression = ArithmeticExpression(allExprs[0])
	}

	// Check AND/OR tokens
	if len(ctx.AllAND()) > 0 {
		out.AndOr = pb.RelationCombinedCondition_AND
	} else if len(ctx.AllOR()) > 0 {
		out.AndOr = pb.RelationCombinedCondition_OR
	}

	// Remaining expressions are right expressions
	for i := 1; i < len(allExprs); i++ {
		out.RightExpressions = append(out.RightExpressions, ArithmeticExpression(allExprs[i]))
	}
	return
}

// RelationalOperator converts the combined token sequence to a RelationalOperator enum
func RelationalOperator(in cobol85.IRelationalOperatorContext) pb.RelationalOperator {
	if in == nil {
		return pb.RelationalOperator_EQUAL
	}
	ctx := in.(*cobol85.RelationalOperatorContext)
	text := strings.TrimSpace(ctx.GetText())
	notPresent := ctx.NOT() != nil

	switch {
	case notPresent && (ctx.GREATER() != nil || ctx.MORETHANCHAR() != nil):
		// NOT GREATER / NOT >
		return pb.RelationalOperator_NOT_GREATER
	case notPresent && (ctx.LESS() != nil || ctx.LESSTHANCHAR() != nil):
		// NOT LESS / NOT <
		return pb.RelationalOperator_NOT_LESS
	case notPresent && (ctx.EQUAL() != nil || ctx.EQUALCHAR() != nil):
		// NOT = / NOT EQUAL
		return pb.RelationalOperator_NOT_EQUAL
	case ctx.NOTEQUALCHAR() != nil:
		return pb.RelationalOperator_NOT_EQUAL
	case ctx.MORETHANOREQUAL() != nil:
		return pb.RelationalOperator_GREATER_OR_EQUAL
	case ctx.LESSTHANOREQUAL() != nil:
		return pb.RelationalOperator_LESS_OR_EQUAL
	case ctx.GREATER() != nil || ctx.MORETHANCHAR() != nil:
		if ctx.OR() != nil || ctx.EQUAL() != nil {
			// GREATER OR EQUAL / GREATER THAN OR EQUAL
			return pb.RelationalOperator_GREATER_OR_EQUAL
		}
		return pb.RelationalOperator_GREATER
	case ctx.LESS() != nil || ctx.LESSTHANCHAR() != nil:
		if ctx.OR() != nil || ctx.EQUAL() != nil {
			// LESS OR EQUAL / LESS THAN OR EQUAL
			return pb.RelationalOperator_LESS_OR_EQUAL
		}
		return pb.RelationalOperator_LESS
	case ctx.EQUAL() != nil || ctx.EQUALCHAR() != nil:
		return pb.RelationalOperator_EQUAL
	}

	// Fallback: parse from text
	switch {
	case strings.Contains(text, ">=") || strings.Contains(text, "MORETHANOREQUAL") ||
		strings.Contains(strings.ToUpper(text), "GREATER") && strings.Contains(strings.ToUpper(text), "EQUAL"):
		return pb.RelationalOperator_GREATER_OR_EQUAL
	case strings.Contains(text, "<=") || strings.Contains(text, "LESSTHANOREQUAL") ||
		strings.Contains(strings.ToUpper(text), "LESS") && strings.Contains(strings.ToUpper(text), "EQUAL"):
		return pb.RelationalOperator_LESS_OR_EQUAL
	case strings.Contains(text, "!=") || strings.Contains(text, "<>") || strings.Contains(text, "NOTEQUAL"):
		return pb.RelationalOperator_NOT_EQUAL
	case strings.Contains(text, ">") || strings.Contains(strings.ToUpper(text), "GREATER"):
		if strings.Contains(text, "NOT") {
			return pb.RelationalOperator_NOT_GREATER
		}
		return pb.RelationalOperator_GREATER
	case strings.Contains(text, "<") || strings.Contains(strings.ToUpper(text), "LESS"):
		if strings.Contains(text, "NOT") {
			return pb.RelationalOperator_NOT_LESS
		}
		return pb.RelationalOperator_LESS
	default:
		return pb.RelationalOperator_EQUAL
	}
}

// Abbreviation converts abbreviated condition parts
func Abbreviation(in cobol85.IAbbreviationContext) (out *pb.Abbreviation) {
	if in == nil {
		return nil
	}
	ctx := in.(*cobol85.AbbreviationContext)
	out = &pb.Abbreviation{}
	if ictx := ctx.RelationalOperator(); ictx != nil {
		out.RelationalOperator = RelationalOperator(ictx)
	}
	if ctx.ArithmeticExpression() != nil {
		out.ArithmeticExpression = ArithmeticExpression(ctx.ArithmeticExpression())
	}

	// Handle recursive abbreviation: ( expression abbreviation )
	if ictx := ctx.Abbreviation(); ictx != nil {
		out.Abbreviation = Abbreviation(ictx)
	}
	return
}
