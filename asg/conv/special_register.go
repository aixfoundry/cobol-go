package conv

import (
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

// SpecialRegister converts a SpecialRegister ANTLR parse context to a pb.SpecialRegister.
// It maps parser token accessors to the SpecialRegister_Type enum and captures
// the nested Identifier for ADDRESS OF / LENGTH OF forms.
func SpecialRegister(in cobol85.ISpecialRegisterContext) (out *pb.SpecialRegister) {
	ctx := in.(*cobol85.SpecialRegisterContext)
	out = &pb.SpecialRegister{}

	switch {
	case ctx.ADDRESS() != nil:
		out.Type = pb.SpecialRegister_ADDRESS_OF
	case ctx.DATE() != nil:
		out.Type = pb.SpecialRegister_DATE
	case ctx.DAY_OF_WEEK() != nil:
		out.Type = pb.SpecialRegister_DAY_OF_WEEK
	case ctx.DAY() != nil:
		out.Type = pb.SpecialRegister_DAY
	case ctx.DEBUG_CONTENTS() != nil:
		out.Type = pb.SpecialRegister_DEBUG_CONTENTS
	case ctx.DEBUG_ITEM() != nil:
		out.Type = pb.SpecialRegister_DEBUG_ITEM
	case ctx.DEBUG_LINE() != nil:
		out.Type = pb.SpecialRegister_DEBUG_LINE
	case ctx.DEBUG_NAME() != nil:
		out.Type = pb.SpecialRegister_DEBUG_NAME
	case ctx.DEBUG_SUB_1() != nil:
		out.Type = pb.SpecialRegister_DEBUG_SUB_1
	case ctx.DEBUG_SUB_2() != nil:
		out.Type = pb.SpecialRegister_DEBUG_SUB_2
	case ctx.DEBUG_SUB_3() != nil:
		out.Type = pb.SpecialRegister_DEBUG_SUB_3
	case ctx.LENGTH() != nil:
		out.Type = pb.SpecialRegister_LENGTH_OF
	case ctx.LINAGE_COUNTER() != nil:
		out.Type = pb.SpecialRegister_LINAGE_COUNTER
	case ctx.LINE_COUNTER() != nil:
		out.Type = pb.SpecialRegister_LINE_COUNTER
	case ctx.PAGE_COUNTER() != nil:
		out.Type = pb.SpecialRegister_PAGE_COUNTER
	case ctx.RETURN_CODE() != nil:
		out.Type = pb.SpecialRegister_RETURN_CODE
	case ctx.SHIFT_IN() != nil:
		out.Type = pb.SpecialRegister_SHIFT_IN
	case ctx.SHIFT_OUT() != nil:
		out.Type = pb.SpecialRegister_SHIFT_OUT
	case ctx.SORT_CONTROL() != nil:
		out.Type = pb.SpecialRegister_SORT_CONTROL
	case ctx.SORT_CORE_SIZE() != nil:
		out.Type = pb.SpecialRegister_SORT_CORE_SIZE
	case ctx.SORT_FILE_SIZE() != nil:
		out.Type = pb.SpecialRegister_SORT_FILE_SIZE
	case ctx.SORT_MESSAGE() != nil:
		out.Type = pb.SpecialRegister_SORT_MESSAGE
	case ctx.SORT_MODE_SIZE() != nil:
		out.Type = pb.SpecialRegister_SORT_MODE_SIZE
	case ctx.SORT_RETURN() != nil:
		out.Type = pb.SpecialRegister_SORT_RETURN
	case ctx.TALLY() != nil:
		out.Type = pb.SpecialRegister_TALLY
	case ctx.TIME() != nil:
		out.Type = pb.SpecialRegister_TIME
	case ctx.WHEN_COMPILED() != nil:
		out.Type = pb.SpecialRegister_WHEN_COMPILED
	}

	// ADDRESS OF identifier and LENGTH [OF] identifier carry a nested Identifier
	if ctx.Identifier() != nil {
		out.Identifier = Identifier(ctx.Identifier())
	}

	return
}
