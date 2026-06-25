package procedure

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ParagraphVisitor struct {
	cobol85.BaseCobol85Visitor
	paragraph *pb.Paragraph
}

func NewParagraphVisitor(paragraph *pb.Paragraph) *ParagraphVisitor {
	return &ParagraphVisitor{
		paragraph: paragraph,
	}
}

func (v *ParagraphVisitor) VisitSentence(ctx *cobol85.SentenceContext) any {
	sentence := &pb.Sentence{}
	if v.paragraph.OneOf == nil {
		v.paragraph.OneOf = &pb.Paragraph_Sentences{
			Sentences: &pb.Sentences{},
		}
	}
	oneof := v.paragraph.OneOf.(*pb.Paragraph_Sentences)
	oneof.Sentences.Sentences = append(oneof.Sentences.Sentences, sentence)
	vr := NewSentenceVisitor(sentence)
	return vr.Visit(ctx)
}

func (v *ParagraphVisitor) VisitParagraph(ctx *cobol85.ParagraphContext) any {
	v.paragraph.ParagraphName = conv.ParagraphName(ctx.ParagraphName())
	if ctx.AlteredGoTo() != nil {
		v.paragraph.OneOf = &pb.Paragraph_AlteredGoTo{}
	}
	return v.VisitChildren(ctx)
}

func (v *ParagraphVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ParagraphVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
