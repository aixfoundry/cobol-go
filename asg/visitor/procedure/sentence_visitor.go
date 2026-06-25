package procedure

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/aixfoundry/cobol-go/asg/conv"
	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/aixfoundry/cobol-go/pb"
)

type SentenceVisitor struct {
	cobol85.BaseCobol85Visitor
	sentence *pb.Sentence
}

func NewSentenceVisitor(sentence *pb.Sentence) *SentenceVisitor {
	return &SentenceVisitor{
		sentence: sentence,
	}
}

func (v *SentenceVisitor) VisitStatement(ctx *cobol85.StatementContext) interface{} {
	v.sentence.Statements = append(v.sentence.Statements, conv.Statement(ctx))
	return v.VisitChildren(ctx)
}

func (v *SentenceVisitor) VisitSentence(ctx *cobol85.SentenceContext) any {
	return v.VisitChildren(ctx)
}

func (v *SentenceVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *SentenceVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}
