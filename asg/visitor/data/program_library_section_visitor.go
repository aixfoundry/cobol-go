package data

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ProgramLibrarySectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ProgramLibrarySection
}

func NewProgramLibrarySectionVisitor(section *pb.ProgramLibrarySection) *ProgramLibrarySectionVisitor {
	return &ProgramLibrarySectionVisitor{
		section: section,
	}
}

func (v *ProgramLibrarySectionVisitor) VisitProgramLibrarySection(ctx *cobol85.ProgramLibrarySectionContext) any {
	for _, ictx := range ctx.AllLibraryDescriptionEntry() {
		cctx := ictx.(*cobol85.LibraryDescriptionEntryContext)
		entry := &pb.LibraryDescriptionEntry{}
		if iformat := cctx.LibraryDescriptionEntryFormat1(); iformat != nil {
			cformat := iformat.(*cobol85.LibraryDescriptionEntryFormat1Context)
			entry.LibraryName = conv.LibraryName(cformat.LibraryName())
			var sharing pb.AttributeClause1_Sharing
			if iac := cformat.LibraryAttributeClauseFormat1(); iac != nil {
				ac := iac.(*cobol85.LibraryAttributeClauseFormat1Context)
				switch {
				case ac.DONTCARE() != nil:
					sharing = pb.AttributeClause1_DONTCARE
				case ac.PRIVATE() != nil:
					sharing = pb.AttributeClause1_PRIVATE
				case ac.SHAREDBYRUNUNIT() != nil:
					sharing = pb.AttributeClause1_SHAREDBYRUNUNIT
				case ac.SHAREDBYALL() != nil:
					sharing = pb.AttributeClause1_SHAREDBYALL
				}
			}
			var programName *pb.ProgramName
			var forValue *pb.Literal
			if ipc := cformat.LibraryEntryProcedureClauseFormat1(); ipc != nil {
				pc := ipc.(*cobol85.LibraryEntryProcedureClauseFormat1Context)
				programName = conv.ProgramName(pc.ProgramName())
				if ipf := pc.LibraryEntryProcedureForClause(); ipf != nil {
					pf := ipf.(*cobol85.LibraryEntryProcedureForClauseContext)
					forValue = conv.Literal(pf.Literal())
				}
			}
			export := &pb.LibraryDescriptionEntry_Export{
				AttributeClause: &pb.AttributeClause1{
					Sharing: sharing,
				},
				ProcedureClause: &pb.ProcedureClause1{
					ProgramName: programName,
					For:         forValue,
				},
			}
			entry.OneOf = &pb.LibraryDescriptionEntry_Export_{
				Export: export,
			}
		} else if iformat := cctx.LibraryDescriptionEntryFormat2(); iformat != nil {
			cformat := iformat.(*cobol85.LibraryDescriptionEntryFormat2Context)
			entry.LibraryName = conv.LibraryName(cformat.LibraryName())
			imp := &pb.LibraryDescriptionEntry_Import{}
			if cformat.LibraryIsGlobalClause() != nil {
				imp.GlobalClause = &pb.GlobalClause{}
			}
			if cformat.LibraryIsCommonClause() != nil {
				imp.IsCommonClause = &pb.IsCommonClause{}
			}
			for _, iac := range cformat.AllLibraryAttributeClauseFormat2() {
				if imp.AttributeClause == nil {
					imp.AttributeClause = &pb.AttributeClause2{}
				}
				ac := iac.(*cobol85.LibraryAttributeClauseFormat2Context)
				if ilaf := ac.LibraryAttributeFunction(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryAttributeFunctionContext)
					imp.AttributeClause.FunctionName = conv.Literal(laf.Literal())
				}
				if ilaf := ac.LibraryAttributeParameter(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryAttributeParameterContext)
					imp.AttributeClause.LibParameter = conv.Literal(laf.Literal())
				}
				if ilaf := ac.LibraryAttributeTitle(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryAttributeTitleContext)
					imp.AttributeClause.Title = conv.Literal(laf.Literal())
				}
				switch {
				case ac.BYFUNCTION() != nil:
					imp.AttributeClause.LibAccess = pb.AttributeClause2_BYFUNCTION
				case ac.BYTITLE() != nil:
					imp.AttributeClause.LibAccess = pb.AttributeClause2_BYTITLE
				}
			}
			for _, iac := range cformat.AllLibraryEntryProcedureClauseFormat2() {
				if imp.ProcedureClause == nil {
					imp.ProcedureClause = &pb.ProcedureClause2{}
				}
				ac := iac.(*cobol85.LibraryEntryProcedureClauseFormat2Context)
				imp.ProcedureClause.ProgramName = conv.ProgramName(ac.ProgramName())
				if ilaf := ac.LibraryEntryProcedureForClause(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryEntryProcedureForClauseContext)
					imp.ProcedureClause.For = conv.Literal(laf.Literal())
				}
				if ilaf := ac.LibraryEntryProcedureForClause(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryEntryProcedureForClauseContext)
					imp.ProcedureClause.For = conv.Literal(laf.Literal())
				}
				if ilaf := ac.LibraryEntryProcedureWithClause(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryEntryProcedureWithClauseContext)
					for _, iwith := range laf.AllLibraryEntryProcedureWithName() {
						cwith := iwith.(*cobol85.LibraryEntryProcedureWithNameContext)
						withName := &pb.ProcedureClause2_WithName{}
						if cwith.LocalName() != nil {
							withName.OneOf = &pb.ProcedureClause2_WithName_LocalName{
								LocalName: conv.LocalName(cwith.LocalName()),
							}
						} else if cwith.FileName() != nil {
							withName.OneOf = &pb.ProcedureClause2_WithName_FileName{
								FileName: conv.FileName(cwith.FileName()),
							}
						}
						imp.ProcedureClause.WithNames = append(imp.ProcedureClause.WithNames, withName)
					}
				}
				if ilaf := ac.LibraryEntryProcedureUsingClause(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryEntryProcedureUsingClauseContext)
					for _, iusing := range laf.AllLibraryEntryProcedureUsingName() {
						cusing := iusing.(*cobol85.LibraryEntryProcedureUsingNameContext)
						usingName := &pb.ProcedureClause2_UsingName{}
						if cusing.DataName() != nil {
							usingName.OneOf = &pb.ProcedureClause2_UsingName_DataName{
								DataName: conv.DataName(cusing.DataName()),
							}
						} else if cusing.FileName() != nil {
							usingName.OneOf = &pb.ProcedureClause2_UsingName_FileName{
								FileName: conv.FileName(cusing.FileName()),
							}
						}
						imp.ProcedureClause.UsingNames = append(imp.ProcedureClause.UsingNames, usingName)
					}
				}
				if ilaf := ac.LibraryEntryProcedureGivingClause(); ilaf != nil {
					laf := ilaf.(*cobol85.LibraryEntryProcedureGivingClauseContext)
					imp.ProcedureClause.Giving = conv.DataName(laf.DataName())
				}
			}
			entry.OneOf = &pb.LibraryDescriptionEntry_Import_{
				Import: imp,
			}
		}
	}
	return v.VisitChildren(ctx)
}

func (v *ProgramLibrarySectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ProgramLibrarySectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
