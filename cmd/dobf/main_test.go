package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"testing"

	"github.com/aixfoundry/cobol-go/gen/cobol85"
	"github.com/antlr4-go/antlr/v4"
)

func TestDobf(t *testing.T) {
	buff, err := os.ReadFile("../../asg/testdata/HelloWorld.cbl.processed")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	processed := string(buff)
	is := antlr.NewInputStream(processed)
	lexer := cobol85.NewCobol85Lexer(is)
	cts := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	cpp := cobol85.NewCobol85Parser(cts)

	listener := NewDobfListener(cts)
	antlr.ParseTreeWalkerDefault.Walk(listener, cpp.StartRule())
	vars := map[string]string{}
	maps.Copy(vars, listener.GetVars())
	maps.Copy(vars, listener.GetFuncs())
	output := map[string]any{
		"vars":   vars,
		"tokens": listener.GetTokens(),
	}
	buff, err = json.Marshal(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s", string(buff))
}
