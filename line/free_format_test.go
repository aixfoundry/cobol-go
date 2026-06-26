package line

import (
	"strings"
	"testing"

	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
)

// ---------------------------------------------------------------------------
// NewLine tests — FREE format
// ---------------------------------------------------------------------------

func TestNewLineFreeFormatNormal(t *testing.T) {
	// In FREE format, indicator is column 1, everything else is source text.
	l := NewLine(" DISPLAY 'HELLO'.", 0, format.FREE, format.ANSI85)
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
	if l.Indicator != " " {
		t.Errorf("Indicator = %q, want space", l.Indicator)
	}
	// ContentA must be empty in FREE format
	if l.ContentA != "" {
		t.Errorf("ContentA = %q, want empty", l.ContentA)
	}
	if !strings.Contains(l.ContentB, "DISPLAY") {
		t.Errorf("ContentB = %q, want it to contain DISPLAY", l.ContentB)
	}
	if l.Comment != "" {
		t.Errorf("Comment = %q, want empty (no comment area in FREE)", l.Comment)
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
	if l.Indicator != "*" {
		t.Errorf("Indicator = %q, want *", l.Indicator)
	}
}

func TestNewLineFreeFormatDebug(t *testing.T) {
	l := NewLine("D   DEBUG-LINE", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != DEBUG {
		t.Errorf("Type = %v, want DEBUG", l.Type)
	}
	if l.Indicator != "D" {
		t.Errorf("Indicator = %q, want D", l.Indicator)
	}
}

func TestNewLineFreeFormatContinuation(t *testing.T) {
	l := NewLine("-   continued text", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != CONTINUATION {
		t.Errorf("Type = %v, want CONTINUATION", l.Type)
	}
	if l.Indicator != "-" {
		t.Errorf("Indicator = %q, want -", l.Indicator)
	}
}

func TestNewLineFreeFormatCompilerDirective(t *testing.T) {
	l := NewLine("$SET SOURCEFORMAT\"FREE\"", 0, format.FREE, format.ANSI85)
	if l == nil {
		t.Fatal("NewLine returned nil")
	}
	if l.Type != COMPILER_DIRECTIVE {
		t.Errorf("Type = %v, want COMPILER_DIRECTIVE", l.Type)
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
		ContentA: "",  // ContentA is always empty in FREE
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
	// Simulate a small FREE format COBOL program.
	src := " IDENTIFICATION DIVISION.\n PROGRAM-ID. HELLO.\n PROCEDURE DIVISION.\n DISPLAY 'HELLO FREE'.\n"
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
	if !strings.Contains(code, "IDENTIFICATION") {
		t.Error("combined code should contain IDENTIFICATION")
	}
	if !strings.Contains(code, "PROGRAM-ID") {
		t.Error("combined code should contain PROGRAM-ID")
	}
	if !strings.Contains(code, "HELLO FREE") {
		t.Error("combined code should contain HELLO FREE")
	}
}

func TestCombineFreeFormatCommentHandling(t *testing.T) {
	src := "*> Top comment\n IDENTIFICATION DIVISION.\n PROGRAM-ID. HELLO.\n"
	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(src), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}

	code := Combine(ll)
	if !strings.Contains(code, "*>") {
		t.Error("comment tag should be preserved")
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
	// A minimal valid COBOL program in free-format style.
	// Note: column 1 is the indicator (space = normal line).
	input := strings.Join([]string{
		" IDENTIFICATION DIVISION.",
		" PROGRAM-ID. HELLO.",
		" PROCEDURE DIVISION.",
		" DISPLAY 'HELLO FROM FREE FORMAT'.",
		" STOP RUN.",
	}, "\n")

	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(input), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}

	code := Combine(ll)

	// Verify all the key pieces survived the round-trip.
	checks := []string{
		"IDENTIFICATION",
		"DIVISION",
		"PROGRAM-ID",
		"HELLO",
		"PROCEDURE",
		"DISPLAY",
		"HELLO FROM FREE FORMAT",
		"STOP RUN",
	}
	for _, c := range checks {
		if !strings.Contains(code, c) {
			t.Errorf("round-trip lost %q", c)
		}
	}
}

// ---------------------------------------------------------------------------
// FREE format continuation line test
// ---------------------------------------------------------------------------

func TestFreeFormatContinuation(t *testing.T) {
	// Simulate a continued string literal in FREE format.
	src := " A = \"HELLO\n- WORLD\"."
	opts := options.NewOptions().SetFormat(format.FREE)
	ll := NewLinkedLine(strings.NewReader(src), opts)
	if ll == nil {
		t.Fatal("NewLinkedLine returned nil")
	}

	// If line 1 ends with open quote, line 2 (continuation) should be handled.
	code := Combine(ll)
	if !strings.Contains(code, "HELLO") || !strings.Contains(code, "WORLD") {
		t.Errorf("continuation may have lost literal content: %q", code)
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
// FREE vs FIXED comparison: same source, different parsing
// ---------------------------------------------------------------------------

func TestFreeVsFixedParsing(t *testing.T) {
	// In FREE format, column 1 is the indicator.
	// In FIXED format, column 7 is the indicator.
	src := " DISPLAY X."

	freeLine := NewLine(src, 0, format.FREE, format.ANSI85)
	if freeLine == nil {
		t.Fatal("FREE NewLine returned nil")
	}
	if freeLine.Indicator != " " {
		t.Errorf("FREE indicator = %q, want space", freeLine.Indicator)
	}

	// In FIXED, the same text has space at col 7 (also valid normal line).
	fixedLine := NewLine("       DISPLAY X.", 0, format.FIXED, format.ANSI85)
	if fixedLine == nil {
		t.Fatal("FIXED NewLine returned nil")
	}
	if fixedLine.Type != NORMAL {
		t.Errorf("FIXED Type = %v, want NORMAL", fixedLine.Type)
	}
}
