package document

import (
	"sort"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/preprocessor"
)

type Context struct {
	replaces   ReplaceStores
	prefixings []string
	buffer     string
}

func NewContext() *Context {
	return &Context{
		replaces:   ReplaceStores{},
		prefixings: []string{},
		buffer:     "",
	}
}

func (ctx *Context) StoreReplace(ctxs []preprocessor.IReplaceClauseContext) {
	for _, v := range ctxs {
		rcc, ok := v.(*preprocessor.ReplaceClauseContext)
		if ok {
			ctx.replaces = append(ctx.replaces, NewReplaceStore(rcc.Replaceable(), rcc.Replacement()))
		}
	}
}

func (ctx *Context) Replace(cts *antlr.CommonTokenStream) {
	if len(ctx.replaces) == 0 {
		return
	}
	sort.Sort(ctx.replaces)
	for _, store := range ctx.replaces {
		ctx.buffer = store.Replace(ctx.buffer, cts)
	}
}

func (ctx *Context) StorePrefixing(str string) {
	ctx.prefixings = append(ctx.prefixings, str)
}

func (ctx *Context) Prefixing(cts *antlr.CommonTokenStream) {
	if len(ctx.prefixings) == 0 {
		return
	}
	// TODO: prefix
}

func (ctx *Context) Read() string {
	return ctx.buffer
}

func (ctx *Context) Write(s string) {
	ctx.buffer += s
}
