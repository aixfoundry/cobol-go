package line

import (
	"strings"
	"testing"

	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
)

// ---------------------------------------------------------------------------
// NewLine tests — FREE format
//
// COBOL 2002/2014 free format has NO indicator column. Line type is decided
// by prefix tokens:
//
//	*>   full-line (or inline) comment
//	D>>  debug line
//	>>   compiler directive (>>SOURCE FORMAT, >>DEFINE, ...)
//	$    Micro-Focus style directive (kept for compatibility)
//
// Everything else — including a line that starts with the letter D such as
// DISPLAY or DATA DIVISION — is normal source. There is no continuation
// indicator; "-" at column 1 is ordinary code.
// ---------------------------------------------------------------------------

func TestNewLineFreeFormatNormal(t *testing.T) {
	// Code begins at column 1: the defining feature of free format.
	l := NewLine("DISPLAY 'HELLO'.", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil for valid FREE format line")
	}
	if l.Format != format.FREE {
		t.Errorf("Format = %v, want FREE", l.Format)
	}
	if l.Type != NORMAL {
		t.Errorf("Type = %v, want NORMAL", l.Type)
	}
	if l.Sequence != "" {
		t.Errorf("Sequence = %q, want empty (no sequence area in FREE)", l.Sequence)
	}
	if l.Indicator != "" {
		t.Errorf("Indicator = %q, want empty (no indicator consumed for normal code)", l.Indicator)
	}
	// The whole line must survive as content — no leading character swallowed.
	if l.Content() != "DISPLAY 'HELLO'." {
		t.Errorf("Content = %q, want %q (first char must not be lost)", l.Content(), "DISPLAY 'HELLO'.")
	}
	if l.Comment != "" {
		t.Errorf("Comment = %q, want empty (no comment area in FREE)", l.Comment)
	}
}

// TestNewLineFreeFormatTopColumnStatements guards the critical regression:
// statements whose first letter is A/B/C/D must NOT be misread as indicator
// characters. Each must parse as NORMAL with its full text intact.
func TestNewLineFreeFormatTopColumnStatements(t *testing.T) {
	cases := []string{
		"DATA DIVISION.",
		"DISPLAY 'HI'.",
		"DIVISION.",
		"ACCEPT X.",
		"ADD 1 TO X.",
		"CALL 'SUB'.",
		"COMPUTE X = 1.",
		"MOVE A TO B.",
		"PERFORM UNTIL X > 0",
		"STOP RUN.",
	}
	for _, src := range cases {
		l := NewLine(src, 0, format.FREE, format.ANSI85)
		if l == nil {
			t.Errorf("NewLine(%q) returned nil", src)
			continue
		}
		if l.Type != NORMAL {
			t.Errorf("NewLine(%q): Type = %v, want NORMAL", src, l.Type)
		}
		if l.Indicator != "" {
			t.Errorf("NewLine(%q): Indicator = %q, want empty", src, l.Indicator)
		}
		if l.Content() != src {
			t.Errorf("NewLine(%q): Content = %q, want identical (no char lost)", src, l.Content())
		}
	}
}

func TestNewLineFreeFormatComment(t *testing.T) {
	l := NewLine("*> This is a comment", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != COMMENT {
		t.Errorf("Type = %v, want COMMENT", l.Type)
	}
	if l.Indicator != "*>" {
		t.Errorf("Indicator = %q, want \"*>\"", l.Indicator)
	}
}

func TestNewLineFreeFormatDebug(t *testing.T) {
	// Free-format debug uses the D>> prefix, not a bare D.
	l := NewLine("D>> DISPLAY 'DEBUG ONLY'", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != DEBUG {
		t.Errorf("Type = %v, want DEBUG", l.Type)
	}
	if l.Indicator != "D>>" {
		t.Errorf("Indicator = %q, want \"D>>\"", l.Indicator)
	}
}

func TestNewLineFreeFormatBareDIsNotDebug(t *testing.T) {
	// A bare D (e.g. DISPLAY) must not be treated as a debug indicator.
	l := NewLine("D", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != NORMAL {
		t.Errorf("Type = %v, want NORMAL (bare D is code, not debug)", l.Type)
	}
	if l.Content() != "D" {
		t.Errorf("Content = %q, want \"D\"", l.Content())
	}
}

func TestNewLineFreeFormatCompilerDirective(t *testing.T) {
	l := NewLine(">>SOURCE FORMAT FREE", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != COMPILER_DIRECTIVE {
		t.Errorf("Type = %v, want COMPILER_DIRECTIVE", l.Type)
	}
	if l.Indicator != ">>" {
		t.Errorf("Indicator = %q, want \">>\"", l.Indicator)
	}
}

func TestNewLineFreeFormatMicroFocusDirective(t *testing.T) {
	// $SET directives (Micro Focus style) are still recognized in FREE format.
	l := NewLine("$SET SOURCEFORMAT\"FREE\"", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != COMPILER_DIRECTIVE {
		t.Errorf("Type = %v, want COMPILER_DIRECTIVE", l.Type)
	}
}

func TestNewLineFreeFormatDashIsNotContinuation(t *testing.T) {
	// Free format does not use the fixed-format "-" continuation character.
	// A "-" at column 1 is ordinary source code.
	l := NewLine("-1 + 2", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != NORMAL {
		t.Errorf("Type = %v, want NORMAL (\"-\" is not a continuation in FREE)", l.Type)
	}
	if l.Content() != "-1 + 2" {
		t.Errorf("Content = %q, want \"-1 + 2\"", l.Content())
	}
}

func TestNewLineFreeFormatAmpersandIsContinuation(t *testing.T) {
	// A leading "&" marks a free-format continuation line (replaces the
	// fixed-format "-" indicator). The "&" is consumed as the indicator and
	// the remainder is the continuation content.
	l := NewLine("&World\" TO X.", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != CONTINUATION {
		t.Errorf("Type = %v, want CONTINUATION", l.Type)
	}
	if l.Indicator != "&" {
		t.Errorf("Indicator = %q, want \"&\"", l.Indicator)
	}
	if l.Content() != "World\" TO X." {
		t.Errorf("Content = %q, want the text following the \"&\"", l.Content())
	}
}

func TestCombineFreeFormatAmpersandLiteralContinuation(t *testing.T) {
	// A multi-line alphanumeric literal continued with a leading "&" must be
	// merged into a single literal with no char lost.
	src := "MOVE \"Hello\n&World\" TO X.\n"
	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(src), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}
	code := Combine(ll)
	// The two halves must be glued into one literal; no newline inside it.
	if !strings.Contains(code, "\"HelloWorld\"") {
		t.Errorf("literal not merged; got:\n%s", code)
	}
}

// ---------------------------------------------------------------------------
// SetContent tests — FREE format
// ---------------------------------------------------------------------------

func TestSetContentFreeFormat(t *testing.T) {
	l := &Line{Format: format.FREE}

	l.SetContent("HELLO WORLD")
	if l.ContentA != "" {
		t.Errorf("ContentA = %q, want empty", l.ContentA)
	}
	if l.ContentB != "HELLO WORLD" {
		t.Errorf("ContentB = %q, want %q", l.ContentB, "HELLO WORLD")
	}
}

func TestSetContentFreeFormatLong(t *testing.T) {
	l := &Line{Format: format.FREE}
	longContent := "DISPLAY 'THIS IS A VERY LONG LINE THAT WOULD NORMALLY SPLIT ACROSS AREA A AND B'"

	l.SetContent(longContent)
	if l.ContentA != "" {
		t.Errorf("ContentA = %q, want empty", l.ContentA)
	}
	if l.ContentB != longContent {
		t.Errorf("ContentB has wrong content")
	}
}

func TestSetContentFreeFormatShort(t *testing.T) {
	l := &Line{Format: format.FREE}

	l.SetContent("A")
	if l.ContentA != "" {
		t.Errorf("ContentA = %q, want empty", l.ContentA)
	}
	if l.ContentB != "A" {
		t.Errorf("ContentB = %q, want %q", l.ContentB, "A")
	}
}

func TestSetContentFixedUnchanged(t *testing.T) {
	// Ensure SetContent for FIXED format is not broken.
	l := &Line{Format: format.FIXED}

	l.SetContent("HELLO WORLD")
	if l.ContentA != "HELL" {
		t.Errorf("ContentA = %q, want first 4 chars", l.ContentA)
	}
	if l.ContentB != "O WORLD" {
		t.Errorf("ContentB = %q, want remainder", l.ContentB)
	}
}

// ---------------------------------------------------------------------------
// Content() tests — FREE format
// ---------------------------------------------------------------------------

func TestContentFreeFormat(t *testing.T) {
	l := &Line{
		Format:   format.FREE,
		ContentA: "", // ContentA is always empty in FREE
		ContentB: "PROCEDURE DIVISION.",
	}
	if got := l.Content(); got != "PROCEDURE DIVISION." {
		t.Errorf("Content() = %q, want %q", got, "PROCEDURE DIVISION.")
	}
}

// ---------------------------------------------------------------------------
// LinePrefix tests
// ---------------------------------------------------------------------------

func TestLinePrefixFree(t *testing.T) {
	if got := LinePrefix(format.FREE); got != "" {
		t.Errorf("LinePrefix(FREE) = %q, want empty (no sequence prefix)", got)
	}
}

func TestLinePrefixTandem(t *testing.T) {
	if got := LinePrefix(format.TANDEM); got != "" {
		t.Errorf("LinePrefix(TANDEM) = %q, want empty", got)
	}
}

func TestLinePrefixFixed(t *testing.T) {
	got := LinePrefix(format.FIXED)
	if len(got) != 6 {
		t.Errorf("LinePrefix(FIXED) length = %d, want 6", len(got))
	}
}

func TestLinePrefixVariable(t *testing.T) {
	got := LinePrefix(format.VARIABLE)
	if len(got) != 6 {
		t.Errorf("LinePrefix(VARIABLE) length = %d, want 6", len(got))
	}
}

// ---------------------------------------------------------------------------
// Combine tests — FREE format
// ---------------------------------------------------------------------------

func TestCombineFreeFormat(t *testing.T) {
	// Simulate a FREE format COBOL program where every line starts at column 1.
	src := "IDENTIFICATION DIVISION.\nPROGRAM-ID. HELLO.\nPROCEDURE DIVISION.\nDISPLAY 'HELLO FREE'.\n"
	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(src), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}

	code := Combine(ll)
	if code == "" {
		t.Fatal("Combine returned empty code")
	}
	// FREE format has no 6-space sequence prefix.
	if strings.Contains(code, "      ") {
		t.Error("FREE format should not have 6-space sequence prefix")
	}
	// Every statement must round-trip intact — no leading character lost.
	for _, want := range []string{
		"IDENTIFICATION",
		"PROGRAM-ID",
		"DISPLAY 'HELLO FREE'.",
	} {
		if !strings.Contains(code, want) {
			t.Errorf("combined code lost %q; got:\n%s", want, code)
		}
	}
}

func TestCombineFreeFormatCommentHandling(t *testing.T) {
	src := "*> Top comment\nIDENTIFICATION DIVISION.\nPROGRAM-ID. HELLO.\n"
	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(src), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}

	code := Combine(ll)
	// The comment marker must survive exactly once — the old bug produced "*> >".
	if !strings.Contains(code, "*>") {
		t.Errorf("comment tag should be preserved; got:\n%s", code)
	}
	if strings.Contains(code, "*> >") {
		t.Errorf("comment tag was duplicated as \"*> >\"; got:\n%s", code)
	}
}

// ---------------------------------------------------------------------------
// MultiLine tests — FREE format
// ---------------------------------------------------------------------------

func TestMultiLineFree(t *testing.T) {
	if format.FREE.MultiLine() {
		t.Error("FREE.MultiLine() should be false (no multi-line comment entries)")
	}
}

// ---------------------------------------------------------------------------
// String tests — FREE format
// ---------------------------------------------------------------------------

func TestFormatStringFree(t *testing.T) {
	if got := format.FREE.String(); got != "FREE" {
		t.Errorf("FREE.String() = %q, want %q", got, "FREE")
	}
}

// ---------------------------------------------------------------------------
// End-to-end: NewLinkedLine → Combine round-trip for FREE format
// ---------------------------------------------------------------------------

func TestFreeFormatRoundTrip(t *testing.T) {
	// A realistic FREE-format program: code at column 1 (no leading spaces).
	input := strings.Join([]string{
		"IDENTIFICATION DIVISION.",
		"PROGRAM-ID. HELLO.",
		"DATA DIVISION.",
		"PROCEDURE DIVISION.",
		"DISPLAY 'HELLO FROM FREE FORMAT'.",
		"ACCEPT X.",
		"STOP RUN.",
	}, "\n")

	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(input), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}

	code := Combine(ll)

	// Verify the key pieces survived the round-trip with NO lost characters.
	// These specifically include statements starting with A and D, which the
	// old col-1-indicator model corrupted.
	checks := []string{
		"IDENTIFICATION DIVISION.",
		"PROGRAM-ID. HELLO.",
		"DATA DIVISION.",
		"PROCEDURE DIVISION.",
		"DISPLAY 'HELLO FROM FREE FORMAT'.",
		"ACCEPT X.",
		"STOP RUN.",
	}
	for _, c := range checks {
		if !strings.Contains(code, c) {
			t.Errorf("round-trip lost or corrupted %q; got:\n%s", c, code)
		}
	}
}

// ---------------------------------------------------------------------------
// FREE format blank line test
// ---------------------------------------------------------------------------

func TestNewLineFreeFormatBlank(t *testing.T) {
	// An empty line in FREE format produces a Line with NORMAL type and
	// empty content — consistent with FIXED format behavior.
	l := NewLine("", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil for empty line (should produce empty Line, consistent with FIXED)")
	}
	if l.Type != NORMAL {
		t.Errorf("Type = %v, want NORMAL", l.Type)
	}
	if l.Content() != "" {
		t.Errorf("Content = %q, want empty", l.Content())
	}
}

// ---------------------------------------------------------------------------
// FREE vs FIXED comparison: same logical line, different parsing
// ---------------------------------------------------------------------------

func TestFreeVsFixedParsing(t *testing.T) {
	// In FREE format, column 1 is source text (no indicator).
	// In FIXED format, column 7 is the indicator.
	freeLine := NewLine("DISPLAY X.", 0, format.FREE, format.ANSI85)
	if freeLine == nil {
		t.Fatal("FREE NewLine returned nil")
	}
	if freeLine.Type != NORMAL {
		t.Errorf("FREE Type = %v, want NORMAL", freeLine.Type)
	}
	if freeLine.Content() != "DISPLAY X." {
		t.Errorf("FREE Content = %q, want \"DISPLAY X.\"", freeLine.Content())
	}

	fixedLine := NewLine("       DISPLAY X.", 0, format.FIXED, format.ANSI85)
	if fixedLine == nil {
		t.Fatal("FIXED NewLine returned nil")
	}
	if fixedLine.Type != NORMAL {
		t.Errorf("FIXED Type = %v, want NORMAL", fixedLine.Type)
	}
}
