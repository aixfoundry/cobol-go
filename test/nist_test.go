package test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/document"
	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/options"
)

type ErrorListener struct {
	antlr.DefaultErrorListener
	errs []string
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		errs: []string{},
	}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("syntax error in line %d : %d %s", line, column, msg)
	l.errs = append(l.errs, err)
}

func (l *ErrorListener) GetErrors() []string {
	return l.errs
}

func TestParse(tt *testing.T) {
	rootdir := "./testdata/nist"
	infos, err := os.ReadDir(rootdir)
	if err != nil {
		tt.Error(err)
		tt.FailNow()
	}
	skips := []string{"ALTL1.CPY", "ALTLB.CPY"}
	opts := options.NewOptions().AddCopyBookDirectory(rootdir).SetFormat(format.FIXED)
FOR:
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		if !strings.HasSuffix(info.Name(), ".CBL") {
			continue FOR
		}
		for _, v := range skips {
			if v == info.Name() {
				continue FOR
			}
		}
		tt.Run(info.Name(), func(t *testing.T) {
			listener := NewErrorListener()

			filepath := path.Join(rootdir, info.Name())
			processedPath := filepath + ".preprocessed"
			var processed string
			_, err := os.Stat(processedPath)
			if err != nil {
				processed = document.ParseFile(filepath, opts)
				os.WriteFile(processedPath, []byte(processed), os.ModePerm)
			} else {
				buf, err := os.ReadFile(processedPath)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				processed = string(buf)
			}

			is := antlr.NewInputStream(processed)
			lexer := cobol85.NewCobol85Lexer(is)
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(listener)

			cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
			cpp := cobol85.NewCobol85Parser(cts)
			cpp.RemoveErrorListeners()
			cpp.AddErrorListener(listener)

			tree := conv.TreesStringTree(cpp.StartRule(), cpp.GetRuleNames(), 0)

			for _, err := range listener.GetErrors() {
				t.Error(err)
			}
			if len(listener.GetErrors()) != 0 {
				t.FailNow()
			}
			treePath := filepath + ".tree"
			_ = os.WriteFile(treePath, []byte(tree), os.ModePerm)
		})
	}
}
