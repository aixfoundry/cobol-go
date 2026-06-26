package options

import (
	"slices"

	"github.com/aixfoundry/cobol-go/format"
)

type Option interface {
	Apply(*Options)
}

type Options struct {
	Format              format.Format
	Dialect             format.Dialect
	CopyBookFiles       []string
	CopyBookDirectories []string
	CopyBookExtensions  []string
}

func NewOptions() (o *Options) {
	return &Options{}
}

// Apply copies all non-zero fields from o into c.
func (o *Options) Apply(c *Options) {
	if o.Dialect != 0 {
		c.Dialect = o.Dialect
	}
	if o.Format != 0 {
		c.Format = o.Format
	}
	if o.CopyBookFiles != nil {
		c.CopyBookFiles = o.CopyBookFiles
	}
	if o.CopyBookDirectories != nil {
		c.CopyBookDirectories = o.CopyBookDirectories
	}
	if o.CopyBookExtensions != nil {
		c.CopyBookExtensions = o.CopyBookExtensions
	}
}

func (o *Options) SetFormat(f format.Format) *Options {
	o.Format = f
	return o
}

func (o *Options) SetDialect(d format.Dialect) *Options {
	o.Dialect = d
	return o
}

func (o *Options) AddCopyBookFile(f string) *Options {
	if slices.Contains(o.CopyBookFiles, f) {
		return o
	}
	o.CopyBookFiles = append(o.CopyBookFiles, f)
	return o
}

func (o *Options) AddCopyBookDirectory(f string) *Options {
	if slices.Contains(o.CopyBookDirectories, f) {
		return o
	}
	o.CopyBookDirectories = append(o.CopyBookDirectories, f)
	return o
}

func (o *Options) AddCopyBookExtension(f string) *Options {
	if slices.Contains(o.CopyBookExtensions, f) {
		return o
	}
	o.CopyBookExtensions = append(o.CopyBookExtensions, f)
	return o
}
