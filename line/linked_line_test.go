package line

import (
	"os"
	"testing"

	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
)

func TestLinkedLine(t *testing.T) {
	f, err := os.Open("./testdata/lbli0420.src")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer f.Close()
	o := options.NewOptions()
	o.SetFormat(format.FIXED)
	ll := NewLinkedLine(f, o)
	source := ll
	for {
		if source == nil {
			break
		}
		t.Log(source)
		source = source.next
	}
	source = ll
	code := Combine(source)
	err = os.WriteFile("./testdata/lbli0420.after", []byte(code), os.ModePerm)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
