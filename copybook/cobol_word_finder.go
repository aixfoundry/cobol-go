package copybook

import (
	"os"
	"path"
	"strings"

	"github.com/aixfoundry/cobol-go/gen/preprocessor"
	"github.com/aixfoundry/cobol-go/options"
)

type CobolWordFinder struct {
}

func NewCobolWordFinder() *CobolWordFinder {
	return &CobolWordFinder{}
}

func (f *CobolWordFinder) GetCopyBook(ctx preprocessor.ICobolWordContext, opts *options.Options) string {
	for _, v := range opts.CopyBookFiles {
		if f.isMatching(ctx, v, opts) {
			return v
		}
	}

	for _, v := range opts.CopyBookDirectories {
		filepath := f.GetCopyBookFromDirectory(ctx, v, opts)
		if filepath != "" {
			return filepath
		}
	}
	return ""
}

func (f *CobolWordFinder) GetCopyBookFromDirectory(ctx preprocessor.ICobolWordContext, dir string, opts *options.Options) (ret string) {
	infos, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	if len(infos) == 0 {
		return
	}
	for _, v := range infos {
		filepath := path.Join(dir, v.Name())
		if f.isMatching(ctx, filepath, opts) {
			return filepath
		}
	}
	return
}

func (f *CobolWordFinder) isMatching(ctx preprocessor.ICobolWordContext, filepath string, opts *options.Options) bool {
	identifier := ctx.GetText()
	if len(opts.CopyBookExtensions) != 0 {
		for _, v := range opts.CopyBookExtensions {
			if f.isMatchingWithExtension(filepath, identifier, v) {
				return true
			}
		}
		return false
	} else {
		return f.isMatchingWithoutExtension(filepath, identifier)
	}
}

func (f *CobolWordFinder) isMatchingWithExtension(filepath, identifier, extension string) bool {
	copyBookName := identifier
	if filepath != "" {
		copyBookName = identifier + "." + extension
	}
	filename := path.Base(filepath)
	return strings.EqualFold(filename, copyBookName)
}

func (f *CobolWordFinder) isMatchingWithoutExtension(filepath, identifier string) bool {
	ext := path.Ext(filepath)
	filename := path.Base(filepath)
	basename := strings.TrimSuffix(filename, ext)
	return strings.EqualFold(basename, identifier)
}
