package line

import (
	"github.com/aixfoundry/cobol-go/format"
)

type Line struct {
	Format    format.Format
	Dialect   format.Dialect
	Type      Type
	No        int // Line No
	Sequence  string
	Indicator string
	ContentA  string
	ContentB  string
	Comment   string
}

func NewLine(text string, no int, f format.Format, dialect format.Dialect) (l *Line) {
	re := f.Regexp()
	groups := re.FindAllStringSubmatch(text, -1)
	if len(groups) == 0 {
		return
	}
	group := groups[0]
	if len(group) != 6 {
		return
	}
	l = &Line{
		Format:    f,
		Type:      ToType(group[2]),
		No:        no,
		Dialect:   dialect,
		Sequence:  group[1],
		Indicator: group[2],
		ContentA:  group[3],
		ContentB:  group[4],
		Comment:   group[5],
	}
	return l
}

func (l *Line) Content() string {
	return l.ContentA + l.ContentB
}

func (l *Line) SetContent(ctt string) {
	if len(ctt) > 4 {
		l.ContentA = ctt[:4]
		l.ContentB = ctt[4:]
	} else {
		l.ContentA = ctt
		l.ContentB = ""
	}
}

func (l *Line) String() string {
	return l.Sequence + l.Indicator + l.ContentA + l.ContentB + l.Comment
}
