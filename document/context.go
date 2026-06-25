package document

import (
	"regexp"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/preprocessor"
)

// cobolKeywords contains COBOL reserved words that should not be prefixed.
var cobolKeywords = map[string]bool{
	"ACCEPT": true, "ADD": true, "ADDRESS": true, "ADVANCING": true, "AFTER": true,
	"ALL": true, "ALPHABETIC": true, "ALPHABETIC-LOWER": true, "ALPHABETIC-UPPER": true,
	"ALPHANUMERIC": true, "ALPHANUMERIC-EDITED": true, "ALSO": true, "ALTER": true,
	"ALTERNATE": true, "AND": true, "ANY": true, "ARE": true, "AREA": true, "AREAS": true,
	"ASCENDING": true, "ASSIGN": true, "AT": true, "AUTHOR": true,
	"BEFORE": true, "BEGINNING": true, "BINARY": true, "BLANK": true, "BLOCK": true,
	"BOTTOM": true, "BY": true,
	"CALL": true, "CANCEL": true, "CHARACTER": true, "CHARACTERS": true, "CLASS": true,
	"CLOSE": true, "COBOL": true, "CODE": true, "CODE-SET": true, "COLLATING": true,
	"COLUMN": true, "COMMA": true, "COMMON": true, "COMMUNICATION": true, "COMP": true,
	"COMPUTATIONAL": true, "COMPUTE": true, "CONFIGURATION": true, "CONTAINS": true,
	"CONTENT": true, "CONTINUE": true, "CONTROL": true, "CONTROLS": true, "CONVERTING": true,
	"COPY": true, "CORR": true, "CORRESPONDING": true, "COUNT": true, "CURRENCY": true,
	"DATA": true, "DATE": true, "DATE-COMPILED": true, "DATE-WRITTEN": true, "DAY": true,
	"DAY-OF-WEEK": true, "DEBUGGING": true, "DECIMAL-POINT": true, "DECLARATIVES": true,
	"DELETE": true, "DELIMITED": true, "DELIMITER": true, "DEPENDING": true,
	"DESCENDING": true, "DESTINATION": true, "DETAIL": true, "DISABLE": true,
	"DISPLAY": true, "DIVIDE": true, "DIVISION": true, "DOWN": true, "DUPLICATES": true,
	"DYNAMIC": true,
	"EJECT": true, "ELSE": true, "ENABLE": true, "END": true, "END-ACCEPT": true,
	"END-ADD": true, "END-CALL": true, "END-COMPUTE": true, "END-DELETE": true,
	"END-DIVIDE": true, "END-EVALUATE": true, "END-IF": true, "END-MULTIPLY": true,
	"END-OF-PAGE": true, "END-PERFORM": true, "END-READ": true, "END-RECEIVE": true,
	"END-RETURN": true, "END-REWRITE": true, "END-SEARCH": true, "END-START": true,
	"END-STRING": true, "END-SUBTRACT": true, "END-UNSTRING": true, "END-WRITE": true,
	"ENTER": true, "ENTRY": true, "ENVIRONMENT": true, "EQUAL": true, "ERROR": true,
	"EVALUATE": true, "EVERY": true, "EXCEPTION": true, "EXEC": true, "EXIT": true,
	"EXTEND": true, "EXTERNAL": true,
	"FALSE": true, "FD": true, "FILE": true, "FILE-CONTROL": true, "FILLER": true,
	"FINAL": true, "FIRST": true, "FOOTING": true, "FOR": true, "FROM": true,
	"FUNCTION": true,
	"GENERATE": true, "GIVING": true, "GLOBAL": true, "GO": true, "GOBACK": true,
	"GREATER": true, "GROUP": true,
	"HEADING": true, "HIGH-VALUE": true, "HIGH-VALUES": true,
	"I-O": true, "I-O-CONTROL": true, "IDENTIFICATION": true, "IF": true, "IN": true,
	"INDEX": true, "INDEXED": true, "INDICATE": true, "INITIAL": true, "INITIALIZE": true,
	"INITIATE": true, "INPUT": true, "INPUT-OUTPUT": true, "INSPECT": true,
	"INSTALLATION": true, "INTO": true, "INVALID": true, "IS": true,
	"JUST": true, "JUSTIFIED": true,
	"KEY": true,
	"LABEL": true, "LAST": true, "LEADING": true, "LEFT": true, "LENGTH": true,
	"LESS": true, "LIMIT": true, "LIMITS": true, "LINAGE": true, "LINAGE-COUNTER": true,
	"LINE": true, "LINE-COUNTER": true, "LINES": true, "LINKAGE": true, "LOCK": true,
	"LOW-VALUE": true, "LOW-VALUES": true,
	"MEMORY": true, "MERGE": true, "MESSAGE": true, "MODE": true, "MODULES": true,
	"MOVE": true, "MULTIPLE": true, "MULTIPLY": true,
	"NATIVE": true, "NEGATIVE": true, "NEXT": true, "NO": true, "NOT": true,
	"NULL": true, "NUMBER": true, "NUMERIC": true, "NUMERIC-EDITED": true,
	"OBJECT-COMPUTER": true, "OCCURS": true, "OF": true, "OFF": true, "OMITTED": true,
	"ON": true, "OPEN": true, "OPTIONAL": true, "OR": true, "ORDER": true,
	"ORGANIZATION": true, "OTHER": true, "OUTPUT": true, "OVERFLOW": true,
	"PACKED-DECIMAL": true, "PADDING": true, "PAGE": true, "PAGE-COUNTER": true,
	"PERFORM": true, "PIC": true, "PICTURE": true, "PLUS": true, "POINTER": true,
	"POSITION": true, "POSITIVE": true, "PROCEDURE": true, "PROCEDURES": true,
	"PROCEED": true, "PROGRAM": true, "PROGRAM-ID": true, "PURGE": true,
	"QUEUE": true, "QUOTE": true, "QUOTES": true,
	"RANDOM": true, "RD": true, "READ": true, "RECEIVE": true, "RECORD": true,
	"RECORDS": true, "REDEFINES": true, "REEL": true, "REFERENCE": true,
	"REFERENCES": true, "RELATIVE": true, "RELEASE": true, "REMAINDER": true,
	"REMOVAL": true, "RENAMES": true, "REPLACE": true, "REPLACING": true, "REPORT": true,
	"REPORTING": true, "REPORTS": true, "RERUN": true, "RESERVE": true, "RESET": true,
	"RETURN": true, "RETURN-CODE": true, "RETURNING": true, "REVERSED": true,
	"REWIND": true, "REWRITE": true, "RIGHT": true, "ROUNDED": true, "RUN": true,
	"SAME": true, "SD": true, "SEARCH": true, "SECTION": true, "SECURE": true,
	"SECURITY": true, "SEGMENT-LIMIT": true, "SELECT": true, "SEND": true,
	"SENTENCE": true, "SEPARATE": true, "SEQUENCE": true, "SEQUENTIAL": true,
	"SET": true, "SIGN": true, "SIZE": true, "SORT": true, "SORT-MERGE": true,
	"SOURCE": true, "SOURCE-COMPUTER": true, "SPACE": true, "SPACES": true,
	"SPECIAL-NAMES": true, "STANDARD": true, "STANDARD-1": true, "STANDARD-2": true,
	"START": true, "STATUS": true, "STOP": true, "STRING": true,
	"SUB-QUEUE": true, "SUBTRACT": true, "SYMBOLIC": true, "SYNC": true,
	"SYNCHRONIZED": true,
	"TABLE": true, "TALLYING": true, "TAPE": true, "TERMINAL": true, "TERMINATE": true,
	"TEST": true, "TEXT": true, "THAN": true, "THEN": true, "THROUGH": true, "THRU": true,
	"TIME": true, "TIMES": true, "TITLE": true, "TO": true, "TOP": true, "TRAILING": true,
	"TRUE": true, "TYPE": true,
	"UNIT": true, "UNSTRING": true, "UNTIL": true, "UP": true, "UPON": true, "USAGE": true,
	"USE": true, "USING": true,
	"VALUE": true, "VALUES": true, "VARYING": true,
	"WHEN": true, "WHEN-COMPILED": true, "WITH": true, "WORDS": true, "WORKING-STORAGE": true,
	"WRITE": true,
	"ZERO": true, "ZEROES": true, "ZEROS": true,
}

// cobolWordPattern matches COBOL identifiers (letters, digits, hyphens).
var cobolWordPattern = regexp.MustCompile(`[A-Z][A-Z0-9-]*[A-Z0-9]|[A-Z]`)

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
	for _, prefix := range ctx.prefixings {
		ctx.buffer = cobolWordPattern.ReplaceAllStringFunc(ctx.buffer, func(word string) string {
			upper := strings.ToUpper(word)
			if cobolKeywords[upper] {
				return word
			}
			// Don't prefix if it already starts with the prefix
			if strings.HasPrefix(upper, strings.ToUpper(prefix)+"-") ||
				strings.HasPrefix(upper, strings.ToUpper(prefix)) {
				return word
			}
			return prefix + "-" + word
		})
	}
}

func (ctx *Context) Read() string {
	return ctx.buffer
}

func (ctx *Context) Write(s string) {
	ctx.buffer += s
}
