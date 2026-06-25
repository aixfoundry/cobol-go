package procedure

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type ParagraphsVisitor struct {
	cobol85.BaseCobol85Visitor
	paragraphs *pb.Paragraphs
}

func NewParagraphsVisitor(paragraphs *pb.Paragraphs) *ParagraphsVisitor {
	if paragraphs.Sentences == nil {
		paragraphs.Sentences = &pb.Sentences{}
	}
	return &ParagraphsVisitor{
		paragraphs: paragraphs,
	}
}

func (v *ParagraphsVisitor) VisitParagraph(ctx *cobol85.ParagraphContext) any {
	paragraph := &pb.Paragraph{}
	v.paragraphs.Paragraphs = append(v.paragraphs.Paragraphs, paragraph)
	vr := NewParagraphVisitor(paragraph)
	return vr.Visit(ctx)
}

func (v *ParagraphsVisitor) VisitSentence(ctx *cobol85.SentenceContext) any {
	sentence := &pb.Sentence{}
	v.paragraphs.Sentences.Sentences = append(v.paragraphs.Sentences.Sentences, sentence)
	vr := NewSentenceVisitor(sentence)
	return vr.Visit(ctx)
}

func (v *ParagraphsVisitor) VisitParagraphs(ctx *cobol85.ParagraphsContext) any {
	return v.VisitChildren(ctx)
}

func (v *ParagraphsVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ParagraphsVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
