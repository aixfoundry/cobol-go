package environment

import (
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
	"github.com/antlr4-go/antlr/v4"
)

type ConfigurationSectionVisitor struct {
	cobol85.BaseCobol85Visitor
	section *pb.ConfigurationSection
}

func NewConfigurationSectionVisitor(section *pb.ConfigurationSection) *ConfigurationSectionVisitor {
	return &ConfigurationSectionVisitor{
		section: section,
	}
}

func (v *ConfigurationSectionVisitor) VisitSegmentLimitClause(ctx *cobol85.SegmentLimitClauseContext) any {
	var integerLiteral *pb.IntegerLiteral
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		integerLiteral = conv.IntegerLiteral(ictx)
	}
	v.section.ObjectComputerParagraph.SegmentLimitClause = &pb.SegmentLimitClause{
		SegmentLimit: integerLiteral,
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitCharacterSetClause(ctx *cobol85.CharacterSetClauseContext) any {
	v.section.ObjectComputerParagraph.CharacterSetClause = &pb.CharacterSetClause{}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitCollatingSequenceClause(ctx *cobol85.CollatingSequenceClauseContext) any {
	csc := &pb.CollatingSequenceClause{}
	for _, v := range ctx.AllAlphabetName() {
		csc.AlphabetNames = append(csc.AlphabetNames, conv.AlphabetName(v))
	}
	if ictx := ctx.CollatingSequenceClauseAlphanumeric(); ictx != nil {
		cctx := ictx.(*cobol85.CollatingSequenceClauseAlphanumericContext)
		csc.Alphanumeric = conv.AlphabetName(cctx.AlphabetName())
	}
	if ictx := ctx.CollatingSequenceClauseNational(); ictx != nil {
		cctx := ictx.(*cobol85.CollatingSequenceClauseNationalContext)
		csc.National = conv.AlphabetName(cctx.AlphabetName())
	}
	v.section.ObjectComputerParagraph.CollatingSequenceClause = csc
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitDiskSizeClause(ctx *cobol85.DiskSizeClauseContext) any {
	unit := pb.DiskSizeClause_WORDS
	switch {
	case ctx.WORDS() != nil:
		unit = pb.DiskSizeClause_WORDS
	case ctx.MODULES() != nil:
		unit = pb.DiskSizeClause_MODULES
	}
	diskSizeClause := &pb.DiskSizeClause{
		Unit: unit,
	}
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		diskSizeClause.DiskSize = &pb.DiskSizeClause_IntegerLiteral{
			IntegerLiteral: conv.IntegerLiteral(ictx),
		}
	} else if ictx := ctx.CobolWord(); ictx != nil {
		diskSizeClause.DiskSize = &pb.DiskSizeClause_CobolWord{
			CobolWord: conv.CobolWord(ictx),
		}
	}
	v.section.ObjectComputerParagraph.DiskSizeClause = diskSizeClause
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitMemorySizeClause(ctx *cobol85.MemorySizeClauseContext) any {
	unit := pb.MemorySizeClause_WORDS
	switch {
	case ctx.WORDS() != nil:
		unit = pb.MemorySizeClause_WORDS
	case ctx.CHARACTERS() != nil:
		unit = pb.MemorySizeClause_CHARACTERS
	case ctx.MODULES() != nil:
		unit = pb.MemorySizeClause_MODULES
	}

	memorySizeClause := &pb.MemorySizeClause{
		Unit: unit,
	}
	if ictx := ctx.IntegerLiteral(); ictx != nil {
		memorySizeClause.MemorySize = &pb.MemorySizeClause_IntegerLiteral{
			IntegerLiteral: conv.IntegerLiteral(ictx),
		}
	} else if ictx := ctx.CobolWord(); ictx != nil {
		memorySizeClause.MemorySize = &pb.MemorySizeClause_CobolWord{
			CobolWord: conv.CobolWord(ictx),
		}
	}
	v.section.ObjectComputerParagraph.MemorySizeClause = memorySizeClause
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitObjectComputerClause(ctx *cobol85.ObjectComputerClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitObjectComputerParagraph(ctx *cobol85.ObjectComputerParagraphContext) any {
	v.section.ObjectComputerParagraph = &pb.ObjectComputerParagraph{
		ComputerName: conv.ComputerName(ctx.ComputerName()),
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitSourceComputerParagraph(ctx *cobol85.SourceComputerParagraphContext) any {
	v.section.SourceComputerParagraph = &pb.SourceComputerParagraph{
		ComputerName:  conv.ComputerName(ctx.ComputerName()),
		DebuggingMode: ctx.DEBUGGING() != nil,
	}
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitConfigurationSectionParagraph(ctx *cobol85.ConfigurationSectionParagraphContext) any {
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) VisitConfigurationSection(ctx *cobol85.ConfigurationSectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *ConfigurationSectionVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ConfigurationSectionVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
