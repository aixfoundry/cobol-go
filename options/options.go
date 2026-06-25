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

func (o *Options) Apply(c *Options) {
	if o != c {
		(*c).Dialect = o.Dialect
		(*c).Format = o.Format
		(*c).CopyBookFiles = o.CopyBookFiles
		(*c).CopyBookDirectories = o.CopyBookDirectories
		(*c).CopyBookExtensions = o.CopyBookExtensions
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
