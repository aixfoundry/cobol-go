package asg

import (
	"os"
	"testing"

	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestAnalyzeFile(t *testing.T) {
	opts := options.NewOptions().SetFormat(format.FIXED)
	program, err := AnalyzeFile("./testdata/HelloWorld.cbl", opts)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := protojson.Marshal(program)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
	if err := os.WriteFile("./testdata/HelloWorld.json", buf, 0o644); err != nil {
		t.Fatal(err)
	}
}
