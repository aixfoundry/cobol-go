package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/document"
	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/options"
	"github.com/antlr4-go/antlr/v4"
)

var (
	formatFlag   = flag.String("format", "FIXED", "cobol file format, FIXED, TANDEM, VARIABLE")
	pathFlag     = flag.String("path", ".", "cobol files path")
	suffixFlag   = flag.String("suffix", ".CBL,src,COB", "cobol cbl, separated by commas")
	copyPathFlag = flag.String("copyPath", "", "cobol copy book directory")
)

var (
	begin = time.Now() // 开始时间
	count = 0          // 开始计数
)

func main() {
	flag.Parse()

	f := format.FIXED
	switch strings.ToUpper(*formatFlag) {
	case "FIXED":
		f = format.FIXED
	case "TANDEM":
		f = format.TANDEM
	case "VARIABLE":
		f = format.VARIABLE
	case "FREE":
		f = format.FREE
	}
	rootPath := *pathFlag
	fi, err := os.Stat(rootPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	if !fi.IsDir() {
		TreesStringTree(rootPath, f)
	} else {
		filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				TreesStringTree(path, f)
			}
			return err
		})
	}
	fmt.Fprintf(os.Stdout, "%d takes %s\n", count, time.Since(begin))
}

func TreesStringTree(path string, f format.Format) {
	if strings.HasSuffix(path, ".tree") {
		return
	}
	suffixes := strings.Split(*suffixFlag, ",")
	has := false
	for _, v := range suffixes {
		trimed := strings.TrimSpace(v)
		r := regexp.MustCompile(fmt.Sprintf(`(?i)%s$`, trimed))
		if r.MatchString(path) {
			has = true
			break
		}
	}
	if !has {
		return
	}

	start := time.Now()
	buff, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	copyPath := *copyPathFlag
	if copyPath == "" {
		copyPath = filepath.Dir(path)
	}
	processed := string(buff)
	if !strings.HasSuffix(path, ".preprocessed") {
		opts := options.NewOptions().SetFormat(f).AddCopyBookDirectory(copyPath)
		processed = document.Parse(processed, opts)
	}

	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	l := NewErrorListener()
	lexer.AddErrorListener(l)

	ctx := cpp.StartRule()

	tree := conv.TreesStringTree(ctx, cpp.GetRuleNames(), 0)
	if err := os.WriteFile(path+".tree", []byte(tree), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "write tree file %s: %s\n", path+".tree", err)
	}
	errs := l.GetErrors()
	if len(errs) != 0 {
		if err := os.WriteFile(path+".error", []byte(strings.Join(errs, "\n")), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write error file %s: %s\n", path+".error", err)
		}
	}
	fmt.Fprintf(os.Stdout, "%s %s %d\n", path, time.Since(start), len(errs))
	count++
	if count%10 == 0 {
		fmt.Fprintf(os.Stdout, "%d takes %s\n", count, time.Since(begin))
	}
}

type ErrorListener struct {
	antlr.DefaultErrorListener
	errs []string
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		errs: []string{},
	}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("syntax error in line %d : %d %s", line, column, msg)
	l.errs = append(l.errs, err)
}

func (l *ErrorListener) GetErrors() []string {
	return l.errs
}
