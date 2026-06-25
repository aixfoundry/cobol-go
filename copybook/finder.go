package copybook

import (
	"regexp"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/preprocessor"
	"github.com/aixfoundry/cobol-go/options"
)

func TrimQuotes(s string) string {
	return regexp.MustCompile("^[\"']|[\"']$").ReplaceAllString(s, "")
}

func GetCopyBook(ctx antlr.ParserRuleContext, opts *options.Options) (ret string) {
	if cctx, ok := ctx.(*preprocessor.CopySourceContext); ok {
		if icctx := cctx.CobolWord(); icctx != nil {
			ret = NewCobolWordFinder().GetCopyBook(icctx, opts)
		} else if icctx := cctx.Filename(); icctx != nil {
			ret = NewFilenameFinder().GetCopyBook(icctx, opts)
		} else if icctx := cctx.Literal(); icctx != nil {
			ret = NewLiteralFinder().GetCopyBook(icctx, opts)
		} else {
			// nothing to do
			ret = ""
		}
	} else if cctx, ok := ctx.(preprocessor.ICobolWordContext); ok {
		ret = NewCobolWordFinder().GetCopyBook(cctx, opts)
	} else if cctx, ok := ctx.(preprocessor.IFilenameContext); ok {
		ret = NewFilenameFinder().GetCopyBook(cctx, opts)
	} else if cctx, ok := ctx.(preprocessor.ILiteralContext); ok {
		ret = NewLiteralFinder().GetCopyBook(cctx, opts)
	} else {
		// nothing to do
		ret = ""
	}
	return
}
