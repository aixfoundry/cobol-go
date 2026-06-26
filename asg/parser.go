package asg

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/asg/visitor"
	"github.com/aixfoundry/cobol-go/document"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/options"
	"github.com/aixfoundry/cobol-go/pb"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AnalyzeFile parses a COBOL source file and returns its full Program AST.
func AnalyzeFile(filename string, opts ...options.Option) (*pb.Program, error) {
	if filename == "" {
		return nil, fmt.Errorf("filename is empty")
	}
	program := &pb.Program{}
	if err := AnalyzeCompilationUnit(filename, program, opts...); err != nil {
		return nil, err
	}
	return program, nil
}

// GetCompilationUnitName derives a compilation unit name from the filename.
func GetCompilationUnitName(filename string) string {
	return cases.Title(language.English).String(strings.TrimSuffix(path.Base(filename), path.Ext(filename)))
}

// AnalyzeCompilationUnit parses a single COBOL compilation unit and populates
// the given Program with the resulting AST. It writes a .tree debug file as
// a side effect.
func AnalyzeCompilationUnit(filename string, program *pb.Program, opts ...options.Option) error {
	name := GetCompilationUnitName(filename)
	processed, err := document.ParseFile(filename, opts...)
	if err != nil {
		return fmt.Errorf("parse file %s: %w", filename, err)
	}

	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	ctx := cpp.StartRule()

	tree := conv.TreesStringTree(ctx, cpp.GetRuleNames(), 0)
	if err := os.WriteFile(filename+".tree", []byte(tree), 0o644); err != nil {
		return fmt.Errorf("write tree file: %w", err)
	}

	vr := visitor.NewCompilationUnitVisitor(name, program)

	vr.Visit(ctx)
	return nil
}
