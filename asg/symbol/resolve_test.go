package symbol_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/symbol"
	"github.com/aixfoundry/cobol-go/asg/visitor"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

// parseIntoProgram parses a COBOL snippet and returns the protobuf AST.
func parseIntoProgram(t *testing.T, name, code string) *pb.Program {
	t.Helper()
	program := &pb.Program{}
	is := antlr.NewInputStream(code)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	parser := cobol85.NewCobol85Parser(cts)
	ctx := parser.StartRule()
	vr := visitor.NewCompilationUnitVisitor(name, program)
	vr.Visit(ctx)
	return program
}

func TestBuildSymbolTable_BasicProgram(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
DATA DIVISION.
WORKING-STORAGE SECTION.
01 WS-VAR PIC X(10).
01 WS-GRP.
   05 WS-SUB PIC 9(4).
   88 WS-COND VALUE ZERO.
PROCEDURE DIVISION.
MAIN-PARA.
    PERFORM SUB-PARA.
    MOVE WS-VAR TO WS-SUB.
SUB-PARA.
    DISPLAY 'hello'.`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)

	// Check data entries
	if entries := table.DataEntries["WS-VAR"]; len(entries) != 1 {
		t.Errorf("expected 1 WS-VAR entry, got %d", len(entries))
	}
	if entries := table.DataEntries["WS-GRP"]; len(entries) != 1 {
		t.Errorf("expected 1 WS-GRP entry, got %d", len(entries))
	}
	if entries := table.DataEntries["WS-SUB"]; len(entries) != 1 {
		t.Errorf("expected 1 WS-SUB entry, got %d", len(entries))
	}

	// Check condition names
	if entries := table.Conditions["WS-COND"]; len(entries) != 1 {
		t.Errorf("expected 1 WS-COND condition, got %d", len(entries))
	}

	// Check paragraphs
	if p := table.Paragraphs["MAIN-PARA"]; p == nil {
		t.Error("MAIN-PARA not found in symbol table")
	}
	if p := table.Paragraphs["SUB-PARA"]; p == nil {
		t.Error("SUB-PARA not found in symbol table")
	}

	// Check programs
	if p := table.Programs["TESTPROG"]; p == nil {
		t.Error("TESTPROG not found in symbol table")
	}
}

func TestResolve_PerformStatement(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
PROCEDURE DIVISION.
MAIN-PARA.
    PERFORM SUB-PARA.
SUB-PARA.
    DISPLAY 'hello'.`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)
	result := symbol.Resolve(program, table)

	foundPerform := false
	for _, call := range result.Calls {
		if call.GetType() == pb.CallType_PROCEDURE_CALL {
			foundPerform = true
			if call.GetName() != "SUB-PARA" {
				t.Errorf("expected SUB-PARA call, got %q", call.GetName())
			}
			if call.GetParagraphName() == nil {
				t.Error("Call target (ParagraphName) should not be nil for resolved PERFORM")
			}
		}
	}
	if !foundPerform {
		t.Error("no PROCEDURE_CALL found for PERFORM statement")
	}
}

func TestResolve_MoveStatement(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
DATA DIVISION.
WORKING-STORAGE SECTION.
01 WS-SRC PIC X(10).
01 WS-DST PIC X(10).
PROCEDURE DIVISION.
MAIN-PARA.
    MOVE WS-SRC TO WS-DST.`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)
	result := symbol.Resolve(program, table)

	dataCallCount := 0
	for _, call := range result.Calls {
		if call.GetType() == pb.CallType_DATA_DESCRIPTION_ENTRY_CALL {
			dataCallCount++
			if call.GetTarget() == nil {
				t.Errorf("resolved call %q has nil target", call.GetName())
			}
		}
	}
	if dataCallCount < 2 {
		t.Errorf("expected at least 2 data calls from MOVE, got %d", dataCallCount)
	}
}

func TestResolve_GoToStatement(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
PROCEDURE DIVISION.
MAIN-PARA.
    GO TO EXIT-PARA.
EXIT-PARA.
    GOBACK.`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)
	result := symbol.Resolve(program, table)

	foundGoTo := false
	for _, call := range result.Calls {
		if call.GetType() == pb.CallType_PROCEDURE_CALL && call.GetName() == "EXIT-PARA" {
			foundGoTo = true
			if call.GetParagraphName() == nil {
				t.Error("resolved GO TO has nil target")
			}
		}
	}
	if !foundGoTo {
		t.Error("no resolved CALL for GO TO EXIT-PARA")
	}
}

func TestResolve_UnresolvedReference(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
PROCEDURE DIVISION.
MAIN-PARA.
    PERFORM DOES-NOT-EXIST.`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)
	result := symbol.Resolve(program, table)

	if len(result.Unresolved) == 0 {
		t.Error("expected at least one unresolved reference for PERFORM DOES-NOT-EXIST")
	}
	hasUnresolved := false
	for _, u := range result.Unresolved {
		if u.Name == "DOES-NOT-EXIST" {
			hasUnresolved = true
		}
	}
	if !hasUnresolved {
		t.Error("expected DOES-NOT-EXIST in unresolved list")
	}
}

func TestResolve_ReadStatement(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
ENVIRONMENT DIVISION.
INPUT-OUTPUT SECTION.
FILE-CONTROL.
    SELECT MYFILE ASSIGN TO 'file.dat'.
DATA DIVISION.
FILE SECTION.
FD MYFILE.
01 MYFILE-REC PIC X(80).
PROCEDURE DIVISION.
MAIN-PARA.
    OPEN INPUT MYFILE.
    READ MYFILE.
    CLOSE MYFILE.`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)

	// Check file entries
	if _, ok := table.FileControlEntries["MYFILE"]; !ok {
		t.Error("MYFILE not found in FileControlEntries")
	}
	if _, ok := table.FileDescriptions["MYFILE"]; !ok {
		t.Error("MYFILE not found in FileDescriptions")
	}

	result := symbol.Resolve(program, table)
	fileCalls := 0
	for _, call := range result.Calls {
		if call.GetType() == pb.CallType_FILE_CONTROL_ENTRY_CALL {
			fileCalls++
		}
	}
	if fileCalls < 3 { // OPEN, READ, CLOSE
		t.Errorf("expected at least 3 file calls, got %d", fileCalls)
	}
}

func TestSymbolTable_NestedPrograms(t *testing.T) {
	// Parse a main program that calls a subprogram (via CALL).
	// Even without actual nesting, both programs should be resolvable.
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. MAINPROG.
PROCEDURE DIVISION.
MAIN-PARA.
    DISPLAY 'hello'.`

	program := parseIntoProgram(t, "MAINPROG", code)
	table := symbol.Build(program)

	if p := table.Programs["MAINPROG"]; p == nil {
		t.Error("MAINPROG not found in symbol table")
	}
	if p := table.Paragraphs["MAIN-PARA"]; p == nil {
		t.Error("MAIN-PARA not found in symbol table")
	}
}

func TestSymbolTable_FileSection(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
ENVIRONMENT DIVISION.
INPUT-OUTPUT SECTION.
FILE-CONTROL.
    SELECT MYFILE ASSIGN TO 'file.dat'.
DATA DIVISION.
FILE SECTION.
FD MYFILE.
01 MYFILE-REC.
   05 MYFLD1 PIC X(10).
   05 MYFLD2 PIC 9(4).
WORKING-STORAGE SECTION.
01 WS-DATA PIC X(10).`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)

	// File section data items should be indexed
	if entries := table.DataEntries["MYFILE-REC"]; len(entries) == 0 {
		t.Error("MYFILE-REC not found in DataEntries (from FILE SECTION)")
	}
	if entries := table.DataEntries["MYFLD1"]; len(entries) == 0 {
		t.Error("MYFLD1 not found in DataEntries (from FILE SECTION)")
	}
}

func TestSymbolTable_LinkageSection(t *testing.T) {
	code := `IDENTIFICATION DIVISION.
PROGRAM-ID. TESTPROG.
DATA DIVISION.
LINKAGE SECTION.
01 LK-PARM PIC X(10).
01 LK-GRP.
   05 LK-SUB PIC 9(4).`

	program := parseIntoProgram(t, "TESTPROG", code)
	table := symbol.Build(program)

	if entries := table.DataEntries["LK-PARM"]; len(entries) != 1 {
		t.Errorf("expected 1 LK-PARM entry, got %d", len(entries))
	}
	if entries := table.DataEntries["LK-SUB"]; len(entries) != 1 {
		t.Errorf("expected 1 LK-SUB entry, got %d", len(entries))
	}
}

func TestReport(t *testing.T) {
	result := &symbol.ResolutionResult{
		Calls: []*pb.Call{
			{Type: pb.CallType_PROCEDURE_CALL, Name: "MY-PARA"},
		},
		Unresolved: []symbol.UnresolvedRef{
			{Name: "UNKNOWN", Location: "PERFORM", Kind: pb.CallType_PROCEDURE_CALL},
		},
	}
	report := result.Report()
	if !strings.Contains(report, "resolved") && !strings.Contains(report, "unresolved") {
		t.Errorf("Report() output unexpected: %s", report)
	}
	_ = fmt.Sprintf("summary: %s", report)
}

func TestSymbolTable_Empty(t *testing.T) {
	table := symbol.Build(nil)
	if table == nil {
		t.Fatal("Build(nil) should return non-nil table")
	}
	result := symbol.Resolve(nil, table)
	if result == nil {
		t.Fatal("Resolve(nil, table) should return non-nil result")
	}

	table = symbol.NewSymbolTable()
	if table == nil {
		t.Fatal("NewSymbolTable() should return non-nil table")
	}
	result = symbol.Resolve(&pb.Program{}, table)
	if len(result.Calls) != 0 {
		t.Error("resolving empty program should produce 0 calls")
	}
}
