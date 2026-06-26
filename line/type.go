package line

import (
	"fmt"

	"github.com/aixfoundry/cobol-go/constant"
)

type Type int32

// String returns a human-readable representation of the line type.
func (t Type) String() string {
	switch t {
	case BLANK:
		return "BLANK"
	case COMMENT:
		return "COMMENT"
	case COMPILER_DIRECTIVE:
		return "COMPILER_DIRECTIVE"
	case CONTINUATION:
		return "CONTINUATION"
	case DEBUG:
		return "DEBUG"
	case NORMAL:
		return "NORMAL"
	default:
		return fmt.Sprintf("Type(%d)", int32(t))
	}
}

const (
	BLANK Type = iota
	COMMENT
	COMPILER_DIRECTIVE
	CONTINUATION
	DEBUG
	NORMAL
)

func ToType(str string) Type {
	switch str {
	// FREE-format prefixes (multi-character; never produced by FIXED/TANDEM/
	// VARIABLE, whose indicator is always a single character).
	case "*>":
		return COMMENT
	case "D>>":
		return DEBUG
	case ">>":
		return COMPILER_DIRECTIVE
	// Fixed/Tandem/Variable single-character indicators.
	case constant.CHAR_D, constant.CHAR_D_:
		return DEBUG
	case constant.CHAR_MINUS:
		return CONTINUATION
	case constant.CHAR_AMPERSAND:
		return CONTINUATION
	case constant.CHAR_DOLLAR_SIGN:
		return COMPILER_DIRECTIVE
	case constant.CHAR_ASTERISK, constant.CHAR_SLASH:
		return COMMENT
	case constant.CHAR_WHITESPACE:
		return NORMAL
	default:
		return NORMAL
	}
}
