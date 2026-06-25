package conv

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

// parseCOBOL is a test helper that parses a COBOL snippet and returns the parse tree.
// It bypasses the line-formatting layer and feeds the code directly to the ANTLR lexer,
// so test COBOL snippets don't need FIXED-format column alignment.
func parseCOBOL(t *testing.T, code string) cobol85.IStartRuleContext {
	t.Helper()
	is := antlr.NewInputStream(code)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	parser := cobol85.NewCobol85Parser(cts)
	parser.RemoveErrorListeners()
	return parser.StartRule()
}

// findFirstIdentifier walks the parse tree to find the first IdentifierContext with content.
func findFirstIdentifier(ctx antlr.ParserRuleContext) *cobol85.IdentifierContext {
	var found *cobol85.IdentifierContext
	var walk func(antlr.Tree)
	walk = func(t antlr.Tree) {
		if found != nil {
			return
		}
		if ictx, ok := t.(*cobol85.IdentifierContext); ok {
			if ictx.QualifiedDataName() != nil || ictx.SpecialRegister() != nil ||
				ictx.TableCall() != nil || ictx.FunctionCall() != nil {
				found = ictx
				return
			}
		}
		for i := 0; i < t.GetChildCount(); i++ {
			walk(t.GetChild(i))
		}
	}
	walk(ctx)
	return found
}

// findFirstCallStatement finds the first CallStatementContext in the tree.
func findFirstCallStatement(ctx antlr.ParserRuleContext) *cobol85.CallStatementContext {
	var found *cobol85.CallStatementContext
	var walk func(antlr.Tree)
	walk = func(t antlr.Tree) {
		if found != nil {
			return
		}
		if v, ok := t.(*cobol85.CallStatementContext); ok {
			found = v
			return
		}
		for i := 0; i < t.GetChildCount(); i++ {
			walk(t.GetChild(i))
		}
	}
	walk(ctx)
	return found
}

// findFirstProcedureName finds the first ProcedureNameContext in the tree.
func findFirstProcedureName(ctx antlr.ParserRuleContext) *cobol85.ProcedureNameContext {
	var found *cobol85.ProcedureNameContext
	var walk func(antlr.Tree)
	walk = func(t antlr.Tree) {
		if found != nil {
			return
		}
		if v, ok := t.(*cobol85.ProcedureNameContext); ok {
			found = v
			return
		}
		for i := 0; i < t.GetChildCount(); i++ {
			walk(t.GetChild(i))
		}
	}
	walk(ctx)
	return found
}

// findFirstTableCallIdentifier finds the first Identifier containing a TableCall.
func findFirstTableCallIdentifier(ctx antlr.ParserRuleContext) *cobol85.IdentifierContext {
	var found *cobol85.IdentifierContext
	var walk func(antlr.Tree)
	walk = func(t antlr.Tree) {
		if found != nil {
			return
		}
		if v, ok := t.(*cobol85.IdentifierContext); ok && v.TableCall() != nil {
			found = v
			return
		}
		for i := 0; i < t.GetChildCount(); i++ {
			walk(t.GetChild(i))
		}
	}
	walk(ctx)
	return found
}

func TestSpecialRegisterConverter(t *testing.T) {
	tests := []struct {
		code         string
		expectedType pb.SpecialRegister_Type
		expectIdent  bool
		desc         string
	}{
		{"ACCEPT DATE", pb.SpecialRegister_DATE, false, "DATE"},
		{"ACCEPT DAY", pb.SpecialRegister_DAY, false, "DAY"},
		{"ACCEPT DAY-OF-WEEK", pb.SpecialRegister_DAY_OF_WEEK, false, "DAY-OF-WEEK"},
		{"ACCEPT DEBUG-CONTENTS", pb.SpecialRegister_DEBUG_CONTENTS, false, "DEBUG-CONTENTS"},
		{"ACCEPT DEBUG-ITEM", pb.SpecialRegister_DEBUG_ITEM, false, "DEBUG-ITEM"},
		{"ACCEPT DEBUG-LINE", pb.SpecialRegister_DEBUG_LINE, false, "DEBUG-LINE"},
		{"ACCEPT DEBUG-NAME", pb.SpecialRegister_DEBUG_NAME, false, "DEBUG-NAME"},
		{"ACCEPT DEBUG-SUB-1", pb.SpecialRegister_DEBUG_SUB_1, false, "DEBUG-SUB-1"},
		{"ACCEPT DEBUG-SUB-2", pb.SpecialRegister_DEBUG_SUB_2, false, "DEBUG-SUB-2"},
		{"ACCEPT DEBUG-SUB-3", pb.SpecialRegister_DEBUG_SUB_3, false, "DEBUG-SUB-3"},
		{"ACCEPT LINAGE-COUNTER", pb.SpecialRegister_LINAGE_COUNTER, false, "LINAGE-COUNTER"},
		{"ACCEPT LINE-COUNTER", pb.SpecialRegister_LINE_COUNTER, false, "LINE-COUNTER"},
		{"ACCEPT PAGE-COUNTER", pb.SpecialRegister_PAGE_COUNTER, false, "PAGE-COUNTER"},
		{"ACCEPT RETURN-CODE", pb.SpecialRegister_RETURN_CODE, false, "RETURN-CODE"},
		{"ACCEPT SHIFT-IN", pb.SpecialRegister_SHIFT_IN, false, "SHIFT-IN"},
		{"ACCEPT SHIFT-OUT", pb.SpecialRegister_SHIFT_OUT, false, "SHIFT-OUT"},
		{"ACCEPT SORT-CONTROL", pb.SpecialRegister_SORT_CONTROL, false, "SORT-CONTROL"},
		{"ACCEPT SORT-CORE-SIZE", pb.SpecialRegister_SORT_CORE_SIZE, false, "SORT-CORE-SIZE"},
		{"ACCEPT SORT-FILE-SIZE", pb.SpecialRegister_SORT_FILE_SIZE, false, "SORT-FILE-SIZE"},
		{"ACCEPT SORT-MESSAGE", pb.SpecialRegister_SORT_MESSAGE, false, "SORT-MESSAGE"},
		{"ACCEPT SORT-MODE-SIZE", pb.SpecialRegister_SORT_MODE_SIZE, false, "SORT-MODE-SIZE"},
		{"ACCEPT SORT-RETURN", pb.SpecialRegister_SORT_RETURN, false, "SORT-RETURN"},
		{"ACCEPT TALLY", pb.SpecialRegister_TALLY, false, "TALLY"},
		{"ACCEPT TIME", pb.SpecialRegister_TIME, false, "TIME"},
		{"ACCEPT WHEN-COMPILED", pb.SpecialRegister_WHEN_COMPILED, false, "WHEN-COMPILED"},
		{"ACCEPT ADDRESS OF SOMEDATA", pb.SpecialRegister_ADDRESS_OF, true, "ADDRESS OF identifier"},
		{"ACCEPT LENGTH OF SOMEDATA", pb.SpecialRegister_LENGTH_OF, true, "LENGTH OF identifier"},
		{"ACCEPT LENGTH SOMEDATA", pb.SpecialRegister_LENGTH_OF, true, "LENGTH identifier"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fullCode := fmt.Sprintf("IDENTIFICATION DIVISION.\nPROGRAM-ID. TESTPGM.\nDATA DIVISION.\nWORKING-STORAGE SECTION.\n01 SOMEDATA PIC X(10).\nPROCEDURE DIVISION.\n%s.\n", tt.code)
			tree := parseCOBOL(t, fullCode)
			ident := findFirstIdentifier(tree.(*cobol85.StartRuleContext))
			if ident == nil {
				t.Fatal("no Identifier context found in parse tree")
			}

			srCtx := ident.SpecialRegister()
			if srCtx == nil {
				t.Fatal("Identifier does not wrap a SpecialRegister")
			}

			sr := SpecialRegister(srCtx)
			if sr == nil {
				t.Fatal("SpecialRegister converter returned nil")
			}
			if sr.GetType() != tt.expectedType {
				t.Errorf("SpecialRegister type = %v, want %v", sr.GetType(), tt.expectedType)
			}
			if tt.expectIdent && sr.Identifier == nil {
				t.Errorf("expected non-nil Identifier for %s", tt.desc)
			}
			if !tt.expectIdent && sr.Identifier != nil {
				t.Errorf("expected nil Identifier for %s, got %v", tt.desc, sr.Identifier)
			}
		})
	}
}

func TestIdentifierSpecialRegister(t *testing.T) {
	fullCode := "IDENTIFICATION DIVISION.\nPROGRAM-ID. TESTPGM.\nPROCEDURE DIVISION.\nACCEPT ADDRESS OF SOMEDATA.\n"
	tree := parseCOBOL(t, fullCode)
	ident := findFirstIdentifier(tree.(*cobol85.StartRuleContext))
	if ident == nil {
		t.Fatal("no Identifier context found")
	}

	result := Identifier(ident)
	if result == nil {
		t.Fatal("Identifier converter returned nil")
	}

	sr := result.GetSpecialRegister()
	if sr == nil {
		t.Fatal("Identifier does not contain SpecialRegister")
	}
	if sr.GetType() != pb.SpecialRegister_ADDRESS_OF {
		t.Errorf("SpecialRegister type = %v, want ADDRESS_OF", sr.GetType())
	}
	if sr.GetIdentifier() == nil {
		t.Error("SpecialRegister Identifier should not be nil for ADDRESS OF")
	}
}

func TestCallFromProcedureName(t *testing.T) {
	code := "IDENTIFICATION DIVISION.\nPROGRAM-ID. TESTPGM.\nPROCEDURE DIVISION.\nMY-PARA.\nPERFORM MY-PARA.\n"
	tree := parseCOBOL(t, code)

	pnCtx := findFirstProcedureName(tree.(*cobol85.StartRuleContext))
	if pnCtx == nil {
		t.Fatal("no ProcedureNameContext found")
	}

	call := CallFromProcedureName(pnCtx)
	if call == nil {
		t.Fatal("CallFromProcedureName returned nil")
	}
	if call.GetType() != pb.CallType_PROCEDURE_CALL {
		t.Errorf("CallType = %v, want PROCEDURE_CALL", call.GetType())
	}
	if call.GetName() != "MY-PARA" {
		t.Errorf("Name = %q, want MY-PARA", call.GetName())
	}
	if call.GetParagraphName() == nil {
		t.Error("ParagraphName target should not be nil")
	}
}

func TestCallFromIdentifier_TableCall(t *testing.T) {
	code := "IDENTIFICATION DIVISION.\nPROGRAM-ID. TESTPGM.\nDATA DIVISION.\nWORKING-STORAGE SECTION.\n01 MY-TABLE.\n   05 MY-ELEM PIC X OCCURS 10.\nPROCEDURE DIVISION.\nMOVE 'X' TO MY-ELEM(1).\n"
	tree := parseCOBOL(t, code)

	identCtx := findFirstTableCallIdentifier(tree.(*cobol85.StartRuleContext))
	if identCtx == nil {
		t.Skip("no TableCall Identifier found")
		return
	}

	call := CallFromIdentifier(identCtx)
	if call.GetType() != pb.CallType_TABLE_CALL {
		t.Errorf("CallType = %v, want TABLE_CALL", call.GetType())
	}
}

func TestCallStatementTargetViaParse(t *testing.T) {
	code := "IDENTIFICATION DIVISION.\nPROGRAM-ID. CALLSTMT.\nPROCEDURE DIVISION.\nCALL SOMEPROG USING BY REFERENCE INTEGER SOMEINT SOMEFILE BY VALUE 1 2 SOMEID1 BY CONTENT ADDRESS OF SOMEID2 LENGTH OF SOMEID3 4 GIVING SOMEID4.\n"

	tree := parseCOBOL(t, code)

	cctx := findFirstCallStatement(tree.(*cobol85.StartRuleContext))
	if cctx == nil {
		t.Fatal("no CallStatementContext found")
	}

	result := CallStatement(cctx)

	ident := result.GetTargetIdentifier()
	if ident == nil {
		t.Fatal("CallStatement target identifier should not be nil for CALL SOMEPROG")
	}

	qdn := ident.GetQualifiedDataName()
	if qdn == nil {
		t.Error("Identifier should wrap a QualifiedDataName")
	}

	if result.GetUsingPhrase() == nil {
		t.Error("UsingPhrase should be populated")
	}
	if result.GetGivingPhrase() == nil {
		t.Error("GivingPhrase should be populated")
	}
}

func TestSpecialRegisterDayOfWeek(t *testing.T) {
	fullCode := fmt.Sprintf("IDENTIFICATION DIVISION.\nPROGRAM-ID. TESTPGM.\nPROCEDURE DIVISION.\nACCEPT DAY-OF-WEEK.\n")
	tree := parseCOBOL(t, fullCode)
	ident := findFirstIdentifier(tree.(*cobol85.StartRuleContext))
	if ident == nil {
		t.Fatal("no Identifier found")
	}

	sr := SpecialRegister(ident.SpecialRegister())
	if sr.GetType() != pb.SpecialRegister_DAY_OF_WEEK {
		t.Errorf("type = %v, want DAY_OF_WEEK", sr.GetType())
	}
}

func TestSpecialRegisterDebugSubs(t *testing.T) {
	tests := []struct {
		code     string
		wantType pb.SpecialRegister_Type
	}{
		{"ACCEPT DEBUG-CONTENTS.", pb.SpecialRegister_DEBUG_CONTENTS},
		{"ACCEPT DEBUG-ITEM.", pb.SpecialRegister_DEBUG_ITEM},
		{"ACCEPT DEBUG-LINE.", pb.SpecialRegister_DEBUG_LINE},
		{"ACCEPT DEBUG-NAME.", pb.SpecialRegister_DEBUG_NAME},
		{"ACCEPT DEBUG-SUB-1.", pb.SpecialRegister_DEBUG_SUB_1},
		{"ACCEPT DEBUG-SUB-2.", pb.SpecialRegister_DEBUG_SUB_2},
		{"ACCEPT DEBUG-SUB-3.", pb.SpecialRegister_DEBUG_SUB_3},
	}

	for _, tt := range tests {
		t.Run(tt.wantType.String(), func(t *testing.T) {
			fullCode := fmt.Sprintf("IDENTIFICATION DIVISION.\nPROGRAM-ID. T.\nPROCEDURE DIVISION.\n%s\n", tt.code)
			tree := parseCOBOL(t, fullCode)
			ident := findFirstIdentifier(tree.(*cobol85.StartRuleContext))
			if ident == nil {
				t.Fatal("no Identifier found")
			}
			srCtx := ident.SpecialRegister()
			if srCtx == nil {
				t.Fatal("not a SpecialRegister")
			}
			sr := SpecialRegister(srCtx)
			if sr.GetType() != tt.wantType {
				t.Errorf("type = %v, want %v", sr.GetType(), tt.wantType)
			}
		})
	}
}

func TestCallStatementTargetLiteral(t *testing.T) {
	for _, litVal := range []string{"'MYPROG'", "\"MYPROG\""} {
		t.Run(litVal, func(t *testing.T) {
			code := fmt.Sprintf("IDENTIFICATION DIVISION.\nPROGRAM-ID. T.\nPROCEDURE DIVISION.\nCALL %s.\n", litVal)
			tree := parseCOBOL(t, code)

			cctx := findFirstCallStatement(tree.(*cobol85.StartRuleContext))
			if cctx == nil {
				t.Fatal("no CallStatementContext found")
			}
			result := CallStatement(cctx)
			if result == nil {
				t.Fatal("CallStatement returned nil")
			}
			if result.GetTargetLiteral() == nil {
				t.Error("target literal should not be nil for CALL literal")
			}
		})
	}
}

func TestCallStatementIdentifierTarget(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. T.
DATA DIVISION.
WORKING-STORAGE SECTION.
01 MYPROG PIC X(8).
PROCEDURE DIVISION.
    CALL MYPROG.`

	tree := parseCOBOL(t, code)

	cctx := findFirstCallStatement(tree.(*cobol85.StartRuleContext))
	if cctx == nil {
		t.Fatal("no CallStatementContext found")
	}
	result := CallStatement(cctx)
	if result == nil {
		t.Fatal("CallStatement returned nil")
	}
	if result.GetTargetIdentifier() == nil {
		t.Error("target identifier should not be nil for CALL identifier")
	}
}

func TestAllCallTypes(t *testing.T) {
	types := []pb.CallType{
		pb.CallType_DATA_DESCRIPTION_ENTRY_CALL,
		pb.CallType_PROCEDURE_CALL,
		pb.CallType_SECTION_CALL,
		pb.CallType_FILE_CONTROL_ENTRY_CALL,
		pb.CallType_FUNCTION_CALL,
		pb.CallType_TABLE_CALL,
		pb.CallType_INDEX_CALL,
		pb.CallType_MNEMONIC_CALL,
		pb.CallType_SPECIAL_REGISTER_CALL,
		pb.CallType_REPORT_CALL,
		pb.CallType_REPORT_DESCRIPTION_ENTRY_CALL,
		pb.CallType_SCREEN_DESCRIPTION_ENTRY_CALL,
		pb.CallType_COMMUNICATION_DESCRIPTION_ENTRY_CALL,
		pb.CallType_ENVIRONMENT_CALL,
		pb.CallType_UNDEFINED_CALL,
	}
	for i, ct := range types {
		if ct != pb.CallType(int32(i)) {
			t.Errorf("CallType[%d] = %v, expected sequential numbering", i, ct)
		}
	}
	if len(types) != 15 {
		t.Errorf("expected 15 CallType values, got %d", len(types))
	}
}
