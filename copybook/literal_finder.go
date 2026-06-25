package copybook

import (
	"io/fs"
	"path"
	fp "path/filepath"
	"strings"

	"github.com/aixfoundry/cobol-go/gen/preprocessor"
	"github.com/aixfoundry/cobol-go/options"
)

type LiteralFinder struct {
}

func NewLiteralFinder() *LiteralFinder {
	return &LiteralFinder{}
}

func (f *LiteralFinder) GetCopyBook(ctx preprocessor.ILiteralContext, opts *options.Options) string {
	for _, v := range opts.CopyBookFiles {
		if f.isMatching(ctx, v, "", opts) {
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

func (f *LiteralFinder) GetCopyBookFromDirectory(ctx preprocessor.ILiteralContext, dir string, opts *options.Options) (ret string) {
	infos := map[string]bool{}
	fp.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			infos[path] = d.IsDir()
		}
		return err
	})
	for filepath := range infos {
		if f.isMatching(ctx, filepath, dir, opts) {
			return filepath
		}
	}
	return
}

func (f *LiteralFinder) isMatching(ctx preprocessor.ILiteralContext, filepath, dir string, opts *options.Options) bool {
	identifier := TrimQuotes(ctx.GetText())
	identifier = strings.ReplaceAll(identifier, "\\", "/")
	if dir != "" {
		return f.isMatchingWithAbsolute(filepath, identifier, dir)
	} else {
		return f.isMatchingWithRelative(filepath, identifier)
	}
}

func (f *LiteralFinder) isMatchingWithAbsolute(filepath, identifier, dir string) bool {
	abspath, err := fp.Abs(filepath)
	if err != nil {
		return false
	}
	identifierPath, err := fp.Abs(path.Join(dir, identifier))
	if err != nil {
		return false
	}
	return strings.EqualFold(abspath, identifierPath)
}

func (f *LiteralFinder) isMatchingWithRelative(filepath, identifier string) bool {
	abspath, err := fp.Abs(filepath)
	if err != nil {
		return false
	}
	relativePath := ""
	if strings.HasPrefix(identifier, "/") ||
		strings.HasPrefix(identifier, "./") ||
		strings.HasPrefix(identifier, "\\") ||
		strings.HasPrefix(identifier, ".\\") {
		relativePath = identifier
	} else {
		relativePath = path.Join("/", identifier)
	}
	abspath = strings.ToLower(abspath)
	relativePath = strings.ToLower(relativePath)
	return strings.HasSuffix(abspath, relativePath)
}
