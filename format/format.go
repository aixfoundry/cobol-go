package format

import "regexp"

type Format int32

const (
	// FIXED Fixed format, standard ANSI / IBM reference. Each line 80 chars.
	//
	//  1-6: sequence area
	//  7: indicator field
	//  8-12: area A
	//  13-72: area B
	//  73-80: comments
	FIXED Format = iota

	// TANDEM HP Tandem format.
	//
	//  1: indicator field
	//  2-5: optional area A
	//  6-132: optional area B
	TANDEM

	// VARIABLE Variable format.
	//
	//  1-6: sequence area
	//  7: indicator field
	//  8-12: optional area A
	//  13-*: optional area B
	VARIABLE

	// FREE COBOL 2002/2014 free format.
	//
	//  1: indicator field
	//  2-*: source text (no fixed columns)
	FREE
)

var (
	formats = map[Format]string{
		FIXED:    `(.{0,6})(?:([ABCdD$\t\-/*# ])(.{0,4})(.{0,61})(.*))?`,
		TANDEM:   `()(?:([ABCdD$\t\-/*# ])(.{0,4})(.*)())?`,
		VARIABLE: `(.{0,6})(?:([ABCdD$\t\-/*# ])(.{0,4})(.*)())?`,
		FREE:     `()(?:([ABCdD$\t\-/*# ]?)()(.*)())`,
	}

	values = map[Format]string{
		FIXED:    "FIXED",
		TANDEM:   "TANDEM",
		VARIABLE: "VARIABLE",
		FREE:     "FREE",
	}

	regexps = map[Format]*regexp.Regexp{
		FIXED:    regexp.MustCompile(FIXED.Format()),
		TANDEM:   regexp.MustCompile(TANDEM.Format()),
		VARIABLE: regexp.MustCompile(VARIABLE.Format()),
		FREE:     regexp.MustCompile(FREE.Format()),
	}

	// if comment entry is multiline
	multilines = map[Format]bool{
		FIXED:    true,
		TANDEM:   false,
		VARIABLE: true,
		FREE:     false,
	}
)

func (f Format) String() string {
	return values[f]
}

func (f Format) Format() string {
	return formats[f]
}

func (f Format) Regexp() *regexp.Regexp {
	return regexps[f]
}

func (f Format) MultiLine() bool {
	return multilines[f]
}
