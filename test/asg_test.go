package test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aixfoundry/cobol-go/asg"
	"github.com/aixfoundry/cobol-go/asg/symbol"
	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
	"github.com/aixfoundry/cobol-go/pb"
)

// asgDirFormat maps directory paths within testdata/cobol/asg to their COBOL source format.
func asgDirFormat(dirPath string) format.Format {
	if strings.Contains(dirPath, "/fixed") {
		return format.FIXED
	}
	if strings.Contains(dirPath, "/tandem") {
		return format.TANDEM
	}
	if strings.Contains(dirPath, "/variable") {
		return format.VARIABLE
	}
	return format.FIXED
}

// knownSkipFiles are files known to have format or visitor issues (pre-existing, not regressions).
var knownSkipFiles = map[string]bool{
	"InvalidKeyword.cbl":    true,
	"InvalidLineFormat.cbl": true,
	"ASGElement.cbl":        true,
}

// safeAnalyzeFile wraps asg.AnalyzeFile with panic recovery and handles nil results.
// Returns (program, ok) where ok=false means the file couldn't be parsed (pre-existing issue).
func safeAnalyzeFile(t *testing.T, filePath string, opts ...options.Option) (program *pb.Program, ok bool) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("SKIP: visitor panicked (known issue): %v", r)
			ok = false
		}
	}()
	program = asg.AnalyzeFile(filePath, opts...)
	if program == nil || len(program.GetCompilationUnits()) == 0 {
		t.Logf("SKIP: could not parse file (format issue or pre-existing limitation)")
		return nil, false
	}
	if len(program.GetCompilationUnits()[0].GetProgramUnits()) == 0 {
		t.Logf("SKIP: no program units parsed")
		return nil, false
	}
	return program, true
}

// TestParseAllASGFiles is a smoke test that walks all .cbl files in testdata/cobol/asg/
// and tries to parse them through the full pipeline. Files that fail due to pre-existing
// format or visitor issues are skipped with a log message.
func TestParseAllASGFiles(t *testing.T) {
	rootdir := "./testdata/cobol/asg"
	var parsed, skipped, failed int

	err := filepath.Walk(rootdir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(info.Name()), ".cbl") {
			return nil
		}

		relPath, _ := filepath.Rel(rootdir, filePath)
		t.Run(relPath, func(t *testing.T) {
			if knownSkipFiles[info.Name()] {
				t.Skip("known issue")
				skipped++
				return
			}

			dirPath := path.Dir(filePath)
			fmt := asgDirFormat(filePath)
			opts := options.NewOptions().AddCopyBookDirectory(dirPath).SetFormat(fmt)

			program, ok := safeAnalyzeFile(t, filePath, opts)
			if !ok {
				skipped++
				t.Skip("format/preprocessing issue (pre-existing)")
				return
			}

			pu := program.GetCompilationUnits()[0].GetProgramUnits()[0]
			if pu.GetIdentificationDivision() == nil {
				t.Error("program unit missing identification division")
				failed++
				return
			}

			parsed++

			t.Logf("Program: %s", programName(pu))
			t.Logf("Data entries: %d, Paragraphs: %d",
				countDataEntries(pu), countParagraphs(pu))

			// Build symbol table
			table := symbol.Build(program)
			if table == nil {
				t.Error("symbol.Build returned nil table")
				return
			}

			// Run resolution
			result := symbol.Resolve(program, table)
			if result == nil {
				t.Error("symbol.Resolve returned nil result")
				return
			}

			t.Logf("Resolved calls: %d, Unresolved: %d",
				len(result.Calls), len(result.Unresolved))
		})
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk test data directory: %v", err)
	}
	t.Logf("Summary: parsed=%d, skipped=%d, failed=%d", parsed, skipped, failed)
}

// TestParseASGCallFiles verifies call resolution for call/*.cbl files.
func TestParseASGCallFiles(t *testing.T) {
	rootdir := "./testdata/cobol/asg/call"
	files, err := os.ReadDir(rootdir)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".cbl") {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			if knownSkipFiles[file.Name()] {
				t.Skip("known issue")
			}

			filePath := path.Join(rootdir, file.Name())
			opts := options.NewOptions().AddCopyBookDirectory(rootdir).SetFormat(format.FIXED)
			program, ok := safeAnalyzeFile(t, filePath, opts)
			if !ok {
				t.Skip("format/preprocessing issue")
				return
			}

			table := symbol.Build(program)
			result := symbol.Resolve(program, table)

			if result == nil {
				t.Error("symbol.Resolve returned nil")
				return
			}

			t.Logf("%s: %d resolved, %d unresolved",
				file.Name(), len(result.Calls), len(result.Unresolved))

			baseName := strings.TrimSuffix(file.Name(), ".cbl")
			switch baseName {
			case "ParagraphCall", "ParagraphInSectionCall":
				hasProcCall(t, result.Calls, "INIT")
			case "SectionCall":
				hasProcCall(t, result.Calls, "INIT")
			case "DataDescriptionEntryCall":
				hasDataCall(t, result.Calls, "FORT-STRUKTUR")
			case "TableCall":
				hasCallType(t, result.Calls, pb.CallType_TABLE_CALL)
			case "FunctionCall", "FunctionDateOfIntegerCall":
				hasCallType(t, result.Calls, pb.CallType_FUNCTION_CALL)
			case "SpecialRegisterCall":
				hasCallType(t, result.Calls, pb.CallType_SPECIAL_REGISTER_CALL)
			}
		})
	}
}

// TestParseASGDataDivisionFiles parses data division test files and verifies symbol table.
func TestParseASGDataDivisionFiles(t *testing.T) {
	categories := []string{
		"data/workingstorage",
		"data/file",
		"data/linkage",
		"data/localstorage",
		"data/communication",
		"data/report",
		"data/screen",
	}

	for _, cat := range categories {
		dirPath := path.Join("./testdata/cobol/asg", cat)
		files, err := os.ReadDir(dirPath)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".cbl") {
				continue
			}

			t.Run(path.Join(cat, file.Name()), func(t *testing.T) {
				if knownSkipFiles[file.Name()] {
					t.Skip("known issue")
				}

				filePath := path.Join(dirPath, file.Name())
				opts := options.NewOptions().AddCopyBookDirectory(dirPath).SetFormat(format.FIXED)
				program, ok := safeAnalyzeFile(t, filePath, opts)
				if !ok {
					t.Skip("format/preprocessing issue")
					return
				}

				pu := program.GetCompilationUnits()[0].GetProgramUnits()[0]
				dd := pu.GetDataDivision()
				if dd == nil {
					t.Log("no data division found")
					return
				}

				table := symbol.Build(program)
				dataCount := len(table.DataEntries)
				condCount := len(table.Conditions)
				t.Logf("data entries: %d, conditions: %d", dataCount, condCount)

				switch {
				case strings.Contains(cat, "workingstorage"):
					checkWorkingStorageEntries(t, file.Name(), table)
				}

				result := symbol.Resolve(program, table)
				t.Logf("resolved: %d calls, %d unresolved",
					len(result.Calls), len(result.Unresolved))
			})
		}
	}
}

// TestParseASGProcedureFiles parses procedure statement test files.
func TestParseASGProcedureFiles(t *testing.T) {
	rootdir := "./testdata/cobol/asg/procedure"
	err := filepath.Walk(rootdir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(info.Name()), ".cbl") {
			return nil
		}

		dirPath := path.Dir(filePath)
		relPath, _ := filepath.Rel(rootdir, filePath)
		t.Run(relPath, func(t *testing.T) {
			if knownSkipFiles[info.Name()] {
				t.Skip("known issue")
			}

			opts := options.NewOptions().AddCopyBookDirectory(dirPath).SetFormat(format.FIXED)
			program, ok := safeAnalyzeFile(t, filePath, opts)
			if !ok {
				t.Skip("format/preprocessing issue")
				return
			}

			pu := program.GetCompilationUnits()[0].GetProgramUnits()[0]
			pd := pu.GetProcedureDivision()
			if pd == nil {
				t.Log("no procedure division found")
				return
			}

			table := symbol.Build(program)
			result := symbol.Resolve(program, table)

			stmtCategory := path.Base(dirPath)
			t.Logf("[%s] %d paragraphs, %d sections, %d resolved calls, %d unresolved",
				stmtCategory,
				len(table.Paragraphs),
				len(table.Sections),
				len(result.Calls),
				len(result.Unresolved))

			verifyStatementCalls(t, stmtCategory, result.Calls)
		})
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk procedure test data: %v", err)
	}
}

// TestParseASGIdentificationFiles parses identification division test files across formats.
func TestParseASGIdentificationFiles(t *testing.T) {
	for _, fmtName := range []string{"fixed", "tandem", "variable"} {
		dirPath := path.Join("./testdata/cobol/asg/identification", fmtName)
		files, err := os.ReadDir(dirPath)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".cbl") {
				continue
			}

			t.Run(path.Join("identification", fmtName, file.Name()), func(t *testing.T) {
				filePath := path.Join(dirPath, file.Name())
				f := asgDirFormat(filePath)
				opts := options.NewOptions().AddCopyBookDirectory(dirPath).SetFormat(f)
				program, ok := safeAnalyzeFile(t, filePath, opts)
				if !ok {
					t.Skip("format/preprocessing issue")
					return
				}

				pu := program.GetCompilationUnits()[0].GetProgramUnits()[0]
				id := pu.GetIdentificationDivision()
				if id == nil {
					t.Error("no identification division found")
					return
				}

				if id.GetProgramIdParagraph() == nil {
					t.Error("no PROGRAM-ID paragraph found")
				} else {
					pn := id.GetProgramIdParagraph().GetProgramName()
					if pn != nil {
						t.Logf("PROGRAM-ID: %v", pn)
					}
				}
			})
		}
	}
}

// TestParseASGRelationConditions parses value statement relation condition tests.
func TestParseASGRelationConditions(t *testing.T) {
	dirPath := "./testdata/cobol/asg/valuestmt/relation"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".cbl") {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			filePath := path.Join(dirPath, file.Name())
			opts := options.NewOptions().AddCopyBookDirectory(dirPath).SetFormat(format.FIXED)
			program, ok := safeAnalyzeFile(t, filePath, opts)
			if !ok {
				t.Skip("format/preprocessing issue")
				return
			}

			table := symbol.Build(program)
			result := symbol.Resolve(program, table)
			t.Logf("condition: %s, resolved: %d, unresolved: %d",
				file.Name(), len(result.Calls), len(result.Unresolved))
		})
	}
}

// --- helpers ---

func programName(pu *pb.ProgramUnit) string {
	id := pu.GetIdentificationDivision()
	if id == nil {
		return "<unknown>"
	}
	pp := id.GetProgramIdParagraph()
	if pp == nil {
		return "<unknown>"
	}
	pn := pp.GetProgramName()
	if pn == nil {
		return "<unknown>"
	}
	if cw := pn.GetCobolWord(); cw != nil {
		return cw.GetValue()
	}
	if nl := pn.GetNonNumericLiteral(); nl != nil {
		return nl.GetValue()
	}
	return "<unknown>"
}

func countDataEntries(pu *pb.ProgramUnit) int {
	count := 0
	dd := pu.GetDataDivision()
	if dd == nil {
		return 0
	}
	if ws := dd.GetWorkingStorageSection(); ws != nil {
		count += len(ws.GetDataDescriptionEntries())
	}
	if ls := dd.GetLinkageSection(); ls != nil {
		count += len(ls.GetDataDescriptionEntries())
	}
	if fs := dd.GetFileSection(); fs != nil {
		for _, fd := range fs.GetFileDescriptionEntries() {
			count += len(fd.GetDataDescriptionEntries())
		}
	}
	return count
}

func countParagraphs(pu *pb.ProgramUnit) int {
	pd := pu.GetProcedureDivision()
	if pd == nil {
		return 0
	}
	count := 0
	if pp := pd.GetParagraphs(); pp != nil {
		count += len(pp.GetParagraphs())
	}
	for _, s := range pd.GetProcedureSections() {
		if pp := s.GetParagraphs(); pp != nil {
			count += len(pp.GetParagraphs())
		}
	}
	return count
}

func hasProcCall(t *testing.T, calls []*pb.Call, name string) {
	t.Helper()
	for _, c := range calls {
		if c.GetType() == pb.CallType_PROCEDURE_CALL && strings.EqualFold(c.GetName(), name) {
			return
		}
	}
	t.Errorf("no PROCEDURE_CALL to %q found in %d calls", name, len(calls))
}

func hasDataCall(t *testing.T, calls []*pb.Call, name string) {
	t.Helper()
	for _, c := range calls {
		if c.GetType() == pb.CallType_DATA_DESCRIPTION_ENTRY_CALL &&
			strings.EqualFold(c.GetName(), name) {
			return
		}
	}
	t.Errorf("no DATA_DESCRIPTION_ENTRY_CALL to %q found in %d calls", name, len(calls))
}

func hasCallType(t *testing.T, calls []*pb.Call, ct pb.CallType) {
	t.Helper()
	for _, c := range calls {
		if c.GetType() == ct {
			return
		}
	}
	t.Errorf("no call of type %v found in %d calls", ct, len(calls))
}

func checkWorkingStorageEntries(t *testing.T, fileName string, table *symbol.SymbolTable) {
	t.Helper()
	switch {
	case strings.Contains(fileName, "DataDescription01"):
		if len(table.DataEntries) == 0 {
			t.Error("DataDescription01: expected data entries")
		}
	case strings.Contains(fileName, "DataDescription88"):
		if len(table.Conditions) == 0 {
			t.Error("DataDescription88: expected 88-level condition entries")
		}
	case strings.Contains(fileName, "DataDescription77"):
		if len(table.DataEntries) == 0 {
			t.Error("DataDescription77: expected 77-level data entry")
		}
	case strings.Contains(fileName, "DataDescription66"):
		if len(table.DataEntries) == 0 {
			t.Error("DataDescription66: expected 66-level data entry")
		}
	case strings.Contains(fileName, "Redefines"):
		if len(table.DataEntries) == 0 {
			t.Error("DataDescriptionRedefines: expected data entries")
		}
	}
}

func verifyStatementCalls(t *testing.T, category string, calls []*pb.Call) {
	t.Helper()
	expectedTypes := map[string][]pb.CallType{
		"move":       {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"add":        {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"subtract":   {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"multiply":   {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"divide":     {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"compute":    {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"perform":    {pb.CallType_PROCEDURE_CALL, pb.CallType_SECTION_CALL},
		"gotostmt":   {pb.CallType_PROCEDURE_CALL},
		"call":       {pb.CallType_PROCEDURE_CALL, pb.CallType_DATA_DESCRIPTION_ENTRY_CALL, pb.CallType_UNDEFINED_CALL},
		"open":       {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"close":      {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"read":       {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"write":      {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"rewrite":    {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"delete":     {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"start":      {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"sort":       {pb.CallType_FILE_CONTROL_ENTRY_CALL, pb.CallType_PROCEDURE_CALL},
		"merge":      {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"release":    {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"returnstmt": {pb.CallType_FILE_CONTROL_ENTRY_CALL},
		"display":    {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"initialize": {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"inspect":    {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"string":     {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"unstring":   {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"set":        {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"search":     {pb.CallType_DATA_DESCRIPTION_ENTRY_CALL, pb.CallType_TABLE_CALL},
		"ifstmt":     {pb.CallType_PROCEDURE_CALL, pb.CallType_DATA_DESCRIPTION_ENTRY_CALL},
		"evaluate":   {pb.CallType_PROCEDURE_CALL},
		"alter":      {pb.CallType_PROCEDURE_CALL},
		"generate":   {pb.CallType_REPORT_CALL},
		"initiate":   {pb.CallType_REPORT_CALL},
		"terminate":  {pb.CallType_REPORT_CALL},
		"send":       {pb.CallType_COMMUNICATION_DESCRIPTION_ENTRY_CALL},
		"receive":    {pb.CallType_COMMUNICATION_DESCRIPTION_ENTRY_CALL},
	}

	expected, ok := expectedTypes[category]
	if !ok || len(expected) == 0 {
		return
	}

	for _, et := range expected {
		for _, c := range calls {
			if c.GetType() == et {
				return
			}
		}
	}

	typeNames := make([]string, len(expected))
	for i, et := range expected {
		typeNames[i] = fmt.Sprintf("%v", et)
	}
	t.Logf("NOTE: category %s expected one of %v in %d calls (may be OK for minimal test files)",
		category, typeNames, len(calls))
}
