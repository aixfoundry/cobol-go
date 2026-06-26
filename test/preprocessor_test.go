package test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/aixfoundry/cobol-go/document"
	"github.com/aixfoundry/cobol-go/format"
	"github.com/aixfoundry/cobol-go/options"
)

var (
	COBOL_EXTS = []string{".cbl", ".cob", ".jcl", ".txt", ""}
	DIRS       = []string{"fixed", "tandem", "variable"}
)

func TestPreprocessor(tt *testing.T) {
	rootdir := "./testdata/cobol/preprocessor"
	for _, dirName := range DIRS {
		tt.Run(dirName, func(t *testing.T) {
			parentDir := path.Join(rootdir, dirName)
			files, err := os.ReadDir(parentDir)
			if err != nil {
				t.Fatal(err)
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				ext := path.Ext(file.Name())
				isCobol := false
				for _, v := range COBOL_EXTS {
					if strings.EqualFold(v, ext) {
						isCobol = true
						break
					}
				}
				if !isCobol {
					continue
				}
				t.Log(file.Name())
				opts := options.NewOptions().AddCopyBookDirectory(parentDir).SetFormat(dir2format(dirName))
				filename := path.Join(parentDir, file.Name())
				processed, err := document.ParseFile(filename, opts)
				if err != nil {
					t.Fatal(err)
				}
				processedBuf, err := os.ReadFile(filename + ".preprocessed")
				if err != nil {
					t.Fatal(err)
				}
				if processed != string(processedBuf) {
					fmt.Println(processed)
					fmt.Println(strings.Repeat("-", 78))
					fmt.Println(string(processedBuf))
					t.FailNow()
				}
			}
		})
	}
}

func dir2format(name string) format.Format {
	f := format.FIXED
	switch name {
	case "fixed":
		f = format.FIXED
	case "tandem":
		f = format.TANDEM
	case "variable":
		f = format.VARIABLE
	default:
		f = format.FIXED
	}
	return f
}

func TestCopy(tt *testing.T) {
	rootdir := "./testdata/cobol/preprocessor/copy"
	dirs, err := os.ReadDir(rootdir)
	if err != nil {
		tt.Fatal(err)
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		tt.Run(dir.Name(), func(t *testing.T) {
			testCopy(path.Join(rootdir, dir.Name()), t)
		})
	}
}

func testCopy(rootdir string, t *testing.T) (err error) {
	files, err := os.ReadDir(rootdir)
	if err != nil {
		t.Fatal(err)
	}
	skips := []string{
		"testdata/cobol/preprocessor/copy/copyof/CopyOf.cbl",
		"testdata/cobol/preprocessor/copy/extension/precedence/variable/CopyPrecedence.cbl",
		"testdata/cobol/preprocessor/copy/extension/precedence/variable/copybooks/SomeCopyBook.cbl",
	}
FOR:
	for _, file := range files {
		filepath := path.Join(rootdir, file.Name())
		if file.IsDir() {
			testCopy(filepath, t)
			continue
		}
		if !strings.HasSuffix(file.Name(), ".cbl") {
			continue
		}
		basepath := path.Dir(filepath)
		copybooksPath := path.Join(basepath, "copybooks")
		parentName := path.Base(basepath)
		for _, v := range skips {
			if v == filepath {
				continue FOR
			}
		}
		t.Log(filepath)
		opts := options.NewOptions().AddCopyBookDirectory(copybooksPath).SetFormat(dir2format(parentName))
		processed, err := document.ParseFile(filepath, opts)
		if err != nil {
			t.Fatal(err)
		}
		processedBuf, err := os.ReadFile(filepath + ".preprocessed")
		if err != nil {
			t.Fatal(err)
		}
		if processed != string(processedBuf) {
			fmt.Println(processed)
			fmt.Println(strings.Repeat("-", 78))
			fmt.Println(string(processedBuf))
			t.FailNow()
		}
	}
	return
}

func TestExtension(t *testing.T) {
	rootdir := "./testdata/cobol/preprocessor/copy/extension/precedence"
	parentName := "variable"
	basepath := path.Join(rootdir, parentName)
	copybooksPath := path.Join(basepath, "copybooks")
	filepath := path.Join(basepath, "CopyPrecedence.cbl")

	opts := options.NewOptions().
		AddCopyBookDirectory(copybooksPath).
		SetFormat(dir2format(parentName)).
		AddCopyBookExtension("someotherextension").
		AddCopyBookExtension("txt").
		AddCopyBookExtension("cbl")
	processed, err := document.ParseFile(filepath, opts)
	if err != nil {
		t.Fatal(err)
	}
	processedBuf, err := os.ReadFile(filepath + ".preprocessed")
	if err != nil {
		t.Fatal(err)
	}
	if processed != string(processedBuf) {
		fmt.Println(processed)
		fmt.Println(strings.Repeat("-", 78))
		fmt.Println(string(processedBuf))
		t.FailNow()
	}
}

func TestPrefix(tt *testing.T) {
	rootdir := "./testdata/cobol/preprocessor/copy/copyprefix"
	filepath := path.Join(rootdir, "lbea0000.cbl")
	opts := options.NewOptions().AddCopyBookDirectory(rootdir).SetFormat(dir2format(rootdir))
	processed, err := document.ParseFile(filepath, opts)
	if err != nil {
		tt.Fatal(err)
	}
	os.WriteFile(filepath+".preprocessed", []byte(processed), os.ModePerm)
}
