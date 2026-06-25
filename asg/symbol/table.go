// Package symbol provides symbol table building and name resolution for the COBOL AST.
package symbol

import (
	"github.com/aixfoundry/cobol-go/pb"
)

// IndexInfo tracks where an index name was defined (INDEXED BY clause on OCCURS).
type IndexInfo struct {
	DataEntry *pb.DataDescriptionEntry
	IndexName string
}

// SymbolTable holds indexed name definitions from a COBOL program AST.
type SymbolTable struct {
	// DataEntries maps data names to their DataDescriptionEntry definitions.
	// Multiple entries may exist for the same name (in different sections).
	DataEntries map[string][]*pb.DataDescriptionEntry
	// Conditions maps 88-level condition names to their DataDescriptionEntry definitions.
	Conditions map[string][]*pb.DataDescriptionEntry
	// Paragraphs maps paragraph names to their Paragraph definitions.
	Paragraphs map[string]*pb.Paragraph
	// Sections maps section names to their ProcedureSection definitions.
	Sections map[string]*pb.ProcedureSection
	// FileControlEntries maps file names from SELECT statements to their FileControlEntry.
	FileControlEntries map[string]*pb.FileControlEntry
	// FileDescriptions maps FD/SD file names to their FileDescriptionEntry.
	FileDescriptions map[string]*pb.FileDescriptionEntry
	// Programs maps PROGRAM-ID names to their ProgramUnit.
	Programs map[string]*pb.ProgramUnit
	// Reports maps report names to their ReportDescription.
	Reports map[string]*pb.ReportDescription
	// Screens maps screen names to their ScreenDescriptionEntry.
	Screens map[string]*pb.ScreenDescriptionEntry
	// CommunicationEntries maps CD names to their CommunicationDescriptionEntry.
	CommunicationEntries map[string]*pb.CommunicationDescriptionEntry
	// Mnemonics maps mnemonic names from SPECIAL-NAMES.
	Mnemonics map[string]*pb.MnemonicName
	// Indices maps index names (INDEXED BY) to their defining data entry.
	Indices map[string]*IndexInfo
}

// NewSymbolTable creates an empty SymbolTable with initialized maps.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		DataEntries:          make(map[string][]*pb.DataDescriptionEntry),
		Conditions:           make(map[string][]*pb.DataDescriptionEntry),
		Paragraphs:           make(map[string]*pb.Paragraph),
		Sections:             make(map[string]*pb.ProcedureSection),
		FileControlEntries:   make(map[string]*pb.FileControlEntry),
		FileDescriptions:     make(map[string]*pb.FileDescriptionEntry),
		Programs:             make(map[string]*pb.ProgramUnit),
		Reports:              make(map[string]*pb.ReportDescription),
		Screens:              make(map[string]*pb.ScreenDescriptionEntry),
		CommunicationEntries: make(map[string]*pb.CommunicationDescriptionEntry),
		Mnemonics:            make(map[string]*pb.MnemonicName),
		Indices:              make(map[string]*IndexInfo),
	}
}

// Build constructs a SymbolTable from a parsed COBOL program's Protobuf AST.
func Build(program *pb.Program) *SymbolTable {
	table := NewSymbolTable()
	if program == nil {
		return table
	}
	for _, cu := range program.GetCompilationUnits() {
		for _, pu := range cu.GetProgramUnits() {
			buildFromProgramUnit(pu, table)
		}
	}
	return table
}

// buildFromProgramUnit indexes all definitions in a single ProgramUnit.
func buildFromProgramUnit(pu *pb.ProgramUnit, table *SymbolTable) {
	if pu == nil {
		return
	}

	// PROGRAM-ID
	if id := pu.GetIdentificationDivision(); id != nil {
		if pp := id.GetProgramIdParagraph(); pp != nil {
			if pn := pp.GetProgramName(); pn != nil {
				if name := nameFromProgramName(pn); name != "" {
					table.Programs[name] = pu
				}
			}
		}
	}

	// Environment Division: File Control, Mnemonics
	if ed := pu.GetEnvironmentDivision(); ed != nil {
		if cfg := ed.GetConfigurationSection(); cfg != nil {
			if sp := cfg.GetSpecialNamesParagraph(); sp != nil {
				// Mnemonic names appear in ChannelClause and EnvironmentSwitchNameClause
				if cc := sp.GetChannelClause(); cc != nil {
					if mn := cc.GetMnemonicName(); mn != nil {
						if name := nameFromCobolWord(mn.GetCobolWord()); name != "" {
							table.Mnemonics[name] = mn
						}
					}
				}
				for _, esc := range sp.GetEnvironmentSwitchNameClauses() {
					if mn := esc.GetMnemonicName(); mn != nil {
						if name := nameFromCobolWord(mn.GetCobolWord()); name != "" {
							table.Mnemonics[name] = mn
						}
					}
				}
			}
		}

		if io := ed.GetInputOutputSection(); io != nil {
			if fcp := io.GetFileControlParagraph(); fcp != nil {
				for _, fce := range fcp.GetFileControlEntries() {
					// FileControlEntry has its own FileName; also check SelectClause
					if selectClause := fce.GetSelectClause(); selectClause != nil {
						if fn := selectClause.GetFileName(); fn != nil {
							if name := nameFromCobolWord(fn.GetCobolWord()); name != "" {
								table.FileControlEntries[name] = fce
								// Also register the FileControlEntry itself by its FileName
								if fceName := nameFromCobolWord(fce.GetFileName().GetCobolWord()); fceName != "" && fceName != name {
									table.FileControlEntries[fceName] = fce
								}
							}
						}
					}
				}
			}
		}
	}

	// Data Division
	if dd := pu.GetDataDivision(); dd != nil {
		// File Section: FD/SD entries + their data descriptions
		if fs := dd.GetFileSection(); fs != nil {
			for _, fd := range fs.GetFileDescriptionEntries() {
				if name := nameFromCobolWord(fd.GetFileName().GetCobolWord()); name != "" {
					table.FileDescriptions[name] = fd
				}
				collectDataEntries(fd.GetDataDescriptionEntries(), table)
			}
		}

		// Working Storage
		if ws := dd.GetWorkingStorageSection(); ws != nil {
			collectDataEntries(ws.GetDataDescriptionEntries(), table)
		}
		// Linkage Section
		if ls := dd.GetLinkageSection(); ls != nil {
			collectDataEntries(ls.GetDataDescriptionEntries(), table)
		}
		// Local Storage Section
		if lss := dd.GetLocalStorageSection(); lss != nil {
			collectDataEntries(lss.GetDataDescriptionEntries(), table)
		}
		// Communication Section
		if cs := dd.GetCommunicationSection(); cs != nil {
			for _, cd := range cs.GetCommunicationDescriptionEntries() {
				if input := cd.GetInput(); input != nil {
					if name := nameFromCobolWord(input.GetCdName().GetCobolWord()); name != "" {
						table.CommunicationEntries[name] = cd
					}
				}
			}
			collectDataEntries(cs.GetDataDescriptionEntries(), table)
		}
		// Report Section
		if rs := dd.GetReportSection(); rs != nil {
			for _, rd := range rs.GetReportDescriptions() {
				if entry := rd.GetReportDescriptionEntry(); entry != nil {
					if rn := entry.GetReportName(); rn != nil {
						if name := qdnName(rn.GetQualifiedDataName()); name != "" {
							table.Reports[name] = rd
						}
					}
				}
			}
		}
		// Screen Section
		if ss := dd.GetScreenSection(); ss != nil {
			for _, se := range ss.GetScreenDescriptionEntries() {
				if sn := se.GetScreenName(); sn != nil {
					if name := nameFromCobolWord(sn.GetCobolWord()); name != "" {
						table.Screens[name] = se
					}
				}
			}
		}
	}

	// Procedure Division: Paragraphs and Sections
	if pd := pu.GetProcedureDivision(); pd != nil {
		// Sections
		for _, s := range pd.GetProcedureSections() {
			if header := s.GetProcedureSectionHeader(); header != nil {
				if sn := header.GetSectionName(); sn != nil {
					if name := nameFromSectionOrParagraph(sn.GetCobolWord(), sn.GetIntegerLiteral()); name != "" {
						table.Sections[name] = s
					}
				}
			}
			// Paragraphs within the section
			if paragraphs := s.GetParagraphs(); paragraphs != nil {
				for _, p := range paragraphs.GetParagraphs() {
					if name := paragraphName(p); name != "" {
						table.Paragraphs[name] = p
					}
				}
			}
		}
		// Top-level paragraphs (not in a section)
		if paragraphs := pd.GetParagraphs(); paragraphs != nil {
			for _, p := range paragraphs.GetParagraphs() {
				if name := paragraphName(p); name != "" {
					table.Paragraphs[name] = p
				}
			}
		}
	}

	// Nested program units
	for _, npu := range pu.GetProgramUnits() {
		buildFromProgramUnit(npu, table)
	}
}

// collectDataEntries indexes data description entries from any section.
// It separates 88-level conditions from regular data items.
func collectDataEntries(entries []*pb.DataDescriptionEntry, table *SymbolTable) {
	for _, dde := range entries {
		if dde == nil {
			continue
		}

		if f1 := dde.GetF1(); f1 != nil {
			// Format1: regular data item (group or elementary)
			if dname := f1.GetDataName(); dname != nil {
				if name := nameFromCobolWord(dname.GetCobolWord()); name != "" {
					table.DataEntries[name] = append(table.DataEntries[name], dde)
				}
			}
			// INDEXED BY indices
			if occurs := f1.GetDataOccursClause(); occurs != nil {
				for _, oi := range occurs.GetIndexes() {
					if oi != nil {
						for _, idx := range oi.GetIndexNames() {
							if idxName := nameFromCobolWord(idx.GetCobolWord()); idxName != "" {
								table.Indices[idxName] = &IndexInfo{
									DataEntry: dde,
									IndexName: idxName,
								}
							}
						}
					}
				}
			}
		} else if f2 := dde.GetF2(); f2 != nil {
			// Format2: 66-level RENAMES
			if dname := f2.GetDataName(); dname != nil {
				if name := nameFromCobolWord(dname.GetCobolWord()); name != "" {
					table.DataEntries[name] = append(table.DataEntries[name], dde)
				}
			}
		} else if f3 := dde.GetF3(); f3 != nil {
			// Format3: 88-level condition name
			if cname := f3.GetConditionName(); cname != nil {
				if name := nameFromCobolWord(cname.GetCobolWord()); name != "" {
					table.Conditions[name] = append(table.Conditions[name], dde)
				}
			}
		}
	}
}

// paragraphName extracts the name from a Paragraph's ParagraphName.
func paragraphName(p *pb.Paragraph) string {
	if p == nil {
		return ""
	}
	pn := p.GetParagraphName()
	if pn == nil {
		return ""
	}
	return nameFromSectionOrParagraph(pn.GetCobolWord(), pn.GetIntegerLiteral())
}

// --- name extraction helpers ---

func nameFromCobolWord(cw *pb.CobolWord) string {
	if cw == nil {
		return ""
	}
	return cw.GetValue()
}

func nameFromSectionOrParagraph(cw *pb.CobolWord, il *pb.IntegerLiteral) string {
	if cw != nil {
		return cw.GetValue()
	}
	if il != nil {
		return il.GetValue()
	}
	return ""
}

func nameFromProgramName(pn *pb.ProgramName) string {
	if pn == nil {
		return ""
	}
	if cw := pn.GetCobolWord(); cw != nil {
		return cw.GetValue()
	}
	if nl := pn.GetNonNumericLiteral(); nl != nil {
		return nl.GetValue()
	}
	return ""
}

// qdnName extracts a data or condition name from a QualifiedDataName.
func qdnName(qdn *pb.QualifiedDataName) string {
	if qdn == nil {
		return ""
	}
	if f1 := qdn.GetF1(); f1 != nil {
		if dn := f1.GetDataName(); dn != nil {
			return nameFromCobolWord(dn.GetCobolWord())
		}
		if cn := f1.GetConditionName(); cn != nil {
			return nameFromCobolWord(cn.GetCobolWord())
		}
	}
	return ""
}
