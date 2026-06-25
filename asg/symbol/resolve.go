package symbol

import (
	"fmt"
	"strings"

	"github.com/aixfoundry/cobol-go/pb"
)

// ResolutionResult holds the output of symbol resolution.
type ResolutionResult struct {
	Calls      []*pb.Call
	Unresolved []UnresolvedRef
}

// UnresolvedRef describes a name reference that could not be resolved.
type UnresolvedRef struct {
	Name     string
	Location string
	Kind     pb.CallType
}

// Resolve walks a parsed COBOL program AST and resolves all name references
// against the given SymbolTable.
func Resolve(program *pb.Program, table *SymbolTable) *ResolutionResult {
	result := &ResolutionResult{}
	if program == nil || table == nil {
		return result
	}
	for _, cu := range program.GetCompilationUnits() {
		for _, pu := range cu.GetProgramUnits() {
			resolveProgramUnit(pu, table, result)
		}
	}
	return result
}

func resolveProgramUnit(pu *pb.ProgramUnit, table *SymbolTable, result *ResolutionResult) {
	if pu == nil {
		return
	}
	if pd := pu.GetProcedureDivision(); pd != nil {
		resolveProcedureDivision(pd, table, pu, result)
	}
	for _, npu := range pu.GetProgramUnits() {
		resolveProgramUnit(npu, table, result)
	}
}

func resolveProcedureDivision(pd *pb.ProcedureDivision, table *SymbolTable, pu *pb.ProgramUnit, result *ResolutionResult) {
	resolveParagraphsBlock(pd.GetParagraphs(), table, pu, result)
	for _, s := range pd.GetProcedureSections() {
		resolveParagraphsBlock(s.GetParagraphs(), table, pu, result)
	}
	// Declaratives
	for _, d := range pd.GetDeclaratives() {
		resolveParagraphsBlock(d.GetParagraphs(), table, pu, result)
	}
}

func resolveParagraphsBlock(paragraphs *pb.Paragraphs, table *SymbolTable, pu *pb.ProgramUnit, result *ResolutionResult) {
	if paragraphs == nil {
		return
	}
	if sents := paragraphs.GetSentences(); sents != nil {
		for _, s := range sents.GetSentences() {
			resolveSentence(s, table, pu, result)
		}
	}
	for _, p := range paragraphs.GetParagraphs() {
		if sents := p.GetSentences(); sents != nil {
			for _, s := range sents.GetSentences() {
				resolveSentence(s, table, pu, result)
			}
		}
	}
}

func resolveSentence(s *pb.Sentence, table *SymbolTable, pu *pb.ProgramUnit, result *ResolutionResult) {
	for _, stmt := range s.GetStatements() {
		resolveStatement(stmt, table, pu, result)
	}
}

func resolveStatement(stmt *pb.Statement, table *SymbolTable, pu *pb.ProgramUnit, result *ResolutionResult) {
	if stmt == nil {
		return
	}
	switch s := stmt.GetOneOf().(type) {
	// --- Procedure-calling ---
	case *pb.Statement_PerformStatement:
		resolvePerform(s.PerformStatement, table, result)
	case *pb.Statement_GoToStatement:
		resolveGoTo(s.GoToStatement, table, result)
	case *pb.Statement_AlterStatement:
		resolveAlter(s.AlterStatement, table, result)
	case *pb.Statement_CallStatement:
		resolveCall(s.CallStatement, table, result)

	// --- Data movement ---
	case *pb.Statement_MoveStatement:
		resolveMove(s.MoveStatement, table, result)
	case *pb.Statement_ComputeStatement:
		resolveCompute(s.ComputeStatement, table, result)

	// --- Arithmetic ---
	case *pb.Statement_AddStatement:
		resolveAdd(s.AddStatement, table, result)
	case *pb.Statement_SubtractStatement:
		resolveSubtract(s.SubtractStatement, table, result)
	case *pb.Statement_MultiplyStatement:
		resolveMultiply(s.MultiplyStatement, table, result)
	case *pb.Statement_DivideStatement:
		resolveDivide(s.DivideStatement, table, result)

	// --- File I/O ---
	case *pb.Statement_ReadStatement:
		resolveRead(s.ReadStatement, table, result)
	case *pb.Statement_WriteStatement:
		resolveWrite(s.WriteStatement, table, result)
	case *pb.Statement_RewriteStatement:
		resolveRewrite(s.RewriteStatement, table, result)
	case *pb.Statement_DeleteStatement:
		resolveDelete(s.DeleteStatement, table, result)
	case *pb.Statement_StartStatement:
		resolveStart(s.StartStatement, table, result)
	case *pb.Statement_OpenStatement:
		resolveOpen(s.OpenStatement, table, result)
	case *pb.Statement_CloseStatement:
		resolveClose(s.CloseStatement, table, result)

	// --- Conditional ---
	case *pb.Statement_IfStatement:
		resolveIf(s.IfStatement, table, pu, result)
	case *pb.Statement_EvaluateStatement:
		resolveEvaluate(s.EvaluateStatement, table, pu, result)
	case *pb.Statement_SearchStatement:
		resolveSearch(s.SearchStatement, table, result)

	// --- Other ---
	case *pb.Statement_InitializeStatement:
		resolveInitialize(s.InitializeStatement, table, result)
	case *pb.Statement_InspectStatement:
		resolveInspect(s.InspectStatement, table, result)
	case *pb.Statement_StringStatement:
		resolveString(s.StringStatement, table, result)
	case *pb.Statement_UnstringStatement:
		resolveUnstring(s.UnstringStatement, table, result)
	case *pb.Statement_SetStatement:
		resolveSet(s.SetStatement, table, result)
	case *pb.Statement_SortStatement:
		resolveSort(s.SortStatement, table, result)
	case *pb.Statement_MergeStatement:
		resolveMerge(s.MergeStatement, table, result)
	case *pb.Statement_DisplayStatement:
		resolveDisplay(s.DisplayStatement, table, result)
	case *pb.Statement_AcceptStatement:
		resolveAccept(s.AcceptStatement, table, result)
	case *pb.Statement_ReleaseStatement:
		resolveRelease(s.ReleaseStatement, table, result)
	case *pb.Statement_ReturnStatement:
		resolveReturn(s.ReturnStatement, table, result)
	}
}

// --- Procedure-calling statements ---

func resolvePerform(stmt *pb.PerformStatement, table *SymbolTable, result *ResolutionResult) {
	ps := stmt.GetProcedureStatement()
	if ps == nil {
		return
	}
	if pn := ps.GetProcedureName(); pn != nil {
		resolveProcName(pn, table, result, "PERFORM")
	}
	if through := ps.GetThrough(); through != nil {
		resolveProcName(through, table, result, "PERFORM THROUGH")
	}
	// Varying phrase for PERFORM VARYING
	// The varying clause is part of PerformFlavors or PerformVarying, not ProcedureStatement
	// TODO: extract varyings from the full perform statement proto
}

func resolveGoTo(stmt *pb.GoToStatement, table *SymbolTable, result *ResolutionResult) {
	if s := stmt.GetSimple(); s != nil {
		resolveProcName(s.GetProcedureName(), table, result, "GO TO")
	}
	if dep := stmt.GetDependingOn(); dep != nil {
		// DependingOnStatement has depending_on which has repeated procedure_names
		if do := dep.GetDependingOn(); do != nil {
			for _, pn := range do.GetProcedureNames() {
				resolveProcName(pn, table, result, "GO TO DEPENDING")
			}
		}
	}
}

func resolveAlter(stmt *pb.AlterStatement, table *SymbolTable, result *ResolutionResult) {
	for _, pt := range stmt.GetProceedTos() {
		resolveProcName(pt.GetFrom(), table, result, "ALTER FROM")
		resolveProcName(pt.GetTo(), table, result, "ALTER TO")
	}
}

func resolveCall(stmt *pb.CallStatement, table *SymbolTable, result *ResolutionResult) {
	if ident := stmt.GetTargetIdentifier(); ident != nil {
		resolveIdent(ident, table, result, "CALL")
	}
}

// --- Data movement ---

func resolveMove(stmt *pb.MoveStatement, table *SymbolTable, result *ResolutionResult) {
	if mt := stmt.GetMoveTo(); mt != nil {
		// Sending area (single value) - oneof { identifier, literal }
		if id := mt.GetIdentifier(); id != nil {
			resolveIdent(id, table, result, "MOVE FROM")
		}
		// To list (repeated identifiers)
		for _, to := range mt.GetTo() {
			resolveIdent(to, table, result, "MOVE TO")
		}
	}
	if mct := stmt.GetMoveCorrespondingTo(); mct != nil {
		resolveIdent(mct.GetSendingArea(), table, result, "MOVE CORR")
		for _, to := range mct.GetTo() {
			resolveIdent(to, table, result, "MOVE CORR TO")
		}
	}
}

func resolveCompute(stmt *pb.ComputeStatement, table *SymbolTable, result *ResolutionResult) {
	for _, store := range stmt.GetStores() {
		if id := store.GetIdentifier(); id != nil {
			resolveIdent(id, table, result, "COMPUTE")
		}
	}
}

// --- Arithmetic ---

func resolveAdd(stmt *pb.AddStatement, table *SymbolTable, result *ResolutionResult) {
	if to := stmt.GetTo(); to != nil {
		for _, id := range to.GetTos() {
			resolveIdent(id, table, result, "ADD TO")
		}
	}
	if tg := stmt.GetToGiving(); tg != nil {
		for _, g := range tg.GetGivings() {
			resolveIdent(g, table, result, "ADD GIVING")
		}
	}
	if corr := stmt.GetCorresponding(); corr != nil {
		resolveIdent(corr.GetCorresponding(), table, result, "ADD CORR FROM")
		resolveIdent(corr.GetTo(), table, result, "ADD CORR TO")
	}
}

func resolveSubtract(stmt *pb.SubtractStatement, table *SymbolTable, result *ResolutionResult) {
	if fs := stmt.GetFromStatement(); fs != nil {
		for _, m := range fs.GetMinuends() {
			resolveIdent(m.GetIdentifier(), table, result, "SUBTRACT FROM")
		}
	}
	if fgs := stmt.GetFromGivingStatement(); fgs != nil {
		for _, g := range fgs.GetGivings() {
			resolveIdent(g.GetIdentifier(), table, result, "SUBTRACT GIVING")
		}
	}
	if cs := stmt.GetCorrespondingStatement(); cs != nil {
		_ = cs
	}
}

func resolveMultiply(stmt *pb.MultiplyStatement, table *SymbolTable, result *ResolutionResult) {
	if id := stmt.GetIdentifier(); id != nil {
		resolveIdent(id, table, result, "MULTIPLY")
	}
	if g := stmt.GetGiving(); g != nil {
		for _, gr := range g.GetGivingResult() {
			resolveIdent(gr.GetIdentifier(), table, result, "MULTIPLY GIVING")
		}
	}
}

func resolveDivide(stmt *pb.DivideStatement, table *SymbolTable, result *ResolutionResult) {
	if is := stmt.GetIntoStatement(); is != nil {
		for _, into := range is.GetIntos() {
			if id := into.GetIdentifier(); id != nil {
				resolveIdent(id, table, result, "DIVIDE INTO")
			}
		}
	}
	if rem := stmt.GetRemainder(); rem != nil {
		if id := rem.GetIdentifier(); id != nil {
			resolveIdent(id, table, result, "DIVIDE REMAINDER")
		}
	}
}

// --- File I/O ---

func resolveRead(stmt *pb.ReadStatement, table *SymbolTable, result *ResolutionResult) {
	if fn := stmt.GetFileName(); fn != nil {
		resolveFileName(fn, table, result, "READ")
	}
}

func resolveWrite(stmt *pb.WriteStatement, table *SymbolTable, result *ResolutionResult) {
	if rn := stmt.GetRecordName(); rn != nil {
		if qdn := rn.GetQualifiedDataName(); qdn != nil {
			if name := qdnName(qdn); name != "" {
				call := &pb.Call{Type: pb.CallType_DATA_DESCRIPTION_ENTRY_CALL, Name: name}
				resolveAndSetTarget(call, name, table, result, "WRITE")
				result.Calls = append(result.Calls, call)
			}
		}
	}
}

func resolveRewrite(stmt *pb.RewriteStatement, table *SymbolTable, result *ResolutionResult) {
	if rn := stmt.GetRecordName(); rn != nil {
		if qdn := rn.GetQualifiedDataName(); qdn != nil {
			if name := qdnName(qdn); name != "" {
				call := &pb.Call{Type: pb.CallType_DATA_DESCRIPTION_ENTRY_CALL, Name: name}
				resolveAndSetTarget(call, name, table, result, "REWRITE")
				result.Calls = append(result.Calls, call)
			}
		}
	}
}

func resolveDelete(stmt *pb.DeleteStatement, table *SymbolTable, result *ResolutionResult) {
	resolveFileName(stmt.GetFileName(), table, result, "DELETE")
}

func resolveStart(stmt *pb.StartStatement, table *SymbolTable, result *ResolutionResult) {
	resolveFileName(stmt.GetFileName(), table, result, "START")
}

func resolveOpen(stmt *pb.OpenStatement, table *SymbolTable, result *ResolutionResult) {
	// OpenStatement contains various input/output/io/extend variants, each with file names
	switch s := stmt.GetOneOf().(type) {
	case *pb.OpenStatement_InputStatement_:
		for _, in := range s.InputStatement.GetInputs() {
			resolveFileName(in.GetFileName(), table, result, "OPEN INPUT")
		}
	case *pb.OpenStatement_OutputStatement_:
		for _, out := range s.OutputStatement.GetOutputs() {
			resolveFileName(out.GetFileName(), table, result, "OPEN OUTPUT")
		}
	case *pb.OpenStatement_IoStatement:
		for _, fn := range s.IoStatement.GetFileNames() {
			resolveFileName(fn, table, result, "OPEN I-O")
		}
	case *pb.OpenStatement_ExtendStatement_:
		for _, fn := range s.ExtendStatement.GetFileNames() {
			resolveFileName(fn, table, result, "OPEN EXTEND")
		}
	}
}

func resolveClose(stmt *pb.CloseStatement, table *SymbolTable, result *ResolutionResult) {
	for _, cf := range stmt.GetCloseFiles() {
		resolveFileName(cf.GetFileName(), table, result, "CLOSE")
	}
}

func resolveRelease(stmt *pb.ReleaseStatement, table *SymbolTable, result *ResolutionResult) {
	if rn := stmt.GetRecordName(); rn != nil {
		if qdn := rn.GetQualifiedDataName(); qdn != nil {
			if name := qdnName(qdn); name != "" {
				call := &pb.Call{Type: pb.CallType_DATA_DESCRIPTION_ENTRY_CALL, Name: name}
				resolveAndSetTarget(call, name, table, result, "RELEASE")
				result.Calls = append(result.Calls, call)
			}
		}
	}
}

func resolveReturn(stmt *pb.ReturnStatement, table *SymbolTable, result *ResolutionResult) {
	resolveFileName(stmt.GetFileName(), table, result, "RETURN")
}

// --- Conditional ---

func resolveIf(stmt *pb.IfStatement, table *SymbolTable, pu *pb.ProgramUnit, result *ResolutionResult) {
	for _, s := range stmt.GetThen().GetStatements() {
		resolveStatement(s, table, pu, result)
	}
	if elseBlock := stmt.GetElse(); elseBlock != nil {
		for _, s := range elseBlock.GetStatements() {
			resolveStatement(s, table, pu, result)
		}
	}
}

func resolveEvaluate(stmt *pb.EvaluateStatement, table *SymbolTable, pu *pb.ProgramUnit, result *ResolutionResult) {
	walkStats := func(stmts []*pb.Statement) {
		for _, s := range stmts {
			resolveStatement(s, table, pu, result)
		}
	}
	// Statements are in WhenPhrase and WhenOther
	for _, wp := range stmt.GetWhenPhrases() {
		walkStats(wp.GetStatements())
	}
	if other := stmt.GetWhenOther(); other != nil {
		walkStats(other.GetStatements())
	}
}

func resolveSearch(stmt *pb.SearchStatement, table *SymbolTable, result *ResolutionResult) {
	if qdn := stmt.GetQualifiedDataName(); qdn != nil {
		if name := qdnName(qdn); name != "" {
			call := &pb.Call{Type: pb.CallType_TABLE_CALL, Name: name}
			resolveAndSetTarget(call, name, table, result, "SEARCH")
			result.Calls = append(result.Calls, call)
		}
	}
}

// --- Other ---

func resolveInitialize(stmt *pb.InitializeStatement, table *SymbolTable, result *ResolutionResult) {
	for _, id := range stmt.GetIdentifiers() {
		resolveIdent(id, table, result, "INITIALIZE")
	}
}

func resolveInspect(stmt *pb.InspectStatement, table *SymbolTable, result *ResolutionResult) {
	if id := stmt.GetIdentifier(); id != nil {
		resolveIdent(id, table, result, "INSPECT")
	}
}

func resolveString(stmt *pb.StringStatement, table *SymbolTable, result *ResolutionResult) {
	for _, sp := range stmt.GetSendingPhrases() {
		for _, snd := range sp.GetSendings() {
			if id := snd.GetIdentifier(); id != nil {
				resolveIdent(id, table, result, "STRING")
			}
		}
	}
	if ip := stmt.GetIntoPhrase(); ip != nil {
		resolveIdent(ip.GetIdentifier(), table, result, "STRING INTO")
	}
}

func resolveUnstring(stmt *pb.UnstringStatement, table *SymbolTable, result *ResolutionResult) {
	if sp := stmt.GetSendingPhrase(); sp != nil {
		resolveIdent(sp.GetIdentifier(), table, result, "UNSTRING FROM")
	}
	if ip := stmt.GetIntoPhrase(); ip != nil {
		for _, into := range ip.GetInto() {
			resolveIdent(into.GetIdentifier(), table, result, "UNSTRING INTO")
		}
	}
}

func resolveSet(stmt *pb.SetStatement, table *SymbolTable, result *ResolutionResult) {
	for _, ts := range stmt.GetToStatements() {
		for _, to := range ts.GetTos() {
			if id := to.GetIdentifier(); id != nil {
				resolveIdent(id, table, result, "SET TO")
			}
		}
		for _, tv := range ts.GetToValues() {
			if id := tv.GetIdentifier(); id != nil {
				resolveIdent(id, table, result, "SET TO VALUE")
			}
		}
	}
}

func resolveSort(stmt *pb.SortStatement, table *SymbolTable, result *ResolutionResult) {
	resolveFileName(stmt.GetFileName(), table, result, "SORT")
}

func resolveMerge(stmt *pb.MergeStatement, table *SymbolTable, result *ResolutionResult) {
	resolveFileName(stmt.GetFileName(), table, result, "MERGE")
}

func resolveDisplay(stmt *pb.DisplayStatement, table *SymbolTable, result *ResolutionResult) {
	for _, op := range stmt.GetOperands() {
		if id := op.GetIdentifier(); id != nil {
			resolveIdent(id, table, result, "DISPLAY")
		}
	}
}

func resolveAccept(stmt *pb.AcceptStatement, table *SymbolTable, result *ResolutionResult) {
	if id := stmt.GetIdentifier(); id != nil {
		resolveIdent(id, table, result, "ACCEPT")
	}
}

// --- Core resolution helpers ---

func resolveIdent(ident *pb.Identifier, table *SymbolTable, result *ResolutionResult, location string) {
	if ident == nil {
		return
	}
	var name string
	callType := pb.CallType_DATA_DESCRIPTION_ENTRY_CALL

	switch id := ident.GetOneOf().(type) {
	case *pb.Identifier_QualifiedDataName:
		name = qdnName(id.QualifiedDataName)
	case *pb.Identifier_TableCall:
		name = qdnName(id.TableCall.GetQualifiedDataName())
		callType = pb.CallType_TABLE_CALL
	case *pb.Identifier_FunctionCall:
		name = id.FunctionCall.GetFunctionName().GetValue()
		callType = pb.CallType_FUNCTION_CALL
	case *pb.Identifier_SpecialRegister:
		callType = pb.CallType_SPECIAL_REGISTER_CALL
		return
	default:
		return
	}

	if name == "" {
		return
	}
	call := &pb.Call{Type: callType, Name: name}
	resolveAndSetTarget(call, name, table, result, location)
	result.Calls = append(result.Calls, call)
}

func resolveProcName(pn *pb.ProcedureName, table *SymbolTable, result *ResolutionResult, location string) {
	if pn == nil {
		return
	}
	var name string
	callType := pb.CallType_PROCEDURE_CALL

	if pg := pn.GetParagraphName(); pg != nil {
		if cw := pg.GetCobolWord(); cw != nil {
			name = cw.GetValue()
		} else if il := pg.GetIntegerLiteral(); il != nil {
			name = il.GetValue()
		}
	} else if sn := pn.GetSectionName(); sn != nil {
		if cw := sn.GetCobolWord(); cw != nil {
			name = cw.GetValue()
		} else if il := sn.GetIntegerLiteral(); il != nil {
			name = il.GetValue()
		}
		callType = pb.CallType_SECTION_CALL
	}

	if name == "" {
		return
	}
	call := &pb.Call{Type: callType, Name: name}
	resolveAndSetTarget(call, name, table, result, location)
	result.Calls = append(result.Calls, call)
}

func resolveFileName(fn *pb.FileName, table *SymbolTable, result *ResolutionResult, location string) {
	if fn == nil {
		return
	}
	name := nameFromCobolWord(fn.GetCobolWord())
	if name == "" {
		return
	}
	call := &pb.Call{Type: pb.CallType_FILE_CONTROL_ENTRY_CALL, Name: name}
	resolveAndSetTarget(call, name, table, result, location)
	result.Calls = append(result.Calls, call)
}


func resolveAndSetTarget(call *pb.Call, name string, table *SymbolTable, result *ResolutionResult, location string) {
	upper := strings.ToUpper(name)

	switch call.GetType() {
	case pb.CallType_PROCEDURE_CALL:
		if p, ok := table.Paragraphs[upper]; ok {
			call.Target = &pb.Call_ParagraphName{ParagraphName: p.GetParagraphName()}
		} else {
			result.Unresolved = append(result.Unresolved, UnresolvedRef{Name: name, Location: location, Kind: pb.CallType_PROCEDURE_CALL})
		}

	case pb.CallType_SECTION_CALL:
		if _, ok := table.Sections[upper]; ok {
			call.Target = &pb.Call_SectionName{
				SectionName: &pb.SectionName{OneOf: &pb.SectionName_CobolWord{CobolWord: &pb.CobolWord{Value: name}}},
			}
		} else {
			result.Unresolved = append(result.Unresolved, UnresolvedRef{Name: name, Location: location, Kind: pb.CallType_SECTION_CALL})
		}

	case pb.CallType_DATA_DESCRIPTION_ENTRY_CALL, pb.CallType_TABLE_CALL:
		if entries, ok := table.DataEntries[upper]; ok && len(entries) > 0 {
			entry := entries[0]
			if f1 := entry.GetF1(); f1 != nil {
				if dn := f1.GetDataName(); dn != nil {
					call.Target = &pb.Call_DataName{DataName: dn}
				}
			}
		} else if entries, ok := table.Conditions[upper]; ok && len(entries) > 0 {
			entry := entries[0]
			if f3 := entry.GetF3(); f3 != nil {
				call.Target = &pb.Call_ConditionName{ConditionName: f3.GetConditionName()}
			}
		} else {
			result.Unresolved = append(result.Unresolved, UnresolvedRef{Name: name, Location: location, Kind: call.GetType()})
		}

	case pb.CallType_FILE_CONTROL_ENTRY_CALL:
		if _, ok := table.FileControlEntries[upper]; ok || table.FileDescriptions[upper] != nil {
			call.Target = &pb.Call_FileName{FileName: &pb.FileName{CobolWord: &pb.CobolWord{Value: name}}}
		} else {
			result.Unresolved = append(result.Unresolved, UnresolvedRef{Name: name, Location: location, Kind: pb.CallType_FILE_CONTROL_ENTRY_CALL})
		}
	}
}

// Report returns a human-readable summary.
func (r *ResolutionResult) Report() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Resolution: %d resolved calls, %d unresolved\n", len(r.Calls), len(r.Unresolved)))
	for _, u := range r.Unresolved {
		b.WriteString(fmt.Sprintf("  UNRESOLVED: %s (%s) at %s\n", u.Name, u.Kind.String(), u.Location))
	}
	return b.String()
}
