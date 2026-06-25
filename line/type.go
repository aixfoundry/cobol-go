package line

import "github.com/aixfoundry/cobol-go/constant"

type Type int32

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
	case constant.CHAR_D, constant.CHAR_D_:
		return DEBUG
	case constant.CHAR_MINUS:
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
